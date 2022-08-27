package logic

import (
	"sort"
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
	Cooldown		 time.Duration
	NeedAchieves     map[int]AchieveElem
	NeedLocations    []int
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

func (a *Achieve) checkConditions(usr *User, scanTime time.Time) bool {

	if scanTime.Sub(usr.CurrentAchieves[a.Id].LastScan) < (5 * time.Minute){
		usr.CurrentAchieves[a.Id].LastScan = scanTime
		return false
	}

	if a.NeedAchieves == nil && a.NeedLocations == nil && usr.haveAchieve(a.Id){
		return true
	} else if a.NeedAchieves == nil && a.NeedLocations == nil && !usr.haveAchieve(a.Id){
		uAch := convertToUserAchieve(*a)
		uAch.LastScan = scanTime
		uAch.ScanCount = 1
		usr.CurrentAchieves[a.Id] = &uAch

		return true
	}

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

	return true

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
		AchieveId:  ach.IdLoc,
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

	for _, achieve := range achieves {
		if isScanInInterval(achieve, scanTime) && achieve.checkConditions(u, scanTime) {
			uAch := u.TempAchieves[locId]
			uAch.ScanCount++
			u.TempAchieves[locId] = uAch
			if uAch.AchieveLvl == achieve.MaxLevel || achieve.ScansCountForLvl[uAch.AchieveLvl+1] < uAch.ScanCount {
				return
			}else if achieve.ScansCountForLvl[uAch.AchieveLvl+1] == uAch.ScanCount{
				uAch.Name = achieve.NameForLvl[uAch.AchieveLvl+1]
				uAch.AchieveLvl++
				u.TempAchieves[locId] = uAch
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

func (u *User) GetAllAchieves() []UserAchieve {
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
