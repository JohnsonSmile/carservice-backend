package service

import (
	"carservice/infra/database"
	"carservice/infra/logger"
	"carservice/model"
	"carservice/service/request"
	"carservice/service/response"
	"carservice/util"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Service) Login(c *gin.Context) {
	req := request.LoginRequest{}
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error("param error", zap.Error(err))
		HandleValidatorError(c, err)
		return
	}

	user, err := database.GetUserByPhone(req.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "phone not exist",
		})
		return
	}

	isVerified := util.VerifyPasswordHash(user.Password, req.Password)
	if !isVerified {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "password not match",
		})
		return
	}
	j := util.NewJWT()
	token, err := j.CreateToken(model.CustomClaims{
		ID:          uint(user.ID),
		NickName:    user.Username,
		AuthorityId: 0, // normal, 1 admin
		StandardClaims: jwt.StandardClaims{
			// 30 days
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "carservice",
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "create token failed",
		})
		return
	}

	// 从链段获取用户信息，主要是score
	userInfo, err := s.chainClient.GetUserInfo(user.ID)
	if err != nil {
		// TODO:通过邮件通知开发人员即可，
		logger.Error("register user failed", err)

		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "create user successfully",
			Data: gin.H{
				"user": gin.H{
					"id":       user.ID,
					"username": user.Username,
					"phone":    user.Phone,
					"avatar":   user.Avatar,
					"bio":      user.Bio,
					"score":    0,
				},
				"token": token,
			},
		})
		return
	}

	score := userInfo.Score.Int64()

	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "create user successfully",
		Data: gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"phone":    user.Phone,
				"avatar":   user.Avatar,
				"bio":      user.Bio,
				"score":    score,
			},
			"token": token,
		},
	})
}

func (s *Service) Register(c *gin.Context) {
	req := request.RegisterRequest{}
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error("param error", zap.Error(err))
		HandleValidatorError(c, err)
		return
	}

	if req.Password != req.RePassword {
		c.JSON(200, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "password not match",
		})
		return
	}

	// password hash
	pwdHash, err := util.GeneratePasswordHash(req.Password)
	if err != nil {
		c.JSON(200, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "generate hash failed",
		})
		return
	}

	// save user in the database
	user, err := database.CreateUser(req.Phone, pwdHash)
	if err != nil {
		c.JSON(200, &response.Response{
			Code: http.StatusInternalServerError,
			Msg:  "create user failed",
		})
		return
	}
	j := util.NewJWT()
	token, err := j.CreateToken(model.CustomClaims{
		ID:          uint(user.ID),
		NickName:    user.Username,
		AuthorityId: 0, // normal, 1 admin
		StandardClaims: jwt.StandardClaims{
			// 30 days
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "carservice",
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, &response.Response{
			Code: http.StatusBadRequest,
			Msg:  "create token failed",
		})
		return
	}

	// 在链端注册用户信息
	// TODO: phone 一定是数字，这里直接忽略转化的错误
	p, _ := strconv.Atoi(user.Phone)
	err = s.chainClient.CreateUser(user.ID, p, 0)
	if err != nil {
		// TODO:这里通过日志或者邮件通知开发者，检查问题即可
		logger.Error("register user failed", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "create user successfully",
			Data: gin.H{
				"user": gin.H{
					"id":       user.ID,
					"username": user.Username,
					"phone":    user.Phone,
					"avatar":   user.Avatar,
					"bio":      user.Bio,
					"score":    0,
				},
				"token": token,
			},
		})
		return
	}

	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "create user successfully",
		Data: gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"phone":    user.Phone,
				"avatar":   user.Avatar,
				"bio":      user.Bio,
				"score":    0,
			},
			"token": token,
		},
	})
}

// 只是获取自己的信息，不存在社交属性
func (s *Service) GetUserInfo(c *gin.Context) {
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
	user, err := database.GetUserById(userId)
	if err != nil {
		logger.Error("get user error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.GET_USER_ERROR,
			Msg:  response.MsgForCode(response.GET_USER_ERROR),
		})
		return
	}

	// 从链端获取用户信息
	userInfo, err := s.chainClient.GetUserInfo(user.ID)
	if err != nil {
		// TODO: 通知开发者邮件信息即可
		logger.Error("get user error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: http.StatusOK,
			Msg:  "create user successfully",
			Data: gin.H{
				"user": gin.H{
					"id":       user.ID,
					"username": user.Username,
					"phone":    user.Phone,
					"avatar":   user.Avatar,
					"bio":      user.Bio,
					"score":    0,
				},
			},
		})
		return
	}

	score := userInfo.Score.Int64()

	c.JSON(http.StatusOK, &response.Response{
		Code: http.StatusOK,
		Msg:  "create user successfully",
		Data: gin.H{
			"user": gin.H{
				"id":       user.ID,
				"username": user.Username,
				"phone":    user.Phone,
				"avatar":   user.Avatar,
				"bio":      user.Bio,
				"score":    score,
			},
		},
	})

}

// TODO:
// 正常来说应该是通过微信支付等方式，给链充值，这里就不考虑那些了，直接给增加score
func (s *Service) ChargeScore(c *gin.Context) {
	req := request.ChargeScoreRequest{}
	if err := c.ShouldBind(&req); err != nil {
		zap.L().Error("param error", zap.Error(err))
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

	// TODO: 直接调用链端的充值接口
	err := s.chainClient.ChargeScore(int(userId), req.Score)
	if err != nil {
		logger.Error("charge score error", err)
		c.JSON(http.StatusOK, &response.Response{
			Code: response.CREATE_ORDER_FAILED,
			Msg:  response.MsgForCode(response.CREATE_ORDER_FAILED),
		})
		return
	}
	c.JSON(http.StatusOK, &response.Response{
		Code: response.SUCCESS,
		Msg:  response.MsgForCode(response.SUCCESS),
	})
}

// TODO: 用户信息更新,可以换头像，修改用户名之类的。。
func (s *Service) UpdateUserInfo(c *gin.Context) {
	logger.Debug("GetUserInfo", nil)
}
