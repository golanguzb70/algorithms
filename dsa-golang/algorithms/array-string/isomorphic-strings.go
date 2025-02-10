package main

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sToT := make(map[byte]byte) // Mapping s -> t
	tToS := make(map[byte]byte) // Mapping t -> s

	for i := 0; i < len(s); i++ {
		sChar := s[i]
		tChar := t[i]

		// Check existing mapping for s -> t
		if val, exists := sToT[sChar]; exists && val != tChar {
			return false
		}

		// Check existing mapping for t -> s
		if val, exists := tToS[tChar]; exists && val != sChar {
			return false
		}

		// Create the mappings
		sToT[sChar] = tChar
		tToS[tChar] = sChar
	}

	return true

}
