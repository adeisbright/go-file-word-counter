package randomness

import (
	"fmt"
	"time"
)

func Time() {
	currentTime := time.Now()
	twoDaysLater := currentTime.AddDate(0, 0, 2)
	oneMonthAhead := currentTime.AddDate(0, 1, 0)
	timeDiff := oneMonthAhead.Sub(twoDaysLater)
	result := timeDiff.Hours() / 24
	fmt.Println("Date : ", twoDaysLater)
	fmt.Println("Time Difference: ", result)
}
