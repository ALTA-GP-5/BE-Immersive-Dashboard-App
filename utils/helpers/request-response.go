package helpers

func SuccessGetResponseData(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"data":    data,
	}
}

func SuccessGetResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": msg,
	}
}

func FailedResponse(message string) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": message,
	}
}
