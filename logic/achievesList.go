package logic

import (
	"fmt"
	"math/rand"
	"time"
)

func recoverFromNil() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func someTestLogic(usr *User, ach *Achieve) bool {
	defer recoverFromNil()
	fmt.Println("ALLO")
	if uAch, ok := usr.TempAchieves[ach.Id]; ok {
		for i, id := range ach.NeedLocations {
			fmt.Println("id: ", id, "  |||  another id: ", uAch.ScannedLocations[i])
			if id != uAch.ScannedLocations[i] {
				return false
			}
		}
		fmt.Println("***** RETURNED TRUE ")
		return true
	} else {
		fmt.Println("WROTTEN")
		uAch := convertToUserAchieve(*ach)
		uAch.ScannedLocations = append(uAch.ScannedLocations, ach.IdLoc)
		usr.TempAchieves[uAch.AchieveId] = &uAch
	}

	return false
}

func TestLogicForMultipleLosc(usr *User, ach *Achieve, locId int, scanTime time.Time, needLocs map[int]struct{}, timeout time.Duration, logCh chan string) bool {

	_, okCur := usr.CurrentAchieves[ach.Id] // проверяем наличие ачивы у юзера в уже полученных
	if okCur {
		return false
	}
	tempAch, okTmp := usr.TempAchieves[ach.Id] // проверяем наличие во временных ачивах
	_, nlOk := needLocs[locId]                 // проверяем, нужную ли локацию отсканили

	if !okTmp && nlOk { //если не нашли во временных ачивах, добавляем туда эту ачиву
		usr.TempAchieves[ach.Id] = &UserAchieve{
			AchieveId:        6,
			AchieveLvl:       0,
			MaxLvl:           1,
			ScanCount:        1,
			Name:             "",
			LastScan:         scanTime,
			ScannedLocations: []int{locId},
		}
	} else if okTmp && nlOk {
		if scanTime.Sub(tempAch.LastScan) < timeout { //если с момента последнего скана подходящей локации меньше 20 минут, ок
			tempAch.LastScan = scanTime

			founded := false //тут мы по бырику перебирем уже отсканированные локации в ачивке
			for _, location := range tempAch.ScannedLocations {
				if location == locId {
					founded = true
				}
			}

			if !founded { //и если не находим, то добавляем в массив айдишник локи
				tempAch.ScannedLocations = append(tempAch.ScannedLocations, locId)
			}

			if len(tempAch.ScannedLocations) == len(needLocs) { //и если вдруг оказалось что у нас все локации собрались, добавляем ачиву в полученные

				tempAch.Name = ach.NameForLvl[1] // тут обновляем имя
				tempAch.AchieveLvl = 1           // тут даем уровень
				logCh <- fmt.Sprintf("%d получил ачивку %s уровня %d", usr.Id, tempAch.Name, tempAch.AchieveLvl)
				usr.CurrentAchieves[ach.Id] = tempAch
				return false
			}

		} else {

			delete(usr.TempAchieves, ach.Id) //если прошло больше 20 минут, удаляем ачивку, обосрамс
		}
	}

	return false
}

var achList = AchieveList{
	0: []Achieve{ //тутт общие ачивый
		{
			Id:               1,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Ранняя пташка"},
			PeriodStart:      time.Time{}.Add(7 * time.Hour),
			PeriodEnd:        time.Time{}.Add(8*time.Hour + 59*time.Minute),
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
		{
			Id:               2,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Добро пожаловать"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
		{
			Id:               3,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Счастливчик"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				s1 := rand.NewSource(time.Now().UnixNano())
				r1 := rand.New(s1)
				if r1.Intn(100) > 90 {
					return true
				}
				return false
			},
		},
		{
			Id:               4,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Бессмертный"},
			PeriodStart:      time.Time{}.Add(6 * time.Hour),
			PeriodEnd:        time.Time{}.Add(7*time.Hour + 30*time.Minute),
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
		{
			Id:               5,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "4-20, тебе нормально?"},
			PeriodStart:      time.Time{}.Add(4*time.Hour + 15*time.Minute),
			PeriodEnd:        time.Time{}.Add(4*time.Hour + 25*time.Minute),
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
		{
			Id:               6,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "локации 2-3-4-5-6"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				needLocs := map[int]struct{}{2: {}, 3: {}, 4: {}, 5: {}, 6: {}} //мапа необходимых локаций
				timeout := 20 * time.Minute
				return TestLogicForMultipleLosc(usr, ach, locId, scanTime, needLocs, timeout, logCh) // в случае успешного прохождения логики ачивка добавится в Current прям из этой функции
			},
		},
		{
			Id:               7,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "локации 1-2-3-4-5-6-7-8-9-10-11-13"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				needLocs := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}, 10: {}, 11: {}, 12: {}, 13: {}} //мапа необходимых локаций
				timeout := 60 * time.Minute
				return TestLogicForMultipleLosc(usr, ach, locId, scanTime, needLocs, timeout, logCh)
			},
		},
		{
			Id:               8,
			IdLoc:            0,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "локации 13-11-10-7-5"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				needLocs := map[int]struct{}{5: {}, 7: {}, 10: {}, 11: {}, 13: {}} //мапа необходимых локаций
				timeout := 20 * time.Minute
				return TestLogicForMultipleLosc(usr, ach, locId, scanTime, needLocs, timeout, logCh)
			},
		},
	},

	2: []Achieve{
		{
			Id:               21,
			IdLoc:            2,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "начало ачивки 2-3"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				if usr.UsrLvl != 2 {
					return false
				}

				if _, okok := usr.CurrentAchieves[31]; okok {
					return false
				}

				_, twok := usr.TempAchieves[83] //проверяем, есть ли отсканированные другие ачивы
				_, eOk := usr.TempAchieves[131]

				if twok {
					delete(usr.TempAchieves, 83) //и удаляем их прогресс, если есть
				}

				if eOk {
					delete(usr.TempAchieves, 131)
				}

				// тут, наверное, ничего

				return true
			},
		},
		{
			Id:               22,
			IdLoc:            2,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "путь маскота ЭРИК"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    []int{2, 3, 4, 7, 8},
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				//предположим нам необходимо пройти локации 2-3-4-7-8
				//сначала проверим, получена ли ачивка кого либо из других масокотов

				_, eOk := usr.TempAchieves[32]
				_, shOk := usr.TempAchieves[42]
				_, grOk := usr.TempAchieves[51]

				if eOk || shOk || grOk || usr.UsrLvl != 3 {
					return false
				}

				//если у юзера нет 22 ачивы, добавляем. Если есть, ничо не делаем, путь выбран
				if _, ok := usr.TempAchieves[22]; !ok {
					usr.TempAchieves[22] = &UserAchieve{
						AchieveId:        22,
						AchieveLvl:       0,
						MaxLvl:           1,
						ScanCount:        1,
						Name:             "",
						LastScan:         scanTime,
						ScannedLocations: []int{2},
					}
				}

				return false
			},
		},
	},

	3: []Achieve{
		{
			Id:               31,
			IdLoc:            3,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "конец ачивки 2-3"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				_, ok := usr.TempAchieves[21]

				if ok {
					_, twok := usr.TempAchieves[131] //проверяем, есть ли отсканированные другие ачивы
					_, eOk := usr.TempAchieves[83]

					if twok {
						delete(usr.TempAchieves, 131) //и удаляем их прогресс, если есть
					}

					if eOk {
						delete(usr.TempAchieves, 83)
					}

					delete(usr.TempAchieves, 21) //удаляем старт прогресса

					usrAch := &UserAchieve{
						AchieveId:        ach.Id,
						AchieveLvl:       1,
						MaxLvl:           1,
						ScanCount:        1,
						Name:             ach.NameForLvl[1],
						LastScan:         scanTime,
						ScannedLocations: nil,
					}

					logCh <- fmt.Sprintf("%d получил ачивку %s уровня %d", usr.Id, usrAch.Name, usrAch.AchieveLvl)
					usr.CurrentAchieves[ach.Id] = usrAch // и добавляем её если всё збс

				}

				return false
			},
		},
		{
			Id:               32,
			IdLoc:            3,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "путь маскота ЙОЛК"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				return false
			},
		},
		{
			Id:               33,
			IdLoc:            3,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "временная ачивка для 22"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				//проверяем, есть ли у юзера ачивка 22, и если последняя отсканеная лока не 2, сбрасываем прогресс
				if aciv, ok := usr.TempAchieves[22]; ok {
					if aciv.ScannedLocations[len(aciv.ScannedLocations)-1] != 2 {
						usr.TempAchieves[22] = &UserAchieve{
							AchieveId:        22,
							AchieveLvl:       0,
							MaxLvl:           1,
							ScanCount:        1,
							Name:             "",
							LastScan:         scanTime,
							ScannedLocations: []int{2},
						}
					} else {
						//если последняя отсканеная лока 2, добавляем эту локу
						usr.TempAchieves[22].ScannedLocations = append(usr.TempAchieves[22].ScannedLocations, 3)
					}
				}
				return false
			},
		},
	},

	4: []Achieve{
		{
			Id:               41,
			IdLoc:            4,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "получена ачивка 13-4"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				_, ok := usr.TempAchieves[131]

				if ok {
					_, twok := usr.TempAchieves[21] //проверяем, есть ли отсканированные другие ачивы
					_, eOk := usr.TempAchieves[83]

					if twok {
						delete(usr.TempAchieves, 21) //и удаляем их прогресс, если есть
					}

					if eOk {
						delete(usr.TempAchieves, 83)
					}

					delete(usr.TempAchieves, 131) //удаляем старт прогресса

					usrAch := &UserAchieve{
						AchieveId:        ach.Id,
						AchieveLvl:       1,
						MaxLvl:           1,
						ScanCount:        1,
						Name:             ach.NameForLvl[1],
						LastScan:         scanTime,
						ScannedLocations: nil,
					}

					logCh <- fmt.Sprintf("%d получил ачивку %s уровня %d", usr.Id, usrAch.Name, usrAch.AchieveLvl)
					usr.CurrentAchieves[ach.Id] = usrAch // и добавляем её если всё збс

				}

				return false
			},
		},
		{
			Id:               42,
			IdLoc:            4,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "путь маскота ШП"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				return false
			},
		},
		{
			Id:               43,
			IdLoc:            4,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "временная ачивка для 22"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				//проверяем, есть ли у юзера ачивка 22, и если последняя отсканеная лока не 3, сбрасываем прогресс
				if aciv, ok := usr.TempAchieves[22]; ok {
					if aciv.ScannedLocations[len(aciv.ScannedLocations)-1] != 3 {
						usr.TempAchieves[22] = &UserAchieve{
							AchieveId:        22,
							AchieveLvl:       0,
							MaxLvl:           1,
							ScanCount:        1,
							Name:             "",
							LastScan:         scanTime,
							ScannedLocations: []int{2},
						}
					} else {
						//если последняя отсканеная лока 3, добавляем эту локу
						usr.TempAchieves[22].ScannedLocations = append(usr.TempAchieves[22].ScannedLocations, 4)
					}
				}
				return false
			},
		},
	},

	5: []Achieve{
		{
			Id:               51,
			IdLoc:            5,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "путь маскота ГРИНЧ"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				return false
			},
		},
	},

	7: []Achieve{ // дед бар
		{
			Id:               71,
			IdLoc:            7,
			MaxLevel:         3,
			BeginLevel:       0,
			ScansCountForLvl: map[int]int{1: 7, 2: 14, 3: 20},
			NameForLvl:       map[int]string{1: "Любитель вкусного", 2: "Специалист миксологии", 3: "Солист затейник"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
		},

		{
			Id:               72,
			IdLoc:            7,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "получена ачивка 8-7"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				_, ok := usr.TempAchieves[83]

				if ok {
					_, twok := usr.TempAchieves[21] //проверяем, есть ли отсканированные другие ачивы
					_, eOk := usr.TempAchieves[131]

					if twok {
						delete(usr.TempAchieves, 21) //и удаляем их прогресс, если есть
					}

					if eOk {
						delete(usr.TempAchieves, 131)
					}

					delete(usr.TempAchieves, 83) //удаляем старт прогресса

					usrAch := &UserAchieve{
						AchieveId:        ach.Id,
						AchieveLvl:       1,
						MaxLvl:           1,
						ScanCount:        1,
						Name:             ach.NameForLvl[1],
						LastScan:         scanTime,
						ScannedLocations: nil,
					}

					logCh <- fmt.Sprintf("%d получил ачивку %s уровня %d", usr.Id, usrAch.Name, usrAch.AchieveLvl)
					usr.CurrentAchieves[ach.Id] = usrAch // и добавляем её если всё збс

				}

				return false
			},
		},
		{
			Id:               73,
			IdLoc:            7,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "временная ачивка для 22"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				//проверяем, есть ли у юзера ачивка 22, и если последняя отсканеная лока не 4, сбрасываем прогресс
				if aciv, ok := usr.TempAchieves[22]; ok {
					if aciv.ScannedLocations[len(aciv.ScannedLocations)-1] != 4 {
						usr.TempAchieves[22] = &UserAchieve{
							AchieveId:        22,
							AchieveLvl:       0,
							MaxLvl:           1,
							ScanCount:        1,
							Name:             "",
							LastScan:         scanTime,
							ScannedLocations: []int{2},
						}
					} else {
						//если последняя отсканеная лока 3, добавляем эту локу
						usr.TempAchieves[22].ScannedLocations = append(usr.TempAchieves[22].ScannedLocations, 7)
					}
				}
				return false
			},
		},
	},
	8: []Achieve{
		{
			Id:               81,
			IdLoc:            8,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Ранний обжора"},
			PeriodStart:      time.Time{}.Add(10 * time.Hour),
			PeriodEnd:        time.Time{}.Add(11 * time.Hour),
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
		},
		{
			Id:               82,
			IdLoc:            8,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Солидный обедарь"},
			PeriodStart:      time.Time{}.Add(14 * time.Hour),
			PeriodEnd:        time.Time{}.Add(15 * time.Hour),
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
		},
		{
			Id:               83,
			IdLoc:            8,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "начало ачивки 8-7"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				/*	_, nineOk := usr.CurrentAchieves[92]
					_, fourOk := usr.CurrentAchieves[41]
					_, sevenOk := usr.CurrentAchieves[72]
					_, threeOk := usr.CurrentAchieves[31]

					if nineOk || fourOk || sevenOk || threeOk { //Если у нас получена какая-либо из финишных ачивок 13-9, 13-4, 2-3, 8-7, то не засчитываем это
						return false
					}*/

				if usr.UsrLvl != 2 {
					return false
				}

				if _, okok := usr.CurrentAchieves[72]; okok {
					return false
				}

				_, twok := usr.TempAchieves[21] //проверяем, есть ли отсканированные другие ачивы
				_, eOk := usr.TempAchieves[131]

				if twok {
					delete(usr.TempAchieves, 21) //и удаляем их прогресс, если есть
				}

				if eOk {
					delete(usr.TempAchieves, 131)
				}

				// тут, наверное, ничего

				return true
			},
		},
		{
			Id:               84,
			IdLoc:            8,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "временная ачивка для 22"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				//проверяем, есть ли у юзера ачивка 22, и если последняя отсканеная лока не 7, сбрасываем прогресс
				if aciv, ok := usr.TempAchieves[22]; ok {
					if aciv.ScannedLocations[len(aciv.ScannedLocations)-1] != 7 {
						usr.TempAchieves[22] = &UserAchieve{
							AchieveId:        22,
							AchieveLvl:       0,
							MaxLvl:           1,
							ScanCount:        1,
							Name:             "",
							LastScan:         scanTime,
							ScannedLocations: []int{2},
						}
					} else {
						//если последняя отсканеная лока 7, последний этап, записываем 22 ачивку в полученные, круто!
						usr.TempAchieves[22].ScannedLocations = append(usr.TempAchieves[22].ScannedLocations, 8)

						aciv.LastScan = scanTime
						aciv.AchieveLvl = 1
						aciv.Name = "ПУТЬ ЭРИКА ПРОЙДЕН ЕБАТЬ "

						usr.CurrentAchieves[22] = aciv
					}
				}
				return false
			},
		},
		{
			Id:               85,
			IdLoc:            8,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Эпичный ужинарь"},
			PeriodStart:      time.Time{}.Add(19 * time.Hour),
			PeriodEnd:        time.Time{}.Add(20*time.Hour + 30*time.Minute),
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
		},
	},
	9: []Achieve{
		{
			Id:               91,
			IdLoc:            9,
			MaxLevel:         3,
			BeginLevel:       1,
			ScansCountForLvl: map[int]int{1: 1, 2: 4, 3: 7},
			NameForLvl:       map[int]string{1: "Фанат", 2: "Коллекционер", 3: "Художник"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
		},
		{
			Id:               92,
			IdLoc:            9,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "получена ачивка 13-9"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				_, ok := usr.TempAchieves[131]

				if ok {
					_, twok := usr.TempAchieves[21] //проверяем, есть ли отсканированные другие ачивы
					_, eOk := usr.TempAchieves[83]

					if twok {
						delete(usr.TempAchieves, 21) //и удаляем их прогресс, если есть
					}

					if eOk {
						delete(usr.TempAchieves, 83)
					}

					delete(usr.TempAchieves, 131) //удаляем старт прогресса

					usrAch := &UserAchieve{
						AchieveId:        ach.Id,
						AchieveLvl:       1,
						MaxLvl:           1,
						ScanCount:        1,
						Name:             ach.NameForLvl[1],
						LastScan:         scanTime,
						ScannedLocations: nil,
					}

					logCh <- fmt.Sprintf("%d получил ачивку %s уровня %d", usr.Id, usrAch.Name, usrAch.AchieveLvl)
					usr.CurrentAchieves[ach.Id] = usrAch // и добавляем её если всё збс

				}

				return false
			},
		},
	},
	10: []Achieve{
		{
			Id:               101,
			IdLoc:            10,
			MaxLevel:         3,
			BeginLevel:       0,
			ScansCountForLvl: map[int]int{1: 3, 2: 6, 3: 9},
			NameForLvl:       map[int]string{1: "Посетитель", 2: "Забиватель", 3: "Твой братюня кальянщик"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
		},
	},
	11: []Achieve{
		{
			Id:               111,
			IdLoc:            11,
			MaxLevel:         2,
			BeginLevel:       0,
			ScansCountForLvl: map[int]int{1: 2, 2: 4},
			NameForLvl:       nil,
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
	},
	12: []Achieve{
		{
			Id:               121,
			IdLoc:            12,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Ух ты, что тут?"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
	},
	13: []Achieve{
		{
			Id:               131,
			IdLoc:            13,
			MaxLevel:         1,
			BeginLevel:       0,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Ух ебать сложная ачивка зависимая 13-9, 13-4"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {

				/*	_, nineOk := usr.CurrentAchieves[92]
					_, fourOk := usr.CurrentAchieves[41]
					_, sevenOk := usr.CurrentAchieves[72]
					_, threeOk := usr.CurrentAchieves[31]

					if nineOk || fourOk || sevenOk || threeOk { //Если у нас получена какая-либо из финишных ачивок 13-9, 13-4, 2-3, 8-7, то не засчитываем это
						return false
					}
				*/

				if usr.UsrLvl != 2 {
					return false
				}

				_, okok := usr.CurrentAchieves[92]
				_, okokk := usr.CurrentAchieves[41]

				if okokk && okok {
					return false
				}

				_, twok := usr.TempAchieves[21] //проверяем, есть ли отсканированные другие ачивы
				_, eOk := usr.TempAchieves[83]

				if twok {
					delete(usr.TempAchieves, 21) //и удаляем их прогресс, если есть
				}

				if eOk {
					delete(usr.TempAchieves, 83)
				}

				// тут, наверное, ничего

				return true
			},
		},
	},
	14: []Achieve{
		{
			Id:               141,
			IdLoc:            14,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       map[int]string{1: "Мокрый, но жаркий"},
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic:     nil,
		},
	},
	20: []Achieve{
		{
			Id:               201,
			IdLoc:            20,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       nil,
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				fmt.Println("chck20 ")
				if achach, ok := usr.TempAchieves[2]; ok {
					fmt.Println("chck20 ")
					achach.ScannedLocations = append(achach.ScannedLocations, 20)
				}
				return true
			},
		},
	},
	30: []Achieve{
		{
			Id:               301,
			IdLoc:            30,
			MaxLevel:         1,
			BeginLevel:       1,
			ScansCountForLvl: nil,
			NameForLvl:       nil,
			PeriodStart:      time.Time{},
			PeriodEnd:        time.Time{},
			Cooldown:         0,
			NeedAchieves:     nil,
			NeedLocations:    nil,
			SpecialLogic: func(usr *User, ach *Achieve, locId int, scanTime time.Time, logCh chan string) bool {
				fmt.Println("chck30 ")
				if achach, ok := usr.TempAchieves[2]; ok {
					fmt.Println("chck30 ")
					achach.ScannedLocations = append(achach.ScannedLocations, 30)
				}
				return true
			},
		},
	},
}

/*
var achList = AchieveList{10: Achieve{
	IdLoc:            10,
	MaxLevel:         1,
	BeginLevel:       1,
	ScansCountForLvl: nil,
	NameForLvl:       nil,
	PeriodStart:      time.Time{},
	PeriodEnd:        time.Time{},
	NeedAchieves:     nil,
},
	11: Achieve{
		IdLoc:            11,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	12: Achieve{
		IdLoc:            12,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	13: Achieve{
		IdLoc:            13,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	14: Achieve{
		IdLoc:            14,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	15: Achieve{
		IdLoc:            15,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	16: Achieve{
		IdLoc:            16,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	17: Achieve{
		IdLoc:            17,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	18: Achieve{
		IdLoc:            18,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	19: Achieve{
		IdLoc:            19,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	20: Achieve{
		IdLoc:            20,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	0: Achieve{
		IdLoc:            0,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       nil,
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	2: Achieve{
		IdLoc:            2,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: nil,
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	3: Achieve{
		IdLoc:            3,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 3, 2: 6, 3: 9},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая ачива простая"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves:     nil,
	},
	4: Achieve{
		IdLoc:            4,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 3, 2: 6, 3: 9},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая ачива СЛОЖНАЯ"},
		PeriodStart:      time.Time{},
		PeriodEnd:        time.Time{},
		NeedAchieves: map[int]AchieveElem{3: {
			NeedId:    3,
			Duration:  0,
			NeedCount: 3,
		}},
	},
	5: Achieve{
		IdLoc:            5,
		MaxLevel:         1,
		BeginLevel:       1,
		ScansCountForLvl: map[int]int{1: 1},
		NameForLvl:       map[int]string{1: "Тестовая одноуровневая простая ачива с временем"},
		PeriodStart:      time.Time{}.Add(10*time.Hour + 10*time.Minute), // from 10:10 AM
		PeriodEnd:        time.Time{}.Add(20 * time.Hour),                   // to 8:00 PM
		NeedAchieves:     nil,
	},
	6: Achieve{
		IdLoc:            6,
		MaxLevel:         3,
		BeginLevel:       0,
		ScansCountForLvl: map[int]int{1: 2, 2: 4, 3: 6},
		NameForLvl:       map[int]string{1: "Тестовая многоуровневая сложная ачива полный сука фарш", 2:"промежуточный фарш", 3:"септолете тотал бля"},
		PeriodStart:      time.Time{}.Add(10*time.Hour + 10*time.Minute), // from 10:10 AM
		PeriodEnd:        time.Time{}.Add(20 * time.Hour),                   // to 8:00 PM
		NeedAchieves: map[int]AchieveElem{3: {
			NeedId:    3,
			Duration:  0,
			NeedCount: 3,
		}},
	},
}
*/
