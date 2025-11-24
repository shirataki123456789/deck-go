package imagepkg

import (
    "image"
    "image/draw"
    "image/color"
    "github.com/disintegration/imaging"
)

// ComposeDeckImage creates a simple placeholder deck image.
// For full feature parity you can extend with gradients, pasted card images, QR, etc.
func ComposeDeckImage(leader image.Image, cards []image.Image, qr image.Image, deckName string) image.Image {
    const W = 2150
    const H = 2048
    canvas := imaging.New(W, H, color.NRGBA{R: 0xee, G: 0xee, B: 0xee, A: 0xff})

    // draw leader at left
    if leader != nil {
        l := imaging.Resize(leader, 400, 600, imaging.Lanczos)
        canvas = imaging.Paste(canvas, l, image.Pt(48, 48))
    }

    // draw deck name center-top
    // for simplicity we won't render fonts in this template; leave as area for future enhancement

    // draw QR at right if provided
    if qr != nil {
        q := imaging.Resize(qr, 400, 400, imaging.Lanczos)
        canvas = imaging.Paste(canvas, q, image.Pt(W-48-400, 48))
    }

    // draw first row of cards (placeholders)
    x := 48
    y := 700
    for i := 0; i < len(cards) && i < 10; i++ {
        c := imaging.Resize(cards[i], 215, 300, imaging.Lanczos)
        canvas = imaging.Paste(canvas, c, image.Pt(x, y))
        x += 215 + 8
    }

    return canvas
}
