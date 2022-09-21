package mentorhandler

import (
	entity "immersive/domains/mentor/entities"
	"immersive/exceptions"
	"immersive/middlewares"
	"immersive/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type mentorHandler struct {
	Usecase entity.IMentorUsecase
}

func New(usecase entity.IMentorUsecase) mentorHandler {
	return mentorHandler{
		Usecase: usecase,
	}
}

func (h *mentorHandler) Create(c echo.Context) error {
	request := request{}
	err := c.Bind(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(request)
	if err != nil {
		return err
	}

	err = h.Usecase.Create(requestToEntity(request))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("mentor add success"))
}

func (h *mentorHandler) ReadAll(c echo.Context) error {
	mentorEntity := entity.MentorEntity{}

	q := c.QueryParam("q")
	if q != "" {
		mentorEntity.GeneralSearch = q
	}

	classStatus := c.QueryParam("class_status")
	if classStatus != "" {
		mentorEntity.ClassStatus = classStatus
	}

	mentorList, err := h.Usecase.ReadAll(mentorEntity)
	if err != nil {
		return err
	}

	var mentorResponseList []response
	for _, mentor := range mentorList {
		mentorResponseList = append(mentorResponseList, EntityToResponse(mentor))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(mentorResponseList))
}

func (h *mentorHandler) ReadById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	mentorEntity := entity.MentorEntity{
		MentorID: uint(id),
	}

	mentor, err := h.Usecase.ReadById(mentorEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(EntityToResponse(mentor)))
}

func (h *mentorHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	uid, _ := middlewares.ExtractToken(c)
	if uid != id {
		return exceptions.NewForbiddenError("forbidden access")
	}

	request := request{}
	err = c.Bind(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(request)
	if err != nil {
		return err
	}

	mentorEntity := requestToEntity(request)
	mentorEntity.MentorID = uint(id)

	err = h.Usecase.Update(mentorEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("mentor update success"))
}

func (h *mentorHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	mentorEntity := entity.MentorEntity{
		MentorID: uint(id),
	}

	err = h.Usecase.Delete(mentorEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("mentor delete success"))
}
