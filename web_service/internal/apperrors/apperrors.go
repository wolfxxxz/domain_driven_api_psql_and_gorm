package apperrors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message  string `json:"message"`
	Code     string
	HTTPCode int
}

func NewAppError() *AppError {
	return &AppError{}
}

var (
	//INIT_ERRORS
	EnvConfigLoadError = AppError{
		Message:  "Failed to parse env file",
		Code:     "ENV_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	EnvConfigParseError = AppError{
		Message:  "Failed to parse env file",
		Code:     "ENV_PARSE_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	InitPostgressErr = AppError{
		Message:  "Failed to InitPostgress",
		Code:     "INIT_CLIENT_POSTGRESS_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	NewLoggerErr = AppError{
		Message:  "Failed to NewLog",
		Code:     "NEW_LOG_AND_SET_LEVEL_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	SetupDatabaseErr = AppError{
		Message:  "Failed to SetupDatabase",
		Code:     "DATABASE_POSTGRES_ERR",
		HTTPCode: http.StatusInternalServerError,
	}
	//REPO
	CreateUserErr = AppError{
		Message:  "Failed to CreateUser",
		Code:     "User_REPO",
		HTTPCode: http.StatusNotFound,
	}
	CreateLectureErr = AppError{
		Message:  "Failed to CreateLecture",
		Code:     "Lecture_REPO",
		HTTPCode: http.StatusNotFound,
	}
	AddStudentToLectureRepoErr = AppError{
		Message:  "Failed to AddStudentToLectureErr",
		Code:     "Lecture_REPO",
		HTTPCode: http.StatusNotFound,
	}
	DropUserFromLectureErr = AppError{
		Message:  "Failed to DropUserFromLectureErr",
		Code:     "Lecture_REPO",
		HTTPCode: http.StatusNotFound,
	}
	GetLecturesStudentsPPErr = AppError{
		Message:  "Failed to GetAllLecturesAndStudentsErr",
		Code:     "Lecture_REPO",
		HTTPCode: http.StatusNotFound,
	}
	//HANDLERS
	CreateUserHandlerErr = AppError{
		Message:  "Failed to createUserHandlerErr",
		Code:     "Server_handlers",
		HTTPCode: http.StatusBadRequest,
	}
	AddStudentToLectureHandlerErr = AppError{
		Message:  "Failed to addUserToLectureHandlerErr",
		Code:     "Server_handlers",
		HTTPCode: http.StatusBadRequest,
	}
	DeleteUserFromLectureHandlerERR = AppError{
		Message:  "Failed to deleteUserFromLectureHandlerERR",
		Code:     "Server_handlers",
		HTTPCode: http.StatusBadRequest,
	}
	CreateLectureHandlerErr = AppError{
		Message:  "Failed to createLectureHandlerErr",
		Code:     "Server_handlers",
		HTTPCode: http.StatusBadRequest,
	}
	GetLecturesPPHandlerErr = AppError{
		Message:  "Failed to getLecturesPPHandlerErr",
		Code:     "Server_handlers",
		HTTPCode: http.StatusBadRequest,
	}
	//SERVICES
	CreateLectureServiceErr = AppError{
		Message:  "Failed to CreateLectureServiceErr",
		Code:     "Lecture_Service",
		HTTPCode: http.StatusInternalServerError,
	}
	AddStudentToLectureServiceErr = AppError{
		Message:  "Failed to AddStudentToLectureServiceErr",
		Code:     "Lecture_Service",
		HTTPCode: http.StatusInternalServerError,
	}
	DeleteUserFromLectureServiceErr = AppError{
		Message:  "Failed to DeleteUserFromLectureServiceErr",
		Code:     "Lecture_Service",
		HTTPCode: http.StatusInternalServerError,
	}
	GetLecturesPPServiceErr = AppError{
		Message:  "Failed to GetAllLecturesAndStudentsServiceErr",
		Code:     "Lecture_Service",
		HTTPCode: http.StatusInternalServerError,
	}
	CreateUserServiceErr = AppError{
		Message:  "Failed to CreateUserServiceErr",
		Code:     "User_Service",
		HTTPCode: http.StatusInternalServerError,
	}
)

func (appError *AppError) HttpCode() int {
	return appError.HTTPCode
}

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message:  fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:     appError.Code,
		HTTPCode: appError.HTTPCode,
	}
}

func IsAppError(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
