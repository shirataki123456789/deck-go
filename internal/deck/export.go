package deck

import "strings"

func ExportDeckText(d Deck) string {
    lines := []string{}
    if d.Name != "" {
        lines = append(lines, "# "+d.Name)
    }
    lines = append(lines, "1x"+d.Leader)
    // simple deterministic order
    for id, cnt := range d.Cards {
        lines = append(lines, strings.TrimSpace(strings.Itoa(cnt)+"x"+id))
    }
    return strings.Join(lines, "\n")
}
