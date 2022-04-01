package main

import (
	"net/http"
	"os"

	aero "github.com/aerospike/aerospike-client-go/v5"
	"github.com/gin-gonic/gin"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	client, err := aero.NewClient(
		getEnv("AEROSPIKE_HOST", "127.0.0.1"),
		3000,
	)
	panicOnError(err)

	key, err := aero.NewKey("test", "users", "1")
	panicOnError(err)

	bins := aero.BinMap{
		"api_key":    "12345",
		"first_name": "John",
		"last_name":  "Doe",
		"company":    "Acme",
	}

	err = client.Put(nil, key, bins)
	panicOnError(err)

	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		api_key := c.Query("api_key")

		req_key, err := aero.NewKey("test", "users", id)
		panicOnError(err)

		rec, err := client.Get(nil, req_key)
		panicOnError(err)

		if rec.Bins["api_key"] == api_key {
			c.JSON(http.StatusOK, gin.H{
				"api_key": api_key,
			})
		}
	})
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
