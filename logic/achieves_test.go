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

func TestIsScanInInterval(t *testing.T) {

	t.Run("zero time diff", func(t *testing.T) {
		tt := time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local) //not in achieve diapason
		result := isScanInInterval(achList[0], tt)

		require.True(t, result)
	})
	t.Run("not in interval", func(t *testing.T) {
		tt := time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local) //not in achieve diapason
		result := isScanInInterval(achList[5], tt)

		require.False(t, result)
	})
	t.Run("not in minute interval", func(t *testing.T) {
			tt := time.Date(2022, time.June, 7, 10, 0, 0, 0, time.Local) //not in achieve diapason
		result := isScanInInterval(achList[5], tt)

		require.False(t, result)
	})
	t.Run("in interval", func(t *testing.T) {
		tt := time.Date(2022, time.June, 7, 10, 30, 0, 0, time.Local) //not in achieve diapason
		result := isScanInInterval(achList[5], tt)

		require.True(t, result)
	})
}

func TestAddAchieve(t *testing.T) {
	t.Run("add one simple achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 2)

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

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		req := 2

		require.Equal(t, req, user.Achieves[3].AchieveLvl)
	})
	t.Run("add max level for multilevel achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		req := 3

		require.Equal(t, req, user.Achieves[3].AchieveLvl)
	})
	t.Run("add first level for multilevel complex achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)

		req := 1

		require.Equal(t, req, user.Achieves[4].AchieveLvl)
	})
	t.Run("add max level for multilevel complex achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)

		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)

		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)
		user.AddAchieve(time.Now(), 4)

		req := 3

		require.Equal(t, req, user.Achieves[4].AchieveLvl)
	})
	t.Run("check time of added achieve", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		tt := time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local) //not in achieve interval

		user.AddAchieve(tt, 5)

		req := achList.convertToUserAchieve(5)
		req.LastScan = tt

		require.NotEqual(t, req, user.Achieves[5])
	})

	t.Run("check time in interval", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		tt := time.Date(2022, time.June, 7, 10, 30, 0, 0, time.Local) //in achieve interval

		user.AddAchieve(tt, 5)

		req := achList.convertToUserAchieve(5)
		req.LastScan = tt

		require.Equal(t, req, user.Achieves[5])
	})

	t.Run("full user achieve data check ok", func(t *testing.T) {
		user := User{
			Id:       1,
			UsrLvl:   0,
			Achieves: map[int]UserAchieve{},
		}

		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)
		user.AddAchieve(time.Now(), 3)

		reqFirstAchieve := achList.convertToUserAchieve(3)
		reqFirstAchieve.Name = "Тестовая многоуровневая ачива простая"
		reqFirstAchieve.AchieveLvl = 1
		reqFirstAchieve.ScanCount = 3
		reqFirstAchieve.LastScan = time.Now()

		require.Equal(t, reqFirstAchieve, user.Achieves[3], "ачивка ID: 3 для начала")

		tt := time.Date(2022, time.June, 7, 10, 30, 0, 0, time.Local) //in achieve interval

		user.AddAchieve(tt, 6)
		user.AddAchieve(tt, 6)

		reqFirstLvl := UserAchieve{
			AchieveId:  6,
			AchieveLvl: 1,
			ScanCount:  2,
			Name:       "Тестовая многоуровневая сложная ачива полный сука фарш",
			LastScan:   tt,
		}

		require.Equal(t, reqFirstLvl, user.Achieves[6], "ачивка ID: 6 первый уровень")

		user.AddAchieve(tt, 6)
		user.AddAchieve(tt, 6)
		user.AddAchieve(tt, 6)
		user.AddAchieve(tt, 6)

		reqMaxLvl := UserAchieve{
			AchieveId:  6,
			AchieveLvl: 3,
			ScanCount:  6,
			Name:       "септолете тотал бля",
			LastScan:   tt,
		}

		require.Equal(t, reqMaxLvl, user.Achieves[6], "ачивка ID: 6 максимальный уровень")
	})
}
