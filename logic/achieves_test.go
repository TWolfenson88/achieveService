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

		t.Skip()

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

		t.Skip()

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

		t.Skip()

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

		t.Skip()

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

	t.Run("check 2-3-4-5-6 achieve sucess", func(t *testing.T) {
		t.Skip()

		lUsr := &User{
			Id:              3,
			UsrLvl:          0,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lTime := time.Date(2022, time.June, 7, 11, 10, 0, 0, time.Local)

		lUsr.AddAchieve(lTime, 2, logCh)

		_, ok := lUsr.CurrentAchieves[6] //проверяем, не получена ли 6я ачивка
		require.False(t, ok)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 3, logCh)

		_, ok = lUsr.CurrentAchieves[6]
		require.False(t, ok)

		lTime = time.Date(2022, time.June, 7, 11, 12, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 5, logCh)

		_, ok = lUsr.CurrentAchieves[6]
		require.False(t, ok)

		lTime = time.Date(2022, time.June, 7, 11, 13, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 4, logCh)

		_, ok = lUsr.CurrentAchieves[6]
		require.False(t, ok)

		lTime = time.Date(2022, time.June, 7, 11, 14, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 6, logCh)

		_, ok = lUsr.CurrentAchieves[6]
		require.True(t, ok)

		//lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		//lUsr.AddAchieve(lTime, 3, logCh)
	})

	t.Run("check 2-3-4-5-6 achieve out of time", func(t *testing.T) {
		t.Skip()

		lUsr := &User{
			Id:              3,
			UsrLvl:          0,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lTime := time.Date(2022, time.June, 7, 11, 10, 0, 0, time.Local)

		lUsr.AddAchieve(lTime, 2, logCh)

		_, ok := lUsr.CurrentAchieves[6] //проверяем, не получена ли 6я ачивка
		require.False(t, ok)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 3, logCh)

		_, ok = lUsr.CurrentAchieves[6]
		require.False(t, ok)

		lTime = time.Date(2022, time.June, 7, 11, 42, 0, 0, time.Local) // + 32 минуты, вне диапазона
		lUsr.AddAchieve(lTime, 5, logCh)

		_, ok = lUsr.CurrentAchieves[6]
		_, okT := lUsr.TempAchieves[6]
		require.False(t, ok)
		require.False(t, okT) //должна быть удалена из временных ачив

		//fmt.Println("__________________ CURRENT USER ACHIEVES: \n", lUsr.CurrentAchieves)
		//fmt.Println("__________________ TEMP USER ACHIEVES: \n", lUsr.TempAchieves)

		//lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		//lUsr.AddAchieve(lTime, 3, logCh)
	})

	t.Run("check 13-9 13-4 2-3 8-7 achieves logic", func(t *testing.T) {
		//t.Skip()

		//логика такая:
		//скаинруем 13 локу, затем 2 локу, потом 9 локу. При скане 13 и 2 локи, у нас начинается прогресс по
		// 13-4, 13-9 и 2-3 соответственно. И после скана 9 локи у нас зачисляется ачивка 13-9, остальные сбрасываются
		//и больше никогда не доступны
		lUsr := &User{
			Id:              3,
			UsrLvl:          2,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lTime := time.Date(2022, time.June, 7, 11, 10, 0, 0, time.Local)

		lUsr.AddAchieve(lTime, 13, logCh)
		_, ok := lUsr.TempAchieves[131]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 2, logCh)
		_, ok = lUsr.TempAchieves[21]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 11, 42, 0, 0, time.Local) // + 32 минуты, вне диапазона
		lUsr.AddAchieve(lTime, 3, logCh)

		_, ok = lUsr.CurrentAchieves[31]

		_, threetenok := lUsr.TempAchieves[131]
		_, twoOk := lUsr.TempAchieves[21]

		//У нас в куррентах должна быть записана 92 ачивка, как конец последовательности 13-9, при этом 13 и 2 удалены
		require.True(t, ok)
		require.False(t, threetenok)
		require.False(t, twoOk)


	})

	t.Run("ERIC way achieve test", func(t *testing.T) {
		//t.Skip()

		//логика такая:
		//скаинруем 13 локу, затем 2 локу, потом 9 локу. При скане 13 и 2 локи, у нас начинается прогресс по
		// 13-4, 13-9 и 2-3 соответственно. И после скана 9 локи у нас зачисляется ачивка 13-9, остальные сбрасываются
		//и больше никогда не доступны
		lUsr := &User{
			Id:              3,
			UsrLvl:          0,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lTime := time.Date(2022, time.June, 7, 11, 10, 0, 0, time.Local)

		lUsr.AddAchieve(lTime, 2, logCh)
		_, ok := lUsr.TempAchieves[22]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 3, logCh)
		_, ok = lUsr.TempAchieves[22]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 4, logCh)
		_, ok = lUsr.TempAchieves[22]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 7, logCh)
		_, ok = lUsr.TempAchieves[22]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		lTime = time.Date(2022, time.June, 7, 11, 11, 0, 0, time.Local)
		lUsr.AddAchieve(lTime, 8, logCh)
		_, ok = lUsr.TempAchieves[22]
		require.True(t, ok)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		_, ok = lUsr.CurrentAchieves[22]

		//_, threetenok := lUsr.TempAchieves[131]
		//_, twoOk := lUsr.TempAchieves[21]

		require.True(t, ok)
		//require.False(t, threetenok)
		//require.False(t, twoOk)

		//fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		//fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)

		//lTime = time.Date(2022, time.June, 7, 11, 42, 0, 0, time.Local) // + 32 минуты, вне диапазона
		//lUsr.AddAchieve(lTime, 13, logCh)

		//_, ok = lUsr.TempAchieves[131]
		//require.False(t, ok)
	})
	
	t.Run("checking 91 achieve", func(t *testing.T) {
		lUsr := &User{
			Id:              3,
			UsrLvl:          0,
			TempAchieves:    map[int]*UserAchieve{},
			CurrentAchieves: map[int]*UserAchieve{},
		}
		lTime := time.Date(2022, time.June, 7, 11, 10, 0, 0, time.Local)

		lUsr.AddAchieve(lTime, 9, logCh)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)



		_, ok := lUsr.CurrentAchieves[91]
		require.True(t, ok)

		lUsr.AddAchieve(lTime, 9, logCh)

		fmt.Println("---- temp achieves ----", lUsr.TempAchieves)
		fmt.Println("---- curr achieves ----", lUsr.CurrentAchieves)



		_, ok = lUsr.CurrentAchieves[91]
		require.True(t, ok)

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
