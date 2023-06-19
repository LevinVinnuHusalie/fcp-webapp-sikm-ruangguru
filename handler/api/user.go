package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAPI interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetUserTaskCategory(c *gin.Context)
}

type userAPI struct {
	userService service.UserService
}

func NewUserAPI(userService service.UserService) *userAPI {
	return &userAPI{userService}
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.UserRegister

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" || user.Fullname == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("register data is empty"))
		return
	}

	var recordUser = model.User{
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
	}

	recordUser, err := u.userService.Register(&recordUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("error internal server"))
		return
	}

	c.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (u *userAPI) Login(c *gin.Context) {
	var user model.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid decode json"))
		return
	}

	var recordUser = model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	token, err := u.userService.Login(&recordUser)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			c.JSON(http.StatusBadRequest, model.NewErrorResponse("unregistered email error"))
			return
		}
		if errors.Is(err, errors.New("invalid password")) {
			c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid password"))
			return
		}
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("login error"))
		return
	}

	expiry15s := 15 * 60
	c.SetCookie("session_token", *token, expiry15s, "", "", false, true)

	c.JSON(http.StatusOK, model.NewSuccessResponse("login success"))
	// TODO: answer here
}

func (u *userAPI) GetUserTaskCategory(c *gin.Context) {
	list, err := u.userService.GetUserTaskCategory()
	if err != nil {
		c.JSON(http.StatusNotFound, model.NewErrorResponse("data not found"))
		return
	}

	c.JSON(http.StatusOK, list)
	// TODO: answer here
}
