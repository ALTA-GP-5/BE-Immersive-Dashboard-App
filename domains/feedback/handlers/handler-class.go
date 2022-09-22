package handlerfeedback

import (
	entity "immersive/domains/feedback/entities"
	"immersive/exceptions"
	"immersive/middlewares"
	"immersive/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	Usecase entity.IFeedBackUseCase
}

func New(usecase entity.IFeedBackUseCase) *feedbackHandler {
	return &feedbackHandler{
		Usecase: usecase,
	}
}

func (h *feedbackHandler) Create(c echo.Context) error {
	var feedBackRequest request

	err := c.Bind(&feedBackRequest)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(&feedBackRequest)

	if err != nil {
		return err
	}
	feedbackEntity := requestToEntity(feedBackRequest)

	uid, _ := middlewares.ExtractToken(c)

	feedbackEntity.MentorID = uint(uid)

	err = h.Usecase.Create(feedbackEntity)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponse("success create feedback"))
}

func (h *feedbackHandler) GetAll(c echo.Context) error {
	feedBackEntity := entity.FeedBackEntity{}

	id, err := strconv.Atoi(c.Param("mentee_id"))
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	feedBackEntity.MenteeID = uint(id)

	feedbackResultList, err := h.Usecase.GetAll(feedBackEntity)
	if err != nil {
		return err
	}

	var feedbackResponseList []response
	for _, feedbackResult := range feedbackResultList {
		feedbackResponseList = append(feedbackResponseList, EntityToResponse(feedbackResult))
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(feedbackResponseList))
}
