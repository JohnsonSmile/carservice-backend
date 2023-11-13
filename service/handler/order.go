package handler

import (
	"carservice/infra/database"
	"carservice/infra/logger"
	"carservice/service/request"
	"carservice/service/response"
	"encoding/base64"
	"log"
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
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "param error",
		})
		return
	}

	// car:highway:start:1:1:1001
	strs := strings.Split(string(original), ":")
	if len(strs) != 6 {
		logger.Error("param error", err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "param error",
		})
		return
	}

	// 从context中获取用户id
	userId := c.GetUint("userId")
	if userId == 0 {
		logger.Error("param error", err)
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "params error",
		})
		return
	}
	// 查找用户是否有订单
	order, err := database.GetLatestUnFinishedHighwayOrderByUserId(userId)
	if err != nil {
		logger.Error("get unfinished highway order", err)
		c.JSON(http.StatusInternalServerError, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "get unfinished highway order error",
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
				Code: http.StatusBadRequest,
				Msg:  "wrong qrcode, you should go to the entrance",
			})
			return
		}
		// TODO: find position
		cityID := strs[3]
		regionID := strs[4]
		placeID := strs[5]
		position, err := database.GetPosition(cityID, regionID, placeID)
		if err != nil {
			logger.Error("get position error", err)
			c.JSON(http.StatusBadRequest, &response.Response{
				Code: http.StatusBadRequest,
				Msg:  "get position error",
			})
			return
		}
		log.Println("error:", err)
		// seperat
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "success",
			Data: gin.H{
				"start_positon": position.Name,
				"end_positon":   "",
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
				Code: http.StatusBadRequest,
				Msg:  "wrong qrcode, you should go to the exit",
			})
			return
		}
		// end status
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "ready to pay",
			Data: gin.H{
				"start_positon": order.StartPosition.Name,
				"end_positon":   order.EndPosition.Name,
				"order_sn":      order.OrderSn,
				"status":        order.OrderStatus,
				"price":         order.Fee,
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
				Code: http.StatusBadRequest,
				Msg:  "you should pay previous order first.",
				Data: gin.H{
					"start_positon": order.StartPosition.Name,
					"end_positon":   order.EndPosition.Name,
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
				Code: http.StatusOK,
				Msg:  "ready to pay",
				Data: gin.H{
					"start_positon": order.StartPosition.Name,
					"end_positon":   order.EndPosition.Name,
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

}

func EndHighWay(c *gin.Context) {

}

// charge

func PreviewCharge(c *gin.Context) {

}

func StartCharge(c *gin.Context) {

}

func EndCharge(c *gin.Context) {

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
