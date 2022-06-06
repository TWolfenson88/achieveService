package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var checkCondTests = []struct{
	achieve Achieve
	usrAchieves map[int]UserAchieve
	exp         bool
	msg         string
}{
	{achieve: Achieve{
		Id:               0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	}, usrAchieves: map[int]UserAchieve{1:{
		AchieveId:  1,
		AchieveLvl: 1,
		ScanCount:  1,
		Name:       "test",
		LastScan:   time.Time{},
	}}, exp: true, msg: "empty needed achieves"},
	{achieve: Achieve{
		Id:               0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves: map[int]AchieveElem{1: {
			NeedId:    1,
			Duration:  15*time.Minute,
			NeedCount: 5,
		}},
	}, usrAchieves: map[int]UserAchieve{1:{
		AchieveId:  1,
		AchieveLvl: 1,
		ScanCount:  5,
		Name:       "test",
		LastScan:   time.Now().Add(-10*time.Minute),
	}}, exp: true, msg: "all conditions ok"},
	{achieve: Achieve{
		Id:               0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves: map[int]AchieveElem{1: {
			NeedId:    1,
			Duration:  0,
			NeedCount: 5,
		}},
	}, usrAchieves: map[int]UserAchieve{1:{
		AchieveId:  1,
		AchieveLvl: 1,
		ScanCount:  5,
		Name:       "test",
		LastScan:   time.Now().Add(-10*time.Minute),
	}}, exp: true, msg: "all conditions ok and duration zero"},
	{achieve: Achieve{
		Id:               0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     map[int]AchieveElem{1: {
			NeedId:    1,
			Duration:  15*time.Minute,
			NeedCount: 5,
		}},
	}, usrAchieves: map[int]UserAchieve{1:{
		AchieveId:  1,
		AchieveLvl: 1,
		ScanCount:  5,
		Name:       "test",
		LastScan:   time.Now().Add(-20*time.Minute),
	}}, exp: false, msg: "too late scan"},
	{achieve: Achieve{
		Id:               0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     map[int]AchieveElem{1: {
			NeedId:    1,
			Duration:  15*time.Minute,
			NeedCount: 5,
		}},
	}, usrAchieves: map[int]UserAchieve{1:{
		AchieveId:  1,
		AchieveLvl: 1,
		ScanCount:  3,
		Name:       "test",
		LastScan:   time.Now().Add(-10*time.Minute),
	}}, exp: false, msg: "not enough scan count"},
	{achieve: Achieve{
		Id:               0,
		MaxLevel:         0,
		BeginLevel:       0,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     map[int]AchieveElem{1: {
			NeedId:    1,
			Duration:  15*time.Minute,
			NeedCount: 5,
		}},
	}, usrAchieves: map[int]UserAchieve{12:{
		AchieveId:  12,
		AchieveLvl: 1,
		ScanCount:  5,
		Name:       "test",
		LastScan:   time.Now().Add(-10*time.Minute),
	}}, exp: false, msg: "user have not needed achieve"},
}

func TestCheckConditions(t *testing.T){
	for _, test := range checkCondTests {
		result := test.achieve.checkConditions(test.usrAchieves)
		require.Equal(t, test.exp, result, test.msg)
	}

}

func TestAddAchieve(t *testing.T) {
	t.Run("add one simple achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 2, 1)

		req := achList.convertToUserAchieve(2)
		req.LastScan = time.Now()

		require.Equal(t, req, user.Achieves[2])
	})
	t.Run("add second level for multilevel achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)

		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)

		req := 2

		require.Equal(t, req, user.Achieves[3].AchieveLvl)
	})
	t.Run("add max level for multilevel achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)

		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)

		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)

		req := 3

		require.Equal(t, req, user.Achieves[3].AchieveLvl)
	})
	t.Run("add first level for multilevel complex achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)
		user.AddAchieve(time.Now(), 3, 1)

		user.AddAchieve(time.Now(), 4, 1)
		user.AddAchieve(time.Now(), 4, 1)
		user.AddAchieve(time.Now(), 4, 1)

		req := 1

		require.Equal(t, req, user.Achieves[4].AchieveLvl)
	})
}
