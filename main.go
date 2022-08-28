package main

import (
	"awesomeProject5/redisDB"
	"fmt"
)

func main()  {

	fmt.Println("Working...")

	client := redisDB.InitClient()

	//client.StreamListner()

}