package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	maindbhost = "localhost"
	maindbport = "5432"
	maindbuser = "postgres"
	maindbpass = "123456"
	maindbname = "Main_Bike_BD"

	historicdbhost = "localhost"
	historicdbport = "5432"
	historicdbuser = "postgres"
	historicdbpass = "123456"
	historicdbname = "Historic_Bike_BD"

	livedbhost = "localhost"
	livedbport = "5432"
	livedbuser = "postgres"
	livedbpass = "123456"
	livedbname = "Live_Bike_BD"

	// maindbhost = "localhost"
	// maindbport = "5432"
	// maindbuser = "postgres"
	// maindbpass = "decentra-tech-db"
	// maindbname = "Main_Bike_DB"

	// historicdbhost = "localhost"
	// historicdbport = "5432"
	// historicdbuser = "postgres"
	// historicdbpass = "decentra-tech-db"
	// historicdbname = "Historic_Bike_DB"

	// livedbhost = "localhost"
	// livedbport = "5432"
	// livedbuser = "postgres"
	// livedbpass = "decentra-tech-db"
	// livedbname = "Live_Bike_DB"
)

var MainDb *sql.DB
var HistoricDb *sql.DB
var LiveDb *sql.DB

func InitDb() {
	config := dbConfig()
	var err error

	//open connection on MainDb
	mainPsqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[maindbhost], config[maindbport],
		config[maindbuser], config[maindbpass], config[maindbname])

	MainDb, err = sql.Open("postgres", mainPsqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	err = MainDb.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MainDb Successfully connected!")

	//open connection on HistoricDb
	HistoricPsqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[historicdbhost], config[historicdbport],
		config[historicdbuser], config[historicdbpass], config[historicdbname])

	HistoricDb, err = sql.Open("postgres", HistoricPsqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	err = HistoricDb.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("HistoricDb Successfully connected!")

	//open connection on LiveDb
	LivePsqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[livedbhost], config[livedbport],
		config[livedbuser], config[livedbpass], config[livedbname])

	LiveDb, err = sql.Open("postgres", LivePsqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	err = LiveDb.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("LiveDb Successfully connected!")
}

func dbConfig() map[string]string {

	conf := make(map[string]string)
	conf[maindbhost] = maindbhost
	conf[maindbport] = maindbport
	conf[maindbuser] = maindbuser
	conf[maindbpass] = maindbpass
	conf[maindbname] = maindbname

	conf[historicdbhost] = historicdbhost
	conf[historicdbport] = historicdbport
	conf[historicdbuser] = historicdbuser
	conf[historicdbpass] = historicdbpass
	conf[historicdbname] = historicdbname

	conf[livedbhost] = livedbhost
	conf[livedbport] = livedbport
	conf[livedbuser] = livedbuser
	conf[livedbpass] = livedbpass
	conf[livedbname] = livedbname
	return conf
}
