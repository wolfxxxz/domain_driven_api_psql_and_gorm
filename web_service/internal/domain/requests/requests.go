package requests

type CreateUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}

type CreateLectureRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Speaker     string `json:"speaker"`
	Date        string `json:"date"`
	Location    string `json:"location"`
	Duration    string `json:"duration"`
}

type AddStudentToLectureReq struct {
	UserId string `json:"user_id"`
}

type DeleteStudentFromLectureRequest struct {
	UserId string `json:"user_id"`
}

type GetLecturesPPRequest struct {
	Page    string `json:"page"`
	PerPage string `json:"per_page"`
}
