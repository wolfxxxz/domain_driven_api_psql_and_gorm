package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"web_service/internal/config"
	"web_service/internal/database"
	"web_service/internal/domain/models"
	"web_service/internal/repositories"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type server struct {
	repoLects repositories.RepoLecture
	repoUsers repositories.UserRepo
	router    Router
	logger    *zap.SugaredLogger
}

func NewServer(repoLects repositories.RepoLecture, repoUsers repositories.UserRepo, logger *zap.SugaredLogger) *server {
	return &server{repoLects: repoLects, repoUsers: repoUsers, router: &router{mux: mux.NewRouter()}, logger: logger}
}

func (srv *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHttp(w, r)
}

func (srv *server) initializeRoutes() {
	srv.logger.Info("server INIT")
	srv.router.Post("/users", srv.contextExpire(srv.createUserHandler()))
	srv.router.Post("/lectures", srv.contextExpire(srv.createLectureHandler()))
	srv.router.Put("/lectures/{lecture_id}/add-student", srv.contextExpire(srv.addUserToLectureHandler()))
	srv.router.Delete("/lectures/{lecture_id}/remove-student", srv.contextExpire(srv.deleteUserFromLectureHandler()))
	srv.router.Get("/lectures", srv.contextExpire(srv.getLecturesPPHandler()))
}

func Run() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	cfg, err := config.NewConfig(logger)
	if err != nil {
		logger.Sugar().Fatal(err)
	}

	ctx := context.Background()
	psglDB := database.NewPostgresDB()
	db, err := psglDB.SetupDatabase(ctx, cfg, logger)
	if err != nil {
		logger.Sugar().Fatal(err)
	}

	if !db.Migrator().HasTable(&models.Lecture{}) {
		err = db.AutoMigrate(&models.Lecture{})
		if err != nil {
			logger.Sugar().Error(err)
			return
		}

		logger.Sugar().Info("Migration success")
	}

	repoLect := repositories.NewRepoLecture(db, logger.Sugar())
	repoUser := repositories.NewUserRepo(db, logger.Sugar())
	srv := NewServer(repoLect, repoUser, logger.Sugar())

	srv.initializeRoutes()
	logger.Sugar().Infof("Listening HTTP service on %s port", cfg.AppPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.AppPort), srv)
	if err != nil {
		logger.Sugar().Fatal(err)
	}
}
