package basic

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// SimpleTest --localhost/basic/simpleTest
func SimpleTest(context *gin.Context) {
	log.Println("@@@SimpleTest is running......")
	context.JSON(http.StatusOK, gin.H{"status": "U r ok"})
}
