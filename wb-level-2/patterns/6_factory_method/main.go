package main

import (
	"errors"
	"fmt"
)

func main() {
	db, _ := GetDB("postgres")
	fmt.Printf("%#v\n", db)

	db, _ = GetDB("mysql")
	fmt.Printf("%#v\n", db)

	db, err := GetDB("bib")
	fmt.Printf("%#v\t%#v\n", db, err)
}

//GetDB ...
func GetDB(name string) (DB, error) {
	switch name {
	case "postgres":
		return NewPostgresDB(), nil
	case "mysql":
		return NewMysqlDB(), nil
	default:
		return nil, errors.New("unknown db name")
	}
}

//DB ...
type DB interface {
	SetValue(value string)
	GetValue() string
}

//PostgresDB ...
type PostgresDB struct {
	val string
}

//NewPostgresDB ...
func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

//SetValue ...
func (db *PostgresDB) SetValue(value string) {
	db.val = value
}

//GetValue ...
func (db *PostgresDB) GetValue() string {
	return db.val
}

//MysqlDB ...
type MysqlDB struct {
	val []byte
}

//NewMysqlDB ...
func NewMysqlDB() *MysqlDB {
	return &MysqlDB{}
}

//SetValue ...
func (db *MysqlDB) SetValue(value string) {
	db.val = []byte(value)
}

//GetValue ...
func (db *MysqlDB) GetValue() string {
	return string(db.val)
}
