package main

import "time"

type Achieve struct {
	Id int
	Name string
	PeriodStart time.Time
	PeriodEnd time.Time
	NeedAchieve int
}

type AchieveElem struct {
	NeedId int
	Duration time.Duration
	NeedCount int
}

type AchieveList map[int][]AchieveElem

type UserAchieve struct {
	AchieveId int
	AchieveCount int
	LastActivation time.Time
}

type User struct {
	Id int
	Achieves []UserAchieve
}

func (u *User) GetAllAchieves() []UserAchieve {
	return u.Achieves
}

func (u *User) AddAchieve(achieve UserAchieve)  {

	for i, userAchieve := range u.Achieves {
		if userAchieve.AchieveId == achieve.AchieveId {
			u.Achieves[i].AchieveCount++
			u.Achieves[i].LastActivation = achieve.LastActivation
			return
		}
	}

	u.Achieves = append(u.Achieves, achieve)
}