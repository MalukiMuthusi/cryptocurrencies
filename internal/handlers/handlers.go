package handlers

import (
	"github.com/MalukiMuthusi/cryptocurrencies/internal/store"
	"github.com/gin-gonic/gin"
)

// SetUpRouter sets up the API router
func SetUpRouter(store store.Store, debugPrintRoute DebugPrintRouteFunc) *gin.Engine {

	r := gin.New()

	// log the endpoints
	gin.DebugPrintRouteFunc = debugPrintRoute

	return r
}

// DebugPrintRouteFunc function that prints the API endpoints served by this service
type DebugPrintRouteFunc func(httpMethod, absolutePath, handlerName string, nuHandlers int)
