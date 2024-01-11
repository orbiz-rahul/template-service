package dal

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	cnf "orbiz.one/template-service/src/config"
)

type DBManager struct {
	db        *sql.DB
	initMutex sync.Mutex
	once      sync.Once
}

// InitDB initializes the PostgreSQL database connection
func (dbm *DBManager) initDB(dbCfg *cnf.DBConfig) error {
	/*
		Do calls the function f if and only if Do is being called for the first time for this instance of Once.
		In other words, given
		var once Once
		if once.Do(f) is called multiple times, only the first call will invoke f, even if f has a different value in each invocation. A new instance of Once is required for each function to execute.

		Do is intended for initialization that must be run exactly once. Since f is niladic, it may be necessary to use a function literal to capture the arguments to a function to be invoked by Do:

		config.once.Do(func() { config.init(filename) })
		Because no call to Do returns until the one call to f returns, if f causes Do to be called, it will deadlock.

		If f panics, Do considers it to have returned; future calls of Do return without calling f.
	*/
	dbm.once.Do(func() {
		// Acquire the lock to prevent concurrent initialization
		dbm.initMutex.Lock()
		defer dbm.initMutex.Unlock()
		// Check again if the database is already initialized to avoid duplication
		if dbm.db == nil {

			// Construct the connection string
			connStr := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
				dbCfg.Host, dbCfg.Port, dbCfg.DBName)
			// Connect to PostgreSQL
			db, err := sql.Open("postgres", connStr)
			if err != nil {
				return
			}
			// Perform a ping to ensure the connection is valid
			err = db.Ping()
			if err != nil {
				fmt.Println("Error pinging the database:", err)
				return
			}
			dbm.db = db
		}
	})

	return nil
	/*mutex.Lock()
	defer mutex.Unlock()

	if isDBReady && dbm.db != nil {

		return nil
	}

	// Construct the connection string
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.DBName)

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	// Attempt a connection ping to check if the connection was successful
	err = db.Ping()
	if err != nil {
		return err
	}

	isDBReady = true
	dbm.db = db
	return nil*/
}

func (dbm *DBManager) GetDBConnector(dbCfg *cnf.DBConfig) (*sql.DB, error) {

	err := dbm.initDB(dbCfg)
	if err != nil {
		return nil, err
	}
	return dbm.db, nil
}

func GetDBManager() *DBManager {
	return &DBManager{}
}
