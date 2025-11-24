package imagepkg

import (
    "image"
    "net/http"
    "time"
    "io"
    "errors"
    "github.com/disintegration/imaging"
)

// DownloadImage downloads an image from URL and returns image.Image (decoded).
func DownloadImage(url string) (image.Image, error) {
    client := http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    if resp.StatusCode != 200 {
        return nil, errors.New("non-200 response")
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    img, err := imaging.Decode(bytesReader(body))
    if err != nil {
        return nil, err
    }
    return img, nil
}

// helper to convert []byte to io.Reader accepted by imaging.Decode
func bytesReader(b []byte) io.Reader {
    return bytes.NewReader(b)
}
