package services

import (
	"context"
	"strconv"
	"web_service/internal/apperrors"
	"web_service/internal/domain/mappers"
	"web_service/internal/domain/models"
	"web_service/internal/domain/requests"
	"web_service/internal/domain/responses"
	"web_service/internal/repositories"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type LectureService struct {
	lectureRepo repositories.RepoLecture
	logger      *zap.SugaredLogger
}

func NewLectureService(lectureRepo repositories.RepoLecture, logger *zap.SugaredLogger) *LectureService {
	return &LectureService{
		lectureRepo: lectureRepo,
		logger:      logger,
	}
}

func (service *LectureService) CreateLecture(ctx context.Context, createLectureRequest *requests.CreateLectureRequest) (*responses.CreateLectureResponse, error) {
	lecture, err := mappers.MapCreateLectureReqToLecture(createLectureRequest)
	if err != nil {
		appErr := apperrors.CreateLectureServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return nil, appErr
	}

	insertedLectureId, err := service.lectureRepo.CreateLecture(ctx, lecture)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return &responses.CreateLectureResponse{LectureId: insertedLectureId}, nil
}

func (service *LectureService) AddUserToLecture(ctx context.Context, lectureId string, userId string) (*responses.AddUserToLecture, error) {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		appErr := apperrors.AddStudentToLectureServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return nil, appErr
	}

	user := &models.User{ID: &userUUID}

	lectureUUID, err := uuid.Parse(lectureId)
	if err != nil {
		appErr := apperrors.AddStudentToLectureServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return nil, appErr
	}

	lecture := &models.Lecture{ID: &lectureUUID}

	err = service.lectureRepo.AddUserToLecture(ctx, lecture, user)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return &responses.AddUserToLecture{LectureId: lecture.ID.String()}, nil
}

func (service *LectureService) DeleteUserFromLecture(ctx context.Context, lectureId string, userId string) error {
	userUUID, err := uuid.Parse(userId)
	if err != nil {
		appErr := apperrors.DeleteUserFromLectureServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return appErr
	}

	user := &models.User{ID: &userUUID}

	lectureUUID, err := uuid.Parse(lectureId)
	if err != nil {
		appErr := apperrors.DeleteUserFromLectureServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return appErr
	}

	lecture := &models.Lecture{ID: &lectureUUID}

	err = service.lectureRepo.DropUserFromLecture(ctx, lecture, user)
	if err != nil {
		service.logger.Error(err)
		return err
	}

	return nil
}

func (service *LectureService) GetLecturesAndStudentsPP(ctx context.Context, page string, perPage string) ([]*responses.GetLecturesAndStudentsPPResponse, error) {
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		appErr := apperrors.GetLecturesPPServiceErr.AppendMessage("strconv err:", page)
		service.logger.Error(appErr)
		return nil, appErr
	}

	perPageNum, err := strconv.Atoi(perPage)
	if err != nil {
		appErr := apperrors.GetLecturesPPServiceErr.AppendMessage("strconv err:", perPage)
		service.logger.Error(appErr)
		return nil, appErr
	}
	lecturesAndStudents, err := service.lectureRepo.GetLecturesAndStudentsPP(ctx, pageNum, perPageNum)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	getLecturesAndStudentsPPResp, err := mappers.MapGetAllLecturesAndStudentsToGetLecturesAndStudentsPPRespResponse(lecturesAndStudents)
	if err != nil {
		appErr := apperrors.GetLecturesPPServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return nil, err
	}

	return getLecturesAndStudentsPPResp, nil
}
