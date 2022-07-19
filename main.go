package main

import (
	corn2 "MyCorn/corn"
	"time"
)

func main() {

	corn := corn2.NewCorn()

	corn.Start()

	corn.AddFunc(time.Now().Add(time.Hour),
		func() {
			print("123")
		})

	corn.Reset()

}
