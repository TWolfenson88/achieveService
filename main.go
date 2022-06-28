package main

import (
	"fmt"
	"strings"
)


func convert(s string, numRows int) string {
	dict := make(map[int][]rune)

	if numRows == 1{
		return s
	}

	counter := 1
	increase := true

	for _, r := range s { //AB
		if increase {
			dict[counter] = append(dict[counter], r)
			if counter == numRows {
				increase = false
				counter--
			}else {
				counter++
			}
		}else{
			dict[counter] = append(dict[counter], r)
			if counter == 1 {
				increase = true
				counter++
			}else {
				counter--
			}
		}
	}

	sb := strings.Builder{}

	for i := 0; i < numRows+1; i++ {
		for _, r := range dict[i] {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}


func main()  {

	fmt.Println("Working...")

 s := "AB"
numRows := 1
//Output: "PAHNAPLSIIGYIR"


fmt.Println(convert(s, numRows))

	//aList := AchieveList{10: []AchieveElem{{
	//	NeedId:    23,
	//	Duration:  15*time.Minute,
	//	NeedCount: 2,
	//},
	//	{
	//		NeedId:    24,
	//		Duration:  10*time.Minute,
	//		NeedCount: 1,
	//	}}}
	//
	//fmt.Println("achieve list: ", aList[10])
	//
	//ach := Achieve{
	//	Id:          1,
	//	Name:        "Ебать ты тестер-тесто-тостер-лоукостер",
	//	PeriodStart: time.Time{},
	//	PeriodEnd:   time.Time{}.Add(20*time.Hour),
	//	NeedAchieve: 0,
	//}
	//
	//fmt.Println("ACHIVE ONE PERIODS: ", ach.PeriodStart.Hour(), ach.PeriodEnd.Hour())
}