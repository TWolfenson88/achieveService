package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type postgre struct {
	db *sqlx.DB
}

type userAchieveDB struct {
	AchieveId        int       `db:"achieve_id"`
	AchieveLvl       int       `db:"achieve_lvl"`
	ScanCount        int       `db:"scan_count"`
	Name             string    `db:"name"`
	LastScan         time.Time `db:"last_scan"`
	ScannedLocations []int     `db:"scanned_locs"`
	tempFl           bool      `db:"temp_fl"`
}

type userDB struct {
	Id              int `db:"usr_id"`
	UsrLvl          int `db:"usr_lvl"`
	TempAchieves    map[int]*userAchieveDB
	CurrentAchieves map[int]*userAchieveDB
}

func InitDB() Saver {
	db, err := sqlx.Connect("pgx", "")
	if err != nil {
		log.Fatalln("db conn err: ", err)
	}

	return &postgre{db: db}
}

type Saver interface {
	SaveUser(user userDB)
}

func (p *postgre) SaveUser(user userDB) {
	tx := p.db.MustBegin()

	tx.MustExec("INSERT INTO table (usr_id, usr_lvl) values (:usr_id, :usr_lvl)", user)
}
