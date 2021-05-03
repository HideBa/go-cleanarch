package infrastructure

import (
	"database/sql"

	"github.com/HideBa/go-cleanarch/app/config"
	"github.com/HideBa/go-cleanarch/app/interfaces/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() database.SqlHandler {
	conn, err := sql.Open("mysql", config.GetConfig().DBConfig.DBUrl)
	driver, _ := mysql.WithInstance(conn, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance("file://app/infrastructure/migrations", "mysql", driver)
	m.Steps(2)
	if err != nil {
		panic(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

// メモ：domainの中からここが呼ばれると、内側が外側に依存することになってしまうため、
// interfaces/database/sql_handlerでインターフェイスを定義して、そいつに依存させ、
// 実態の実装はここで行うようにする
func (h *SqlHandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := h.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (h *SqlHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := h.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRows), err
	}
	row := new(SqlRows)
	row.Rows = rows
	return row, err
}

type SqlResult struct {
	Result sql.Result
}

func (s SqlResult) LastInsertId() (int64, error) {
	return s.Result.LastInsertId()
}

func (s SqlResult) RowsAffected() (int64, error) {
	return s.Result.RowsAffected()
}

type SqlRows struct {
	Rows *sql.Rows
}

func (s SqlRows) Scan(dest ...interface{}) error {
	return s.Rows.Scan(dest...)
}

func (s SqlRows) Next() bool {
	return s.Rows.Next()
}

func (s SqlRows) Close() error {
	return s.Rows.Close()
}
