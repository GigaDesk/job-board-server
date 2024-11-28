package postgreutils

import (
	"log"

	"github.com/GigaDesk/eardrum-server/graph/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresInstance struct {
	Dborm db.AutoGqlDB
}

func (p *PostgresInstance) Init(url string) {
	log.Println("Initializing PostgreSQL at: ", url)
	dbInstance, dbError := gorm.Open(postgres.Open(url), &gorm.Config{})
	if dbError != nil {
		panic(dbError)
	}
	dborm := db.NewAutoGqlDB(dbInstance)
	dborm.Init()
	p.Dborm = dborm
}
