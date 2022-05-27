package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uere/grafana-backup/models"
)

func SaveGrafanaDashboards(c *gin.Context) {
	// var dashboards []models.DashboardMeta
	var g models.Grafana
	if err := c.ShouldBindJSON(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	if err := models.ValidateGrafana(&g); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	ListDashboards := models.ListDashboards(&g)
	c.JSON(200, ListDashboards)
}
