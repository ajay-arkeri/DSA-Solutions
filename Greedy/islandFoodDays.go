package greedy

func BuyDays(S int, N int, M int) int {
	total_food := S * M

	total_sundays := S / 7

	buyDays := S - total_sundays
	resultDays := 0

	if total_food%N == 0 {
		resultDays = total_food / N
	} else {
		resultDays = (total_food / N) + 1
	}

	if resultDays <= buyDays {
		return resultDays
	} else {
		return -1
	}
}
