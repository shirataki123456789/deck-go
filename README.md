# DeckApp — Go migration of deck.py

This repository is an initial, runnable scaffold of the `deck.py` Streamlit app translated to **Go**, intended to be developed further.

## Features included
- CSV loader (best-effort; expects `data/*.csv`)
- Filter API endpoint (`POST /api/filter`)
- QR PNG generator (`GET /api/qr?text=...`)
- Deck image generation endpoint (`POST /api/deck/image`) — composes leader, cards and QR into a PNG (basic)

## Quick start (GitHub Codespaces)
1. Open this repo in GitHub Codespaces (recommended).
2. Codespace will run `go mod tidy` as configured in devcontainer.
3. Start the server:
   ```bash
   cd /workspaces/deckapp
   go run ./cmd/server
   ```
4. Open `http://localhost:8080/api/health`

## Notes
- This scaffold uses:
  - `github.com/gin-gonic/gin` for API
  - `github.com/skip2/go-qrcode` for QR generation
  - `github.com/disintegration/imaging` for image processing
- The original uploaded Python file has been copied into `data/original_deck.py` for reference.

