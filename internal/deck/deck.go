package deck

type Deck struct {
    Name   string            `json:"name"`
    Leader string            `json:"leader"`
    Cards  map[string]int    `json:"cards"` // card_id -> count
}
