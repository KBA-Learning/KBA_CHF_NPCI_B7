package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type user struct {
    Id       string `json:"id"`
    Username string `json:"username"`
     }


	 var users = make(map[string]user)

func main(){

	router := gin.Default()

	router.GET("/", Hello)

	router.POST("create", func(c *gin.Context) {
        var user user
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
            return
        }

        users[user.Id] = user// map user structure to users map with key as user.id, mapping is a key valu pair without key mapping not working

        c.JSON(http.StatusOK, user)
    })

	router.GET("read/:id", func(c *gin.Context) {
        key := c.Param("id")

        value := users[key]

        c.JSON(http.StatusOK, gin.H{"data": value})

    })


	router.Run(":3000")

}

func Hello(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"Message": "Hello World"})
}