package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectSqlistDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("task3.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func main() {
	{
		db := ConnectSqlistDB()

		AutoMigrateStudent(db)

		SqlInsertStudent(db)
		SqlQueryStudent(db)
		SqlUpdateStudent(db)
		SqlDeleteStudent(db)
	}

	{
		db := ConnectSqlistDB()

		InitTransactionData(db)
		TransferMoney(db)

	}
}
