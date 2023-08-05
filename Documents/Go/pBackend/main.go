package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type element struct {
	Repo string;
	Title string;
	Description string;
	Img string
	
}

type elementList struct {
	List []element
} 
var List elementList

func main() {
		// Create individual element instances
	element1 := element{"repo1", "Title 1", "Description 1", "image1.jpg"}
	element2 := element{"repo2", "Title 2", "Description 2", "image2.jpg"}
	element3 := element{"repo3", "Title 3", "Description 3", "image3.jpg"}

	// Create an elementList instance and add elements to it
	List = elementList{}
	List.List = append(List.List, element1)
	List.List = append(List.List, element2)
	List.List = append(List.List, element3)
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": List,
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}