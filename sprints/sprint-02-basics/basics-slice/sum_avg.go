package problems

import "fmt"

// SumAndAverage считает сумму и среднее значение элементов целочисленного слайса.
//
// Вход:
//   - nums []int: входной слайс, например []int{1, 2, 3}
//
// Выход:
//   - sum int: сумма элементов
//   - avg float64: среднее арифметическое (с плавающей точностью)
//
// Пример:
//   - nums = []int{1, 2, 3} -> sum = 6, avg = 2.0
func SumAndAverage(nums []int) (sum int, avg float64) {
	if len(nums) == 0 {
		return 0, 0.0
	}
	sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	avg = float64(sum) / float64(len(nums))
	fmt.Printf("sum = %v, avg = %v\n", sum, avg)

	return sum, avg
}
