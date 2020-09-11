package gostraw

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

type Straw struct {
	Driver        string
	ReadDatabase  *gorm.DB
	WriteDatabase *gorm.DB
}

func Init(driver string) *Straw {
	straw := Straw{
		Driver: driver,
	}

	return &straw
}

func (s *Straw) SetWrite(host string, port string, user string, password string, database string) {
	log.Print("initialize write db...")
	db, err := gorm.Open(s.Driver, fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, database))
	if err != nil {
		panic("Failed to connect to write database")
	}
	log.Print("successfuly connected to write db")

	s.WriteDatabase = db

	if s.ReadDatabase == nil {
		s.ReadDatabase = db
	}
}

func (s *Straw) SetRead(host string, port string, user string, password string, database string) {
	log.Print("initialize read db...")
	db, err := gorm.Open(s.Driver, fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, database))
	if err != nil {
		panic("failed to connect to read database")
	}
	log.Print("successfuly connected to read db")

	s.ReadDatabase = db

	if s.WriteDatabase == nil {
		s.WriteDatabase = db
	}
}

func (s *Straw) Where(query interface{}, args ...interface{}) *gorm.DB {
	return s.ReadDatabase.Where(query, args)
}

func (s *Straw) Save(value interface{}) *gorm.DB {
	return s.WriteDatabase.Save(value)
}

func (s *Straw) Create(value interface{}) *gorm.DB {
	return s.WriteDatabase.Create(value)
}

func (s *Straw) Exec(sql string, values ...interface{}) *gorm.DB {
	return s.WriteDatabase.Exec(sql, values)
}

func (s *Straw) Conn() *gorm.DB {
	return s.WriteDatabase
}
