package userhandler

import (
	entity "immersive/domains/users/entities"
	"immersive/exceptions"
	"immersive/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	Usecase entity.IusecaseUser
}

func New(usecase entity.IusecaseUser) *userHandler {
	return &userHandler{
		Usecase: usecase,
	}
}

func (h *userHandler) Store(c echo.Context) error {
	request := Request{}
	err := c.Bind(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(&request)
	if err != nil {
		return err
	}

	h.Usecase.Store(RequestToEntity(request))

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success create data"))
}

func (h *userHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	request := Request{}
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Bind(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(request)
	if err != nil {
		return err
	}

	userEntity := RequestToEntity(request)
	userEntity.UID = uint(id)

	err = h.Usecase.Update(userEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success update data"))
}

func (h *userHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = h.Usecase.Delete(entity.UserEntity{
		UID: uint(id),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success delete data"))
}

func (h *userHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	user, err := h.Usecase.GetById(entity.UserEntity{
		UID: uint(id),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(EntityToResponse(user)))
}
