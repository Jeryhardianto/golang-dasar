// * container Ring

package main
import (
	"fmt"
	"container/ring"
	"strconv"
)

func main() {

	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data "+ strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}
	data.Do(func(element interface{}) {
		fmt.Println(element)
	})
	// fmt.Println(*data)

}