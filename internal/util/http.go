package util

import (
    "io"
    "net/http"
    "time"
)

func GetBytes(url string) ([]byte, error) {
    client := http.Client{Timeout: 12 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}
