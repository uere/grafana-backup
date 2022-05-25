package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uere/grafana-backup/models"
)

func ListDashboardsHandleFunc(c *gin.Context) {
	var backup models.Backup
	if err := c.ShouldBindJSON(&backup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidaBackup(&backup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	ListDashboards := models.ListDashboards(&backup)
	models.GetDashboards(&backup, ListDashboards)
	c.JSON(200, ListDashboards)
}
