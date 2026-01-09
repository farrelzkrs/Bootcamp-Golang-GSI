package ring

import (
	"container/ring"
	"fmt"
	"strconv"
)

func Cincin() {
	var data *ring.Ring = ring.New(5)
	fmt.Print("-> ")
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(val interface{}) {
		fmt.Print(val, " -> ")
	})
}