package routes

import (
	"github.com/labstack/echo/v4"
	"webapp/pkg/controllers"
)

func Register(srv *echo.Echo) {
	srv.GET("/", controllers.HomeGET)
	srv.GET("/search", controllers.SearchGET)
	srv.GET("/gallery", controllers.GalleryGET)
	srv.GET("/login", controllers.LoginGET)
	srv.GET("/dashboard/galleries", controllers.DashboardGalleriesGET)
	srv.GET("/dashboard/galleries/new", controllers.DashboardGalleriesNewGET)
}
