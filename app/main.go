package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	router := gin.Default()

	router.GET("/wait/:sec", func(c *gin.Context) {
		i, err := strconv.Atoi(c.Param("sec"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("%v", err),
			})
			return
		}

		sec := max(0, i)
		time.Sleep(time.Second * time.Duration(sec))

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("complete(wait %d)", sec),
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	server.ListenAndServe()
}
