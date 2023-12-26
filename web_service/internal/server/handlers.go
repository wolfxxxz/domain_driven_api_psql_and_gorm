package server

import (
	"encoding/json"
	"net/http"

	"web_service/internal/apperrors"
	"web_service/internal/domain/requests"
	"web_service/internal/domain/responses"
	"web_service/internal/services"

	"github.com/gorilla/mux"
)

func (srv *server) createUserHandler() http.HandlerFunc {
	srv.logger.Info("createUserHandler has been initiated.")
	return func(w http.ResponseWriter, r *http.Request) {
		createUserRequest := &requests.CreateUserRequest{}
		err := srv.decode(r, createUserRequest)
		if err != nil {
			appErr := apperrors.CreateUserHandlerErr.AppendMessage("DECODE ERR: ", err)
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		srv.logger.Infof("createUserHandler has been invoked. Request: %+v", createUserRequest)

		userService := services.NewUserService(srv.repoUsers, srv.logger)
		srv.logger.Info("services.NewUserService")

		createUserResponse, err := userService.CreateUser(r.Context(), createUserRequest)
		if err != nil {
			srv.logger.Error(err)
			appErrors := err.(*apperrors.AppError)
			srv.respond(w, appErrors.Message, http.StatusInternalServerError)
			return
		}

		srv.logger.Infof("createUserHandler has been processed. Response: %+v", createUserResponse)
		srv.respond(w, createUserResponse, http.StatusCreated)
	}
}

func (srv *server) getLecturesPPHandler() http.HandlerFunc {
	srv.logger.Info("getLecturesHandler has been initiated.")
	return func(w http.ResponseWriter, r *http.Request) {
		getLectsPPRequest := &requests.GetLecturesPPRequest{}
		err := srv.decode(r, getLectsPPRequest)
		if err != nil {
			appErr := apperrors.GetLecturesPPHandlerErr.AppendMessage("Bind lecture id")
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		srv.logger.Infof("getLecturesHandler has been invoked. Page %v, PerPage %v", getLectsPPRequest.Page, getLectsPPRequest.PerPage)
		lectureService := services.NewLectureService(srv.repoLects, srv.logger)
		getLecturesAndStudentsPPResp, err := lectureService.GetLecturesAndStudentsPP(r.Context(), getLectsPPRequest.Page, getLectsPPRequest.PerPage)
		if err != nil {
			appErr := err.(*apperrors.AppError)
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusInternalServerError)
			return
		}

		srv.logger.Infof("getLecturesHandler has been processed. Response: %+v", getLecturesAndStudentsPPResp)
		srv.respond(w, getLecturesAndStudentsPPResp, http.StatusOK)
	}
}

func (srv *server) addUserToLectureHandler() http.HandlerFunc {
	srv.logger.Info("addUserToLectureHandler has been initiated.")
	return func(w http.ResponseWriter, r *http.Request) {
		addStudentToLectureRequest := &requests.AddStudentToLectureReq{}
		err := srv.decode(r, addStudentToLectureRequest)
		if err != nil {
			appErr := apperrors.AddStudentToLectureHandlerErr.AppendMessage("Bind user_id")
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		lectureId, ok := mux.Vars(r)["lecture_id"]
		if !ok {
			appErr := apperrors.AddStudentToLectureHandlerErr.AppendMessage("Vars lecture_id")
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		srv.logger.Infof("addUserToLectureHandler has been invoked. Response: %+v, and lecture_id:", addStudentToLectureRequest, lectureId)

		lectureService := services.NewLectureService(srv.repoLects, srv.logger)
		addUserToLectureResp, err := lectureService.AddUserToLecture(r.Context(), lectureId, addStudentToLectureRequest.UserId)
		if err != nil {
			srv.logger.Error(err)
			appErr := err.(*apperrors.AppError)
			srv.respond(w, appErr.Message, http.StatusInternalServerError)
			return
		}

		srv.logger.Infof("addUserToLectureHandler has been processed. Response: %+v", addUserToLectureResp)
		srv.respond(w, addUserToLectureResp, http.StatusOK)
	}
}

func (srv *server) deleteUserFromLectureHandler() http.HandlerFunc {
	srv.logger.Info("deleteUserHandler has been initiated.")

	return func(w http.ResponseWriter, r *http.Request) {
		deleteStudentFromLectureRequest := &requests.DeleteStudentFromLectureRequest{}
		err := srv.decode(r, deleteStudentFromLectureRequest)
		if err != nil {
			appErr := apperrors.DeleteUserFromLectureHandlerERR.AppendMessage(err)
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		lectureId, ok := mux.Vars(r)["lecture_id"]
		if !ok {
			appErr := apperrors.DeleteUserFromLectureHandlerERR.AppendMessage("Bind lecture id")
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		srv.logger.Infof("deleteUserFromLectureHandler has been invoked. Response: %+v, and lecture_id:", deleteStudentFromLectureRequest, lectureId)

		lectureService := services.NewLectureService(srv.repoLects, srv.logger)
		err = lectureService.DeleteUserFromLecture(r.Context(), lectureId, deleteStudentFromLectureRequest.UserId)
		if err != nil {
			srv.logger.Error(err)
			appErr := err.(*apperrors.AppError)
			srv.respond(w, appErr.Message, http.StatusInternalServerError)
			return
		}

		DeleteStudentFromLectureResp := &responses.DeleteStudentFromLectureResponse{Result: "Success"}
		srv.logger.Infof("deleteStudentFromLectureHandler has been processed. Response: %+v", DeleteStudentFromLectureResp)
		srv.respond(w, DeleteStudentFromLectureResp, http.StatusOK)
	}
}

func (srv *server) createLectureHandler() http.HandlerFunc {
	srv.logger.Info("createLectureHandler has been initiated.")

	return func(w http.ResponseWriter, r *http.Request) {
		createLectureRequest := &requests.CreateLectureRequest{}
		err := srv.decode(r, createLectureRequest)
		if err != nil {
			appErr := apperrors.CreateLectureHandlerErr.AppendMessage(err)
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusBadRequest)
			return
		}

		srv.logger.Infof("createLectureHandler has been invoked. Request: %+v", createLectureRequest)

		lectureService := services.NewLectureService(srv.repoLects, srv.logger)
		createLectResp, err := lectureService.CreateLecture(r.Context(), createLectureRequest)
		if err != nil {
			appErr := err.(*apperrors.AppError)
			srv.logger.Error(appErr)
			srv.respond(w, appErr.Message, http.StatusInternalServerError)
			return
		}

		srv.logger.Infof("createLectureHandler has been processed. Response: %+v", createLectResp)
		srv.respond(w, createLectResp, http.StatusCreated)
	}
}

func (srv *server) decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (srv *server) respond(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		srv.logger.Error(err)
	}
}
