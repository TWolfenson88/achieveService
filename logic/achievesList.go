package logic

import (
	"fmt"
	"time"
)

func recoverFromNil() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func someTestLogic(usr *User, ach *Achieve) bool {
	defer recoverFromNil()
	fmt.Println("ALLO")
	if uAch, ok := usr.TempAchieves[ach.Id]; ok {
		for i, id := range ach.NeedLocations {
			fmt.Println("id: ", id, "  |||  another id: ", uAch.ScannedLocations[i])
			if id != uAch.ScannedLocations[i] {
				return false
			}
		}
		fmt.Println("***** RETURNED TRUE ")
		return true
	} else {
		fmt.Println("WROTTEN")
		uAch := convertToUserAchieve(*ach)
		uAch.ScannedLocations = append(uAch.ScannedLocations, ach.IdLoc)
		usr.TempAchieves[uAch.AchieveId] = &uAch
	}

	return false
}

var achList = AchieveList{0: []Achieve{ //тутт общие ачивый
	{
		Id:               0,
		IdLoc:            0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		Cooldown:         0,
		NeedAchieves:     nil,
		NeedLocations:    nil,
		SpecialLogic:     nil,
	},
}, 10: []Achieve{{ //Массив ачив для 10й локации
	Id:               1,
	IdLoc:            10,
	MaxLevel:         1,
	BeginLevel:       1,
	ScansCountForLvl: nil,
	NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
	PeriodStart:      time.Time{},
	PeriodEnd:        time.Time{},
	Cooldown:         0,
	NeedAchieves:     nil,
	NeedLocations:    nil,
},
	{
		Id:               2,
		IdLoc:            10,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива на несколько локаций"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		Cooldown:         0,
		NeedAchieves:     nil,
		NeedLocations:    []int{10, 20, 30},
		SpecialLogic:     someTestLogic,
	},
	{
		Id:               3,
		IdLoc:            10,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая ачива с спецусловием"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		Cooldown:         0,
		NeedAchieves:     nil,
		NeedLocations:    nil,
		SpecialLogic: func(usr *User, ach *Achieve) bool {
			if len(usr.CurrentAchieves) == 0 {
				return true
			}
			return false
		},
	},
},
	11: []Achieve{
		{
			Id:               32,
			IdLoc:            11,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: map[int]int{1: 2, 2: 4},
			NameForLvl:       nil,
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
	},
	20: []Achieve{
		{
			Id:               33,
			IdLoc:            20,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       nil,
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve) bool {
				fmt.Println("chck20 ")
				if achach, ok := usr.TempAchieves[2]; ok {
					fmt.Println("chck20 ")
					achach.ScannedLocations = append(achach.ScannedLocations, 20)
				}
				return true
			},
		},
	},
	30: []Achieve{
		{
			Id:               34,
			IdLoc:            30,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       nil,
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve) bool {
				fmt.Println("chck30 ")
				if achach, ok := usr.TempAchieves[2]; ok {
					fmt.Println("chck30 ")
					achach.ScannedLocations = append(achach.ScannedLocations, 30)
				}
				return true
			},
		},
	},
}

/*
var achList = AchieveList{10: Achieve{
	IdLoc:            10,
	MaxLevel:         1,
	BeginLevel:       1,
	ScansCountForLvl: nil,
	NameForLvl:       nil,
	PeriodStart:      time.Time{},
	PeriodEnd:        time.Time{},
	NeedAchieves:     nil,
},
	11: Achieve{
		IdLoc:            11,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	12: Achieve{
		IdLoc:            12,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	13: Achieve{
		IdLoc:            13,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	14: Achieve{
		IdLoc:            14,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	15: Achieve{
		IdLoc:            15,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	16: Achieve{
		IdLoc:            16,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	17: Achieve{
		IdLoc:            17,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	18: Achieve{
		IdLoc:            18,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	19: Achieve{
		IdLoc:            19,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	20: Achieve{
		IdLoc:            20,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	0: Achieve{
		IdLoc:            0,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	2: Achieve{
		IdLoc:            2,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	3: Achieve{
		IdLoc:            3,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 3, 2: 6, 3: 9},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	4: Achieve{
		IdLoc:            4,
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
		IdLoc:            5,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: map[int]int{1: 1},
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая простая ачива с временем"},
		PeriodStart:      time.Time{}.Add(10*time.Hour + 10*time.Minute), // from 10:10 AM
		PeriodEnd:        time.Time{}.Add(20 * time.Hour),                   // to 8:00 PM
		NeedAchieves:     nil,
	},
	6: Achieve{
		IdLoc:            6,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 2, 2: 4, 3: 6},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая сложная ачива полный сука фарш", 2:"промежуточный фарш", 3:"септолете тотал бля"},
		PeriodStart:      time.Time{}.Add(10*time.Hour + 10*time.Minute), // from 10:10 AM
		PeriodEnd:        time.Time{}.Add(20 * time.Hour),                   // to 8:00 PM
		NeedAchieves: map[int]AchieveElem{3: {
			NeedId:    3,
			Duration:  0,
			NeedCount: 3,
		}},
	},
}
*/
