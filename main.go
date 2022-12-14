package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct{
	Id string `json:"id"`
	Title string 	`json:"title"`
	Completed bool 	`json:"completed"`
}

var todos= []todo{
	{Id: "1", Title: "Learn Go", Completed: false},
	{Id: "2", Title: "Learn Rust", Completed: false},
	{Id: "3", Title: "Learn C", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}
func postTodos(context *gin.Context){
	var newTodo todo
	if err := context.BindJSON(&newTodo); err!=nil{
		return
	}
	todos=append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context){
	id := context.Param("id")
	todo,err:= getTodoById(id)

	if err !=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo) 
	
}
func getTodoById(id string)(*todo,error){
	for i,t := range todos{
		if t.Id==id{
			return &todos[i],nil
		}
	}
	return nil,errors.New("not found")

}
func updateById(context *gin.Context){
	id := context.Param("id")
	todo,err:= getTodoById(id)

	if err !=nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message":"not found"})
		return
	}
	todo.Completed=!todo.Completed
	context.IndentedJSON(http.StatusOK, todo) 
}
func main(){
	router:= gin.Default()
	router.GET("/todos",getTodos)
	router.GET("/todos/:id",getTodo)
	router.POST("/todos",postTodos)
	router.PATCH("/todos/:id",updateById)
	router.Run("localhost:5001")
}