package main

import (
	"awesomeProject5/db"
	"awesomeProject5/logic"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func Test_handleAchieveInfo(t *testing.T) {
	t.Run("Checking last minute functionality ", func(t *testing.T) {

		user := logic.User{
			Id:           1,
			UsrLvl:       1,
			TempAchieves: nil,
			CurrentAchieves: map[int]*logic.UserAchieve{1: {
				AchieveId:        1,
				AchieveLvl:       1,
				MaxLvl:           1,
				ScanCount:        1,
				Name:             "one ",
				LastScan:         time.Now().Add(-10 * time.Second),
				ScannedLocations: nil,
			},
				2: {
					AchieveId:        2,
					AchieveLvl:       2,
					MaxLvl:           3,
					ScanCount:        22,
					Name:             "two ",
					LastScan:         time.Now(),
					ScannedLocations: nil,
				},
				3: {
					AchieveId:        3,
					AchieveLvl:       2,
					MaxLvl:           2,
					ScanCount:        21,
					Name:             "three ",
					LastScan:         time.Now().Add(-2 * time.Minute),
					ScannedLocations: nil,
				}},
		}

		users := Users{}
		users[1] = &user

		requset, err := http.NewRequest("GET", "/getUserInfo", nil)
		if err != nil {
			fmt.Println("request creation error : ", err)
		}

		q := requset.URL.Query()
		q.Add("user", "1")
		q.Add("secret", "PosholNahuySuka")
		requset.URL.RawQuery = q.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(users.handleAchieveInfo)

		handler.ServeHTTP(rr, requset)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		fmt.Println("response body", rr.Body.String())

	})

	t.Run("test getting acieve from BiDe", func(t *testing.T) {

		//перед запуском теста положить в биде тестового юзера с айди 10 и пару-тройку ачивок

		conn := db.InitDB()

		logCh := make(chan string, 5)
		go SendLogs(logCh)

		users := Users{}
		users = conn.InitUserData()

		requset, err := http.NewRequest("GET", "/getUserInfo", nil)
		if err != nil {
			fmt.Println("request creation error : ", err)
		}

		q := requset.URL.Query()
		q.Add("user", "10")
		q.Add("secret", "PosholNahuySuka")
		requset.URL.RawQuery = q.Encode()

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(users.handleAchieveInfo)

		handler.ServeHTTP(rr, requset)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		fmt.Println("response body", rr.Body.String())

	})
}
