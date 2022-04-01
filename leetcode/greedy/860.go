package greedy

func lemonadeChange(bills []int) bool {
	tenDollar, fiveDoller := 0, 0
	for _, bill := range bills {
		switch {
		case bill == 5:
			fiveDoller++
		case bill == 10:
			fiveDoller--
			if fiveDoller < 0 {
				return false
			}
			tenDollar++
		case bill == 20:
			if fiveDoller > 0 && tenDollar > 0 {
				fiveDoller--
				tenDollar--
			} else if fiveDoller >= 3 {
				fiveDoller -= 3
			} else {
				return false
			}
		}
	}
	return true
}
