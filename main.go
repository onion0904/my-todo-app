package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
	"os"
)

// type User struct {
//     gorm.Model
//     Name  string
//     Email string
//     Age   int
// }

var (
	dbUser := os.Getenv("USERNAME")
	dbPassword := os.Getenv("USERPASS")
	dbDatabase := os.Getenv("DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
)

func main() {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "John Doe", Email: "john@example.com", Age: 25})

	Todo := gin.Default()

	method := Todo.Group("Todo"){
		method.POST("/add",controller.Add)
		method.GET("/list",controller.List)
		method.PUT("/update",controller.Update)
		method.DELETE("/delete",controller.Delete)
	}


	//test

    // r.GET("/user/:id", func(c *gin.Context) {
    //     var user User
    //     id := c.Param("id")
        
    //     result := db.First(&user, id)
    //     if result.Error != nil {
    //         if result.Error == gorm.ErrRecordNotFound {
    //             c.JSON(404, gin.H{"error": "User not found"})
    //             return
    //         }
    //         c.JSON(500, gin.H{"error": result.Error.Error()})
    //         return
    //     }

    //     c.JSON(200, user)
    // })

    r.Run(":8080")
}