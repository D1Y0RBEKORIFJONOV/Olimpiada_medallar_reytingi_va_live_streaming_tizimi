package connections

import (
	"database/sql"
	"log"
	"medals/internal/clients/athlete"
	athletefunctions "medals/internal/clients/athlete/functions"
	"medals/internal/clients/event"
	eventfunctions "medals/internal/clients/event/functions"
	dbmethods "medals/internal/database/methods"
	databaseservice "medals/internal/database/service"
	interface17 "medals/internal/interface"
	"medals/internal/services"

	_ "github.com/lib/pq"
)

func NewDatabase()(interface17.MedalsService,error){
	db,err:=sql.Open("postgres","postgres://postgres:2005@localhost:5432/hackathon?sslmode=disable")
	if err!=nil{
		return nil,err
	}
	if err:=db.Ping();err!=nil{
		return nil,err
	}
	return &dbmethods.Database{Db: db},nil
}

func NewService()(*databaseservice.DatabaseService,error){
	res,err:=NewDatabase()
	if err!=nil{
		return nil,err
	}
	return &databaseservice.DatabaseService{Psql: res},nil
}

func NewServer()(*services.Server,error){
	res,err:=NewService()
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	athlete:=NewAthlete()
	event:=NewEvent()
	return &services.Server{S: res,A: athlete,E: event},nil
}

func NewAthlete()*athletefunctions.Athlete{
	a:=athlete.UserClinet()
	return &athletefunctions.Athlete{C: a}
}

func NewEvent()*eventfunctions.Event{
	a:=event.EventClient()
	return &eventfunctions.Event{E: a}
}