package logic

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCheckConditions(t *testing.T) {
	t.Run("check all conditions", func(t *testing.T) {
		achiv := Achieve{
			Id:          10,
			Name:        "testAch",
			PeriodStart: time.Time{}.Add(10*time.Hour + 10*time.Minute),
			PeriodEnd:   time.Time{}.Add(22*time.Hour),
			OpenedAchieves: []AchieveElem{{
				NeedId:    12,
				Duration:  15*time.Minute,
				NeedCount: 1,
			}},
		}

		usr := User{
			Id:       1488,
			Achieves: map[int]UserAchieve{12: {
				AchieveId:      12,
				AchieveCount:   1,
				LastActivation: time.Now().Add(-10*time.Minute),
			}},
		}

		require.True(t, achiv.CheckConditons(usr))
	})

	t.Run("have not needed achieve", func(t *testing.T) {
		achiv := Achieve{
			Id:          10,
			Name:        "testAch",
			PeriodStart: time.Time{}.Add(10*time.Hour + 10*time.Minute),
			PeriodEnd:   time.Time{}.Add(22*time.Hour),
			OpenedAchieves: []AchieveElem{{
				NeedId:    12,
				Duration:  15*time.Minute,
				NeedCount: 1,
			}},
		}

		usr := User{
			Id:       1488,
			Achieves: nil,
		}

		require.False(t, achiv.CheckConditons(usr))
	})

}

func TestAddAchieve(t *testing.T) {

	t.Run("add achieve into empty user", func(t *testing.T) {
		usr := User{
			Id: 1,
		}

		achiv := Achieve{
			Id:          10,
			Name:        "testAch",
			PeriodStart: time.Time{}.Add(10*time.Hour + 10*time.Minute),
			PeriodEnd:   time.Time{}.Add(22*time.Hour),
			OpenedAchieves: nil,
		}

		expected := UserAchieve{
			AchieveId:      10,
			AchieveCount:   1,
			LastActivation: time.Now(),
		}

		usr.AddAchieve(achiv)
		require.Equal(t, expected, usr.Achieves[10] )
	})

	t.Run("increase achieve counter", func(t *testing.T) {
		usr := User{
			Id: 1,
		}

		achiv := Achieve{
			Id:          10,
			Name:        "testAch",
			PeriodStart: time.Time{}.Add(10*time.Hour + 10*time.Minute),
			PeriodEnd:   time.Time{}.Add(22*time.Hour),
			OpenedAchieves: nil,
		}

		usr.AddAchieve(achiv)
		usr.AddAchieve(achiv)
		require.Equal(t, 2, usr.Achieves[10].AchieveCount )
	})

/*	t.Run("increase achieve counter", func(t *testing.T) {
		usr := User{
			Id: 1,
		}

		uAch := UserAchieve{
			AchieveId:      1,
			AchieveCount:   1,
			LastActivation: time.Now(),
		}

		usr.AddAchieve(uAch)
		usr.AddAchieve(uAch)
		require.Equal(t, 2, usr.Achieves[0].AchieveCount )
	})

	t.Run("two different achieves", func(t *testing.T) {
		usr := User{
			Id: 1,
		}

		uAch := UserAchieve{
			AchieveId:      1,
			AchieveCount:   1,
			LastActivation: time.Now(),
		}
		uAch2 := UserAchieve{
			AchieveId:      1488,
			AchieveCount:   12,
			LastActivation: time.Now().Add(time.Hour),
		}

		usr.AddAchieve(uAch)
		usr.AddAchieve(uAch2)
		require.Equal(t, []UserAchieve{uAch, uAch2}, usr.Achieves )
	})*/

}