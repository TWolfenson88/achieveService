package logic

import (
	"fmt"
	"time"
)

type Achieve struct {
	Id               int
	IdLoc            int
	MaxLevel         int
	BeginLevel       int
	ScansCountForLvl map[int]int
	NameForLvl       map[int]string
	PeriodStart      time.Time
	PeriodEnd        time.Time
	Cooldown         time.Duration
	NeedAchieves     map[int]AchieveElem
	NeedLocations    []int
	SpecialLogic     func(usr *User, ach *Achieve) bool
}

type AchieveElem struct {
	NeedId    int
	Duration  time.Duration
	NeedCount int
}

type AchieveList map[int][]Achieve

type UserAchieve struct {
	AchieveId        int
	AchieveLvl       int
	MaxLvl           int
	ScanCount        int
	Name             string
	LastScan         time.Time
	ScannedLocations []int
}

type User struct {
	Id              int
	UsrLvl          int
	TempAchieves    map[int]*UserAchieve
	CurrentAchieves map[int]*UserAchieve
}

func checkCooldown(lastScan, scanTime time.Time) bool {
	return scanTime.Sub(lastScan) < (5 * time.Minute)
}

func (a *Achieve) checkConditions(usr *User, scanTime time.Time) bool {

	fmt.Println("rere", usr)

	ach, ok := usr.CurrentAchieves[a.Id]
	if !ok {
		ach = &UserAchieve{
			AchieveId:        0,
			AchieveLvl:       0,
			MaxLvl:           0,
			ScanCount:        0,
			Name:             "",
			LastScan:         scanTime,
			ScannedLocations: nil,
		}
	}

	fmt.Println("фчива^ \n", a)

	if a.SpecialLogic == nil && a.NeedAchieves == nil && checkCooldown(ach.LastScan, scanTime) {
		fmt.Println("очень просто ")
		return true
	} else if a.SpecialLogic == nil && a.NeedAchieves != nil && checkCooldown(ach.LastScan, scanTime) {
		fmt.Println("чуть сложнее")
		for _, elem := range a.NeedAchieves {
			if uAch, ok := usr.CurrentAchieves[elem.NeedId]; !ok {
				return false
			} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
				return false
			} else {
				fmt.Println("чуть сложнее2")
				return true
			}
		}
	} else if a.SpecialLogic != nil && a.NeedAchieves == nil && checkCooldown(ach.LastScan, scanTime) {
		fmt.Println("спешщл логик")
		return a.SpecialLogic(usr, a)
	} else if a.SpecialLogic != nil && a.NeedAchieves != nil && checkCooldown(ach.LastScan, scanTime) {
		fmt.Println("довольно сложно")
		for _, elem := range a.NeedAchieves {
			if uAch, ok := usr.CurrentAchieves[elem.NeedId]; !ok {
				return false
			} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
				return false
			} else {
				fmt.Println("спешщл логик")
				return a.SpecialLogic(usr, a)
			}
		}
	}

	return false

}

func convertToUserAchieve(ach Achieve) UserAchieve {
	return UserAchieve{
		AchieveId:  ach.Id,
		AchieveLvl: ach.BeginLevel,
		MaxLvl:     ach.MaxLevel,
		ScanCount:  1,
		Name:       ach.NameForLvl[ach.BeginLevel],
		LastScan:   time.Time{},
	}
}

func (u *User) haveAchieve(achId int) bool {
	_, ok := u.CurrentAchieves[achId]
	return ok
}

func (u *User) haveAchieves() bool {
	if len(u.CurrentAchieves) != 0 {
		return true
	}

	return false
}

func isScanInInterval(a Achieve, t time.Time) bool {
	if a.PeriodEnd.IsZero() && a.PeriodStart.IsZero() {
		return true
	}

	hr, min, _ := t.Clock()
	tt := time.Time{}.Add(time.Duration(hr)*time.Hour + time.Duration(min)*time.Minute)

	if a.PeriodStart.Before(tt) && a.PeriodEnd.After(tt) {
		return true
	}

	return false
}

func (u *User) AddAchieve(scanTime time.Time, locId int, logCh chan string) {

	//Получаем и фильтруем все ачивки по локации и по времени скана
	achieves := achList[locId]

	fmt.Println("ach arr len ", len(achieves))

	for _, achieve := range achieves {
		fmt.Println("re", achieve.Id)
		if isScanInInterval(achieve, scanTime) && achieve.checkConditions(u, scanTime) {
			fmt.Println("KEKE")
			uAch := convertToUserAchieve(achieve)

			tempUsrAch, ok := u.TempAchieves[uAch.AchieveId]
			if ok && tempUsrAch.ScanCount+1 == achieve.ScansCountForLvl[tempUsrAch.AchieveLvl+1] {
				fmt.Println("HERE ")
				tempUsrAch.Name = achieve.NameForLvl[tempUsrAch.AchieveLvl+1]
				tempUsrAch.ScanCount++
				tempUsrAch.AchieveLvl++
				u.CurrentAchieves[tempUsrAch.AchieveId] = tempUsrAch

				logCh <- fmt.Sprintf("%d получил ачивку %s", u.Id, tempUsrAch.Name)

			} else if ok && achieve.ScansCountForLvl != nil {
				fmt.Println("TEMP PLUSPLUS ")
				tempUsrAch.ScanCount++
			} else {
				fmt.Println("ELSE")
				if uAch.AchieveLvl > 0 {
					fmt.Println("CURRENT ")
					fmt.Println("ACHIEVE: ", uAch, "\n", "USER: ", u.CurrentAchieves)
					u.CurrentAchieves[uAch.AchieveId] = &uAch

					logCh <- fmt.Sprintf("%d получил ачивку %s", u.Id, uAch.Name)

				} else {
					u.TempAchieves[uAch.AchieveId] = &uAch
				}
			}

		}
	}

}

func (u *User) RemoveAchieve(achId int) {
	delete(u.TempAchieves, achId)
}
