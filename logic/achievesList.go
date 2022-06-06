package main

import "time"

var achList = AchieveList{1:Achieve{
	Id:               0,
	MaxLevel:         0,
	BeginLevel:       0,
	ScansCountForLvl: nil,
	NameForLvl:       nil,
	PeriodStart:      time.Time{},
	PeriodEnd:        time.Time{},
	NeedAchieves:     nil,
},
	2:Achieve{
		Id:               2,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl: map[int]string{1:"Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	3:Achieve{
		Id:               3,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1:3, 2:6, 3:9},
		NameForLvl: map[int]string{1:"Тестовая многоуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	4:Achieve{
		Id:               4,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1:3, 2:6, 3:9},
		NameForLvl: map[int]string{1:"Тестовая многоуровневая ачива СЛОЖНАЯ"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves: map[int]AchieveElem{3: {
			NeedId:    3,
			Duration:  0,
			NeedCount: 3,
		}},
	}}
