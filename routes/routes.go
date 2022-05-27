package routes

import (
	"github.com/uere/grafana-backup/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.POST("/dashboards", controllers.SaveGrafanaDashboards)
	r.Run()
}
