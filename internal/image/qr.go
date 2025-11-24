package imagepkg

import (
    "image"
    "image/png"
    "bytes"

    qrcode "github.com/skip2/go-qrcode"
)

// GenerateQRPNG returns PNG bytes of a QR code for the given text.
func GenerateQRPNG(text string, size int) ([]byte, error) {
    pngBytes, err := qrcode.Encode(text, qrcode.Medium, size)
    if err != nil {
        return nil, err
    }
    // validate png decode
    _, err = png.Decode(bytes.NewReader(pngBytes))
    if err != nil {
        return nil, err
    }
    return pngBytes, nil
}

// GenerateQRImage returns an image.Image for further composition.
func GenerateQRImage(text string, size int) (image.Image, error) {
    b, err := GenerateQRPNG(text, size)
    if err != nil {
        return nil, err
    }
    img, err := png.Decode(bytes.NewReader(b))
    return img, err
}
