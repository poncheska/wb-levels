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

func GetDB(name string) (DB,error) {
	switch name {
	case "postgres":
		return NewPostgresDB(), nil
	case "mysql":
		return NewMysqlDB(), nil
	default:
		return nil, errors.New("unknown db name")
	}
}

type DB interface {
	SetValue(value string)
	GetValue() string
}

type PostgresDB struct {
	val string
}

func NewPostgresDB() *PostgresDB {
	return &PostgresDB{}
}

func (db *PostgresDB) SetValue(value string) {
	db.val = value
}

func (db *PostgresDB) GetValue() string {
	return db.val
}

type MysqlDB struct {
	val []byte
}

func NewMysqlDB() *MysqlDB {
	return &MysqlDB{}
}

func (db *MysqlDB) SetValue(value string) {
	db.val = []byte(value)
}

func (db *MysqlDB) GetValue() string {
	return string(db.val)
}
