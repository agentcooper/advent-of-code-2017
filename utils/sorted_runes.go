package utils

// SortedRunes provides Sort interface for []rune
type SortedRunes []rune

func (s SortedRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortedRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortedRunes) Len() int {
	return len(s)
}
