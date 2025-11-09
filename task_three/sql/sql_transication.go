package main

import (
	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
)

// 假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
type Account struct {
	Id      int `gorm:column:id;primary_key;auto_increment`
	Name    string `gorm:column:name;size:255;not null;index`
	Balance int `gorm:column:balance;not null`
}
type Transaction struct {
	Id            int `gorm:column:id;primary_key;auto_increment`
	FromAccountId int `gorm:column:from_account_id;not null`
	ToAccountId   int `gorm:column:to_account_id;not null`
	Amount        int `gorm:column:amount;not null`
}

func InitTransactionData(db *gorm.DB) {
	db.Migrator().DropTable(&Account{}, &Transaction{})

	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
	db.Create(&Account{Name: "A", Balance: 1000})
	db.Create(&Account{Name: "B", Balance: 10})
	db.Create(&Account{Name: "C", Balance: 100})
}

// 要求 ：
// 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
func TransferMoney(db *gorm.DB) {
	// 开启事务
	tx := db.Begin()
	// 获取账户 A 的余额
	var accountA Account
	err := tx.Where("Name = ?", "A").First(&accountA).Error
	if err != nil {
		tx.Rollback()
		return
	}
	var accountB Account
	err = tx.Where("Name = ?", "B").First(&accountB).Error
	if err != nil {
		tx.Rollback()
		return
	}
	// 检查账户 A 的余额是否充足
	if accountA.Balance >= 100 {
		// 扣除账户 A 的 100 元
		accountA.Balance -= 100
		accountB.Balance += 100
		tx.Save(accountA)
		tx.Save(accountB)
		tx.Create(&Transaction{FromAccountId: accountA.Id, ToAccountId: accountB.Id, Amount: 100})
		tx.Commit()
	}else{
		tx.Rollback()
	}
}
