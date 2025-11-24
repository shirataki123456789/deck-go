package cards

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)

func parseListCell(s string) []string {
    s = strings.ReplaceAll(s, "／", "/")
    parts := strings.Split(s, "/")
    out := []string{}
    for _, p := range parts {
        t := strings.TrimSpace(p)
        if t != "" && t != "-" {
            out = append(out, t)
        }
    }
    return out
}

// LoadCardsFromDataDir loads CSV files from a data directory (best-effort).
// It expects at least cardlist_filtered.csv; custom and parallel CSVs are optional.
func LoadCardsFromDataDir(dataDir string) ([]Card, error) {
    files := []string{
        filepath.Join(dataDir, "cardlist_filtered.csv"),
        filepath.Join(dataDir, "custom_cards.csv"),
        filepath.Join(dataDir, "cardlist_p_only.csv"),
    }

    var all []Card
    var found bool
    for _, f := range files {
        if _, err := os.Stat(f); err != nil {
            // skip missing files
            continue
        }
        found = true
        cs, err := loadSingleCSV(f)
        if err != nil {
            return nil, fmt.Errorf("loading %s: %w", f, err)
        }
        all = append(all, cs...)
    }
    if !found {
        return nil, fmt.Errorf("no input CSVs found in %s", dataDir)
    }
    return all, nil
}

func loadSingleCSV(path string) ([]Card, error) {
    fp, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer fp.Close()

    r := csv.NewReader(fp)
    r.FieldsPerRecord = -1
    rows, err := r.ReadAll()
    if err != nil {
        return nil, err
    }
    if len(rows) < 1 {
        return nil, fmt.Errorf("csv %s has no header", path)
    }
    header := rows[0]
    cols := map[string]int{}
    for i, h := range header {
        cols[h] = i
    }

    get := func(row []string, name string) string {
        if idx, ok := cols[name]; ok && idx < len(row) {
            return row[idx]
        }
        return ""
    }

    out := []Card{}
    for _, row := range rows[1:] {
        c := Card{
            CardID:   get(row, "カードID"),
            Name:     get(row, "カード名"),
            Color:    get(row, "色"),
            Type:     get(row, "タイプ"),
            Counter:  get(row, "カウンター"),
            Text:     get(row, "テキスト"),
            Trigger:  get(row, "トリガー"),
            BlockIcon: get(row, "ブロックアイコン"),
            ImageURL: get(row, "画像URL"),
        }
        costStr := get(row, "コスト")
        if costStr == "-" || costStr == "" {
            c.Cost = 0
        } else {
            v, _ := strconv.Atoi(costStr)
            c.Cost = v
        }
        // features / attributes
        c.Features = parseListCell(get(row, "特徴"))
        c.Attributes = parseListCell(get(row, "属性"))

        // is_parallel: infer from filename or column "is_parallel"
        isPar := get(row, "is_parallel")
        if isPar == "True" || isPar == "true" || isPar == "1" {
            c.IsParallel = true
        } else {
            c.IsParallel = false
        }

        // series id from 入手情報's 【】
        info := get(row, "入手情報")
        series := "その他"
        if strings.Contains(info, "【") && strings.Contains(info, "】") {
            a := strings.Index(info, "【")
            b := strings.Index(info, "】")
            if b > a+1 {
                series = strings.TrimSpace(info[a+1:b])
            }
        } else if strings.TrimSpace(info) == "-" || info == "" {
            series = "-"
        }
        c.SeriesID = series

        out = append(out, c)
    }
    return out, nil
}
