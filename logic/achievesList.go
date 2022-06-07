package main

import "time"

var achList = AchieveList{0: Achieve{
	Id:               0,
	MaxLevel:         0,
	BeginLevel:       0,
	ScansCountForLvl: nil,
	NameForLvl:       nil,
	PeriodStart:      time.Time{},
	PeriodEnd:        time.Time{},
	NeedAchieves:     nil,
},
	2: Achieve{
		Id:               2,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	3: Achieve{
		Id:               3,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 3, 2: 6, 3: 9},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	4: Achieve{
		Id:               4,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 3, 2: 6, 3: 9},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая ачива СЛОЖНАЯ"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves: map[int]AchieveElem{3: {
			NeedId:    3,
			Duration:  0,
			NeedCount: 3,
		}},
	},
	5: Achieve{
		Id:               5,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: map[int]int{1: 1},
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая простая ачива с временем"},
		PeriodStart:       time.Time{}.Add(10*time.Hour + 10*time.Minute), // from 10:10 AM
		PeriodEnd:         time.Time{}.Add(20 * time.Hour),                   // to 8:00 PM
		NeedAchieves:     nil,
	},
	6: Achieve{
		Id:               6,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 2, 2: 4, 3: 6},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая сложная ачива полный сука фарш", 2:"промежуточный фарш", 3:"септолете тотал бля"},
		PeriodStart:       time.Time{}.Add(10*time.Hour + 10*time.Minute), // from 10:10 AM
		PeriodEnd:         time.Time{}.Add(20 * time.Hour),                   // to 8:00 PM
		NeedAchieves: map[int]AchieveElem{3: {
			NeedId:    3,
			Duration:  0,
			NeedCount: 3,
		}},
	},
}
