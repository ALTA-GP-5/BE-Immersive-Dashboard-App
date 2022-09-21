package handlerclass

import (
	entity "immersive/domains/class/entities"
	"immersive/exceptions"
	"immersive/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type classHandler struct {
	Usecase entity.IClassUseCase
}

func New(usecase entity.IClassUseCase) *classHandler {
	return &classHandler{
		Usecase: usecase,
	}
}

func (h *classHandler) Create(c echo.Context) error {
	var classRequest request

	err := c.Bind(&classRequest)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(&classRequest)

	if err != nil {
		return err
	}

	err = h.Usecase.Create(requestToEntity(classRequest))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success create class"))
}

func (h *classHandler) Update(c echo.Context) error {
	var classRequest request

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Bind(&classRequest)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(&classRequest)

	if err != nil {
		return err
	}

	clasEntity := requestToEntity(classRequest)
	clasEntity.ClassID = uint(id)

	err = h.Usecase.Update(clasEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success update class"))
}

func (h *classHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = h.Usecase.Delete(entity.ClassEntity{
		ClassID: uint(id),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success delete class"))
}

func (h *classHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	classResult, err := h.Usecase.Get(entity.ClassEntity{
		ClassID: uint(id),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(EntityToResponse(classResult)))
}

func (h *classHandler) GetAll(c echo.Context) error {
	classEntity := entity.ClassEntity{}

	q := c.QueryParam("q")
	if q != "" {
		classEntity.GeneralSearch = q
	}

	classResultList, err := h.Usecase.GetAll(classEntity)
	if err != nil {
		return err
	}

	var classResponseList []response
	for _, classResult := range classResultList {
		classResponseList = append(classResponseList, EntityToResponse(classResult))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(classResponseList))
}
