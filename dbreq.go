package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"

	"github.com/mattn/go-sqlite3"
)

type sqlDB struct {
	db *sql.DB
}

type user struct {
	id    uint64
	admin bool
	// name string
	// user string
}

type fanta struct {
	owner int
	name  string
	code  string
}

type sqlDBInterface interface {
	addUser(id uint64)
	close()
	getUserFantas(id uint64) []fanta
	createFanta(id uint64, nameFanta string)
	joinFanta(id uint64, inviteCode string) bool
	printUsers()
}

func connectDB() sqlDBInterface {
	var newDB sqlDB
	db, err := sql.Open("sqlite3", "./fantaf1.db?_fk=true")
	if err != nil {
		log.Fatal(err)
	}
	newDB.db = db
	var ret sqlDBInterface = newDB
	return ret
}

func prepareQuery(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}

func prepareQueryTransiction(tx *sql.Tx, query string) *sql.Stmt {
	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}

func prepareTransiction(db *sql.DB) *sql.Tx {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func (container sqlDB) addUser(id uint64) {
	db := container.db
	stmt := prepareQuery(db, "INSERT OR IGNORE INTO user (id) VALUES (?)")
	defer stmt.Close()
	_, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func (container sqlDB) createFanta(id uint64, nameFanta string) {
	db := container.db
	tx := prepareTransiction(db)
	var idBytes [8]byte
	binary.LittleEndian.PutUint64(idBytes[:], id)
	var toHash []byte = append([]byte(nameFanta), idBytes[:]...)
	var md5Hash [16]byte = md5.Sum(toHash)
	var inviteCode string = hex.EncodeToString(md5Hash[:])

	stmt := prepareQueryTransiction(tx, "INSERT INTO fanta (name, owner, code) values (?,?,?)")
	defer stmt.Close()
	_, err := stmt.Exec(nameFanta, id, inviteCode)
	if err != nil {
		log.Fatal(err)
	}

	stmt2 := prepareQueryTransiction(tx, "INSERT INTO fanta_user (user, fanta) values (?,?)")
	defer stmt2.Close()
	_, err2 := stmt2.Exec(id, inviteCode)
	if err2 != nil {
		log.Fatal(err2)
	}
	txErr := tx.Commit()
	if txErr != nil {
		log.Fatal(txErr)
	}
}

func (container sqlDB) getUserFantas(id uint64) []fanta {
	db := container.db
	stmt := prepareQuery(db, "SELECT f.name, f.code, f.owner FROM fanta f, fanta_user fu WHERE fu.user=? and fu.fanta=f.code")
	rows, err := stmt.Query(id)
	var fantas []fanta
	var temp fanta
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		rows.Scan(&temp.name, &temp.code, &temp.owner)
		fantas = append(fantas, temp)
	}
	return fantas
}

// return false if foreign key constraint fail
func (container sqlDB) joinFanta(id uint64, inviteCode string) bool {
	db := container.db
	stmt := prepareQuery(db, "INSERT INTO fanta_user (user, fanta) VALUES (?,?)")
	_, err := stmt.Exec(id, inviteCode)
	var sqliteErr sqlite3.Error
	if err != nil {
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintForeignKey) {
				return false
			}
		}
		log.Fatal(err)
	}
	return true
}

func (container sqlDB) printUsers() {
	db := container.db
	rows, err := db.Query("SELECT id from user")
	if err != nil {
		log.Fatal(err)
	}
	var buff int64
	for rows.Next() {
		rows.Scan(&buff)
		fmt.Println(buff)
	}

}

func (container sqlDB) close() {
	container.db.Close()
}
