package leetcode

func reverseDegree(s string) int {
	output := 0
	for pos, char := range s {
		index := (26 - char + 'a')
		output += int(index) * (pos + 1)
	}
	return output
}
