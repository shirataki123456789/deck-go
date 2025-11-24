package cards

type Card struct {
    CardID     string   `json:"card_id"`
    Name       string   `json:"name"`
    Color      string   `json:"color"`
    Type       string   `json:"type"`
    Cost       int      `json:"cost"`
    Counter    string   `json:"counter"`
    Features   []string `json:"features"`
    Attributes []string `json:"attributes"`
    Text       string   `json:"text"`
    Trigger    string   `json:"trigger"`
    BlockIcon  string   `json:"block_icon"`
    IsParallel bool     `json:"is_parallel"`
    ImageURL   string   `json:"image_url"`
    SeriesID   string   `json:"series_id"`
}
