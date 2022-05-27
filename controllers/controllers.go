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
	models.GetDashboards(&g, ListDashboards)
	c.JSON(http.StatusOK, gin.H{
		"grafana":    g.Url,
		"dashboards": g.Dashboards,
		"project":    g.Project})
}
