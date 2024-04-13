package router

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/mfarrasml/template-authorization-app/config"
	"github.com/mfarrasml/template-authorization-app/handler"
	"github.com/mfarrasml/template-authorization-app/repository"
	"github.com/mfarrasml/template-authorization-app/usecase"
	"github.com/mfarrasml/template-authorization-app/util"
)

func ServeRouter(db *sql.DB, config config.Config) error {
	pwdUtil := util.NewBcryptHasherUtil(config.BcryptCost())
	tokenUtil := util.NewJwtTokenUtil(util.JwtTokenOpts{
		Secret:     config.JwtSecret(),
		Issuer:     config.JwtIssuer(),
		ExpMinutes: config.JwtExpiry(),
	})

	userRepo := repository.NewUserRepoPostgres(db)
	userUc := usecase.NewUserUcImpl(userRepo, pwdUtil, tokenUtil)
	userHandler := handler.NewUserHandler(userUc)

	opt := HandlerOpt{
		userHandler: userHandler,
		tokenUtil:   tokenUtil,
	}

	router := NewRouter(opt)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.ApiHost(), config.ApiPort()),
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return errors.New("error serving API")
	}

	return nil
}
