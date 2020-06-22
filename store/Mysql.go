package store

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDB struct {
	db *gorm.DB
}

func (ptr *MysqlDB) Open(dbUri string) {
	var err error
	ptr.db, err = gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	} else {
		fmt.Println("connect database success")
		ptr.db.SingularTable(true)
	}
	defer ptr.db.Close()
}

func (ptr *MysqlDB) Insert(data interface{}) {
	ptr.db.Create(data)
}

func (ptr *MysqlDB) Update(data interface{}, where interface{}) {
	//search the 1st matched
	err := ptr.db.First(data, where).Error
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ptr.db.Save(data) //commit change
	}
}

func (ptr *MysqlDB) List(line interface{}, data []interface{}) {
	ptr.db.Limit(line).Find(&data)
}

func (ptr *MysqlDB) GetById(id interface{}, data interface{}) {
	err := ptr.db.First(&data, id).Error
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (ptr *MysqlDB) Delete(id interface{}, data interface{}) {
	ptr.db.Where("id = ?", id).Delete(&data)
}

func (ptr *MysqlDB) Clear(table interface{}) {
	ptr.db.Delete(&table)
}
