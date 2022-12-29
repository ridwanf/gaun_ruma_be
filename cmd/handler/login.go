package handler

import (
	"gaunRumaRestApi/config"
	"gaunRumaRestApi/helpers"
	"gaunRumaRestApi/repository"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (u *UserHandler) GenerateHasPassword(c echo.Context) error {
	req := new(UserLoginRequest)
	hash, _ := helpers.HashPassword(req.UserPassword)

	return c.JSON(http.StatusOK, hash)
}

func (u *UserHandler) Login(c echo.Context) error {
	req := new(UserLoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	userLogin, err := u.userRepository.GetUserByUserName(req.UserName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}

	match, _ := helpers.CheckPasswordHash(req.UserPassword, userLogin.UserPassword)
	if !match {
		return echo.ErrUnauthorized

	}

	// generate tokens
	claims := &JwtCustomClaims{
		req.UserName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response
	conf, err := config.GetConfig(".")
	if err != nil {
		log.Fatal("cannot load config ", err)
	}

	t, err := token.SignedString([]byte(conf.SIGNING_KEY))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (u *UserHandler) Register(c echo.Context) error {
	req := new(UserLoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	hash, _ := helpers.HashPassword(req.UserPassword)

	userLogin, err := u.userRepository.CreateUser(req.UserName, hash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

	}
	return c.JSON(http.StatusOK, userLogin)
}
