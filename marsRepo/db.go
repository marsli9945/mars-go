package marsRepo

import (
	"context"
	"database/sql"
	"github.com/marsli9945/mars-go/marsLog"
	"github.com/marsli9945/mars-go/marsSql"
)

type DBRepository struct {
	DB       *sql.DB
	FieldTag string
}

func (repository *DBRepository) Execute(sql string, args ...any) error {
	return marsSql.ExecuteContext(context.Background(), repository.DB, sql, args...)
}

func (repository *DBRepository) ExecuteContext(ctx context.Context, sql string, args ...any) error {
	return marsSql.ExecuteContext(ctx, repository.DB, sql, args...)
}

func (repository *DBRepository) Select(results any, sentence string, args ...any) error {
	return marsSql.SelectContext(context.Background(), repository.DB, repository.FieldTag, results, sentence, args...)
}
func (repository *DBRepository) SelectContext(ctx context.Context, results any, sentence string, args ...any) error {
	return marsSql.SelectContext(ctx, repository.DB, repository.FieldTag, results, sentence, args...)
}

func (repository *DBRepository) PrepareBatch(query string, rows [][]interface{}) error {
	return repository.PrepareBatchContext(context.Background(), query, rows)
}
func (repository *DBRepository) PrepareBatchContext(ctx context.Context, query string, rows [][]interface{}) error {
	// 开始事务
	tx, err := repository.DB.BeginTx(ctx, nil)
	if err != nil {
		marsLog.Logger().ErrorF("BeginTx error: %v, query: %s", err, query)
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			err = tx.Rollback()
			if err != nil {
				marsLog.Logger().ErrorF("Rollback error: %v", err)
				return
			}
			marsLog.Logger().ErrorF("PrepareBatch error: %v", p)
		}
	}()

	// 准备插入语句
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		marsLog.Logger().ErrorF("PrepareContext error: %v, query: %s", err, query)
		return err
	}
	defer func(stmt *sql.Stmt) {
		err = stmt.Close()
		if err != nil {
			marsLog.Logger().ErrorF("Close stmt error: %v", err)
			return
		}
	}(stmt)

	for _, row := range rows {
		if _, err = stmt.ExecContext(ctx, row...); err != nil {
			marsLog.Logger().ErrorF("ExecContext error: %v, query: %s, params: %v", err, query, row)
			if rbErr := tx.Rollback(); rbErr != nil {
				marsLog.Logger().ErrorF("Rollback error: %v", rbErr)
			}
			return err
		}
	}

	// 提交事务
	if err = tx.Commit(); err != nil {
		marsLog.Logger().ErrorF("CommitTx error: %v, query: %s", err, query)
		return err
	}

	return nil
}
