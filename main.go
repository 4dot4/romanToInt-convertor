func romanToInt(s string) int {
    
	roman := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var output int
	for i := 0; i < len(s); i++ {
		if i == 0 {
			output = roman[string(s[i])]
		} else {
            next := 0
			if i < len(s)-1 {
				if roman[string(s[i])] == 1 || roman[string(s[i])] == 10 || roman[string(s[i])] == 100 {
					next = roman[string(s[i+1])]
				}
			}
			if output > roman[string(s[i])] && next > roman[string(s[i])] {
				output += next - roman[string(s[i])]
				i++
			} else if output >= roman[string(s[i])] {
				output += roman[string(s[i])]
			} else {
				output = roman[string(s[i])] - output
			}
		}

	}
	return output
}
