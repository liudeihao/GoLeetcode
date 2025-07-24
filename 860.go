package main

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, b := range bills {
		switch b {
		case 5:
			fives++
		case 10:
			fives--
			tens++
			if fives < 0 {
				return false
			}
		case 20:
			if tens > 0 {
				tens--
				fives--
			} else {
				fives -= 3
			}
			if tens < 0 || fives < 0 {
				return false
			}
		}
	}
	return true
}
