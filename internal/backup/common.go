package backup

import (
	"database/sql"
	"errors"
	"fmt"

	cfg "github.com/KostyaBagr/duple-duple/internal/config"
	"github.com/KostyaBagr/duple-duple/internal/utils"
	_ "github.com/lib/pq"
)

// Generates full path. It local type of storage was selected all dump files will be uploaded
// there and will not be deleted. Instead in tmpDirPostres
func dumpFullPath(fileName string) (path string) {
	localStoragePath := cfg.AppConfig.Storage.Local.Path

	isLocalCfgEmpty := utils.IsEmpty(localStoragePath)
	if !isLocalCfgEmpty {
		_, err := utils.PathExists(localStoragePath, true)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return localStoragePath + fileName

	}

	return cfg.TmpBackupPath + fileName
}

type DumpDBMS interface {
	dump() (*DumpFileStats, string, error)
	ping() error
}

type dumpPostgres struct {
}

// Create a postgres dump methdo
func (p dumpPostgres) dump() (*DumpFileStats, string, error) {
	stat, path, err := PostgresDump(
		cfg.AppConfig.DBMS.Postgres.Host,
		cfg.AppConfig.DBMS.Postgres.User,
		cfg.AppConfig.DBMS.Postgres.Password,
		cfg.AppConfig.DBMS.Postgres.DB,
		cfg.AppConfig.DBMS.Postgres.Port,
	)
	if err != nil {
		fmt.Printf("Error during postgres dump: %s", err)
		return stat, "", err
	}
	return stat, path, nil
}

// Ping method
func (p dumpPostgres) ping() error {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=disable",
		cfg.AppConfig.DBMS.Postgres.Host,
		cfg.AppConfig.DBMS.Postgres.Port,
		cfg.AppConfig.DBMS.Postgres.User,
		cfg.AppConfig.DBMS.Postgres.Password,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Failed to open database connection: %v", err)
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Printf("Database is not ready: %v", err)
		return err
	}

	fmt.Println("Database is ready!")
	return nil
}

var dbmsStructMapper = map[string]DumpDBMS{
	cfg.Postgres.String(): dumpPostgres{},
}

// This is a discpatcher for dumping.
// Based on dbms var it runs a piplelin
// Returns DumpFileStats (such as time complexity,
// file size and so on),
// path to file (for storagies) and error (optional)
func DumpDispatcher(dbms string) (*DumpFileStats, string, error) {
	dbmsDumpProcessor, ok := dbmsStructMapper[dbms]
	if !ok {
		fmt.Printf("Invalid DBMS type %s", dbms)
		return &DumpFileStats{}, "", errors.New("Undefined dbms")
	}
	err := dbmsDumpProcessor.ping()
	if err != nil {
		return &DumpFileStats{}, "", err
	}
	return dbmsDumpProcessor.dump()

}
