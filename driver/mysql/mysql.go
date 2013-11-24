package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/brendensoares/storm"
	"github.com/brendensoares/storm/driver"
	"fmt"
)

type MysqlDriver struct {
	db *sql.DB
	config string
}


func (self *MysqlDriver) Open(config string) (openError error)  {
	if db, dbError := sql.Open("mysql", config); dbError != nil {
		// Failure
		return
	} else {
		// Success
		self.db = db
		self.config = config
		return nil
	}
}

func (self *MysqlDriver) Name() string {
	return "mysql"
}

func (self *MysqlDriver) Config() string {
	return self.config
}

func (self *MysqlDriver) Create(container string, query driver.Query) (id string, requestError error)  {
	fmt.Println(container, query)
	return
}

func (self *MysqlDriver) ReadOne(container string, id interface{}) (result interface{}, requestError error)  {
	return
}

func (self *MysqlDriver) ReadMany(container string, conditions string) (results interface{}, requestError error)  {
	return
}

func (self *MysqlDriver) Update(container string, id interface{}, query driver.Query) (requestError error)  {
	return
}

func (self *MysqlDriver) Delete(container string, id interface{}) (requestError error)  {
	return
}


func init() {
	storm.RegisterDriver(&MysqlDriver{})
}
