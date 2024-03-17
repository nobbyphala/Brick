package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nobbyphala/Brick/adapter/rest_api"
	"github.com/nobbyphala/Brick/config"
	"github.com/nobbyphala/Brick/external/database"
	"github.com/nobbyphala/Brick/external/http_request"
	"github.com/nobbyphala/Brick/usecase"
	"github.com/nobbyphala/Brick/usecase/api"
	"github.com/nobbyphala/Brick/usecase/repository"
	"log"
)

func main() {
	// init driver or framework
	db, err := database.NewPostgresDB(database.ConnectionOption{
		Host:     config.DB_HOST,
		Port:     config.DB_PORT,
		User:     config.DB_USER,
		Password: config.DB_PASSWORD,
		Database: config.DB_DATABASE,
	})
	if err != nil {
		log.Panicln(err)
	}

	postgresSql := database.NewPostgresSqlClient(database.PostgresSQLOpts{
		DB: db,
	})

	// repository
	disbursementRepository := repository.NewDisbursement(repository.DisbursementDeps{
		DB: postgresSql,
	})
	utilsRepository := repository.NewRepositoryUtils(repository.UtilsOpts{
		DB: db,
	})

	// api
	bankApi := api.NewBankApiClient(api.BankApiClientOpts{
		BaseUrl:     config.BankBaseURL,
		HttpRequest: http_request.NewHttpRequest(),
	})

	// usecase
	disbursementUsecase := usecase.NewDisbursement(usecase.DisbursementDeps{
		BankApi:                bankApi,
		UtilsRepository:        utilsRepository,
		DisbursementRepository: disbursementRepository,
	})

	// controller
	disbursementController := rest_api.NewDisbursementController(rest_api.DisbursementControllerDeps{
		DisbursementUsecase: disbursementUsecase,
	})

	// init http server
	r := gin.Default()
	rest_api.RegisterRouter(r, rest_api.RouteController{
		DisbursementController: disbursementController,
	})

	r.Run("127.0.0.1:8080")
}
