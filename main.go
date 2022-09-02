package main

import (
	"awesomeProject5/db"
	"awesomeProject5/logic"
	"awesomeProject5/redisDB"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Working...")

	conn := db.InitDB()

	//conn.SaveUserData(logic.User{})

	users := map[int]*logic.User{}

	client := redisDB.InitRedis()

	go redisDB.StreamListener(client, users, conn)

	http.HandleFunc("/testHandle", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("USERS LIST: ", users[1010], '\n', users[1020])
	})

	fmt.Println("HANDLING!")
	err := http.ListenAndServe(":7981", nil)

	if err != nil {
		log.Fatalln("SOSI!@", err)
	}

}
