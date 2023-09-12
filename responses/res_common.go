package responses

import (
	"golang_boilerplate_with_gin/utils"
	"net/http"
)

// COMMON RESPONSES

//var UnauthorizedResponse = []byte(`{"detail": "You are not authorized to perform this request. You have to login first."}`)

var UnauthorizedResponseData = utils.ResponseState{
	StatusCode: http.StatusUnauthorized,
	Code:       "UNAUTHORIZED_REQUEST_401",
	MessageEn:  "You are not authorized to perform this request. You have to login first.",
	MessageBn:  "",
}

var ResponseSuccess = utils.ResponseState{
	StatusCode: http.StatusOK,
	Code:       "REQUEST_SUCCESS_200",
	MessageEn:  "Your request is successful.",
	MessageBn:  "",
}

var ResponseFailed = utils.ResponseState{
	StatusCode: http.StatusBadRequest,
	Code:       "REQUEST_FAILED_400",
	MessageEn:  "Due to some unexpected issues your request cannot be processed now. Please try again later.",
	MessageBn:  "",
}

var ResponseDataHealthCheck = utils.ResponseState{
	StatusCode: http.StatusOK,
	Code:       "DATABASE_CONNECTION_SUCCESS_200",
	MessageEn:  "Your request is successful.",
	MessageBn:  "",
}

var ResponseDataHealthCheckFailed = utils.ResponseState{
	StatusCode: http.StatusInternalServerError,
	Code:       "DATABASE_CONNECTION_FAILED_500",
	MessageEn:  "Your request is failed.",
	MessageBn:  "",
}

var ResponseInvalidRequest = utils.ResponseState{
	StatusCode: http.StatusBadRequest,
	Code:       "INVALID_REQUEST_400",
	MessageEn:  "The provided information is invalid. Please recheck and try again.",
	MessageBn:  "",
}

var ResponseLogoutSuccess = utils.ResponseState{
	StatusCode: http.StatusOK,
	Code:       "REQUEST_SUCCESS_200",
	MessageEn:  "Logout Success!",
	MessageBn:  "",
}
