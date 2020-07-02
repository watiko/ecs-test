package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
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

	ln, err := listen(server.Addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(-1)
	}

	server.Serve(ln)
}

func listen(address string) (net.Listener, error) {
	lc := net.ListenConfig{
		KeepAlive: 2 * time.Minute,
	}
	ln, err := lc.Listen(context.Background(), "tcp", address)

	if err != nil {
		return nil, err
	}
	return ln, nil
}
