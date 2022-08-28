package controller

import (
	"net/http"

	"gomysql/model"
	"gomysql/storage"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c echo.Context) error {
	db := storage.GetDBInstance()
	users := []model.User{}

	if err := db.Find(&users).Error; err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	db := storage.GetDBInstance()
	user := model.User{}

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	db := storage.GetDBInstance()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := db.Create(user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	db := storage.GetDBInstance()
	if err := db.Save(user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	db := storage.GetDBInstance()
	if err := db.Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "")
}


