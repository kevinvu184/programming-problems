package lib

func Sort(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}
