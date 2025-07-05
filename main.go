package main

import (
	"fmt"

	"multifinancetest/apps/middlewares/security"
	"multifinancetest/apps/repositories/psql"
	"multifinancetest/apps/repositories/s3storage"
	routerRest "multifinancetest/apps/router/rest"
	authsvc "multifinancetest/apps/service/auth"
	errorhandler "multifinancetest/helpers/error_handler"

	"github.com/go-playground/validator/v10"
	"github.com/vizucode/gokit/adapter/dbc"
	"github.com/vizucode/gokit/config"
	"github.com/vizucode/gokit/factory/server"
	"github.com/vizucode/gokit/factory/server/rest"
	"github.com/vizucode/gokit/utils/constant"
	"github.com/vizucode/gokit/utils/env"
)

func main() {

	/*
		Library
	*/
	serviceName := env.GetString("SERVICE_NAME")
	config.Load(serviceName, ".")
	validator10 := validator.New()

	connectionPath := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.GetString("DB_USER"),
		env.GetString("DB_PASSWORD"),
		env.GetString("DB_HOST"),
		env.GetInteger("DB_PORT"),
		env.GetString("DB_NAME"),
	)

	dbConnection := dbc.NewGormConnection(
		dbc.SetGormURIConnection(connectionPath),
		dbc.SetGormDriver(constant.MySQL),
		dbc.SetGormMaxIdleConnection(2),
		dbc.SetGormMaxPoolConnection(50),
		dbc.SetGormMinPoolConnection(10),
		dbc.SetGormSkipTransaction(true),
		dbc.SetGormServiceName(serviceName),
	)

	/*
		Repositories
	*/
	postgreDB := psql.NewPsql(dbConnection.DB)

	/*
		Service Mapping
	*/
	storageBucket := s3storage.NewAwsS3Implement(
		env.GetString("AWS_REGION"),
		env.GetString("AWS_SECRET_KEY"),
		env.GetString("AWS_ACCESS_KEY"),
		env.GetString("AWS_BUCKET_NAME"),
		env.GetString("AWS_ENDPOINT"),
	)

	authSvc := authsvc.NewAuth(
		postgreDB,
		storageBucket,
		validator10,
	)

	restRouter := routerRest.NewRest(
		security.NewSecurity(
			postgreDB,
		),
		authSvc,
	)

	app := server.NewService(
		server.SetServiceName(serviceName),
		server.SetRestHandler(restRouter),
		server.SetRestHandlerOptions(
			rest.SetHTTPHost(env.GetString("APP_HOST")),
			rest.SetHTTPPort(env.GetInteger("APP_PORT")),
			rest.SetErrorHandler(errorhandler.FiberErrHandler),
		),
	)

	appServer := server.New(app)
	appServer.Run()
}
