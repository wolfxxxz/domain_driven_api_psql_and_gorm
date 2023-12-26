package mappers

import (
	"strconv"
	"time"
	"web_service/internal/domain/models"
	"web_service/internal/domain/requests"
	"web_service/internal/domain/responses"

	"github.com/google/uuid"
)

func MapCreateUserRequestToUser(createUserRequest *requests.CreateUserRequest) *models.User {
	bid := uuid.New()
	return &models.User{
		ID:        &bid,
		Email:     createUserRequest.Email,
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Role:      createUserRequest.Role,
		Password:  createUserRequest.Password,
	}
}

func MapCreateLectureReqToLecture(createLectureReq *requests.CreateLectureRequest) (*models.Lecture, error) {
	durationNum, err := strconv.Atoi(createLectureReq.Duration)
	if err != nil {
		return nil, err
	}

	bid := uuid.New()
	dateTime, err := time.Parse(time.RFC3339, createLectureReq.Date)
	if err != nil {
		return nil, err
	}

	return &models.Lecture{
		ID:          &bid,
		Title:       createLectureReq.Title,
		Description: createLectureReq.Description,
		Speaker:     createLectureReq.Speaker,
		Location:    createLectureReq.Location,
		Duration:    durationNum,
		Date:        dateTime,
	}, nil
}

func MapGetAllLecturesAndStudentsToGetLecturesAndStudentsPPRespResponse(lectures []*models.Lecture) ([]*responses.GetLecturesAndStudentsPPResponse, error) {
	lecturesResp := []*responses.GetLecturesAndStudentsPPResponse{}
	for _, lecture := range lectures {
		studentsResp := []*responses.StudentResp{}

		for _, student := range lecture.Students {
			studentResp := &responses.StudentResp{
				ID:    student.ID.String(),
				Email: student.Email,
			}

			studentsResp = append(studentsResp, studentResp)
		}

		lectureGetAllLectsResp := &responses.GetLecturesAndStudentsPPResponse{
			ID:       lecture.ID.String(),
			Title:    lecture.Title,
			Students: studentsResp,
		}

		lecturesResp = append(lecturesResp, lectureGetAllLectsResp)
	}

	return lecturesResp, nil
}
