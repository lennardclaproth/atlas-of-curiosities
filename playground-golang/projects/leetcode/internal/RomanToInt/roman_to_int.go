package romantoint

func main() {

}

var m = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// converts roman numerals as a string to an integer
func romanToInt(s string) int {
	// convert input string to list of characters (runes in golang)
	rs := []rune(s)

	// Get the max length so that we can reverse over it
	max := len(rs) - 1
	res := 0

	// Get the range and assign it to i
	var prev rune = 0
	var cur rune
	for i := range rs {
		// gets the runes in reverse order
		cur = rs[max-i]
		if prev == 0 {
			res += m[cur]
			prev = cur
			continue
		}

		// if cur is less than prev
		if m[cur] < m[prev] {
			res -= m[cur]
		} else if m[cur] >= m[prev] {
			res += m[cur]
		}

		prev = cur
	}

	return res
}
