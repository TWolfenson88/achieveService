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
	//fmt.Println("CHECK COOLDWON", scanTime)
	//fmt.Println("CHECK COOLDWON", lastScan)
	//fmt.Println("CHECK COOLDWON", scanTime.Sub(lastScan))
	return scanTime.Sub(lastScan) < (5 * time.Minute)
}

func (a *Achieve) checkConditions(usr *User, scanTime time.Time) bool {

	fmt.Println("rere", usr)

	ach, ok := usr.CurrentAchieves[a.Id]
	if !ok {
		ach = &UserAchieve{
			AchieveId:        0,
			AchieveLvl:       0,
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

	/*	if a.NeedAchieves != nil {

			for _, elem := range a.NeedAchieves {
				if uAch, ok := usr.CurrentAchieves[elem.NeedId]; !ok {
					return false
				} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
					return false
				}
			}
		}

		if a.SpecialLogic != nil {
			return a.SpecialLogic(usr)
		}*/

	/*	if usr.haveAchieves() && scanTime.Sub(usr.CurrentAchieves[a.Id].LastScan) < (5 * time.Minute){
			usr.CurrentAchieves[a.Id].LastScan = scanTime
			return false
		}

		fmt.Println("rere")

		if a.NeedAchieves == nil && a.NeedLocations == nil && usr.haveAchieve(a.Id){
			return true
		} else if a.NeedAchieves == nil && a.NeedLocations == nil && !usr.haveAchieve(a.Id){
			uAch := convertToUserAchieve(*a)
			uAch.LastScan = scanTime
			uAch.ScanCount = 1
			usr.CurrentAchieves[a.Id] = &uAch

			return true
		}

		fmt.Println("rere")

		if a.NeedLocations == nil && a.NeedAchieves != nil {
			for _, elem := range a.NeedAchieves {
				if uAch, ok := usr.CurrentAchieves[elem.NeedId]; !ok {
					return false
				} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
					return false
				}
			}

			return true

		}

		fmt.Println("rere")

		if a.NeedLocations != nil {
			tempAch, ok := usr.TempAchieves[a.Id]
			if !ok {
				uAch := convertToUserAchieve(*a)
				uAch.LastScan = scanTime
				uAch.ScanCount = 1
				uAch.ScannedLocations = append(uAch.ScannedLocations, a.IdLoc)
				usr.TempAchieves[a.Id] = &uAch
				return true
			} else if tempAch.ScannedLocations[len(tempAch.ScannedLocations)-1] != a.NeedLocations[len(tempAch.ScannedLocations)-1]{
				delete(usr.TempAchieves, a.Id)
				return false
			}
		}

		fmt.Println("rere")*/

	return false

	/*
		if a.NeedAchieves == nil {
			return true
		}

		for _, elem := range a.NeedAchieves {
			if uAch, ok := usrAchs[elem.NeedId]; !ok {
				return false
			} else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
				return false
			}
		}
		return true
	*/
}

func convertToUserAchieve(ach Achieve) UserAchieve {
	return UserAchieve{
		AchieveId:  ach.Id,
		AchieveLvl: ach.BeginLevel,
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

func (u *User) AddAchieve(scanTime time.Time, locId int) {

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
			} else if ok && achieve.ScansCountForLvl != nil {
				fmt.Println("TEMP PLUSPLUS ")
				tempUsrAch.ScanCount++
			} else {
				fmt.Println("ELSE")
				if uAch.AchieveLvl > 0 {
					fmt.Println("CURRENT ")
					fmt.Println("ACHIEVE: ", uAch, "\n", "USER: ", u.CurrentAchieves)
					u.CurrentAchieves[uAch.AchieveId] = &uAch
				} else {
					u.TempAchieves[uAch.AchieveId] = &uAch
				}
			}

		}
	}

	/*
		if !u.haveAchieve(locId) && isScanInInterval(achList[locId], scanTime){
			uAch := achList.convertToUserAchieve(locId)
			uAch.LastScan = scanTime
			uAch.ScanCount = 1
			u.TempAchieves[locId] = uAch
		}else if isScanInInterval(achList[locId], scanTime) {
			uAch := u.TempAchieves[locId]
			ach := achList[locId]
			uAch.ScanCount++
			u.TempAchieves[locId] = uAch
			if uAch.AchieveLvl == ach.MaxLevel || ach.ScansCountForLvl[uAch.AchieveLvl+1] < uAch.ScanCount {
				return
			}else if ach.checkConditions(u.TempAchieves) && ach.ScansCountForLvl[uAch.AchieveLvl+1] == uAch.ScanCount{
				uAch.Name = ach.NameForLvl[uAch.AchieveLvl+1]
				uAch.AchieveLvl++
				u.TempAchieves[locId] = uAch
			}
		}
	*/
}

func (u *User) RemoveAchieve(achId int) {
	delete(u.TempAchieves, achId)
}

/*func (u *User) GetAllAchieves() []UserAchieve {
	var result []UserAchieve

	for _, achieve := range u.TempAchieves {
		result = append(result, achieve)
	}

	return result
}

func GetAllLastAchieves(users []User, n int) []UserAchieve {
	var result []UserAchieve

	for _, user := range users {
		for _, achieve := range user.TempAchieves {
			result = append(result, achieve)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].LastScan.Before(result[j].LastScan)
	})

	return result[:n]
}
*/
