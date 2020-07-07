package main

import (
	"fmt"
	"time"

	pie "github.com/mrmiguu/Pies"
)

func app() {
	countptr := pie.IntPtr(0)

	fmt.Printf("app(): countptr %v\n", *countptr)

	count, setCount := pie.IntVar(1)
	thirdCount := count == 3

	pie.Do(func() {
		fmt.Println("app(): this will only show up once; it's a side effect with no dependencies")
	})

	pie.Do(func() {
		fmt.Println("app(): count", count)

		time.Sleep(time.Duration(count) * time.Second)

		if count == 4 {
			*countptr = 123
		}

		if count < 5 {
			setCount(count + 1)
		}
	}, count)

	pie.Do(func() {
		if !thirdCount {
			return
		}
		fmt.Println("app(): thirdCount := count == 3")
	}, thirdCount)

	postApp()
}

func postApp() {
	fmt.Println("postApp()")

	pie.Do(func() {
		fmt.Println("postApp(): this will only show up once; it's a side effect with no dependencies")
	})
}

func init() {
	// pie.Debug = true
}

func main() {
	pie.Mount(app)
}
