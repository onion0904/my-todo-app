package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string
    Email string
    Age   int
}


//テスト
func main() {

    dbUser := "onion"
	dbPassword := "noino"
	dbDatabase := "Todo_db"
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?parseTime=true", dbUser,dbPassword, dbDatabase)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "John Doe", Email: "john@example.com", Age: 25})

	r := gin.Default()

    r.GET("/user/:id", func(c *gin.Context) {
        var user User
        id := c.Param("id")
        
        result := db.First(&user, id)
        if result.Error != nil {
            if result.Error == gorm.ErrRecordNotFound {
                c.JSON(404, gin.H{"error": "User not found"})
                return
            }
            c.JSON(500, gin.H{"error": result.Error.Error()})
            return
        }

        c.JSON(200, user)
    })

    r.Run()
}