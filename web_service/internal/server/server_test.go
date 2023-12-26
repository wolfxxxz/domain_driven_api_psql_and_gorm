package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"web_service/internal/apperrors"
	"web_service/internal/domain/mappers"
	"web_service/internal/domain/models"
	"web_service/internal/domain/requests"
	"web_service/internal/domain/responses"
	"web_service/internal/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateUser(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()
	logger.Info("logger inited")
	createUserRequest := &requests.CreateUserRequest{
		Email:     "har@name.one",
		FirstName: "Third",
		LastName:  "last name",
		Password:  "BoBEEEEEEER3",
		Role:      "student3",
	}

	requestBody, err := json.Marshal(createUserRequest)
	if err != nil {
		t.Fatal(err)
	}

	user := mappers.MapCreateUserRequestToUser(createUserRequest)
	userUUID, err := uuid.Parse("c616fed8-e6d2-45f5-80e5-d2eacfd8e4bf")
	if err != nil {
		logger.Error(err.Error())
	}

	user.ID = &userUUID

	createUserResp := &responses.CreateUserResponse{
		UserId: userUUID.String(),
	}

	testTable := []struct {
		scenario      string
		inputCreateUR []byte
		user          *models.User
		contentType   string
		response      *responses.CreateUserResponse
		expectedErr   error
		httpCode      int
	}{
		{
			"create_user_decode_err",
			[]byte("invalid json"),
			user,
			"json",
			nil,
			&apperrors.CreateLectureHandlerErr,
			apperrors.CreateLectureHandlerErr.HTTPCode,
		},
		{
			"create_user_Created",
			requestBody,
			user,
			"json",
			createUserResp,
			nil,
			http.StatusCreated,
		},

		{
			"create_user_service_err",
			requestBody,
			user,
			"json",
			nil,
			&apperrors.CreateUserServiceErr,
			apperrors.CreateUserServiceErr.HTTPCode,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.Info("ctrl inited")

	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctx := context.Background()

			usersRepoMock := mock.NewMockUserRepo(ctrl)
			logger.Info("mocks inited")
			srv := &server{repoUsers: usersRepoMock, logger: logger.Sugar()}
			logger.Info("server inited")

			logger.Info("reqBody inited")
			var reqCreateUserArr io.Reader
			if tc.scenario != "create_user_decode_err" {
				reqCreateUserArr = bytes.NewReader(tc.inputCreateUR)
			}

			if tc.scenario == "create_user_decode_err" {
				reqCreateUserArr = strings.NewReader("invalid json")
			}

			req := httptest.NewRequest(http.MethodPost, "/users", reqCreateUserArr)

			req.Header.Set("Content-Type", tc.contentType)
			logger.Info("httptest.NewRequest inited")
			rec := httptest.NewRecorder()

			usersRepoMock.EXPECT().CreateUser(ctx, gomock.Any()).Return(tc.user.ID.String(), tc.expectedErr).AnyTimes()
			logger.Info("mock.EXPECT inited")

			createUser := srv.createUserHandler()
			createUser(rec, req)

			if rec.Code != tc.httpCode {
				t.Errorf("expected HTTP status code %d, got %d", tc.httpCode, rec.Code)
				return
			}

			if rec.Code != http.StatusCreated {
				apperrors.IsAppError(err, tc.expectedErr.(*apperrors.AppError))
				assert.Equal(t, tc.httpCode, rec.Code)
				return
			}

			assert.Equal(t, tc.httpCode, rec.Code)
			marshalledResponse, err := json.Marshal(tc.response)
			if assert.NoError(t, err) {
				assert.Equal(t, string(marshalledResponse), strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		})
	}
}

func TestCreateLecture(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()
	logger.Info("logger inited")
	createLectRequest := &requests.CreateLectureRequest{
		Title:       "newYear",
		Description: "how to celebrate",
		Speaker:     "Santa Claus",
		Date:        "2024-12-25T08:00:00Z",
		Location:    "Christmas tree",
		Duration:    "60",
	}

	requestBody, err := json.Marshal(createLectRequest)
	if err != nil {
		t.Fatal(err)
	}

	lecture, err := mappers.MapCreateLectureReqToLecture(createLectRequest)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	lectureID, err := uuid.Parse("c616fed8-e6d2-45f5-80e5-d2eacfd8e4bf")
	if err != nil {
		logger.Error(err.Error())
	}

	lecture.ID = &lectureID

	createLectResp := &responses.CreateLectureResponse{
		LectureId: lectureID.String(),
	}

	testTable := []struct {
		scenario      string
		inputCreateUR []byte
		user          *models.Lecture
		contentType   string
		response      *responses.CreateLectureResponse
		expectedErr   error
		httpCode      int
	}{
		{
			"create_lecture_decode_err",
			[]byte("invalid json"),
			lecture,
			"json",
			nil,
			&apperrors.CreateLectureHandlerErr,
			apperrors.CreateLectureHandlerErr.HTTPCode,
		},
		{
			"create_user_Created",
			requestBody,
			lecture,
			"json",
			createLectResp,
			nil,
			http.StatusCreated,
		},

		{
			"create_user_service_err",
			requestBody,
			lecture,
			"json",
			nil,
			&apperrors.CreateUserServiceErr,
			apperrors.CreateUserServiceErr.HTTPCode,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.Info("ctrl inited")

	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctx := context.Background()

			lectureRepoMock := mock.NewMockRepoLecture(ctrl)
			logger.Info("mocks inited")
			srv := &server{repoLects: lectureRepoMock, logger: logger.Sugar()}
			logger.Info("server inited")

			logger.Info("reqBody inited")
			var reqCreateLectureArr io.Reader
			if tc.scenario != "create_user_decode_err" {
				reqCreateLectureArr = bytes.NewReader(tc.inputCreateUR)
			}

			if tc.scenario == "create_user_decode_err" {
				reqCreateLectureArr = strings.NewReader("invalid json")
			}

			req := httptest.NewRequest(http.MethodPost, "/lecture", reqCreateLectureArr)

			req.Header.Set("Content-Type", tc.contentType)
			logger.Info("httptest.NewRequest inited")
			rec := httptest.NewRecorder()

			lectureRepoMock.EXPECT().CreateLecture(ctx, gomock.Any()).Return(tc.user.ID.String(), tc.expectedErr).AnyTimes()
			logger.Info("mock.EXPECT inited")

			createLect := srv.createLectureHandler()
			createLect(rec, req)

			if rec.Code != tc.httpCode {
				t.Errorf("expected HTTP status code %d, got %d", tc.httpCode, rec.Code)
				return
			}

			if rec.Code != http.StatusCreated {
				apperrors.IsAppError(err, tc.expectedErr.(*apperrors.AppError))
				assert.Equal(t, tc.httpCode, rec.Code)
				return
			}

			assert.Equal(t, tc.httpCode, rec.Code)
			marshalledResponse, err := json.Marshal(tc.response)
			if assert.NoError(t, err) {
				assert.Equal(t, string(marshalledResponse), strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		})
	}
}

func TestGetLecturesPPHandler(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()
	logger.Info("logger inited")
	getLectsPPRequest := &requests.GetLecturesPPRequest{
		Page:    "hi",
		PerPage: "10",
	}

	requestBody, err := json.Marshal(getLectsPPRequest)
	if err != nil {
		t.Fatal(err)
	}

	getLectsPPRequestPos := &requests.GetLecturesPPRequest{
		Page:    "1",
		PerPage: "10",
	}

	requestBodyPos, err := json.Marshal(getLectsPPRequestPos)
	if err != nil {
		t.Fatal(err)
	}

	lectID, _ := uuid.Parse("318f38ad-76dc-41d9-8ce5-7900559264dd")
	userID, _ := uuid.Parse("9ead1870-0962-4f24-ac0b-c1901af0899b")
	user := &models.User{ID: &userID, Email: "har@name.one4"}

	lecture := &models.Lecture{
		ID:       &lectID,
		Title:    "newYear",
		Students: []*models.User{user},
	}

	lectures := []*models.Lecture{lecture}

	getLectsPPResp, _ := mappers.MapGetAllLecturesAndStudentsToGetLecturesAndStudentsPPRespResponse(lectures)

	testTable := []struct {
		scenario              string
		inputGetLectReq       []byte
		contentType           string
		response              []*responses.GetLecturesAndStudentsPPResponse
		expectedSetOfLectures []*models.Lecture
		expectedErr           error
		httpCode              int
	}{
		{
			"get_lecture_decode_err",
			[]byte("invalid json"),
			"json",
			nil,
			lectures,
			apperrors.GetLecturesPPHandlerErr.AppendMessage("Bind lecture id"),
			apperrors.GetLecturesPPHandlerErr.HTTPCode,
		},
		{
			"get_lecture_service_err",
			requestBody,
			"json",
			nil,
			lectures,
			apperrors.GetLecturesPPServiceErr.AppendMessage("strconv err: hi"),
			apperrors.GetLecturesPPServiceErr.HTTPCode,
		},
		{
			"get_lecture_service_POSITIVE",
			requestBodyPos,
			"json",
			getLectsPPResp,
			lectures,
			nil,
			http.StatusOK,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.Info("ctrl inited")

	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			ctx := context.Background()

			lectureRepoMock := mock.NewMockRepoLecture(ctrl)
			logger.Info("mocks inited")
			srv := &server{repoLects: lectureRepoMock, logger: logger.Sugar()}
			logger.Info("server inited")

			logger.Info("reqBody inited")
			var reqCreateLectureArr io.Reader
			if tc.scenario != "get_lecture_decode_err" {
				reqCreateLectureArr = bytes.NewReader(tc.inputGetLectReq)
			}

			if tc.scenario == "get_lecture_decode_err" {
				reqCreateLectureArr = strings.NewReader("invalid json")
			}

			req := httptest.NewRequest(http.MethodGet, "/lectures", reqCreateLectureArr)

			req.Header.Set("Content-Type", tc.contentType)
			logger.Info("httptest.NewRequest inited")
			rec := httptest.NewRecorder()

			lectureRepoMock.EXPECT().GetLecturesAndStudentsPP(ctx, gomock.Any(), gomock.Any()).Return(tc.expectedSetOfLectures, tc.expectedErr).AnyTimes()
			logger.Info("mock.EXPECT inited")

			getLectsPP := srv.getLecturesPPHandler()
			getLectsPP(rec, req)

			if rec.Code != tc.httpCode {
				t.Errorf("expected HTTP status code %d, got %d", tc.httpCode, rec.Code)
				return
			}

			if rec.Code != http.StatusOK {
				apperrors.IsAppError(err, tc.expectedErr.(*apperrors.AppError))
				assert.Equal(t, tc.httpCode, rec.Code)
				return
			}

			assert.Equal(t, tc.httpCode, rec.Code)
			marshalledResponse, err := json.Marshal(tc.response)
			fmt.Println(tc.response)
			logger.Info(string(marshalledResponse))
			logger.Info(strings.TrimSuffix(rec.Body.String(), "\n"))
			if assert.NoError(t, err) {
				assert.Equal(t, string(marshalledResponse), strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		})
	}
}

func TestAddUserToLectureHandler(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()
	logger.Info("logger inited")
	addUserToLectReq := requests.AddStudentToLectureReq{UserId: "9ead1870-0962-4f24-ac0b-c1901af0899b"}

	requestBody, err := json.Marshal(addUserToLectReq)
	if err != nil {
		t.Fatal(err)
	}

	lectIdFail := "22"
	lectureID := "c616fed8-e6d2-45f5-80e5-d2eacfd8e4bf"

	createLectResp := &responses.CreateLectureResponse{
		LectureId: lectureID,
	}

	testTable := []struct {
		scenario       string
		inputUserID    []byte
		inputLectureID string
		contentType    string
		response       *responses.CreateLectureResponse
		message        string
		expectedErr    error
		httpCode       int
	}{
		{
			"add_user_to_lecture_decode_err",
			[]byte("invalid json"),
			"lecture_id",
			"json",
			nil,
			"Failed to addUserToLectureHandlerErr : [Bind user_id id]",
			apperrors.AddStudentToLectureHandlerErr.AppendMessage("Bind user_id id"),
			apperrors.AddStudentToLectureHandlerErr.HTTPCode,
		},
		{
			"add_user_to_lecture_MUX_VARS_err",
			requestBody,
			"",
			"json",
			nil,
			"Failed to addUserToLectureHandlerErr : [Vars lecture_id]",
			apperrors.AddStudentToLectureHandlerErr.AppendMessage("Vars lecture_id"),
			apperrors.AddStudentToLectureHandlerErr.HTTPCode,
		},
		{
			"add_user_to_lecture_SERVICE_err",
			requestBody,
			lectIdFail,
			"json",
			nil,
			"Failed to AddStudentToLectureServiceErr : [invalid UUID length: 2]",
			&apperrors.AddStudentToLectureServiceErr,
			apperrors.AddStudentToLectureServiceErr.HTTPCode,
		},
		{
			"add_user_to_lecture_POSITIVE",
			requestBody,
			lectureID,
			"json",
			createLectResp,
			"",
			nil,
			http.StatusOK,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.Info("ctrl inited")

	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			lectureRepoMock := mock.NewMockRepoLecture(ctrl)
			logger.Info("mocks inited")
			srv := &server{repoLects: lectureRepoMock, logger: logger.Sugar()}
			logger.Info("server inited")

			logger.Info("reqBody inited")
			var reqCreateLectureArr io.Reader
			if tc.scenario != "add_user_to_lecture_decode_err" {
				reqCreateLectureArr = bytes.NewReader(tc.inputUserID)
			}

			if tc.scenario == "add_user_to_lecture_decode_err" {
				reqCreateLectureArr = strings.NewReader("invalid json")
			}

			req := httptest.NewRequest(http.MethodPut, "/lectures/{lecture_id}/add-student", reqCreateLectureArr)
			if tc.scenario != "add_user_to_lecture_MUX_VARS_err" {
				req = mux.SetURLVars(req, map[string]string{"lecture_id": tc.inputLectureID})
			}

			req.Header.Set("Content-Type", tc.contentType)
			logger.Info("httptest.NewRequest inited")
			rec := httptest.NewRecorder()

			lectureRepoMock.EXPECT().AddUserToLecture(req.Context(), gomock.Any(), gomock.Any()).Return(tc.expectedErr).AnyTimes() //CreateLecture(ctx, gomock.Any()).Return(tc.user.ID.String(), tc.expectedErr).AnyTimes()
			logger.Info("mock.EXPECT inited")

			addUserToLect := srv.addUserToLectureHandler()
			logger.Info("addUserToLect")
			addUserToLect(rec, req)
			logger.Info("addUserToLect(rec, req)")

			if rec.Code != http.StatusOK {
				apperrors.IsAppError(err, tc.expectedErr.(*apperrors.AppError))
				assert.Equal(t, tc.httpCode, rec.Code)
				return
			}

			assert.Equal(t, tc.httpCode, rec.Code)
			marshalledResponse, err := json.Marshal(tc.response)
			if assert.NoError(t, err) {
				assert.Equal(t, string(marshalledResponse), strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		})
	}
}

func TestDeleteUserFromLectureHandler(t *testing.T) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()
	logger.Info("logger inited")
	dropUserFromLectReq := requests.DeleteStudentFromLectureRequest{UserId: "9ead1870-0962-4f24-ac0b-c1901af0899b"}

	requestBody, err := json.Marshal(dropUserFromLectReq)
	if err != nil {
		t.Fatal(err)
	}

	lectIdFail := "22"
	lectureID := "c616fed8-e6d2-45f5-80e5-d2eacfd8e4bf"

	type ReqSuc struct {
		Result string `json:"result"`
	}

	respSuccess := &ReqSuc{Result: "Success"}

	testTable := []struct {
		scenario       string
		inputUserID    []byte
		inputLectureID string
		contentType    string
		response       *ReqSuc
		message        string
		expectedErr    error
		httpCode       int
	}{
		{
			"drop_user_from_lecture_decode_err",
			[]byte("invalid json"),
			"lecture_id",
			"json",
			nil,
			"Failed to addUserToLectureHandlerErr : [Bind user_id]",
			apperrors.DeleteUserFromLectureHandlerERR.AppendMessage("Bind user_id"),
			apperrors.DeleteUserFromLectureHandlerERR.HTTPCode,
		},
		{
			"add_user_to_lecture_MUX_VARS_err",
			requestBody,
			"",
			"json",
			nil,
			"Failed to addUserToLectureHandlerErr : [Vars lecture_id]",
			apperrors.AddStudentToLectureHandlerErr.AppendMessage("Vars lecture_id"),
			apperrors.AddStudentToLectureHandlerErr.HTTPCode,
		},
		{
			"add_user_to_lecture_SERVICE_err",
			requestBody,
			lectIdFail,
			"json",
			nil,
			"Failed to AddStudentToLectureServiceErr : [invalid UUID length: 2]",
			&apperrors.AddStudentToLectureServiceErr,
			apperrors.AddStudentToLectureServiceErr.HTTPCode,
		},
		{
			"add_user_to_lecture_POSITIVE",
			requestBody,
			lectureID,
			"json",
			respSuccess,
			"",
			nil,
			http.StatusOK,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger.Info("ctrl inited")

	for _, tc := range testTable {
		t.Run(tc.scenario, func(t *testing.T) {
			lectureRepoMock := mock.NewMockRepoLecture(ctrl)
			logger.Info("mocks inited")
			srv := &server{repoLects: lectureRepoMock, logger: logger.Sugar()}
			logger.Info("server inited")

			logger.Info("reqBody inited")
			var dropUserFromLectReq io.Reader
			if tc.scenario != "drop_user_from_lecture_decode_err" {
				dropUserFromLectReq = bytes.NewReader(tc.inputUserID)
			}

			if tc.scenario == "drop_user_from_lecture_decode_err" {
				dropUserFromLectReq = strings.NewReader("invalid json")
			}

			req := httptest.NewRequest(http.MethodDelete, "/lectures/{lecture_id}/remove-student", dropUserFromLectReq)
			if tc.scenario != "add_user_to_lecture_MUX_VARS_err" {
				req = mux.SetURLVars(req, map[string]string{"lecture_id": tc.inputLectureID})
			}

			req.Header.Set("Content-Type", tc.contentType)
			logger.Info("httptest.NewRequest inited")
			rec := httptest.NewRecorder()

			lectureRepoMock.EXPECT().DropUserFromLecture(req.Context(), gomock.Any(), gomock.Any()).Return(tc.expectedErr).AnyTimes() //CreateLecture(ctx, gomock.Any()).Return(tc.user.ID.String(), tc.expectedErr).AnyTimes()
			logger.Info("mock.EXPECT inited")

			removeUserFromLect := srv.deleteUserFromLectureHandler()
			logger.Info("addUserToLect")
			removeUserFromLect(rec, req)
			logger.Info("addUserToLect(rec, req)")

			if rec.Code != http.StatusOK {
				apperrors.IsAppError(err, tc.expectedErr.(*apperrors.AppError))
				assert.Equal(t, tc.httpCode, rec.Code)
				return
			}

			assert.Equal(t, tc.httpCode, rec.Code)
			marshalledResponse, err := json.Marshal(tc.response)
			if assert.NoError(t, err) {
				assert.Equal(t, string(marshalledResponse), strings.TrimSuffix(rec.Body.String(), "\n"))
			}
		})
	}
}
