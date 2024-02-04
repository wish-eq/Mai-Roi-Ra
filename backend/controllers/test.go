package controllers

import (
	"log"
	"net/http"

	"github.com/2110366-2566-2/Mai-Roi-Ra/backend/services"
	_ "github.com/2110366-2566-2/Mai-Roi-Ra/backend/swagger/docs" // Import the auto-generated docs file

	"github.com/gin-gonic/gin"
)

// @Summary GetTest
// @Description Get a test message
// @Tags Test
// @Accept json
// @Produce json
// @Success 200 {object} TestResponse
// @Router /test [get]
func GetTest(c *gin.Context) () {
	log.Println("[CTRL: GetTest] Called")

	res, err := services.GetTest()

	if err != nil {
		log.Fatal("[CTRL: GetTest] error: ", err)
	}
	c.JSON(http.StatusOK, res)
}