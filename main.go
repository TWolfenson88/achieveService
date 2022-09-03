package main

import (
	"awesomeProject5/db"
	"awesomeProject5/logic"
	"awesomeProject5/redisDB"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Users map[int]*logic.User

type HandlerAchieves struct {
	AchieveName  string `json:"achieve_name"`
	CurrentLevel int    `json:"current_level"`
	MaxLevel     int    `json:"max_level"`
}

type UserInfo struct {
	UserLevel          int               `json:"user_level"`
	LastMinuteAchieves []string          `json:"last_minute_achieves"`
	CurrentAchieves    []HandlerAchieves `json:"current_achieves"`
	TopUsers           []int             `json:"top_users"`
}

func (u Users) findTopUsers() []int {
	result := []int{}
	currentCount := 0

	for _, user := range u {
		if len(user.CurrentAchieves) > currentCount {
			currentCount = len(user.CurrentAchieves)
			result = append(result, user.Id)
		}
	}

	return result[len(result)-5:]
}

func findLastAchieves(user *logic.User) []string {
	result := []string{}

	for _, achieve := range user.CurrentAchieves {
		if time.Now().Sub(achieve.LastScan) < 1*time.Minute {
			result = append(result, achieve.Name)
		}
	}

	return result
}

func (u Users) handleAchieveInfo(w http.ResponseWriter, r *http.Request) {
	usrIdS := r.URL.Query().Get("user")
	usrId, _ := strconv.Atoi(usrIdS)

	secret := r.URL.Query().Get("secret")

	if secret != "PosholNahuySuka" {
		return
	}

	user := u[usrId]
	hndlAchs := []HandlerAchieves{}

	for _, achieve := range user.CurrentAchieves {
		hAch := HandlerAchieves{
			AchieveName:  achieve.Name,
			CurrentLevel: achieve.AchieveLvl,
			MaxLevel:     achieve.MaxLvl,
		}
		hndlAchs = append(hndlAchs, hAch)
	}

	usrInfo := &UserInfo{
		UserLevel:          user.UsrLvl,
		LastMinuteAchieves: findLastAchieves(user),
		CurrentAchieves:    hndlAchs,
		TopUsers:           u.findTopUsers(),
	}

	marshallUserInfo, err := json.Marshal(usrInfo)
	if err != nil {
		log.Println("MARSHALL ERR : ", err)
	}

	_, err = w.Write(marshallUserInfo)
	if err != nil {
		log.Println("WRITE ERR : ", err)
	}

}

func SendLogs(ch chan string) {
	for {
		select {
		case rslt := <-ch:
			fmt.Println("LOGGER WORKING", rslt)

			client := &http.Client{}

			req, err := http.NewRequest("GET", "http://app:8000/achievements/log_message", nil)
			if err != nil {
				fmt.Println(")))")
			}

			req.URL.Query().Add("log_message", rslt)
			req.URL.Query().Add("secret", "PosholNahuySuka")

			_, err = client.Do(req)
			if err != nil {
				fmt.Println(")))")
			}

		}
	}
}

func main() {

	fmt.Println("Working...")

	conn := db.InitDB()

	logCh := make(chan string, 5)
	go SendLogs(logCh)

	users := Users{}

	//users := map[int]*logic.User{}

	client := redisDB.InitRedis()

	go redisDB.StreamListener(client, users, conn, logCh)

	http.HandleFunc("/getUserInfo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("USERS LIST: ", users[1010], '\n', users[1020])
	})

	fmt.Println("HANDLING!")
	err := http.ListenAndServe(":7981", nil)

	if err != nil {
		log.Fatalln("SOSI!@", err)
	}

}
