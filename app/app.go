package app

import (
	"fmt"
	"gomp/auth"
	"gomp/domain"
	"gomp/logger"
	"gomp/service"
	"gomp/users"
	"net/http"
	"os"
	"strings"

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

	// init repo
	jobsRepositoryDB := domain.NewJobsRepositoryDB(dbClient)
	userRepositoryDB := users.NewUserRepositoryDB(dbClient)

	// init service
	jobsService := service.NewJobsService(&jobsRepositoryDB)
	userService := users.NewUserService(&userRepositoryDB)
	authService := auth.NewService()

	// init handlers
	jh := JobsHandler{jobsService}
	uh := users.NewUserHandler(userService, authService)

	router := gin.Default()

	// api routes
	router.GET("/jobs", authMiddleware(authService, userService), jh.getAll)

	router.GET("/jobs/:id", authMiddleware(authService, userService), jh.getJobsByID)

	router.POST("/jobs", authMiddleware(authService, userService), jh.createJob)

	router.PUT("/jobs/:id", authMiddleware(authService, userService), jh.updateJob)

	router.DELETE("/jobs/:id", authMiddleware(authService, userService), jh.deleteJob)

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

// token authorization
func authMiddleware(auth auth.Service, user users.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenString := ""
		tokenArray := strings.Split(authHeader, " ")
		if len(tokenArray) == 2 {
			tokenString = tokenArray[1]
		}
		result, userId, err := auth.ValidateToken(tokenString)
		if err != nil && !result && userId == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		} else {
			user, err := user.GetUserByID(userId)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
				return
			}
			c.Set("currentUser", user)
		}
	}
}
