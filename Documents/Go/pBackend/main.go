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
	element1 := element{"repo1", "Portfolio", "Description 1", "Astro"}
	element2 := element{"repo2", "Image Finder", "Description 2", "Go"}
	element3 := element{"repo3", "User Finder", "Description 3", "Go"}
	element4 := element{"repo3", "Title 3", "Description 3", "Next"}

	// Create an elementList instance and add elements to it
	List = elementList{}
	List.List = append(List.List, element1)
	List.List = append(List.List, element2)
	List.List = append(List.List, element3)
	List.List = append(List.List, element4)
  r := gin.Default()

  r.Use(func(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")
})

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": List,
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}