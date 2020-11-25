package banco

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

// ConectCar - Abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	connectionString := "sqlserver://sa:@stefany@1994@localhost:1433?database=GoLangCurso"

	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
