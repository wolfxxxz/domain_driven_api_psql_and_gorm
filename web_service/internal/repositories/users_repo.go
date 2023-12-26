package repositories

import (
	"context"

	"web_service/internal/apperrors"
	"web_service/internal/domain/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) (string, error)
}

type userRepo struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewUserRepo(db *gorm.DB, logger *zap.SugaredLogger) UserRepo {
	return &userRepo{
		db:     db,
		logger: logger,
	}
}

func (repo *userRepo) CreateUser(ctx context.Context, user *models.User) (string, error) {
	if user == nil {
		appErr := apperrors.CreateUserErr.AppendMessage("user is nil")
		repo.logger.Error(appErr)
		return "", appErr
	}

	tx := repo.db.WithContext(ctx)
	if tx.Error != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(tx.Error)
		repo.logger.Error(appErr)
		return "", appErr
	}

	result := tx.Create(user)
	if result.Error != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(result.Error)
		repo.logger.Error(appErr)
		return "", appErr
	}

	if result.RowsAffected == 0 {
		appErr := apperrors.CreateUserErr.AppendMessage("no rows affected")
		repo.logger.Error(appErr)
		return "", appErr
	}

	createdUser := &models.User{}
	if err := tx.First(createdUser, "id = ?", user.ID).Error; err != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return "", appErr
	}

	return createdUser.ID.String(), nil
}
