package services

import (
	"context"
	"web_service/internal/apperrors"
	"web_service/internal/domain/mappers"
	"web_service/internal/domain/requests"
	"web_service/internal/domain/responses"
	"web_service/internal/repositories"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.UserRepo
	logger   *zap.SugaredLogger
}

func NewUserService(urRepo repositories.UserRepo, logger *zap.SugaredLogger) *UserService {
	return &UserService{
		userRepo: urRepo,
		logger:   logger,
	}
}

func (service *UserService) CreateUser(ctx context.Context, createUserRequest *requests.CreateUserRequest) (*responses.CreateUserResponse, error) {
	user := mappers.MapCreateUserRequestToUser(createUserRequest)
	userHashPassword, err := hashPassword(createUserRequest.Password)
	if err != nil {
		appErr := apperrors.CreateUserServiceErr.AppendMessage(err)
		service.logger.Error(appErr)
		return nil, err
	}

	user.Password = userHashPassword
	insertedUserID, err := service.userRepo.CreateUser(ctx, user)
	if err != nil {
		service.logger.Error(err)
		return nil, err
	}

	return &responses.CreateUserResponse{UserId: insertedUserID}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
