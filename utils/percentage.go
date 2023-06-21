package utils

func CalcPercentageDiff(a, b int32) float64 {
	oldValue := float64(a)
	newValue := float64(b)

	difference := newValue - oldValue
	average := (newValue + oldValue) / 2
	percentage := (difference / average) * 100

	return percentage
}
