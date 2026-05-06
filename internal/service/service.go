package service

// GetOutcome converts scores to 1 (Home Win), 2 (Away Win), or 0 (Draw)
func GetOutcome(a, b int32) int {
	if a > b {
		return 1
	}
	if b > a {
		return 2
	}
	return 0
}

// CheckIfWinnerGuessed returns 1 point if the user correctly guessed the outcome
func CheckIfWinnerGuessed(actualA, actualB, predA, predB int32) int {
	if GetOutcome(actualA, actualB) == GetOutcome(predA, predB) {
		return 1
	}
	return 0
}
