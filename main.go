package main

import (
	"fmt"
	"time"
)

func main()  {


	usr := User{
		Id:       1,
	}

	uAch1 := UserAchieve{
		AchieveId:      1,
		AchieveCount:   1,
		LastActivation: time.Now(),
	}

	uAch2 := UserAchieve{
		AchieveId:      2,
		LastActivation: time.Now(),
	}

	usr.AddAchieve(uAch1)
	usr.AddAchieve(uAch1)

	fmt.Println(usr)

	usr.AddAchieve(uAch2)
	usr.AddAchieve(uAch2)
	fmt.Println(usr)






	fmt.Println("Working...")

	aList := AchieveList{10: []AchieveElem{{
		NeedId:    23,
		Duration:  15*time.Minute,
		NeedCount: 2,
	},
		{
			NeedId:    24,
			Duration:  10*time.Minute,
			NeedCount: 1,
		}}}

	fmt.Println("achieve list: ", aList[10])

	ach := Achieve{
		Id:          1,
		Name:        "Ебать ты тестер-тесто-тостер-лоукостер",
		PeriodStart: time.Time{},
		PeriodEnd:   time.Time{}.Add(20*time.Hour),
		NeedAchieve: 0,
	}

	fmt.Println("ACHIVE ONE PERIODS: ", ach.PeriodStart.Hour(), ach.PeriodEnd.Hour())
}