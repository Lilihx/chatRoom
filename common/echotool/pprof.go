package echotool

import (
	"net/http"
	"net/http/pprof"

	"github.com/labstack/echo/v4"
)

func RegisterPProf(e *echo.Echo) {
	router := e.Group("")
	router.Use()

	router.GET("/debug/pprof", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/allocs", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/block", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/goroutine", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/heap", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/mutex", echo.WrapHandler(http.HandlerFunc(pprof.Index)))
	router.GET("/debug/pprof/cmdline", echo.WrapHandler(http.HandlerFunc(pprof.Cmdline)))
	router.GET("/debug/pprof/profile", echo.WrapHandler(http.HandlerFunc(pprof.Profile)))
	router.GET("/debug/pprof/symbol", echo.WrapHandler(http.HandlerFunc(pprof.Symbol)))
	router.GET("/debug/pprof/trace", echo.WrapHandler(http.HandlerFunc(pprof.Trace)))
}
