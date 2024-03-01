package randomness

import (
	"fmt"
	"time"
)

func Time() {
	currentTime := time.Now()
	fmt.Println("Working with time ", currentTime)
}
