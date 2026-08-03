package play

import "golang.org/x/exp/slices"

func FilterPlays(plays []Play, onlyTags, skipTags []string) []Play {

	var filteredPlays []Play
	for _, play := range plays {
		skip := false
		for _, tag := range skipTags {
			if slices.Contains(play.Tags, tag) {
				skip = true
			}
		}
		if skip {
			continue
		}
		if len(onlyTags) > 0 {
			skip = true
			for _, tag := range onlyTags {
				if slices.Contains(play.Tags, tag) {
					skip = false
				}
			}
		}
		if skip {
			continue
		}
		// If play passes all filters, add it to the filteredPlays list
		filteredPlays = append(filteredPlays, play)
	}
	return filteredPlays
}
