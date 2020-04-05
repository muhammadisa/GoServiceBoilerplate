package mapsmerger

// MapsMerger merging maps into one maps
func MapsMerger(maps ...map[string]string) map[string]string {
	result := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}