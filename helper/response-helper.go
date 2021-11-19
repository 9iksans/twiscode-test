package helper

type Response struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildSuccessResponse(data interface{}) Response {
	response := Response{
		Status:  true,
		Code:    200,
		Message: "OK!",
		Data:    data,
	}
	return response
}

func BuildErrorResponse(code int, message string, data interface{}) Response {
	response := Response{
		Status:  false,
		Code:    code,
		Message: message,
		Data:    data,
	}
	return response
}
