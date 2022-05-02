package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomeGET(c echo.Context) error {
	return c.Render(http.StatusOK, "index.page.html", nil)
}

func SearchGET(c echo.Context) error {
	return c.Render(http.StatusOK, "search.page.html", nil)
}

func GalleryGET(c echo.Context) error {
	return c.Render(http.StatusOK, "gallery.page.html", nil)
}

func LoginGET(c echo.Context) error {
	return c.Render(http.StatusOK, "login.page.html", nil)
}

func DashboardGalleriesGET(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard_galleries.page.html", nil)
}

func DashboardGalleriesNewGET(c echo.Context) error {
	return c.Render(http.StatusOK, "dashboard_galleries_new.page.html", nil)
}
