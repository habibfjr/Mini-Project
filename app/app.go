package app

import (
	"fmt"
	"gomp/auth"
	"gomp/domain"
	"gomp/logger"
	"gomp/service"
	"gomp/users"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"DB_USER",
		"DB_PASSWD",
		"DB_ADDR",
		"DB_PORT",
		"DB_NAME",
	}

	for _, envKey := range envProps {
		if os.Getenv(envKey) == "" {
			logger.Fatal(fmt.Sprintf("environment variable %s not defined. terminating application...", envKey))
		}
	}

	logger.Info("environment variables loaded...")

}

func Start() {

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("error loading .env file")
	}
	logger.Info("load environment variables...")

	sanityCheck()

	dbClient := getClientDB()

	jobsRepositoryDB := domain.NewJobsRepositoryDB(dbClient)
	userRepositoryDB := users.NewUserRepositoryDB(dbClient)
	// authRepositoryDB := domain.NewAuthRepositoryDB(dbClient)

	jobsService := service.NewJobsService(&jobsRepositoryDB)
	userService := users.NewUserService(&userRepositoryDB)
	authService := auth.NewService()
	// authService := service.NewAuthService(authRepositoryDB)

	jh := JobsHandler{jobsService}
	uh := users.NewUserHandler(userService, authService)
	// ah := AuthHandler{authService}

	router := gin.Default()

	router.GET("/jobs", jh.getAll)

	router.GET("/jobs/:id", jh.getJobsByID)

	router.POST("/jobs", jh.createJob)

	router.PUT("/jobs/:id", jh.updateJob)

	router.DELETE("/jobs/:id", jh.deleteJob)

	router.POST("/users", uh.CreateUser)

	router.POST("/login", uh.LoginUser)

	router.Run(":8000")
}
func getClientDB() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// dsn := "host=localhost user=postgres password=postgres dbname=mini_jobs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("success connect to database...")

	return db
}
