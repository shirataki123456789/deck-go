package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/health", health)
        api.POST("/filter", filterHandler)
        api.POST("/deck/image", deckImageHandler)
        api.GET("/qr", qrHandler)
    }
}
