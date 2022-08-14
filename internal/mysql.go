package mysql

import (
	"github.com/ErfanMomeniii/Shop-Backend/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func Withretry(createdb func(dbconfig config.Mysql) (*gorm.DB, error), configdb config.Mysql, attemptlimit int) *gorm.DB {
	for i := 0; i < attemptlimit; i++ {
		db, err := createdb(configdb)
	
		if err != nil {
			log.Error(err)
		}
	
		return db
	}

	log.Panic("we can not connect to database")
	
	return nil
}
