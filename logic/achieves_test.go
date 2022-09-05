package logic

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

//go test -v ./logic
//go test ./logic -bench=BenchmarkAddAchieve -benchmem

func TestUser_AddAchieve(t *testing.T) {

	usr := &User{
		Id:              1,
		UsrLvl:          0,
		TempAchieves:    map[int]*UserAchieve{},
		CurrentAchieves: map[int]*UserAchieve{},
	}

	logCh := make(chan string, 500)

	t.Run("Just some add", func(t *testing.T) {

		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 10, logCh)

		fmt.Println(usr.TempAchieves[2], "NOOT NEELLLLLLLLLLLLL")

		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 20, logCh)
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 30, logCh)
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 10, logCh)

		fmt.Println("FINISH! ________")

		fmt.Println(usr.CurrentAchieves, usr.TempAchieves, "это не нил")

		require.NotNil(t, usr.CurrentAchieves)
	})

	t.Run("Check lvl increase", func(t *testing.T) {
		//t.Skip()
		fmt.Println("test print 0 :", usr.CurrentAchieves[32], usr.TempAchieves[32])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11, logCh)
		fmt.Println("test print 1 :", usr.CurrentAchieves[32], usr.TempAchieves[32])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11, logCh)
		fmt.Println("test print 2 :", usr.CurrentAchieves[32], usr.TempAchieves[32])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11, logCh)
		fmt.Println("test print 3 :", usr.CurrentAchieves[32], usr.TempAchieves[32])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11, logCh)
		//fmt.Println("test print 14 :", usr.CurrentAchieves[21], usr.TempAchieves[21])

		fmt.Println("test print:", usr.CurrentAchieves[32], usr.TempAchieves[32])

		require.NotNil(t, usr.CurrentAchieves[32])
	})

	t.Run("Check bar achievement", func(t *testing.T) {
		lUsr := &User{
			Id:              2,
			UsrLvl:          0,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lUsr.AddAchieve(time.Date(2022, time.June, 7, 7, 1, 0, 0, time.Local), 7, logCh)
		fmt.Println("Added bar one time ", lUsr.TempAchieves[71].AchieveLvl)
		for count := 1; count < 6; count++ {
			l_time := time.Date(2022, time.June, 7, 7, 10+5*count, 0, 0, time.Local)
			lUsr.AddAchieve(l_time, 7, logCh)
		}
		//TODO Here we need check, that user not have achievement

		lUsr.AddAchieve(time.Date(2022, time.June, 7, 8, 1, 0, 0, time.Local), 7, logCh)
		// we get lvl1
		for count := 1; count < 8; count++ {
			lUsr.AddAchieve(time.Date(2022, time.June, 7, 8, 10+5*count, 0, 0, time.Local), 7, logCh)
		}
		// lvl2

		for count := 1; count < 8; count++ {
			lUsr.AddAchieve(time.Date(2022, time.June, 7, 9, 10+5*count, 0, 0, time.Local), 7, logCh)
		}
		// lvl3
		fmt.Println(lUsr.CurrentAchieves[71], lUsr.TempAchieves[71])
	})

	t.Run("Check id 1-3 achievements", func(t *testing.T) {
		lUsr := &User{
			Id:              3,
			UsrLvl:          0,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lTime := time.Date(2022, time.June, 7, 11, 10, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 0, logCh)
		// we must have only id 2 or 2 and 3
		fmt.Println(lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 7, 34, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 0, logCh)
		// we must get early bird achieve id 1
		fmt.Println(lUsr.CurrentAchieves)

		count := 0
		for lUsr.CurrentAchieves[3] == nil && count < 100 {
			lUsr.AddAchieve(lTime, 0, logCh)
			count++
		}
		fmt.Println("Count for lucky ", count)
		fmt.Println(lUsr.CurrentAchieves)
	})
}

/*func BenchmarkAddAchieve(b *testing.B) {

	users := generateUserArr(100)
	tt:= time.Now()

	b.Run("bench for 100 users with 18 adds", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, user := range users {
				user.AddAchieve(tt, 3)
				user.AddAchieve(tt, 3)
				user.AddAchieve(tt, 3)

				user.AddAchieve(tt, 6)
				user.AddAchieve(tt, 6)
				user.AddAchieve(tt, 6)
				user.AddAchieve(tt, 6)
				user.AddAchieve(tt, 6)
				user.AddAchieve(tt, 6)

				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)
				user.AddAchieve(tt, 4)


			}
		}
	})

	usr := User{
		Id:           1,
		UsrLvl:       0,
		TempAchieves: map[int]UserAchieve{},
	}

	b.Run("bench for 1 users with 3 scans", func(b *testing.B) {
		for i := 0; i < b.N; i++ {

			usr.AddAchieve(time.Now(), 3)
			usr.AddAchieve(time.Now(), 3)
			usr.AddAchieve(time.Now(), 3)
		}
	})
}
*/
