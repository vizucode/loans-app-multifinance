package formatting

import "time"

func CheckAge(dateOfBirth time.Time) int {
	return time.Now().Year() - dateOfBirth.Year()
}
