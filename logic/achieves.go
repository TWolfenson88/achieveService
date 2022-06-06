package main

import (
	"time"
)

type Achieve struct {
	Id int
	MaxLevel int
	BeginLevel int
	ScansCountForLvl map[int]int
	NameForLvl map[int]string
	PeriodStart time.Time
	PeriodEnd time.Time
	NeedAchieves map[int]AchieveElem
}

type AchieveElem struct {
	NeedId int
	Duration time.Duration
	NeedCount int
}

type AchieveList map[int]Achieve

type UserAchieve struct {
	AchieveId int
	AchieveLvl int
	ScanCount int
	Name string
	LastScan time.Time
}

type User struct {
	Id int
	UsrLvl int
	Achieves map[int]UserAchieve
}

func (a Achieve) checkConditions(usrAchs map[int]UserAchieve) bool {
	
	if a.NeedAchieves == nil {
		return true
	}

/*	for _, uAch := range usrAchs {
		fmt.Println(uAch.AchieveId)
		if ach, ok := a.NeedAchieves[uAch.AchieveId]; !ok {
			fmt.Println("need ach ne ok")
			return false
		}else if (time.Now().Sub(uAch.LastScan) > ach.Duration && ach.Duration > 0) || (uAch.ScanCount) < ach.NeedCount{
			fmt.Println("this ne ok")
			return false
		}
	}
	return true*/

	for _, elem := range a.NeedAchieves {
		if uAch, ok := usrAchs[elem.NeedId]; !ok {
			return false
		}else if (time.Now().Sub(uAch.LastScan) > elem.Duration && elem.Duration != 0) || elem.NeedCount > uAch.ScanCount {
			return false
		}
	}
return true
}

func (al AchieveList) convertToUserAchieve(achId int) UserAchieve {
	ach := al[achId]
	return UserAchieve{
		AchieveId:  ach.Id,
		AchieveLvl: ach.BeginLevel,
		ScanCount:  1,
		Name:       ach.NameForLvl[ach.BeginLevel],
		LastScan:   time.Time{},
	}
}

func (u User) haveAchieve(achId int) bool {
	_, ok := u.Achieves[achId]
	return ok
}

func (u User) AddAchieve(scanTime time.Time, achId, usrId int) {
	if !u.haveAchieve(achId){
		uAch := achList.convertToUserAchieve(achId)
		uAch.LastScan = scanTime
		uAch.ScanCount = 1
		u.Achieves[achId] = uAch
	}else {
		uAch := u.Achieves[achId]
		ach := achList[achId]
		uAch.ScanCount++
		u.Achieves[achId] = uAch
		//fmt.Println("here here ", ach.checkConditions(u.Achieves))
		if uAch.AchieveLvl == ach.MaxLevel || ach.ScansCountForLvl[uAch.AchieveLvl+1] < uAch.ScanCount {
			//fmt.Println("here here ")
			return
		}else if ach.checkConditions(u.Achieves) && ach.ScansCountForLvl[uAch.AchieveLvl+1] == uAch.ScanCount{
			uAch.Name = ach.NameForLvl[uAch.AchieveLvl+1]
			uAch.AchieveLvl++
			u.Achieves[achId] = uAch
		}
	}
}
