package logic

import (
	"time"
)

type Achieve struct {
	Id int
	Name string
	PeriodStart time.Time
	PeriodEnd time.Time
	OpenedAchieves []AchieveElem
}

type AchieveElem struct {
	NeedId int
	Duration time.Duration
	NeedCount int
}

type AchieveList map[int]Achieve

type UserAchieve struct {
	AchieveId int
	AchieveCount int
	LastActivation time.Time
}

type User struct {
	Id int
	Achieves map[int]UserAchieve
}

func (a *Achieve) CheckConditons(user User) bool {
	t := time.Now()

	if !(a.PeriodStart.Hour() < t.Hour() && a.PeriodEnd.Hour() > t.Hour()){
		return false
	}

	if a.OpenedAchieves == nil {
		return true
	}else{
		for _, achieve := range a.OpenedAchieves {
			uAch, ok := user.Achieves[achieve.NeedId]
			if !ok || uAch.AchieveCount != achieve.NeedCount || (time.Now().Sub(uAch.LastActivation)) > achieve.Duration {
				return false
			}
		}
	}

	return true
}

func (u *User) GetAllAchieves() []UserAchieve {
	var allAchieves []UserAchieve
	for _, achieve := range u.Achieves {
		allAchieves = append(allAchieves, achieve)
	}
	return allAchieves
}

func (u *User) AddAchieve(achieve Achieve)  {

	if uAch, ok := u.Achieves[achieve.Id]; achieve.CheckConditons(*u) && !ok{
		u.Achieves = map[int]UserAchieve{}
		u.Achieves[achieve.Id] = UserAchieve{
			AchieveId:      achieve.Id,
			AchieveCount:   1,
			LastActivation: time.Now(),
		}
	}else if achieve.CheckConditons(*u) && ok {
		uAch.AchieveCount++
		u.Achieves[achieve.Id] = UserAchieve{
			AchieveId:      achieve.Id,
			AchieveCount:   uAch.AchieveCount,
			LastActivation: time.Now(),
		}
	}

}