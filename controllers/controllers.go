package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uere/grafana-backup/models"
)

func GetDashboards(c *gin.Context) {
	// var dashboards []models.DashboardMeta
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
	c.JSON(200, ListDashboards)
}
