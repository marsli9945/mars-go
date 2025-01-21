package marsRepo

import (
	"database/sql"
	"errors"
	"github.com/marsli9945/mars-go/marsLog"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestPrepareBatch_TransactionBeginFails_ReturnsError(t *testing.T) {
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

	mock.ExpectBegin().WillReturnError(errors.New("transaction begin error"))

	repo := &DBRepository{DB: db}
	err = repo.PrepareBatch("INSERT INTO table (col1, col2) VALUES (?, ?)", [][]any{{1, "value1"}})
	if err == nil {
		t.Errorf("expected an error, got nil")
	} else {
		if err.Error() != "transaction begin error" {
			t.Errorf("expected error 'transaction begin error', got '%s'", err.Error())
		}
	}
}

func TestPrepareBatch_PrepareStatementFails_ReturnsError(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO table (col1, col2) VALUES (?, ?)").WillReturnError(errors.New("prepare statement error"))
	mock.ExpectRollback()

	repo := &DBRepository{DB: db}
	err = repo.PrepareBatch("INSERT INTO table (col1, col2) VALUES (?, ?)", [][]any{{1, "value1"}})
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
	if err.Error() != "prepare statement error" {
		t.Errorf("expected error 'prepare statement error', got '%s'", err.Error())
	}
}

// todo mock.ExpectPrepare总失败
func TestPrepareBatch_ExecStatementFails_ReturnsError(t *testing.T) {
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

	mock.ExpectBegin()

	stmt := mock.ExpectPrepare(`INSERT INTO table\s*\(col1, col2\)\s*VALUES\s*\(\?, \?\);`)
	stmt.ExpectExec().WithArgs(1, "value1").WillReturnResult(sqlmock.NewResult(1, 1))
	stmt.ExpectExec().WithArgs(2, "value2").WillReturnResult(sqlmock.NewResult(2, 1))
	stmt.WillReturnError(errors.New("exec statement error"))
	mock.ExpectCommit()

	repo := &DBRepository{DB: db}
	err = repo.PrepareBatch("INSERT INTO table (col1, col2) VALUES (?, ?)", [][]any{{1, "value1"}, {2, "value2"}})
	if err == nil {
		t.Errorf("expected an error, got nil")
	} else {
		if err.Error() != "exec statement error" {
			t.Errorf("expected error 'exec statement error', got '%s'", err.Error())
		}
	}
}

func TestPrepareBatch_CommitTransactionFails_ReturnsError(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO table (col1, col2) VALUES (\\?, \\?)").ExpectExec().WithArgs(1, "value1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO table (col1, col2) VALUES (?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(errors.New("commit transaction error"))

	repo := &DBRepository{DB: db}
	err = repo.PrepareBatch("INSERT INTO table (col1, col2) VALUES (?, ?)", [][]any{{1, "value1"}})
	if err == nil {
		t.Errorf("expected an error, got nil")
	} else {
		if err.Error() != "commit transaction error" {
			t.Errorf("expected error 'commit transaction error', got '%s'", err.Error())
		}
	}
}

func TestPrepareBatch_SuccessfulExecution_ReturnsNoError(t *testing.T) {
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

	mock.ExpectBegin()
	mock.ExpectPrepare("INSERT INTO table (col1, col2) VALUES (?, ?)").ExpectExec().WithArgs(1, "value1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO table (col1, col2) VALUES (?, ?)").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	repo := &DBRepository{DB: db}
	err = repo.PrepareBatch("INSERT INTO table (col1, col2) VALUES (?, ?)", [][]any{{1, "value1"}})
	if err != nil {
		t.Errorf("expected no error, got '%s'", err.Error())
	}
}
