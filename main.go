package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct{
	Id int `json:"id"`
	Title string 	`json:"title"`
	Completed bool 	`json:completed`
}

var todos= []todo{
	{Id: 1, Title: "Learn Go", Completed: false},
	{Id: 2, Title: "Learn Rust", Completed: false},
	{Id: 3, Title: "Learn C", Completed: false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}


func main(){
	router:= gin.Default()
	router.GET("/todos",getTodos)
		router.Run("localhost:5001")
}