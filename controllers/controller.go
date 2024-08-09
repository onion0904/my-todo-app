package controllers

import (
	"github.com/gin-gonic/gin"
	"TodoApp/models"
	"TodoApp/repository"
	"gorm.io/gorm"
	"strconv"
	"net/http"
	"log"
)

type TodoController struct {
	db *gorm.DB
}

func NewTodoController(db *gorm.DB) *TodoController {
    return &TodoController{db: db}
}

func (con TodoController) Add(c *gin.Context) {
	todo := models.Todo{}
	err := c.Bind(&todo)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
        return
	}

	repo := repository.NewTodoRepository(con.db)
	err = repo.Add(&todo)
	if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"status": "ok"})
}

func (con TodoController) List(c *gin.Context) {
	repo := repository.NewTodoRepository(con.db)
	TodoList := repo.List()
	c.JSON(200, gin.H{
		"message": "ok",
		"list": TodoList,
	})
}

func (con TodoController) Update(c *gin.Context) {
	todo := models.Todo{}
	err := c.Bind(&todo)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	repo := repository.NewTodoRepository(con.db)
	err = repo.Update(&todo)
	if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"status": "ok"})
}

func (con TodoController) Delete(c *gin.Context) {
	id := c.Query("id")
	intId, err := strconv.Atoi(id)
    if err != nil {
		log.Printf("Error converting ID to integer: %v", err)
		c.String(http.StatusBadRequest, "Bad request: invalid ID")
		return
	}
	
	repo := repository.NewTodoRepository(con.db)
	err = repo.Delete(intId)
	if err!= nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"status": "ok"})
}