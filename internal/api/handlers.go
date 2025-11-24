package api

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/youruser/deckapp/internal/cards"
    imagepkg "github.com/youruser/deckapp/internal/image"
    "image/png"
    "bytes"
    "log"
    "strconv"
)

// health
func health(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// simplified filter handler
func filterHandler(c *gin.Context) {
    var opt cards.FilterOptions
    if err := c.BindJSON(&opt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    all, err := cards.LoadCardsFromDataDir("data")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    out := cards.Filter(all, opt)
    c.JSON(http.StatusOK, gin.H{"count": len(out), "cards": out})
}

// qr endpoint returns a PNG of a QR for "text" query param
func qrHandler(c *gin.Context) {
    text := c.Query("text")
    if text == "" {
        text = "deck:example"
    }
    sizeStr := c.Query("size")
    size := 400
    if sizeStr != "" {
        if v, err := strconv.Atoi(sizeStr); err == nil {
            size = v
        }
    }
    b, err := imagepkg.GenerateQRPNG(text, size)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Data(http.StatusOK, "image/png", b)
}

// deck image: accepts minimal JSON with leader_url and card_urls (array) and deck_name
func deckImageHandler(c *gin.Context) {
    var req struct {
        LeaderURL string   `json:"leader_url"`
        CardURLs  []string `json:"card_urls"`
        DeckName  string   `json:"deck_name"`
        QRText    string   `json:"qr_text"`
    }
    if err := c.BindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // fetch leader and cards (best-effort)
    var leaderImg image.Image
    if req.LeaderURL != "" {
        img, err := imagepkg.DownloadImage(req.LeaderURL)
        if err == nil {
            leaderImg = img
        } else {
            log.Println("leader download error:", err)
        }
    }
    var cardImgs []image.Image
    for i, u := range req.CardURLs {
        if i >= 10 { break }
        img, err := imagepkg.DownloadImage(u)
        if err == nil {
            cardImgs = append(cardImgs, img)
        }
    }
    var qrImg image.Image
    if req.QRText != "" {
        q, err := imagepkg.GenerateQRImage(req.QRText, 400)
        if err == nil {
            qrImg = q
        }
    }
    out := imagepkg.ComposeDeckImage(leaderImg, cardImgs, qrImg, req.DeckName)
    buf := new(bytes.Buffer)
    if err := png.Encode(buf, out); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Data(http.StatusOK, "image/png", buf.Bytes())
}
