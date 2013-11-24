package mysql

import (
	"github.com/brendensoares/storm"
	"testing"
)

type User struct {
	*storm.Model
	FirstName string
	LastName string
	Email string
}

func NewUser() *User {
	return storm.Factory(&User{}).(*User)
}

func TestMysqlCreate(t *testing.T) {
	if commError := storm.Connect("mysql", "user:@unix(/var/run/mysqld/mysqld.sock)/database"); commError != nil {
		// Failure
		t.Fatal("Database error")
	} else {
		// Success, create new user
		newUser := NewUser()
		newUser.Email = "brenden@test.com"
		if saveError, _ := newUser.Save(); saveError != nil {
			// Failure
			t.Fatalf("%s %s", "Query error", saveError)
		}
	}
}
