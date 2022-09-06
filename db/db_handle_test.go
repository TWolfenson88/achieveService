package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"testing"
)

func TestPostgre_InitUserData(t *testing.T) {

	t.Run("simple test", func(t *testing.T) {
		conn, err := sqlx.Connect("pgx", "host='localhost' port=5432 user='AchUser' password='A4Pass' dbname='postgres' sslmode=disable")
		if err != nil {
			log.Println("db conn err: ", err) //таймауты покрасивее придумать как
		}

		bb := postgre{db: conn}

		result := bb.InitUserData()
		fmt.Println("result! : ", result)
	})

}
