package inventory

import (
	"errors"
	"strings"

	inventoryV1 "github.com/mercuryqa/rocket-lab/inventory/pkg/proto/inventory_v1"
)

// helper to check if filter is empty
func isFilterEmpty(f *inventoryV1.PartsFilter) bool {
	if f == nil {
		return true
	}
	return len(f.Uuids) == 0 && len(f.Names) == 0 && len(f.Categories) == 0 && len(f.ManufacturerCountries) == 0 && len(f.Tags) == 0
}

// Filter helpers
func filterByUUIDs(parts []*inventoryV1.Part, uuids []string) ([]*inventoryV1.Part, error) {
	set := make(map[string]struct{}, len(uuids))
	for _, u := range uuids {
		if strings.TrimSpace(u) == "" {
			continue
		}
		set[u] = struct{}{}
	}
	if len(set) == 0 {
		return nil, errors.New("provided uuids are empty")
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if _, ok := set[p.Uuid]; ok {
			out = append(out, p)
		}
	}
	return out, nil
}

func filterByNames(parts []*inventoryV1.Part, names []string) []*inventoryV1.Part {
	set := make(map[string]struct{}, len(names))
	for _, n := range names {
		set[n] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if _, ok := set[p.Name]; ok {
			out = append(out, p)
		}
	}
	return out
}

func filterByCategories(parts []*inventoryV1.Part, cats []inventoryV1.Category) []*inventoryV1.Part {
	set := make(map[inventoryV1.Category]struct{}, len(cats))
	for _, c := range cats {
		set[c] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if _, ok := set[p.Category]; ok {
			out = append(out, p)
		}
	}
	return out
}

func filterByCountries(parts []*inventoryV1.Part, countries []string) []*inventoryV1.Part {
	set := make(map[string]struct{}, len(countries))
	for _, c := range countries {
		set[c] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		if p.Manufacturer != nil {
			if _, ok := set[p.Manufacturer.Country]; ok {
				out = append(out, p)
			}
		}
	}
	return out
}

func filterByTags(parts []*inventoryV1.Part, tags []string) []*inventoryV1.Part {
	set := make(map[string]struct{}, len(tags))
	for _, t := range tags {
		set[t] = struct{}{}
	}
	out := make([]*inventoryV1.Part, 0)
	for _, p := range parts {
		for _, pt := range p.Tags {
			if _, ok := set[pt]; ok {
				out = append(out, p)
				break
			}
		}
	}
	return out
}
