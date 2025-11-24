package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/youruser/deckapp/internal/api"
    "github.com/youruser/deckapp/internal/cards"
)

func main() {
    // Load cards at startup (best-effort)
    _, err := cards.LoadCardsFromDataDir("data")
    if err != nil {
        log.Println("Warning: failed to load CSVs at startup:", err)
    }

    r := gin.Default()
    api.RegisterRoutes(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Println("starting server on http://localhost:" + port)
    if err := r.Run(":" + port); err != nil && err != http.ErrServerClosed {
        log.Fatal(err)
    }
}
