package routes

import (
	"github.com/uere/grafana-backup/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/dashboards", controllers.GetDashboards)
	r.Run()
}
