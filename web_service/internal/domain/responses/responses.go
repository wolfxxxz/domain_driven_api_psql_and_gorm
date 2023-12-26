package responses

type CreateUserResponse struct {
	UserId string `json:"user_id"`
}

type CreateLectureResponse struct {
	LectureId string `json:"lecture_id"`
}

type AddUserToLecture struct {
	LectureId string `json:"lecture_id"`
}

type DeleteStudentFromLectureResponse struct {
	Result string `json:"result"`
}

type GetLecturesAndStudentsPPResponse struct {
	ID       string         `json:"lecture_id"`
	Title    string         `json:"title"`
	Students []*StudentResp `json:"students"`
}

type StudentResp struct {
	ID    string `json:"student_id"`
	Email string `json:"user_email"`
}
