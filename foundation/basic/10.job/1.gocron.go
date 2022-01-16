package main

import (
	"fmt"
	"github.com/claudiu/gocron"
)

func main() {
	//gocron.Every(1).Day().Do(Job1)
	//gocron.Every(7).Days().Do(Job2)

	gocron.Every(10).Second().Do(Job1)
	//gocron.Every(5).Second().Do(Job2)
	gocron.Start()
}

func Job1()  {
	fmt.Println("@@@Job1 is running...")
}

func Job2()  {
	fmt.Println("@@@Job2 is running...")
}