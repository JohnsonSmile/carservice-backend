package handler

import (
	"carservice/infra/database"
	"carservice/infra/logger"
	"carservice/model"
	"carservice/service/request"
	"carservice/service/response"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 不需要考虑并发问题
func PreviewHighWay(c *gin.Context) {
	req := request.PreviewRequest{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// TODO：查找信息
	original, err := base64.RawStdEncoding.DecodeString(req.Data)
	if err != nil {
		logger.Error("param error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PARAM_ERROR,
			Msg:  response.MsgForCode(response.PARAM_ERROR),
		})
		return
	}

	// car:highway:start:1:1:1001
	strs := strings.Split(string(original), ":")
	if len(strs) != 6 {
		logger.Error("param error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PARAM_ERROR,
			Msg:  response.MsgForCode(response.PARAM_ERROR),
		})
		return
	}

	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.UNAUTHORIZED_ERROR,
			Msg:  response.MsgForCode(response.UNAUTHORIZED_ERROR),
		})
		return
	}
	// 查找用户是否有订单
	order, err := database.GetLatestUnFinishedHighwayOrderByUserId(userId)
	if err != nil {
		logger.Error("get unfinished highway order", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_UNFINISHED_ORDER_ERROR,
			Msg:  response.MsgForCode(response.GET_UNFINISHED_ORDER_ERROR),
		})
		return
	}

	// start or end
	action := strs[2]
	// 如果没有 Highway 订单，那么当前只能是start，如果是end那就是走错口了，甚至可能是逆行。
	if order == nil {
		// action == start
		if action != "start" {
			c.JSON(http.StatusOK, &response.Response{
				Code: response.WRONG_QRCODE_ENTRANCE_ERROR,
				Msg:  response.MsgForCode(response.WRONG_QRCODE_ENTRANCE_ERROR),
			})
			return
		}
		// TODO: find position
		cityID := strs[3]
		regionID := strs[4]
		placeID := strs[5]
		position, err := database.GetPosition(1, cityID, regionID, placeID)

		if err != nil {
			logger.Error("get position error", err)
			c.JSON(http.StatusOK, &response.Response{
				Code: response.GET_POSITION_ERROR,
				Msg:  response.MsgForCode(response.GET_POSITION_ERROR),
			})
			return
		}
		// seperat
		c.JSON(http.StatusOK, &response.Response{
			Code: response.SUCCESS,
			Msg:  response.MsgForCode(response.SUCCESS),
			Data: gin.H{
				"start_positon": position.Name,
				"start_id":      position.ID,
				"end_positon":   "",
				"end_id":        0,
				"order_sn":      "",
				"status":        -1, // -1 表示没有订单,not created
				"price":         0,
				"start_at":      "",
				"end_at":        "",
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	// 查看最近的 Highway 订单的状态，如果是start（0）的，则当前的应该是end才对，如果不是end，提示状态信息
	// start status
	if order.OrderStatus == 0 {
		if action != "end" {
			c.JSON(http.StatusOK, &response.Response{
				Code: response.WRONG_QRCODE_EXIT_ERROR,
				Msg:  response.MsgForCode(response.WRONG_QRCODE_EXIT_ERROR),
			})
			return
		}

		// get current end position

		// TODO: find position
		cityID := strs[3]
		regionID := strs[4]
		placeID := strs[5]
		position, err := database.GetPosition(1, cityID, regionID, placeID)
		if err != nil {
			logger.Error("get position error", err)
			c.JSON(http.StatusOK, &response.Response{
				Code: response.GET_POSITION_ERROR,
				Msg:  response.MsgForCode(response.GET_POSITION_ERROR),
			})
			return
		}

		// query fee
		fee, err := database.GetHighwayFee(order.StartPositionID, position.ID)
		if err != nil {
			logger.Error("get fee error", err)
			c.JSON(http.StatusOK, &response.Response{
				Code: response.GET_FEE_ERROR,
				Msg:  response.MsgForCode(response.GET_FEE_ERROR),
			})
			return
		}

		// end status
		c.JSON(http.StatusOK, &response.Response{
			Code: response.SUCCESS,
			Msg:  response.MsgForCode(response.SUCCESS),
			Data: gin.H{
				"start_positon": order.StartPosition.Name,
				"start_id":      order.StartPosition.ID,
				"end_positon":   position.Name,
				"end_id":        position.ID,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         fee,
				"start_at":      order.StartAt,
				"end_at":        time.Now(), // 只是preview，并不更改数据库
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	// 如果最近的 Highway 订单的状态是end（1），但是不是payed（2）状态，当前如果是end则返回未支付的订单状态，如果扫描的是start，则提示还有未支付的订单
	if order.OrderStatus == 1 {
		if action == "start" {
			c.JSON(http.StatusOK, &response.Response{
				Code: response.PREVIOUS_ORDER_FIRST_ERROR,
				Msg:  response.MsgForCode(response.PREVIOUS_ORDER_FIRST_ERROR),
				Data: gin.H{
					"start_positon": order.StartPosition.Name,
					"start_id":      order.StartPosition.ID,
					"end_positon":   order.EndPosition.Name,
					"end_id":        order.EndPosition.ID,
					"order_sn":      order.OrderSn,
					"status":        order.OrderStatus,
					"price":         order.Fee,
					"start_at":      order.StartAt,
					"end_at":        order.EndAt,
					// TODO: 数字人民币的支付
				},
			})
			return
		} else {
			// TODO: 其实应该看看当前的id和之前的订单的id是不是相同，这里就不做这么复杂了
			// 准备支付
			c.JSON(http.StatusOK, &response.Response{
				Code: response.SUCCESS,
				Msg:  response.MsgForCode(response.SUCCESS),
				Data: gin.H{
					"start_positon": order.StartPosition.Name,
					"start_id":      order.StartPosition.ID,
					"end_positon":   order.EndPosition.Name,
					"end_id":        order.EndPosition.ID,
					"order_sn":      order.OrderSn,
					"status":        order.OrderStatus,
					"price":         order.Fee,
					"start_at":      order.StartAt,
					"end_at":        order.EndAt,
					// TODO: 数字人民币的支付
				},
			})
			return
		}

	}

}

func StartHighWay(c *gin.Context) {
	req := request.StartRequest{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// user id
	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.UNAUTHORIZED_ERROR,
			Msg:  response.MsgForCode(response.UNAUTHORIZED_ERROR),
		})
		return
	}

	// 查看当前用户有没有未完成的highway订单，如果有则状态异常,提示用户完成之前的订单，或者可能他走错口了？？
	order, err := database.GetLatestUnFinishedHighwayOrderByUserId(userId)
	if err != nil {
		logger.Error("get unfinished highway order", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_UNFINISHED_ORDER_ERROR,
			Msg:  response.MsgForCode(response.GET_UNFINISHED_ORDER_ERROR),
		})
		return
	}
	if order != nil {
		// not finished order
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PREVIOUS_ORDER_FIRST_ERROR,
			Msg:  response.MsgForCode(response.PREVIOUS_ORDER_FIRST_ERROR),
			Data: gin.H{
				"start_positon": order.StartPosition.Name,
				"start_id":      order.StartPosition.ID,
				"end_positon":   order.EndPosition.Name,
				"end_id":        order.EndPosition.ID,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         order.Fee,
				"start_at":      order.StartAt,
				"end_at":        order.EndAt,
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	now := time.Now()
	// ETC 2023(年) 11(月) 11(日) 06(时) 10(分) 12(秒) XXXX(用户id.hex())
	userIdHex := hex.EncodeToString([]byte(fmt.Sprintf("%d", userId)))
	orderSn := fmt.Sprintf("ETC%d%d%d%d%d%d%06s", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), userIdHex)

	newOrder := &model.Order{
		OrderSn:         orderSn, // TODO: generate order sn with type
		OrderTypeID:     1,
		OrderStatus:     0,
		UserID:          int(userId),
		StartAt:         &now,
		EndAt:           nil,
		Fee:             0,
		StartPositionID: req.PositionID,
		EndPositionID:   0,
	}

	position, err := database.GetPositionByID(req.PositionID)
	if err != nil {
		logger.Error("get position failed", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_POSITION_ERROR,
			Msg:  response.MsgForCode(response.GET_POSITION_ERROR),
		})
		return
	}

	if err := database.CreateOrder(newOrder); err != nil {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.CREATE_ORDER_FAILED,
			Msg:  response.MsgForCode(response.CREATE_ORDER_FAILED),
		})
		return
	}

	c.JSON(http.StatusOK, &response.Response{
		Code: response.SUCCESS,
		Msg:  response.MsgForCode(response.SUCCESS),
		Data: gin.H{
			"start_positon": position.Name,
			"start_id":      position.ID,
			"end_positon":   "",
			"end_id":        0,
			"order_sn":      orderSn,
			"status":        0,
			"price":         0,
			"start_at":      now,
			"end_at":        "",
			// TODO: 数字人民币的支付
		},
	})

}

func EndHighWay(c *gin.Context) {
	req := request.EndRequest{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// user id
	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "params error",
		})
		return
	}

	// get order with order sn
	order, err := database.GetHighwayOrderByOrderSn(req.OrderSn)
	if err != nil {
		logger.Error("get highway order by order sn", err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "get highway order by order sn error",
		})
		return
	}

	if order.EndPositionID > 0 {
		// 说明已经end过了，提示前端支付
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: gin.H{
				"start_positon": order.StartPosition.Name,
				"start_id":      order.StartPosition.ID,
				"end_positon":   order.EndPosition.Name,
				"end_id":        order.EndPosition.ID,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         order.Fee,
				"start_at":      order.StartAt,
				"end_at":        order.EndAt, // 应该已经写入了
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	// update order status and end

	position, err := database.GetPositionByID(req.PositionID)
	if err != nil {
		logger.Error("get position error", err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "get position error",
		})
		return
	}
	order.EndPositionID = position.ID

	// query fee
	fee, err := database.GetHighwayFee(order.StartPositionID, position.ID)

	if err != nil {
		logger.Error("get fee error", err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "get fee error",
		})
		return
	}
	order.Fee = fee
	now := time.Now()
	order.EndAt = &now
	err = database.EndOrder(order)
	if err != nil {
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "end order failed",
		})
		return
	}

	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: gin.H{
			"start_positon": order.StartPosition.Name,
			"start_id":      order.StartPosition.ID,
			"end_positon":   position.Name,
			"end_id":        position.ID,
			"order_sn":      order.OrderSn,
			"status":        1,
			"price":         order.Fee,
			"start_at":      order.StartAt,
			"end_at":        now,
			// TODO: 数字人民币的支付

		},
	})
}

func GetHighWayOrders(c *gin.Context) {
	req := request.HighwayOrdersRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// user id
	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "params error",
		})
		return
	}

	orders, total, err := database.GetHighwayOrders(req.Page, req.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "get highway orders failed",
		})
		return
	}
	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: gin.H{
			"orders": orders,
			"total":  total,
		},
	})
}

// charge

func PreviewCharge(c *gin.Context) {

	req := request.PreviewRequest{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// TODO：查找信息
	original, err := base64.RawStdEncoding.DecodeString(req.Data)
	if err != nil {
		logger.Error("param error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PARAM_ERROR,
			Msg:  response.MsgForCode(response.PARAM_ERROR),
		})
		return
	}

	// car:charge:start:1:1:11
	strs := strings.Split(string(original), ":")
	if len(strs) != 6 {
		logger.Error("param error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PARAM_ERROR,
			Msg:  response.MsgForCode(response.PARAM_ERROR),
		})
		return
	}

	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.UNAUTHORIZED_ERROR,
			Msg:  response.MsgForCode(response.UNAUTHORIZED_ERROR),
		})
		return
	}
	// 查找用户是否有订单
	order, err := database.GetLatestUnFinishedChargeOrderByUserId(userId)
	if err != nil {
		logger.Error("get unfinished charge order", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_UNFINISHED_ORDER_ERROR,
			Msg:  response.MsgForCode(response.GET_UNFINISHED_ORDER_ERROR),
		})
		return
	}

	// start or end
	action := strs[2]
	// 如果没有 charge 订单，那么当前只能是start，如果是end那就是扫错二维码了。
	if order == nil {
		if action != "start" {
			c.JSON(http.StatusOK, &response.Response{
				Code: response.WRONG_QRCODE_CHARGE_START_ERROR,
				Msg:  response.MsgForCode(response.WRONG_QRCODE_CHARGE_START_ERROR),
			})
			return
		}
		// TODO: find position
		cityID := strs[3]
		regionID := strs[4]
		placeID := strs[5]
		position, err := database.GetPosition(2, cityID, regionID, placeID)

		if err != nil {
			logger.Error("get position error", err)
			c.JSON(http.StatusOK, &response.Response{
				Code: response.GET_POSITION_ERROR,
				Msg:  response.MsgForCode(response.GET_POSITION_ERROR),
			})
			return
		}

		orderSn := fmt.Sprintf("CG%04d%04d%04d", position.CityID, position.RegionID, position.PlaceID)

		// 获取fee
		fee, err := database.GetChargeFee(position.ID)
		if err != nil {
			logger.Error("get position error", err)
			c.JSON(http.StatusOK, &response.Response{
				Code: response.GET_FEE_ERROR,
				Msg:  response.MsgForCode(response.GET_FEE_ERROR),
			})
			return
		}

		// seperat
		c.JSON(http.StatusOK, &response.Response{
			Code: response.SUCCESS,
			Msg:  response.MsgForCode(response.SUCCESS),
			Data: gin.H{
				"id":            0,
				"start_positon": position.Name,
				"start_id":      position.ID,
				"end_positon":   position.Name,
				"end_id":        position.ID,
				"order_sn":      orderSn,
				"status":        -1,  // -1 表示没有订单,not created
				"price":         fee, // fee
				"degree":        0,
				"start_at":      "",
				"end_at":        "",
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	// 查看最近的 charge 订单的状态，如果是start（0）的，则当前的应该是end才对，如果不是end，提示状态信息
	// start status
	if order.OrderStatus == 0 {
		if action != "end" {
			c.JSON(http.StatusOK, &response.Response{
				Code: response.WRONG_QRCODE_CHARGE_END_ERROR,
				Msg:  response.MsgForCode(response.WRONG_QRCODE_CHARGE_END_ERROR),
			})
			return
		}

		// end status
		c.JSON(http.StatusOK, &response.Response{
			Code: response.SUCCESS,
			Msg:  response.MsgForCode(response.SUCCESS),
			Data: gin.H{
				"id":            order.ID,
				"start_positon": order.StartPosition.Name,
				"start_id":      order.StartPosition.ID,
				"end_positon":   order.StartPosition.Name,
				"end_id":        order.StartPosition.ID,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         order.Fee,
				"start_at":      order.StartAt,
				"end_at":        time.Now(), // 只是preview，并不更改数据库
				"degree":        500,        // TODO:这里固定给5度得了，正常应该是充电桩和后台有一个mqtt订阅或者ws传输的过程
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	// 如果最近的 charge 订单的状态是end（1），但是不是payed（2）状态，当前如果是end则返回未支付的订单状态，如果扫描的是start，则提示还有未支付的订单
	if order.OrderStatus == 1 {
		if action == "start" {
			c.JSON(http.StatusOK, &response.Response{
				Code: response.PREVIOUS_ORDER_FIRST_ERROR,
				Msg:  response.MsgForCode(response.PREVIOUS_ORDER_FIRST_ERROR),
				Data: gin.H{
					"id":            order.ID,
					"start_positon": order.StartPosition.Name,
					"start_id":      order.StartPosition.ID,
					"end_positon":   order.EndPosition.Name,
					"end_id":        order.EndPosition.ID,
					"order_sn":      order.OrderSn,
					"status":        order.OrderStatus,
					"price":         order.Fee,
					"start_at":      order.StartAt,
					"end_at":        order.EndAt,
					"degree":        1050, // TODO: 这里给一个固定数字就行， 105=10.5度电
					// TODO: 数字人民币的支付
				},
			})
			return
		} else {
			// TODO: 其实应该看看当前的id和之前的订单的id是不是相同，这里就不做这么复杂了
			// 准备支付
			c.JSON(http.StatusOK, &response.Response{
				Code: response.SUCCESS,
				Msg:  response.MsgForCode(response.SUCCESS),
				Data: gin.H{
					"id":            order.ID,
					"start_positon": order.StartPosition.Name,
					"start_id":      order.StartPosition.ID,
					"end_positon":   order.EndPosition.Name,
					"end_id":        order.EndPosition.ID,
					"order_sn":      order.OrderSn,
					"status":        order.OrderStatus,
					"price":         order.Fee,
					"start_at":      order.StartAt,
					"end_at":        order.EndAt,
					"degree":        1050, // TODO: 这里给一个固定数字就行，
					// TODO: 数字人民币的支付
				},
			})
			return
		}

	}

}

func StartCharge(c *gin.Context) {
	req := request.StartRequest{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// user id
	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.UNAUTHORIZED_ERROR,
			Msg:  response.MsgForCode(response.UNAUTHORIZED_ERROR),
		})
		return
	}

	// 查看当前用户有没有未完成的charge订单，如果有则状态异常,提示用户完成之前的订单
	order, err := database.GetLatestUnFinishedChargeOrderByUserId(userId)
	if err != nil {
		logger.Error("get unfinished charge order", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_UNFINISHED_ORDER_ERROR,
			Msg:  response.MsgForCode(response.GET_UNFINISHED_ORDER_ERROR),
		})
		return
	}

	if order != nil {
		// not finished order
		c.JSON(http.StatusOK, &response.Response{
			Code: response.PREVIOUS_ORDER_FIRST_ERROR,
			Msg:  response.MsgForCode(response.PREVIOUS_ORDER_FIRST_ERROR),
			Data: gin.H{
				"id":            order.ID,
				"start_positon": order.StartPosition.Name,
				"start_id":      order.StartPosition.ID,
				"end_positon":   order.EndPosition.Name,
				"end_id":        order.EndPosition.ID,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         order.Fee,
				"start_at":      order.StartAt,
				"end_at":        order.EndAt,
				"degree":        order.UniteCount, // 精度100,1度表示为100
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	position, err := database.GetPositionByID(req.PositionID)
	if err != nil {
		logger.Error("get position failed", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_POSITION_ERROR,
			Msg:  response.MsgForCode(response.GET_POSITION_ERROR),
		})
		return
	}

	// 获取fee
	fee, err := database.GetChargeFee(position.ID)
	if err != nil {
		logger.Error("get position error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_FEE_ERROR,
			Msg:  response.MsgForCode(response.GET_FEE_ERROR),
		})
		return
	}

	// create order
	now := time.Now()
	orderSn := fmt.Sprintf("CG%04d%04d%04d", position.CityID, position.RegionID, position.PlaceID)

	newOrder := &model.Order{
		OrderSn:         orderSn, // TODO: generate order sn with type
		OrderTypeID:     2,
		OrderStatus:     0,
		UserID:          int(userId),
		StartAt:         &now,
		EndAt:           nil,
		Fee:             fee,
		UniteCount:      0,
		StartPositionID: req.PositionID,
		EndPositionID:   req.PositionID,
	}

	if err := database.CreateOrder(newOrder); err != nil {
		c.JSON(http.StatusOK, &response.Response{
			Code: response.CREATE_ORDER_FAILED,
			Msg:  response.MsgForCode(response.CREATE_ORDER_FAILED),
		})
		return
	}

	c.JSON(http.StatusOK, &response.Response{
		Code: response.SUCCESS,
		Msg:  response.MsgForCode(response.SUCCESS),
		Data: gin.H{
			"id":            newOrder.ID,
			"start_positon": position.Name,
			"start_id":      position.ID,
			"end_positon":   position.Name,
			"end_id":        position.ID,
			"order_sn":      orderSn,
			"status":        0,   // start
			"price":         fee, // fee
			"degree":        0,
			"start_at":      now,
			"end_at":        "",
			// TODO: 数字人民币的支付
		},
	})

}

func EndCharge(c *gin.Context) {

	req := request.ChargeEndRequest{}
	if err := c.ShouldBind(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// user id
	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "params error",
		})
		return
	}

	// 根据order id 获取order
	order, err := database.GetChargeOrderByID(req.ID)
	if err != nil {
		logger.Error("get order error", err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "get order error",
		})
		return
	}

	if order.UniteCount > 0 {
		// 说明已经end过了，提示前端支付
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: gin.H{
				"id":            order.ID,
				"start_positon": order.StartPosition.Name,
				"start_id":      order.StartPosition.ID,
				"end_positon":   order.EndPosition.Name,
				"end_id":        order.EndPosition.ID,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         order.Fee,
				"start_at":      order.StartAt,
				"end_at":        order.EndAt, // 应该已经写入了
				"degree":        order.UniteCount,
				// TODO: 数字人民币的支付
			},
		})
		return
	}

	// update order status and endtime and unite count
	now := time.Now()
	order.EndAt = &now
	order.UniteCount = 1050 // TODO: 假数据，就默认消耗了10.5度
	err = database.EndChargeOrder(order)
	if err != nil {
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "end order failed",
		})
		return
	}

	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: gin.H{
			"id":            order.ID,
			"start_positon": order.StartPosition.Name,
			"start_id":      order.StartPosition.ID,
			"end_positon":   order.StartPosition.Name,
			"end_id":        order.StartPosition.ID,
			"order_sn":      order.OrderSn,
			"status":        1,
			"price":         order.Fee,
			"start_at":      order.StartAt,
			"end_at":        now,
			"degree":        order.UniteCount,
			// TODO: 数字人民币的支付
		},
	})

}

func GetChargeOrders(c *gin.Context) {
	req := request.HighwayOrdersRequest{}
	if err := c.ShouldBindQuery(&req); err != nil {
		logger.Error("param error", err)
		HandleValidatorError(c, err)
		return
	}

	// user id
	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "params error",
		})
		return
	}

	orders, total, err := database.GetChargeOrders(req.Page, req.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "get highway orders failed",
		})
		return
	}
	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: gin.H{
			"orders": orders,
			"total":  total,
		},
	})
}

// park

func PreviewPark(c *gin.Context) {

}

func StartPark(c *gin.Context) {

}

func EndPark(c *gin.Context) {

}

func PayOrder(c *gin.Context) {

}
