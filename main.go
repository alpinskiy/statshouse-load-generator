package main

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/vkcom/statshouse-go"
)

func main() {
	m := map[statshouse.Tags]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			m[statshouse.Tags{
				1: strconv.Itoa(10 * (i + 1)),
				2: strconv.Itoa(j + 1),
			}] = 10*(i+1) + (j + 1)
		}
	}
	fmt.Printf("%d\n", len(m))
	for {
		v := float64(math.Sin(2 * math.Pi * float64(time.Now().Second()) / 60.))
		for k, n := range m {
			for i := 0; i < n; i++ {
				statshouse.Metric("load_generator_value", k).Value(v)
			}
		}
		time.Sleep(time.Nanosecond)
	}
}
