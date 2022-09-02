package db

import (
	"awesomeProject5/logic"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type postgre struct {
	db *sqlx.DB
}

type userAchieveDB struct {
	Uid              int       `db:"uid"`
	AchieveId        int       `db:"achieve_id"`
	AchieveLvl       int       `db:"achieve_lvl"`
	ScanCount        int       `db:"scan_count"`
	Name             string    `db:"name"`
	LastScan         time.Time `db:"last_scan"`
	ScannedLocations []int     `db:"scanned_locs"`
	TempFl           bool      `db:"temp_fl"`
}

type UserDB struct {
	Id              int `db:"usr_id"`
	UsrLvl          int `db:"usr_lvl"`
	TempAchieves    map[int]*userAchieveDB
	CurrentAchieves map[int]*userAchieveDB
}

func InitDB() Saver {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "192.168.10.205", 5432, "adminpsg", "PadminSGG", "server_gredit_db")

	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		log.Fatalln("db conn err: ", err)
	}

	return &postgre{db: db}
}

type Saver interface {
	SaveUserData(user logic.User)
}

func achLogicToDB(achMap map[int]*logic.UserAchieve, uid int, fl bool) map[int]*userAchieveDB {
	result := make(map[int]*userAchieveDB)
	for i, achieve := range achMap {
		result[i] = &userAchieveDB{
			Uid:              uid,
			AchieveId:        achieve.AchieveId,
			AchieveLvl:       achieve.AchieveLvl,
			ScanCount:        achieve.ScanCount,
			Name:             achieve.Name,
			LastScan:         achieve.LastScan,
			ScannedLocations: achieve.ScannedLocations,
			TempFl:           fl,
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

	usrSaveStmt, err := p.db.PrepareNamed(`INSERT INTO test_ach.users values(:usr_id, :usr_lvl)`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	achSaveStmt, err := p.db.PrepareNamed(`INSERT INTO test_ach.user_achieves values(:uid, :achieve_id, :achieve_lvl, :scan_count, :name, :last_scan, :scanned_locs, :temp_fl)`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	usrDel, err := p.db.Prepare(`DELETE FROM test_ach.users where usr_id = $1`)
	if err != nil {
		log.Fatalln("stmt  err : ", err)
	}

	achDel, err := p.db.Prepare(`DELETE FROM test_ach.user_achieves where uid = $1`)
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
