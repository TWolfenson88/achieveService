package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestAddAchieve(t *testing.T) {
	t.Run("add achieve", func(t *testing.T) {
		usr := User{
			Id: 1,
		}

		uAch := UserAchieve{
			AchieveId:      1,
			AchieveCount:   1,
			LastActivation: time.Now(),
		}

		usr.AddAchieve(uAch)
		require.Equal(t, uAch, usr.Achieves[0] )
	})

	t.Run("increase achieve counter", func(t *testing.T) {
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
	})

}