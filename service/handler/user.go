package handler

import (
	"carservice/infra/database"
	"carservice/model"
	"carservice/service/request"
	"carservice/service/response"
	"carservice/util"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
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
			},
			"token": token,
		},
	})
}

func Register(c *gin.Context) {
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
			},
			"token": token,
		},
	})
}

func GetUserInfo(c *gin.Context) {
	zap.L().Debug("GetUserInfo")
}
