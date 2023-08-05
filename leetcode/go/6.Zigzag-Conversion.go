package leetcode

// https://leetcode.com/problems/zigzag-conversion/
// PAY PAL ISH IRI NG
func Convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	colStep := numRows + numRows - 2

	ss := make([]uint8, len(s))

	diagStep := colStep - 2

	i := 0
	for row := 0; row < numRows; row++ {
		// does this row have diagonal cells?
		diag := row > 0 && row < numRows-1
		for j := row; j < len(ss); j += colStep {
			ss[i] = s[j] // column value
			i += 1
			if diag && j+diagStep < len(s) {
				ss[i] = s[j+diagStep] // diagonal value
				i += 1
			}
		}
		if diag {
			diagStep -= 2
		}
	}

	return string(ss)
}
