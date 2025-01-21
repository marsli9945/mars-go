package marsSql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/marsli9945/mars-go/marsLog"
	"reflect"
	"testing"
)

type User struct {
	ID   int    `row:"id"`
	Name string `row:"name"`
}

func TestExecute_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			marsLog.Logger().ErrorF("Close db error: %v", err)
		}
	}(db)

	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectClose()

	err = ExecuteContext(context.Background(), db, "INSERT INTO users (name) VALUES (?)", "John")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestExecute_Failure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			marsLog.Logger().ErrorF("Close db error: %v", err)
		}
	}(db)

	mock.ExpectExec("INSERT INTO users").WillReturnError(fmt.Errorf("SQL error"))
	mock.ExpectClose()

	err = ExecuteContext(context.Background(), db, "INSERT INTO users (name) VALUES (?)", "John")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestSelect_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			marsLog.Logger().ErrorF("Close db error: %v", err)
		}
	}(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "John").
		AddRow(2, "Jane")

	mock.ExpectQuery("SELECT id, name FROM users").WillReturnRows(rows)
	mock.ExpectClose()

	var users []User
	err = SelectContext(context.Background(), db, "json", &users, "SELECT id, name FROM users")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestSelect_Failure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			marsLog.Logger().ErrorF("Close db error: %v", err)
		}
	}(db)

	mock.ExpectQuery("SELECT id, name FROM users").WillReturnError(fmt.Errorf("SQL error"))
	mock.ExpectClose()

	var users []User
	err = SelectContext(context.Background(), db, "json", &users, "SELECT id, name FROM users")
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestMapToAllSlice_StructSlice_Success(t *testing.T) {
	data := []map[string]any{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}

	var users []User
	err := mapToAllSlice("row", data, &users)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}

	// 验证用户数据是否正确
	if users[0].ID != 1 || users[0].Name != "John" {
		t.Errorf("expected user 1 to be John, got %v", users[0])
	}
	if users[1].ID != 2 || users[1].Name != "Jane" {
		t.Errorf("expected user 2 to be Jane, got %v", users[1])
	}
}

func TestMapToAllSlice_MapSlice_Success(t *testing.T) {
	data := []map[string]any{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}

	var users []map[string]any
	err := mapToAllSlice("json", data, &users)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestMapToAllSlice_UnsupportedType(t *testing.T) {
	data := []map[string]any{
		{"id": 1, "name": "John"},
	}

	var unsupportedType int
	err := mapToAllSlice("json", data, &unsupportedType)
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestMapToStructSlice_Success(t *testing.T) {
	data := []map[string]any{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}

	var users []User
	mapToStructSlice("row", data, reflect.TypeOf(&users).Elem(), reflect.ValueOf(&users))
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}

	// 验证用户数据是否正确
	if users[0].ID != 1 || users[0].Name != "John" {
		t.Errorf("expected user 1 to be John, got %v", users[0])
	}
	if users[1].ID != 2 || users[1].Name != "Jane" {
		t.Errorf("expected user 2 to be Jane, got %v", users[1])
	}
}

func TestMapToMapSlice_Success(t *testing.T) {
	data := []map[string]any{
		{"id": 1, "name": "John"},
		{"id": 2, "name": "Jane"},
	}

	var users []map[string]any
	mapToMapSlice(data, reflect.TypeOf(&users).Elem(), reflect.ValueOf(&users))
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}
