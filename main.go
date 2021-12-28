package main

import (
	"context"
	"fmt"
	"halill/ent"
	"halill/ent/migrate"
	"halill/handler"
	"halill/repository"
	"halill/security"
	"halill/service"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		log.Println("RUN on DEBUG mode")
	}
}

func InitializeUser(e *echo.Group, db *ent.Client, jwtSecret string) (*handler.UserHandler, error) {
	userRepository := repository.NewUserRepository(db)
	jwtProvider := security.NewJWTProvider(jwtSecret)
	userService := service.NewUserSerice(userRepository, jwtProvider)
	userHandler := handler.NewUserHandler(e, userService)
	return userHandler, nil
}

func InitializeTodo(e *echo.Group, db *ent.Client, jwtSecret string) (*handler.TodoHandler, error) {
	todoRepository := repository.NewTodoRepository(db)
	userRepository := repository.NewUserRepository(db)
	todoService := service.NewTodoService(todoRepository, userRepository)
	todoHandler := handler.NewTodoHandler(e, todoService, jwtSecret)
	return todoHandler, nil
}

func main() {
	dbDriver := viper.GetString("database.driver")
	dbHost := viper.GetString("database.host")
	// dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	secret := viper.GetString("jwt.secret")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	client, err := ent.Open(dbDriver, connection)
	if err != nil {
		log.Fatal(errors.WithStack(err))
	}
	defer client.Close()

	if err = client.Schema.Create(
		context.TODO(),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed createing schema resources: %v", err)
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1"},
		AllowMethods: []string{"*"},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	user := e.Group("")
	_, err = InitializeUser(user, client, secret)
	if err != nil {
		e.Logger.Fatal(err)
	}

	todo := e.Group("/todo")
	_, err = InitializeTodo(todo, client, secret)
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":5000"))
}
