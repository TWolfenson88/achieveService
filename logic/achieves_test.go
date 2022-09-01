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

	t.Run("Just some add", func(t *testing.T) {

		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 10)

		fmt.Println(usr.TempAchieves[2], "NOOT NEELLLLLLLLLLLLL")

		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 20)
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 30)
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 10)

		fmt.Println("FINISH! ________")

		fmt.Println(usr.CurrentAchieves, usr.TempAchieves, "это не нил")

		require.NotNil(t, usr.CurrentAchieves)
	})

	t.Run("Check lvl increase", func(t *testing.T) {
		t.Skip()
		fmt.Println("test print 0 :", usr.CurrentAchieves[21], usr.TempAchieves[21])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11)
		fmt.Println("test print 1 :", usr.CurrentAchieves[21], usr.TempAchieves[21])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11)
		fmt.Println("test print 2 :", usr.CurrentAchieves[21], usr.TempAchieves[21])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11)
		fmt.Println("test print 3 :", usr.CurrentAchieves[21], usr.TempAchieves[21])
		usr.AddAchieve(time.Date(2022, time.June, 7, 7, 10, 0, 0, time.Local), 11)
		//fmt.Println("test print 14 :", usr.CurrentAchieves[21], usr.TempAchieves[21])

		fmt.Println("test print:", usr.CurrentAchieves[21], usr.TempAchieves[21])

		require.NotNil(t, usr.CurrentAchieves[21])
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
