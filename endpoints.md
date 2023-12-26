# 1.Create user profile
    URL: /users
    method: POST
        Request Body:
          {
              "email": "user_email",
              "first_name": "John",
              "last_name": "Doe",
              "password": "secure_password",
              "role": "student",
          }
        Response:
          201 Created
              Response body:
              {
              "id": 1
              }
          400 Bad Request
              Response body:
              {"error": "Incorrect field"}
          409 Conflict
              Response body:
              {"error": "Invalid request body"}
# 2.Create lecture
    URL: /lecture
    method: POST
        Request Body:
          {
              "title": "IDE",
              "description": "integrated development environment",
              "speaker": "Mat Ryer",
              "date": "2023-12-31T12:00:00",
              "location": "White house",
              "duration": 60, 
          }
        Response:
          201 Created
              Response body:
              {
              "lecture_id": 1
              }
          400 Bad Request
              Response body:
              {"error": "Incorrect field"}
          409 Conflict
              Response body:
              {"error": "Invalid request body"}
# 3.AddStudentToLecture
    URL: /lectures/:lecture_id/add-student
    method: PUT
        Path Parameters:
          user_id
        Request Body:
              {
              "user_id": 1,
              }
        Response:
          200 OK
              Response Body:
              {
              "count_of_registered_students": 1,
              }
          400 Bad Request
              Response body:
              {"error": "Incorrect field"}
          404 Not Found
              Response body:
              {"error": "Lecture not found"}
          500 InternalServerError
              Response body:
              {"error": "Internal server error"}
# 4.RemoveAStudentFromLecture
    URL: /lectures/:lecture_id/drop-student
    method: PUT
        Path Parameters:
          user_id
        Request Body:
              {
              "user_id": 1,
              }
        Response:
          200 OK +
              Response Body:
              {
              "count_of_registered_students": 1
              }
          400 Bad Request
              Response body:
              {"error": "Incorrect field"}
          404 Not Found
              Response body:
              {"error": "Lecture not found"}
          500 InternalServerError
              Response body:
              {"error": "Internal server error"}
# 5.ListLecturesWithStudents
    URL: /lectures?page=1&per_page=10
    method: GET
        Query Parameters:
          page
          per_page
        Response:
          200 OK
              Response Body:
              {
                  "lectures": [
                      {
                          "lecture_id": "1",
                          "title": "IDE",
                          "description": "integrated development environment",
                          "speaker": "Mat Ryer",
                          "date": "2023-12-31T12:00:00",
                          "location": "White house",
                          "duration": 60, 
                          "count_of_registered_students": 1,
                          "students":[
                            {
                               "user_id": "121jsd31",
                               "email": "user_email",
                               "user_name": "user_name",
                               "first_name": "John",
                               "last_name": "Doe",
                            },
                          ]
                      },
                      {
                          "lecture_id": "2",
                          "title": "HTCNY",
                          "description": "how to celebrate new year",
                          "speaker": "Santa Claus",
                          "date": "2023-12-31T12:00:00",
                          "location": "Christmas tree",
                          "duration": time.Duration(the whole night), 
                          "count_of_registered_students": 2,
                          "students":[
                            {
                               "user_id": "121jsd33",
                               "email": "user_email",
                               "user_name": "user_name",
                               "first_name": "John",
                               "last_name": "Doe",
                            },
                            {
                               "user_id": "121jsd34",
                               "email": "user_email",
                               "user_name": "user_name",
                               "first_name": "Bob",
                               "last_name": "Doe",
                            },
                          ]
                      }
                  ],
                  "page": 1,
                  "per_page": 10,
                  "total_users": 2
              }
          400 Bad request
              Response Body:
              {"error": "wrong parameters page, per_page"}
          500 InternalServerError
              Response body:
              {"error": "Internal server error"}