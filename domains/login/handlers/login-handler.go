package loginhandler

import (
	entity "immersive/domains/login/entities"
	"immersive/exceptions"
	"immersive/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type loginHandler struct {
	Usecase entity.ILoginUsecase
}

func New(usecase entity.ILoginUsecase) loginHandler {
	return loginHandler{
		Usecase: usecase,
	}
}

func (h *loginHandler) Login(c echo.Context) error {
	request := request{}
	err := c.Bind(&request)
	if err != nil {
		return exceptions.NewBadRequestError(err.Error())
	}

	err = c.Validate(request)
	if err != nil {
		return err
	}

	token, role, err := h.Usecase.Login(requestToEntity(request))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, helpers.SuccessGetResponseData(map[string]interface{}{
		"token": token,
		"role":  role,
	}))
}
