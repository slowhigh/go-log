package calculate_insufficient_amount_of_money

func solution(price int, money int, count int) int64 {
	var totalPrice, result int64
	totalPrice = 0

	for i := 1; i <= count; i++ {
		totalPrice += int64(price) * int64(i)
	}

	result = totalPrice - int64(money)

	if result <= 0 {
		return 0
	} else {
		return result
	}
}

func solution2(price int, money int, count int) int64 {
	totalPrice := int64(price) * (int64(count) * (int64(count) + 1) / 2)
	result := totalPrice - int64(money)

	if result <= 0 {
		return 0
	} else {
		return result
	}
}
