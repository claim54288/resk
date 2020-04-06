package test

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
)

var db *dbx.Database

func init() {
	settings := dbx.Settings{
		DriverName: "mysql",
		User:       "root",
		Password:   "root",
		Database:   "po",
		Host:       "127.0.0.1:3306",
		Options: map[string]string{
			"parseTime": "true",
		},
	}
	var err error
	db, err = dbx.Open(settings)
	if err != nil {
		fmt.Println(err)
	}
	db.RegisterTable(&GoodsSigned{}, "goods")
	db.RegisterTable(&GoodsUnsigned{}, "goods_unsigned")
}

//事务锁方案
func UpdateForLock(g Goods) {
	//通过db.tx函数构建事务锁代码块
	err := db.Tx(func(runner *dbx.TxRunner) error {
		//第一步：锁定需要修改的资源，也就是数据行
		//编写事务锁查询语句，使用for update子句来锁定资源
		query := "SELECT * FROM goods WHERE envelope_no = ?" +
			" FOR UPDATE"
		out := &GoodsSigned{}
		ok, err := runner.Get(out, query, g.EnvelopeNo)
		if !ok || err != nil {
			return err
		}
		//第二步：计算剩余金额和剩余数量
		subAmount := decimal.NewFromFloat(0.01)
		remainAmount := out.RemainAmount.Sub(subAmount)
		remainQuantity := out.ReaminQuantity - 1
		//第三步：执行更新
		update := "UPDATE goods SET remain_amount=?,remain_quantity=?" +
			" WHERE envelope_no=?"
		_, rowsAffected, err := db.Execute(update, remainAmount, remainQuantity, g.EnvelopeNo)
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return errors.New("库存扣减失败")
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

//数据库无符号类型+直接更新方案
func UpdateForUnsigned(g Goods) {
	update := "UPDATE goods_unsigned SET remain_amount=remain_amount-?,remain_quantity=remain_quantity-? " +
		"WHERE envelope_no=?"
	_, rowsAffected, err := db.Execute(update, 0.01, 1, g.EnvelopeNo)
	if err != nil {
		fmt.Println(err)
	}
	if rowsAffected < 1 {
		fmt.Println("扣减失败")
	}
}

//乐观锁方案
func UpdateDorOptimistic(g Goods) {
	update := "update goods set remain_amount=remain_amount-?,remain_quantity=remain_quantity-?" +
		" WHERE envelope_no=? AND remain_amount>=? AND remain_quantity>=?"
	_, rowsAffected, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if rowsAffected < 1 {
		fmt.Println("扣减失败")
	}

}

//乐观锁+无符号方案
func UpdateDorOptimisticAndUnsigned(g Goods) {
	update := "update goods_unsigned set remain_amount=remain_amount-?,remain_quantity=remain_quantity-?" +
		" WHERE envelope_no=? AND remain_amount>=? AND remain_quantity>=?"
	_, rowsAffected, err := db.Execute(update, 0.01, 1, g.EnvelopeNo, 0.01, 1)
	if err != nil {
		fmt.Println(err)
	}
	if rowsAffected < 1 {
		fmt.Println("扣减失败")
	}

}
