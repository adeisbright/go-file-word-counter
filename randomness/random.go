package randomness

import (
	"fmt"
	"math/rand"
)

func GenerateNumber() {
	number := rand.Intn(4)
	fmt.Printf("the random number generated is %d \n", number)
	var floater = rand.Float64()
	fmt.Printf("the random float generated is %f \n", floater)
}
