package dal

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	cnf "orbiz.one/template-service/src/config"
)

type DBManager struct {
	Config *cnf.DBConfig
}

var mutex sync.Mutex
var isDBReady bool

func (dbm *DBManager) initDB() (db *sql.DB, err error) {
	mutex.Lock()
	defer mutex.Unlock()

	if isDBReady {

		return db, nil
	}

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
		dbm.Config.Host, dbm.Config.Port, dbm.Config.DBName)

	// Connect to PostgreSQL
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Attempt a connection ping to check if the connection was successful
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	isDBReady = true

	return db, nil
}

func (dbm *DBManager) GetDBConnector() (*sql.DB, error) {

	db, err := dbm.initDB()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDBManager(dbCfg *cnf.DBConfig) *DBManager {
	return &DBManager{
		Config: dbCfg,
	}

}
