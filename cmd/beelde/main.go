package main

import (
	"github.com/gin-gonic/gin"
	orm "github.com/jinzhu/gorm"
	cors "github.com/rs/cors/wrapper/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("/v1")

	{
		v1.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})

		v1.GET("/db", func(c *gin.Context) {
			log.Info("Hello, World!")
			db, err := orm.Open("sqlite3", "test.db")
			if err != nil {
				panic("failed to connect database")
			}
			defer db.Close()
			c.Status(http.StatusNoContent)
		})
	}

	router.Run(":8080")

}
