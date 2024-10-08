package main

import(
	"github.com/gin-contrib/cors"
	"fmt"
	"github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"os"
	"log"
	"TodoApp/models"
	"TodoApp/controllers"
	"github.com/joho/godotenv"
	"time"
)

func main() {
	err := godotenv.Load()  // Load .env file
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	err = db.AutoMigrate(&models.Todo{})
    if err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

	Todo := gin.Default()

	con := controllers.NewTodoController(db)

	method := Todo.Group("/Todo")
	{
		method.POST("/add",con.Add)
		method.GET("/list",con.List)
		method.PUT("/update",con.Update)
		method.DELETE("/delete",con.Delete)
	}

	Todo.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        AllowOriginFunc: func(origin string) bool {
            return origin == "http://localhost:3000"
        },
        MaxAge: 12 * time.Hour,
    }))

    Todo.Run(":8080")
}