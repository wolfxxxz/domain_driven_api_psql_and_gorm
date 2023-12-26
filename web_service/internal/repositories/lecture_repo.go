package repositories

import (
	"context"
	"web_service/internal/apperrors"
	"web_service/internal/domain/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type RepoLecture interface {
	CreateLecture(ctx context.Context, lecture *models.Lecture) (string, error)
	AddUserToLecture(ctx context.Context, lecture *models.Lecture, user *models.User) error
	DropUserFromLecture(ctx context.Context, lecture *models.Lecture, user *models.User) error
	GetLecturesAndStudentsPP(ctx context.Context, page int, perPage int) ([]*models.Lecture, error)
}

type repoLecture struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func NewRepoLecture(db *gorm.DB, log *zap.SugaredLogger) RepoLecture {
	return &repoLecture{db: db, logger: log}
}

func (repo *repoLecture) CreateLecture(ctx context.Context, lecture *models.Lecture) (string, error) {
	if lecture == nil {
		appErr := apperrors.CreateLectureErr.AppendMessage("lecture is nil")
		repo.logger.Error(appErr)
		return "", appErr
	}

	tx := repo.db.WithContext(ctx)
	if tx.Error != nil {
		appErr := apperrors.CreateLectureErr.AppendMessage(tx.Error)
		repo.logger.Error(appErr)
		return "", appErr
	}

	result := tx.Create(lecture)
	if result.Error != nil {
		appErr := apperrors.CreateLectureErr.AppendMessage(result.Error)
		repo.logger.Error(appErr)
		return "", appErr
	}

	if result.RowsAffected == 0 {
		appErr := apperrors.CreateLectureErr.AppendMessage("no rows affected")
		repo.logger.Error(appErr)
		return "", appErr
	}

	createdLecture := &models.Lecture{}
	if err := tx.First(createdLecture, "id = ?", lecture.ID).Error; err != nil {
		appErr := apperrors.CreateLectureErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return "", appErr
	}

	return createdLecture.ID.String(), nil
}

func (repo *repoLecture) AddUserToLecture(ctx context.Context, lecture *models.Lecture, user *models.User) error {
	if err := repo.db.First(user, "id = ?", user.ID).Error; err != nil {
		appErr := apperrors.AddStudentToLectureRepoErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return appErr
	}

	if err := repo.db.First(lecture, "id = ?", lecture.ID).Error; err != nil {
		appErr := apperrors.AddStudentToLectureRepoErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return appErr
	}

	if err := repo.db.Model(&lecture).Association("Students").Append([]*models.User{user}); err != nil {
		appErr := apperrors.AddStudentToLectureRepoErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return appErr
	}

	if err := repo.db.Save(lecture).Error; err != nil {
		appErr := apperrors.AddStudentToLectureRepoErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return appErr
	}

	return nil
}

func (repo *repoLecture) DropUserFromLecture(ctx context.Context, lecture *models.Lecture, user *models.User) error {
	association := repo.db.Model(&models.Lecture{}).Association("Students")
	if association.Error != nil {
		appErr := apperrors.DropUserFromLectureErr.AppendMessage(association.Error)
		repo.logger.Error(appErr)
		return appErr
	}

	if err := association.Delete(user); err != nil {
		appErr := apperrors.DropUserFromLectureErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return appErr
	}

	return nil
}

func (repo *repoLecture) GetLecturesAndStudentsPP(ctx context.Context, page int, perPage int) ([]*models.Lecture, error) {
	var lectures []*models.Lecture
	offset := (page - 1) * perPage
	if err := repo.db.Preload("Students").Order("created_at DESC").Offset(offset).Limit(perPage).Find(&lectures).Error; err != nil {
		appErr := apperrors.GetLecturesStudentsPPErr.AppendMessage(err)
		repo.logger.Error(appErr)
		return nil, appErr
	}

	return lectures, nil
}
