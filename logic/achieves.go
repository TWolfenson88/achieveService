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
	SpecialLogic     func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool
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

func (u *UserAchieve) String() string {
	return fmt.Sprintf("%#v \n", u)
}

type User struct {
	Id              int
	UsrLvl          int
	TempAchieves    map[int]*UserAchieve
	CurrentAchieves map[int]*UserAchieve
}

func checkCooldown(lastScan, scanTime time.Time) bool {
	if lastScan.IsZero() {
		return true
	}
	return scanTime.Sub(lastScan) < (5 * time.Minute)
}

func (a *Achieve) checkConditions(usr *User, scanTime time.Time, locId int, logCh chan string) bool {

	fmt.Println("User checking conditions: ", usr)

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

	fmt.Println("Current achieves: ^ \n", a)

	if !checkCooldown(ach.LastScan, scanTime) {
		fmt.Println("Soo less cooldown, ", ach.LastScan, scanTime)
		return false
	}

	if a.SpecialLogic == nil && a.NeedAchieves == nil {
		fmt.Println("Added simple achievement ")
		return true
	} else if a.SpecialLogic == nil && a.NeedAchieves != nil {
		fmt.Println("Achieve with dependances")
		for _, elem := range a.NeedAchieves {
			if uAch, ok := usr.CurrentAchieves[elem.NeedId]; !ok {
				return false
			} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
				return false
			} else {
				fmt.Println("Achieve with dependances and counts")
				return true
			}
		}
	} else if a.SpecialLogic != nil && a.NeedAchieves == nil {
		fmt.Println("Achieve with special logic")
		return a.SpecialLogic(usr, a, locId, scanTime, logCh)
	} else if a.SpecialLogic != nil && a.NeedAchieves != nil {
		fmt.Println("Achieve with special logic and dependancies")
		for _, elem := range a.NeedAchieves {
			if uAch, ok := usr.CurrentAchieves[elem.NeedId]; !ok {
				return false
			} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
				return false
			} else {
				fmt.Println("Checking special logic")
				return a.SpecialLogic(usr, a, locId, scanTime, logCh)
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

	achieves = append(achieves, achList[0]...) // вот тут к проверяемому массиву аппенжу все общие ачивы

	fmt.Println("Achieves number to check:  ", len(achieves))

	for _, achieve := range achieves {
		fmt.Println("Checking achieve ", achieve.Id)
		if isScanInInterval(achieve, scanTime) && achieve.checkConditions(u, scanTime, locId, logCh) {
			fmt.Println("KEKE")
			uAch := convertToUserAchieve(achieve)

			tempUsrAch, ok := u.TempAchieves[uAch.AchieveId]
			if ok && tempUsrAch.ScanCount+1 == achieve.ScansCountForLvl[tempUsrAch.AchieveLvl+1] {
				fmt.Println("HERE ")
				tempUsrAch.Name = achieve.NameForLvl[tempUsrAch.AchieveLvl+1]
				tempUsrAch.ScanCount++
				tempUsrAch.AchieveLvl++

				chac, okok := u.CurrentAchieves[uAch.AchieveId]

				if okok && chac.AchieveLvl == tempUsrAch.AchieveLvl {
					continue
				}

				u.CurrentAchieves[tempUsrAch.AchieveId] = tempUsrAch

				logCh <- fmt.Sprintf("%d получил ачивку %s уровня %d", u.Id, tempUsrAch.Name, tempUsrAch.AchieveLvl)

			} else if ok && achieve.ScansCountForLvl != nil {
				fmt.Println("Increasing achieve scan count ")
				tempUsrAch.ScanCount++
			} else {
				fmt.Println("ELSE")

				_, okok := u.CurrentAchieves[uAch.AchieveId]

				if uAch.AchieveLvl > 0 {
					fmt.Println("CURRENT ")
					fmt.Println("ACHIEVE: ", uAch, "\n", "USER: ", u.CurrentAchieves)

					if !okok {
						u.CurrentAchieves[uAch.AchieveId] = &uAch

						logCh <- fmt.Sprintf("%d получил ачивку %s", u.Id, uAch.Name)
					}

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
