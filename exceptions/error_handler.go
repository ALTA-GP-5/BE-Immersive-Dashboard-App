package exceptions

import (
	"fmt"
	"immersive/utils/helpers"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

/*
	Fungsi untuk membuat custom error.
*/
func CustomErrorHandling(err error, c echo.Context) {
	if notFoundError(err, c) {
		return
	} else if badRequestError(err, c) {
		return
	} else if forbiddenError(err, c) {
		return
	} else if validationError(err, c) {
		return
	} else {
		internalServerError(err, c)
		return
	}
}

func internalServerError(err error, c echo.Context) bool {
	response, ok := err.(*InternalServerErrorStruct)
	if ok {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})
		return true
	}
	return false
}

func badRequestError(err error, c echo.Context) bool {
	response, ok := err.(*BadRequestErrorStruct)
	if ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})
		return true
	}
	return false
}

func notFoundError(err error, c echo.Context) bool {
	response, ok := err.(*NotFoundErrorStruct)
	if ok {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})
		return true
	}
	return false
}

func forbiddenError(err error, c echo.Context) bool {
	response, ok := err.(*ForbiddenErrorStruct)
	if ok {
		c.JSON(http.StatusForbidden, map[string]interface{}{
			"success": false,
			"message": response.Error(),
		})
		return true
	}
	return false
}

func validationError(err error, c echo.Context) bool {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		var report string
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				report = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}
		}
		c.JSON(http.StatusBadRequest, helpers.FailedResponse(report))
		return true
	}
	return false
}
