package banco

import (
	"database/sql"
	// import implícito
	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	////string padrao para mysql
	////urlConexao := "usuario:senha@/banco"
	stringConexao := "root:root@tcp(localhost:3306)/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
