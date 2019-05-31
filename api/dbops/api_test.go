package dbops

import (
	"log"
	"testing"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comment")
	dbConn.Exec("truncate sessions")
}
func TestMain(m *testing.M) {
	clearTables()
	log.Printf("cao")
	m.Run()
	clearTables()
}
func TestUserWorkFlow(t *testing.T) {
	t.Run("ADD", testAddUser)
	t.Run("GET", testGetUser)
	t.Run("DEL", testDelUser)
	t.Run("REGET", testReGetUser)
}
func testAddUser(t *testing.T) {
	err := AddUserCredential("terry", "123")
	if err != nil {
		t.Errorf("Error of AddUser : %v", err)
	}
}
func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("terry")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}
func testDelUser(t *testing.T) {
	err := DeleteUser("terry", "123")
	if err != nil {
		t.Errorf("Error of DelUser : %v", err)
	}
}
func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("terry")
	if err != nil {
		t.Errorf("Error of ReGetUser : %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
