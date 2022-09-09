// 可以导出的type
package main

import (
	"fmt"
	"ultimate/language/exporting/example1/counters"
)

func main() {

	counter := counters.AlertCounter(1)
	fmt.Printf("exporting field %d\n", counter)
}
