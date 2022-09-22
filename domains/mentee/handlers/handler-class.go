package handlerclass

import (
	"fmt"
	entity "immersive/domains/mentee/entities"
	"immersive/exceptions"
	"immersive/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type menteeHandler struct {
	Usecase entity.IMenteeUseCase
}

func New(usecase entity.IMenteeUseCase) *menteeHandler {
	return &menteeHandler{
		Usecase: usecase,
	}
}

func (h *menteeHandler) Create(c echo.Context) error {
	var menteeRequest request

	err := c.Bind(&menteeRequest)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(&menteeRequest)

	if err != nil {
		return err
	}

	err = h.Usecase.Create(requestToEntity(menteeRequest))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success create mentee"))
}

func (h *menteeHandler) Update(c echo.Context) error {
	var menteeRequest request

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Bind(&menteeRequest)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	fmt.Println(menteeRequest)

	err = c.Validate(&menteeRequest)

	if err != nil {
		return err
	}

	menteeEntity := requestToEntity(menteeRequest)
	menteeEntity.MenteeID = uint(id)

	err = h.Usecase.Update(menteeEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success update mentee"))
}

func (h *menteeHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = h.Usecase.Delete(entity.MenteeEntity{
		MenteeID: uint(id),
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success delete mentee"))
}

func (h *menteeHandler) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	classResult, err := h.Usecase.Get(entity.MenteeEntity{
		MenteeID: uint(id),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(EntityToResponse(classResult)))
}

func (h *menteeHandler) GetAll(c echo.Context) error {
	MenteeEntity := entity.MenteeEntity{}

	q := c.QueryParam("q")
	if q != "" {
		MenteeEntity.GeneralSearch = q
	}

	class := c.QueryParam("class")
	if class != "" {
		id, err := strconv.Atoi(c.QueryParam("class"))
		if err != nil {
			return exceptions.NewBadRequestError(err.Error())
		}
		MenteeEntity.ClassID = uint(id)
	}

	status := c.QueryParam("status")
	if status != "" {
		MenteeEntity.Status = status
	}

	category := c.QueryParam("category")
	if status != "" {
		MenteeEntity.Status = category
	}

	menteeResultList, err := h.Usecase.GetAll(MenteeEntity)
	if err != nil {
		return err
	}

	var menteeResponseList []response
	for _, menteeResult := range menteeResultList {
		menteeResponseList = append(menteeResponseList, EntityToResponse(menteeResult))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(menteeResponseList))
}
