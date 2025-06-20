package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SimulatedService(name string, delay time.Duration, ch chan<- string) {
	time.Sleep(delay)
	ch <- fmt.Sprintf("%s finished after %v", name, delay)
}

func ConcurrentFetch(c *gin.Context) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start goroutines for both services
	go SimulatedService("Service A", 2*time.Second, ch1)
	go SimulatedService("Service B", 3*time.Second, ch2)

	// Receive results from both channels
	res1 := <-ch1
	res2 := <-ch2

	c.JSON(http.StatusOK, gin.H{
		"result1": res1,
		"result2": res2,
	})
}
