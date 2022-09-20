package exceptions

type BadRequestErrorStruct struct {
	ErrorMsg string
}

func NewBadRequestError(msg string) *InternalServerErrorStruct {
	return &InternalServerErrorStruct{
		ErrorMsg: msg,
	}
}

func (e *BadRequestErrorStruct) Error() string {
	return e.ErrorMsg
}