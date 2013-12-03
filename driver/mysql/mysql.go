package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/brendensoares/storm"
	"github.com/brendensoares/storm/driver"
	"fmt"
	"strings"
)

type MysqlDriver struct {
	db *sql.DB
	config string
}


func (self *MysqlDriver) Open(config string) (openError error)  {
	if self.db, openError = sql.Open("mysql", config); openError != nil {
		// Failure
		return
	} else {
		// Success
		// Check that db is active
		if openError = self.db.Ping(); openError != nil {
			// Failure
			fmt.Println("OPEN ERROR:", openError)
			return
		}
		self.config = config
		fmt.Println("mysql driver open:", self)
		return nil
	}
}

func (self *MysqlDriver) Name() string {
	return "mysql"
}

func (self *MysqlDriver) Config() string {
	return self.config
}

func (self *MysqlDriver) Create(container string, query driver.Query) (id interface{}, requestError error)  {
	fmt.Println("mysql driver create:", container, query)
	// Create mysql insert string from given query
	var (
		queryFields []string
		queryValues []string
	)
	for k, v := range query {
		fmt.Println("key/value:", k, v)
		// TODO: check for sql injection
		queryFields = append(queryFields, k)
		queryValues = append(queryValues, fmt.Sprintf("'%s'", v))
	}
	fmt.Println("fields:", queryFields)
	fmt.Println("values:", queryValues)
	var result sql.Result
	if result, requestError = self.db.Exec(fmt.Sprintf("INSERT INTO `users`(%s) VALUES(%s)", strings.Join(queryFields,","), strings.Join(queryValues,","))); requestError != nil {
		// Failure
		return
	} else {
		// Success
		if id, requestError = result.LastInsertId(); requestError != nil {
			// Failure
			return
		}
	}
	// Success
	fmt.Println("success?:", id, requestError)
	return
}

func (self *MysqlDriver) ReadOne(container string, id interface{}) (result interface{}, requestError error)  {
	fmt.Println("mysql driver read one:", container, id)
	return
}

func (self *MysqlDriver) ReadMany(container string, conditions []interface{}) (results interface{}, requestError error)  {
	fmt.Println("mysql driver read many:", container, conditions)
	return
}

func (self *MysqlDriver) Update(container string, id interface{}, query driver.Query) (requestError error)  {
	fmt.Println("mysql driver update:", container, id, query)
	return
}

func (self *MysqlDriver) Delete(container string, id interface{}) (requestError error)  {
	fmt.Println("mysql driver delete:", container, id)
	return
}


func init() {
	storm.RegisterDriver(&MysqlDriver{})
}
