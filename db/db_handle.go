package db

import (
	"awesomeProject5/logic"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

type Users map[int]*logic.User

type postgre struct {
	db *sqlx.DB
}

type userAchieveDB struct {
	Uid              int       `db:"uid"`
	AchieveId        int       `db:"achieve_id"`
	AchieveLvl       int       `db:"achieve_lvl"`
	MaxLvl           int       `db:"max_lvl"`
	ScanCount        int       `db:"scan_count"`
	Name             string    `db:"name"`
	LastScan         time.Time `db:"last_scan"`
	ScannedLocations string    `db:"scanned_locs"`
	TempFl           bool      `db:"temp_fl"`
}

type UserDB struct {
	Id              int `db:"usr_id"`
	UsrLvl          int `db:"usr_lvl"`
	TempAchieves    map[int]*userAchieveDB
	CurrentAchieves map[int]*userAchieveDB
}

func InitDB() Saver {
	port, _ := strconv.Atoi(getEnv("POSTGRES_PORT", "5432")) //хуйня какая-то, потом покрасивее надо сделать

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		getEnv("POSTGRES_HOST", "postgres"),
		port,
		getEnv("POSTGRES_USER", "postgres"),
		getEnv("POSTGRES_PASS", "postgres"),
		getEnv("POSTGRES_DB", "postgres"))

	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		log.Println("db conn err: ", err) //таймауты покрасивее придумать как
		time.Sleep(15 * time.Second)
		log.Fatalln("db conn err: ", err)
	}

	return &postgre{db: db}
}

type Saver interface {
	SaveUserData(user logic.User)
	InitUserData() map[int]*logic.User
}

func (p *postgre) InitUserData() map[int]*logic.User {
	users := make(map[int]*logic.User)

	//dbUsers := []UserDB{}
	rowsUsers, err := p.db.Queryx("SELECT * from ach_service.users")
	if err != nil {
		log.Println("users init data err : ", err)
	}

	achStmt, err := p.db.Preparex("select * from ach_service.user_achieves where uid = $1")

	for rowsUsers.Next() {
		user := logic.User{}
		err = rowsUsers.Scan(&user.Id, &user.UsrLvl)
		if err != nil {
			log.Println("struct users scan err : ", err)
		}

		achRows, err := achStmt.Queryx(user.Id)
		if err != nil {
			log.Println("ach queryX err : ", err)
		}

		user.TempAchieves = map[int]*logic.UserAchieve{}
		user.CurrentAchieves = map[int]*logic.UserAchieve{}

		for achRows.Next() {
			achieve := logic.UserAchieve{}

			var scannedLocsNull sql.NullString
			var uid int
			var tempFl bool

			err = achRows.Scan(&uid, &achieve.AchieveId, &achieve.AchieveLvl, &achieve.MaxLvl, &achieve.ScanCount, &achieve.Name, &achieve.LastScan, &scannedLocsNull, &tempFl)
			if err != nil {
				log.Println("struct achieves scan err : ", err)
			}

			var scnLocsInts []int
			var scannedLocs string

			if scannedLocsNull.Valid {
				scannedLocs = scannedLocsNull.String
			} else {
				scannedLocs = ""
			}

			fmt.Println("ach arr", scannedLocs)
			if len(scannedLocs) > 0 {
				strArr := strings.Split(scannedLocs[1:len(scannedLocs)-1], ",")

				for _, s := range strArr {
					res, err := strconv.Atoi(s)
					if err != nil {
						log.Println("str arr parse err", err)
					}
					scnLocsInts = append(scnLocsInts, res)
				}
			}

			achieve.ScannedLocations = scnLocsInts

			if tempFl {
				user.TempAchieves[achieve.AchieveId] = &achieve
			} else {
				user.CurrentAchieves[achieve.AchieveId] = &achieve
			}

		}

		users[user.Id] = &user
	}

	fmt.Println(users[15].CurrentAchieves[1])

	return users
}

func achLogicToDB(achMap map[int]*logic.UserAchieve, uid int, fl bool) map[int]*userAchieveDB {
	result := make(map[int]*userAchieveDB)
	for i, achieve := range achMap {
		result[i] = &userAchieveDB{
			Uid:        uid,
			AchieveId:  achieve.AchieveId,
			AchieveLvl: achieve.AchieveLvl,
			ScanCount:  achieve.ScanCount,
			Name:       achieve.Name,
			LastScan:   achieve.LastScan,
			//ScannedLocations: achieve.ScannedLocations,
			TempFl: fl,
		}
	}

	return result
}

func logicToDB(user logic.User) UserDB {
	return UserDB{
		Id:              user.Id,
		UsrLvl:          user.UsrLvl,
		TempAchieves:    achLogicToDB(user.TempAchieves, user.Id, true),
		CurrentAchieves: achLogicToDB(user.CurrentAchieves, user.Id, false),
	}
}

func (p *postgre) SaveUserData(user logic.User) {

	usr := logicToDB(user)

	usrSaveStmt, err := p.db.PrepareNamed(`INSERT INTO ach_service.users values(:usr_id, :usr_lvl)`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	achSaveStmt, err := p.db.PrepareNamed(`INSERT INTO ach_service.user_achieves values(:uid, :achieve_id, :achieve_lvl, :max_lvl, :scan_count, :name, :last_scan, :scanned_locs, :temp_fl)`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	usrDel, err := p.db.Prepare(`DELETE FROM ach_service.users where usr_id = $1`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	achDel, err := p.db.Prepare(`DELETE FROM ach_service.user_achieves where uid = $1`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	tx, err := p.db.BeginTxx(context.Background(), &sql.TxOptions{})
	if err != nil {
		log.Fatalln("transaction init err : ", err)
	}

	_, err = tx.Stmt(achDel).Exec(usr.Id)
	if err != nil {
		log.Fatalln("ach del stmt exec err : ", err)
	}

	_, err = tx.Stmt(usrDel).Exec(usr.Id)
	if err != nil {
		log.Fatalln("usr del stmt exec err : ", err)
	}

	_, err = tx.NamedStmt(usrSaveStmt).Exec(usr)
	if err != nil {
		log.Fatalln("usr stmt exec err : ", err)
	}

	for _, ach := range usr.TempAchieves {
		ach.Uid = usr.Id
		_, err = tx.NamedStmt(achSaveStmt).Exec(ach)
		if err != nil {
			log.Fatalln("achieve stmt exec err : ", err)
		}
	}

	for _, ach := range usr.CurrentAchieves {
		ach.Uid = usr.Id
		_, err = tx.NamedStmt(achSaveStmt).Exec(ach)
		if err != nil {
			log.Fatalln("achieve stmt exec err : ", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalln("transaction commit err : ", err)
	}

}
