package cards

import "strings"

type FilterOptions struct {
    Colors       []string
    Types        []string
    Costs        []int
    Counters     []string
    Attributes   []string
    Blocks       []string
    Features     []string
    FreeWords    string
    SeriesIDs    []string
    LeaderColors []string
    ParallelMode string // "normal", "parallel", "both"
}

func containsAny(hay []string, needles []string) bool {
    for _, n := range needles {
        for _, h := range hay {
            if strings.Contains(h, n) {
                return true
            }
        }
    }
    return false
}

func Filter(cards []Card, opt FilterOptions) []Card {
    var out []Card
    for _, c := range cards {
        // parallel filter
        if opt.ParallelMode == "normal" && c.IsParallel {
            continue
        }
        if opt.ParallelMode == "parallel" && !c.IsParallel {
            continue
        }
        // leader color filtering (exclude LEADER type in deck addition context)
        if len(opt.LeaderColors) > 0 {
            if c.Type == "LEADER" {
                continue
            }
            ok := false
            for _, lc := range opt.LeaderColors {
                if strings.Contains(c.Color, lc) {
                    ok = true
                    break
                }
            }
            if !ok {
                continue
            }
        }
        if len(opt.Colors) > 0 {
            if !containsAny([]string{c.Color}, opt.Colors) {
                continue
            }
        }
        if len(opt.Types) > 0 {
            matched := false
            for _, t := range opt.Types {
                if c.Type == t {
                    matched = true
                    break
                }
            }
            if !matched {
                continue
            }
        }
        if len(opt.Costs) > 0 {
            matched := false
            for _, co := range opt.Costs {
                if c.Cost == co {
                    matched = true
                    break
                }
            }
            if !matched {
                continue
            }
        }
        if len(opt.Attributes) > 0 {
            if !containsAny(c.Attributes, opt.Attributes) {
                continue
            }
        }
        if len(opt.Features) > 0 {
            if !containsAny(c.Features, opt.Features) {
                continue
            }
        }
        if len(opt.SeriesIDs) > 0 {
            matched := false
            for _, s := range opt.SeriesIDs {
                if c.SeriesID == s {
                    matched = true
                    break
                }
            }
            if !matched {
                continue
            }
        }
        if opt.FreeWords != "" {
            kw := strings.Fields(opt.FreeWords)
            ok := true
            for _, k := range kw {
                k = strings.ToLower(k)
                if !strings.Contains(strings.ToLower(c.Name), k) &&
                    !strings.Contains(strings.ToLower(c.Text), k) &&
                    !strings.Contains(strings.ToLower(strings.Join(c.Features, " ")), k) &&
                    !strings.Contains(strings.ToLower(c.Trigger), k) {
                    ok = false
                    break
                }
            }
            if !ok {
                continue
            }
        }
        out = append(out, c)
    }
    return out
}
