package repository

type Todo struct {
    gorm.Model
    Name  string
    text string 
    TimeLimit int
}
