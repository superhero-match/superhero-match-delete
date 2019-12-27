package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ctrl "github.com/superhero-match-delete/cmd/api/model"
	"go.uber.org/zap"
)

// Delete publishes delete match message on Kafka for it to be marked as deleted in DB.
func (ctl *Controller) Delete(c *gin.Context) {
	var req ctrl.Request

	err := c.BindJSON(&req)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type Request",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	err = ctl.Service.DeleteMatch(ctrl.Match{
		ID:                 req.ID,
	})
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.DeleteMatch()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}

func checkError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})

		return true
	}

	return false
}
