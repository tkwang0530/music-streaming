package utils

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, status int, err error) {
	log.Printf("ERROR: %v", err)
	c.AbortWithStatusJSON(status, gin.H{"error": err.Error()})
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database?charset=utf8")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
