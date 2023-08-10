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
	element1 := element{"https://github.com/CHenrique97/chenrique97.github.io", "Portfolio", "This is my re-re-re-made portfolio, where I showcase the projects that I consider are my best work", "Astro"}
	element2 := element{"https://github.com/CHenrique97/ImageFinder", "Image Finder","This is a microservice that uses gRPC to comunicate with the bff of a pet-project of a friend of mine. It is responsible for finding the profile pic of the users registered in our databank, thus the name", "Go"}
	element3 := element{"https://github.com/CHenrique97/UserFinder", "User Finder","This is a microservice that uses gRPC to comunicate with the bff of a pet-project of a friend of mine. It is responsible for finding the Users registered in our databank, thus the name", "Go"}
	element4 := element{"https://github.com/CHenrique97/firula", "Firula", "This is a service that facilitates the renting of courts and fields for group games,made using next, still WIP", "Next"}

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