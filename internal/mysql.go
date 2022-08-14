package mysql

import (
	"database/sql"
	"fmt"
	"github.com/ErfanMomeniii/Shop-Backend/config"
	"github.com/ErfanMomeniii/Shop-Backend/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Company{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Delivery{})
	db.AutoMigrate(&model.Order{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.User{})
}

func Create(configdb config.Mysql) (*gorm.DB, error) {
	dbconnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		configdb.Dbusername, configdb.Dbpassword, configdb.Dbhost, configdb.Dbport)
	sqldb, err := sql.Open("mysql", dbconnect)

	if err != nil {
		log.Fatal(err)
	}

	_, err = sqldb.Exec("CREATE DATABASE " + configdb.Dbname)

	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/authform?charset=%s&parseTime=True&loc=Local",
		configdb.Dbusername, configdb.Dbpassword, configdb.Dbhost, configdb.Dbport, configdb.Dbcharset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	Migrate(db)

	return db, nil
}

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
