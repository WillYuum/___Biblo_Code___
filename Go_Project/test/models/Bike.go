package models

import (
	"Go_Project/test/database"
	"fmt"
)

type Bike struct {
	Id        int    `json:"_id"`
	Bike_id   string `json:"bike_id"`
	Lat       string `json:"lat"`
	Long      string `json:"long"`
	Battery   string `json:"battery"`
	Sat_count string `json:"sat_count"`
	Date      string `json:"date"`
}

type Bikes struct {
	ListBikes []Bike
}

func (e Bike) ToJSON() string {
	return fmt.Sprintf(`{"id": "%d","bike_id": "%s","lat": "%s","long": "%s","battery": "%s","sat_count": "%s","date": "%s"}`, e.Id, e.Bike_id, e.Lat, e.Long, e.Battery, e.Sat_count, e.Date)
}

func QueryBikes() Bikes {
	bikes := Bikes{}
	rows, err := database.MainDb.Query(`
		SELECT "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", battery, "DateTime" FROM public."Raw_Bike" `)
	if err != nil {
		return bikes
	}
	defer rows.Close()
	for rows.Next() {
		bike := Bike{}
		err = rows.Scan(
			&bike.Id,
			&bike.Bike_id,
			&bike.Lat,
			&bike.Long,
			&bike.Sat_count,
			&bike.Battery,
			&bike.Date,
		)
		if err != nil {
			return bikes
		}
		bikes.ListBikes = append(bikes.ListBikes, bike)
	}
	err = rows.Err()
	if err != nil {
		return bikes
	}
	return bikes
}

func AddBike(bike Bike) {
	sqlStatement := `
INSERT INTO public."Raw_Bike"(
	 "Bike_ID", "Latitude", "Longitude", "DateTime", "Sat_Count", battery)
	VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := database.MainDb.Exec(sqlStatement, bike.Bike_id, bike.Lat, bike.Long, bike.Date, bike.Sat_count, bike.Battery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully added")
}
