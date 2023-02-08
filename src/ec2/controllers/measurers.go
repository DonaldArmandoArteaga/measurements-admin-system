package controllers

import (
	"input-system/config"
	"input-system/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	path = "/measurers"
)

type MeasurersController struct {
	measurerService services.Measurers
}

func CreateMeasurersController(r *gin.Engine, measurerService services.Measurers) {
	measurersController := &MeasurersController{
		measurerService: measurerService,
	}
	r.GET(path+"/:id", measurersController.getById)
	r.GET(path, measurersController.getBySerial)
}

func (ms *MeasurersController) getById(c *gin.Context) {
	measurer, err := ms.measurerService.GetById(c.Param("id"))

	if err != nil {
		config.ErrorLogger.Println("Got an error while trying to retrieve measurers by Id:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, measurer)
}

func (ms *MeasurersController) getBySerial(c *gin.Context) {

	measurers, count, err := ms.measurerService.GetBySerial(c.Query("serial"))

	if err != nil {
		config.ErrorLogger.Println("Got an error while trying to retrieve measurers by Serial :", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":     count,
		"measurers": measurers,
	})
}
