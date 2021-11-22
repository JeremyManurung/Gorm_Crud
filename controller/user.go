package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"strconv"
	"res-task/db"
	"res-task/model"
)


func GetUserC(echoContext echo.Context) error {
	users, err := db.GetUsers()
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}

func GetUserbyIDC(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	users, err := db.GetbyIDUser(id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}


func SaveBookC(echoContext echo.Context) error {
	var userReq model.User
	echoContext.Bind(&userReq)

	result, err := db.StoreUser(userReq)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func UpdateUserbyIDC(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	var userReq model.User
	echoContext.Bind(&userReq)

	result, err := db.UpdateUser(id, userReq)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   result,
	})
}

func DeleteUserbyIDC(echoContext echo.Context) error {

	id, _ := strconv.Atoi(echoContext.Param("id"))

	users, err := db.DeleteUser(id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":   "err",
			"messages": err,
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}