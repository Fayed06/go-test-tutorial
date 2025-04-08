package lesson2

func Keys(m map[string]struct{}) []string {
	keys := []string{}
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func Dedup(input []string) []string {
	m := map[string]struct{}{}
	for _, s := range input {
		m[s] = struct{}{}
	}
	return Keys(m)
}
