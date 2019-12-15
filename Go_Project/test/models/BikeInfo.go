package models

import (
	"Go_Project/test/database"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/umahmood/haversine"
)

type BikeInfo struct {
	Id        int    `json:"_id"`
	Bike_id   string `json:"bike_id"`
	Lat       string `json:"lat"`
	Long      string `json:"long"`
	Battery   string `json:"battery"`
	Sat_count string `json:"sat_count"`
	Date      string `json:"date"`
	Speed     string `json:"speed"`
	Distance  string `json:"distance"`
}

type BikeData struct {
	id        int32
	lat       string
	long      string
	battery   string
	sat_count string
	date      string
	bike_id   string
}
type BikesInfo struct {
	ListBikesInfo []BikeInfo
}

func (e BikeInfo) ToJSON() string {
	return fmt.Sprintf(`{"_id": "%s","bike_id": "%s","lat": "%s","long": "%s","battery": "%s","sat_count": "%s","date": "%s","speed": "%s","distance": "%s"}`, e.Id, e.Bike_id, e.Lat, e.Long, e.Battery, e.Sat_count, e.Date, e.Speed, e.Distance)
}
func AddBikeToHistoric(bike Bike) {
	sql := `
	INSERT INTO public."Historic_Raw_Bike"(
		"Bike_ID", "Latitude", "Longitude", "DateTime", "Sat_Count", "Speed", "Distance", "Battery")
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	var bikeHInfo BikeInfo
	bikeHInfo = GetLastRowFromHistoricDb(bike.Bike_id)

	if bikeHInfo.Speed == "0" && bikeHInfo.Id == -1 {

		database.HistoricDb.Exec(sql, bike.Bike_id, bike.Lat, bike.Long, bike.Date, bike.Sat_count, 0, 0, bike.Battery)

	} else {
		bikee := GetLastRowFromMainDb(bike.Bike_id)

		lat1, err := strconv.ParseFloat(bike.Lat, 32)
		long1, err := strconv.ParseFloat(bike.Long, 32)
		lat2, err := strconv.ParseFloat(bikeHInfo.Lat, 32)
		long2, err := strconv.ParseFloat(bikeHInfo.Long, 32)

		date1, err := time.Parse(time.RFC3339, bikeHInfo.Date)
		date2, err := time.Parse(time.RFC3339, bikee.Date)

		// minutes := date2.Sub(date1).Minutes()
		// hours := date2.Sub(date1).Hours()
		// seconds := date2.Sub(date1).Seconds()
		// var t float64 = hours + minutes/60
		loc1 := haversine.Coord{Lat: lat1, Lon: long1} // Oxford, UK
		loc2 := haversine.Coord{Lat: lat2, Lon: long2} // Turin, Italy
		mi, _ := haversine.Distance(loc1, loc2)
		var speed float64 = 0
		fmt.Println("__________________")
		fmt.Println(date1)
		fmt.Println(date2)
		hr1, min1, sec1 := date1.Clock()
		hr2, min2, sec2 := date2.Clock()

		// minutes := min2 - min1
		// seconds := sec2 - sec1
		hours := hr2 - hr1
		totalTime := (min2*60 + sec2) - (min1*60 + sec1)
		f := float64(totalTime)

		fmt.Println(totalTime)
		if totalTime > 0 && totalTime <= 30 && hours == 0 {
			speed = mi / f
		} else {
			mi = 0
			speed = 0
		}
		_, err = database.HistoricDb.Exec(sql, bike.Bike_id, bike.Lat, bike.Long, bike.Date, bike.Sat_count, speed, mi, bike.Battery)

		if err != nil {
			panic(err)
		}
	}

}
func AddBikeToLive(bike Bike) {
	sqlStatement := `
INSERT INTO public."Live_Raw_Bike"(
	"Bike_ID", "Latitude", "Longitude", "DateTime", "Sat_Count", "Speed", "Distance", "Battery")
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`

	var bikeInfo BikeInfo

	bikeInfo = GetLastRowFromLiveDb(bike.Bike_id)
	if bikeInfo.Speed == "0" && bikeInfo.Id == -1 {

		_, err := database.LiveDb.Exec(sqlStatement, bike.Bike_id, bike.Lat, bike.Long, bike.Date, bike.Sat_count, 0, 0, bike.Battery)
		if err != nil {
			panic(err)
		}
	} else {

		bikee := GetLastRowFromMainDb(bike.Bike_id)

		lat1, err := strconv.ParseFloat(bike.Lat, 32)
		long1, err := strconv.ParseFloat(bike.Long, 32)
		lat2, err := strconv.ParseFloat(bikeInfo.Lat, 32)
		long2, err := strconv.ParseFloat(bikeInfo.Long, 32)

		date1, err := time.Parse(time.RFC3339, bikeInfo.Date)
		date2, err := time.Parse(time.RFC3339, bikee.Date)

		// minutes := date2.Sub(date1).Minutes()
		// hours := date2.Sub(date1).Hours()
		// seconds := date2.Sub(date1).Seconds()
		// var t float64 = hours + minutes/60
		loc1 := haversine.Coord{Lat: lat1, Lon: long1} // Oxford, UK
		loc2 := haversine.Coord{Lat: lat2, Lon: long2} // Turin, Italy
		mi, _ := haversine.Distance(loc1, loc2)
		var speed float64 = 0
		fmt.Println("__________________")
		fmt.Println(date1)
		fmt.Println(date2)
		hr1, min1, sec1 := date1.Clock()
		hr2, min2, sec2 := date2.Clock()

		// minutes := min2 - min1
		// seconds := sec2 - sec1
		hours := hr2 - hr1
		totalTime := (min2*60 + sec2) - (min1*60 + sec1)
		f := float64(totalTime)

		fmt.Println(totalTime)
		if totalTime > 0 && totalTime <= 30 && hours == 0 {
			speed = mi / f
		} else {
			mi = 0
			speed = 0
		}
		_, err = database.LiveDb.Exec(sqlStatement, bike.Bike_id, bike.Lat, bike.Long, bike.Date, bike.Sat_count, speed, mi, bike.Battery)

		if err != nil {
			fmt.Println("Error", err)
		}
		//fmt.Println(bikeInfo.Date, bikee.Date)
		fmt.Println(mi)

	}

	fmt.Println("Successfully added")
}

func GetLastRowFromLiveDb(bike_id string) BikeInfo {
	bike := BikeInfo{}
	bike.Speed = "0"
	bike.Id = -1
	sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Live_Raw_Bike" where "Bike_ID" like $1 ORDER BY "ID" DESC`
	rows, err := database.LiveDb.Query(sql, bike_id)
	if err != nil {
		return bike
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&bike.Id,
			&bike.Bike_id,
			&bike.Lat,
			&bike.Long,
			&bike.Sat_count,
			&bike.Speed,
			&bike.Distance,
			&bike.Battery,
			&bike.Date,
		)
		return bike

	}
	return bike
}

func GetLastRowFromHistoricDb(bike_id string) BikeInfo {

	bike := BikeInfo{}
	bike.Speed = "0"
	bike.Id = -1
	sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Historic_Raw_Bike" where "Bike_ID" like $1 ORDER BY "ID" DESC`
	rows, err := database.HistoricDb.Query(sql, bike_id)
	if err != nil {
		return bike
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&bike.Id,
			&bike.Bike_id,
			&bike.Lat,
			&bike.Long,
			&bike.Sat_count,
			&bike.Speed,
			&bike.Distance,
			&bike.Battery,
			&bike.Date,
		)
		return bike

	}
	return bike
}

func GetLastRowFromMainDb(bike_id string) Bike {

	bike := Bike{}
	sql := `SELECT "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", battery, "DateTime" FROM public."Raw_Bike" where "Bike_ID" like $1 ORDER BY "ID" DESC`
	rows, err := database.MainDb.Query(sql, bike_id)
	if err != nil {
		return bike
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(
			&bike.Id,
			&bike.Bike_id,
			&bike.Lat,
			&bike.Long,
			&bike.Sat_count,
			&bike.Battery,
			&bike.Date,
		)

		fmt.Println(bike)
		return bike

	}
	return bike
}
func GetHistoricSpeedAndDistance(request Request) BikesInfo {
	bikesInfo := BikesInfo{}

	for i := 0; i < len(request.BikeIds); i++ {
		fmt.Println("'" + strconv.Itoa(request.BikeIds[i]) + "'")
		sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Historic_Raw_Bike" where "Bike_ID" like $1 AND "DateTime" BETWEEN $2 AND $3`
		rows, err := database.HistoricDb.Query(sql, strconv.Itoa(request.BikeIds[i]), request.From, request.To)
		fmt.Println(err)

		if err != nil {
			return bikesInfo
		}
		defer rows.Close()
		for rows.Next() {

			bikeInfo := BikeInfo{}
			err = rows.Scan(
				&bikeInfo.Id,
				&bikeInfo.Bike_id,
				&bikeInfo.Lat,
				&bikeInfo.Long,
				&bikeInfo.Sat_count,
				&bikeInfo.Speed,
				&bikeInfo.Distance,
				&bikeInfo.Battery,
				&bikeInfo.Date,
			)
			fmt.Println(bikeInfo)
			bikesInfo.ListBikesInfo = append(bikesInfo.ListBikesInfo, bikeInfo)

		}
	}
	return bikesInfo

}

func GetTotalHours(request Request) TotalHoursForBikes {
	totalHoursForBikes := TotalHoursForBikes{}
	fmt.Println(request.BikeIds)
	for i := 0; i < len(request.BikeIds); i++ {
		totalhours := TotalHours{}
		var hours float64 = 0
		fmt.Println("'" + strconv.Itoa(request.BikeIds[i]) + "'")
		sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Historic_Raw_Bike" where "Bike_ID" like $1 AND "DateTime" BETWEEN $2 AND $3`
		rows, err := database.HistoricDb.Query(sql, strconv.Itoa(request.BikeIds[i]), request.From, request.To)

		if err != nil {
			return totalHoursForBikes
		}
		defer rows.Close()
		totalhours.BikeId = strconv.Itoa(request.BikeIds[i])
		for rows.Next() {
			bikeInfo := BikeInfo{}
			err = rows.Scan(
				&bikeInfo.Id,
				&bikeInfo.Bike_id,
				&bikeInfo.Lat,
				&bikeInfo.Long,
				&bikeInfo.Sat_count,
				&bikeInfo.Speed,
				&bikeInfo.Distance,
				&bikeInfo.Battery,
				&bikeInfo.Date,
			)
			spe, _ := strconv.ParseFloat(bikeInfo.Speed, 64)

			if spe != 0 {
				dis, _ := strconv.ParseFloat(bikeInfo.Distance, 64)
				sec := dis / spe
				hrs := sec / 3600
				hours = hours + hrs
			}

			totalhours.BikeId = bikeInfo.Bike_id
		}
		fmt.Println(hours)

		totalhours.Hours = hours
		totalHoursForBikes.ListTotalHours = append(totalHoursForBikes.ListTotalHours, totalhours)
	}
	return totalHoursForBikes
}
func GetAllSaved(request Request) AllSavedForBikes {
	allSavedForBikes := AllSavedForBikes{}
	fmt.Println(request.BikeIds)
	for i := 0; i < len(request.BikeIds); i++ {
		allSavedForBike := AllSavedForBike{}
		var hours float64 = 0
		fmt.Println("'" + strconv.Itoa(request.BikeIds[i]) + "'")
		sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Historic_Raw_Bike" where "Bike_ID" like $1 AND "DateTime" BETWEEN $2 AND $3`
		rows, err := database.HistoricDb.Query(sql, strconv.Itoa(request.BikeIds[i]), request.From, request.To)

		if err != nil {
			return allSavedForBikes
		}
		defer rows.Close()
		allSavedForBike.BikeId = strconv.Itoa(request.BikeIds[i])
		for rows.Next() {
			bikeInfo := BikeInfo{}
			err = rows.Scan(
				&bikeInfo.Id,
				&bikeInfo.Bike_id,
				&bikeInfo.Lat,
				&bikeInfo.Long,
				&bikeInfo.Sat_count,
				&bikeInfo.Speed,
				&bikeInfo.Distance,
				&bikeInfo.Battery,
				&bikeInfo.Date,
			)
			spe, _ := strconv.ParseFloat(bikeInfo.Speed, 64)

			if spe != 0 {
				dis, _ := strconv.ParseFloat(bikeInfo.Distance, 64)
				sec := dis / spe
				hrs := sec / 3600
				hours = hours + hrs
			}
		}
		fmt.Println(hours)
		allSavedForBike.BikeId = strconv.Itoa(request.BikeIds[i])
		allSavedForBike.Distance = GetTotalDistancePerBike(strconv.Itoa(request.BikeIds[i]), request.From, request.To)
		allSavedForBike.SavedCarbon = GetTotalDistancePerBike(strconv.Itoa(request.BikeIds[i]), request.From, request.To) * 45
		allSavedForBike.SavedCost = hours * 10
		allSavedForBike.Hours = hours
		allSavedForBikes.ListAllSavedForBike = append(allSavedForBikes.ListAllSavedForBike, allSavedForBike)
	}
	return allSavedForBikes

}
func GetTotalDistancePerBike(bike_id string, from string, to string) float64 {
	var dist float64 = 0
	sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Historic_Raw_Bike" where "Bike_ID" like $1 AND "DateTime" BETWEEN $2 AND $3`
	rows, err := database.HistoricDb.Query(sql, bike_id, from, to)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		bikeInfo := BikeInfo{}
		err = rows.Scan(
			&bikeInfo.Id,
			&bikeInfo.Bike_id,
			&bikeInfo.Lat,
			&bikeInfo.Long,
			&bikeInfo.Sat_count,
			&bikeInfo.Speed,
			&bikeInfo.Distance,
			&bikeInfo.Battery,
			&bikeInfo.Date,
		)
		spe, _ := strconv.ParseFloat(bikeInfo.Speed, 64)

		if spe != 0 {
			dis, _ := strconv.ParseFloat(bikeInfo.Distance, 64)
			dist = dist + dis
		}

	}
	fmt.Println("dist ", dist)
	return dist
}
func GetActiveBikes() BikesInfo {
	bike := BikeInfo{}
	bikesInfo := BikesInfo{}
	bike.Speed = "0"
	bike.Id = -1
	sql := `SELECT "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "DateTime", "Battery"
	FROM public."Raw_Speed_Data_Historic" WHERE ("Bike_ID","DateTime") in (SELECT "Bike_ID",max("DateTime") FROM public."Historic_Raw_Bike" group by "Bike_ID")`

	rows, err := database.HistoricDb.Query(sql)
	if err != nil {
		return bikesInfo
	}
	now := time.Now()
	_, min1, sec1 := now.Clock()

	fmt.Println("Now: ", now)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&bike.Id,
			&bike.Bike_id,
			&bike.Lat,
			&bike.Long,
			&bike.Sat_count,
			&bike.Speed,
			&bike.Distance,
			&bike.Date,
			&bike.Battery,
		)
		date, _ := time.Parse(time.RFC3339, bike.Date)
		_, min2, sec2 := date.Clock()
		totalTime := (min1*60 + sec1) - (min2*60 + sec2)
		fmt.Println("TotalTime in live: ", totalTime)
		if totalTime > 0 && totalTime < 30 {
			bikesInfo.ListBikesInfo = append(bikesInfo.ListBikesInfo, bike)
		}

	}
	return bikesInfo
}
func SavedCost(request Request) TotalSavedCostForBikes {
	totalSavedCostForBikes := TotalSavedCostForBikes{}

	totalHoursForBikes := GetTotalHours(request)

	for index := 0; index < len(totalHoursForBikes.ListTotalHours); index++ {
		totalSavedCost := TotalSavedCost{}
		totalSavedCost.BikeId = totalHoursForBikes.ListTotalHours[index].BikeId
		totalSavedCost.SavedCost = totalHoursForBikes.ListTotalHours[index].Hours * 10
		totalSavedCostForBikes.ListTotalSavedCost = append(totalSavedCostForBikes.ListTotalSavedCost, totalSavedCost)

	}
	return totalSavedCostForBikes
}
func SavedCarbon(request Request) TotalSavedCarbonForBikes {
	totalSavedCarbonForBikes := TotalSavedCarbonForBikes{}
	for i := 0; i < len(request.BikeIds); i++ {
		totalSavedCarbon := TotalSavedCarbon{}
		totalSavedCarbon.BikeId = strconv.Itoa(request.BikeIds[i])
		totalSavedCarbon.SavedCarbon = GetTotalDistancePerBike(strconv.Itoa(request.BikeIds[i]), request.From, request.To) * 45

		totalSavedCarbonForBikes.ListTotalSavedCarbon = append(totalSavedCarbonForBikes.ListTotalSavedCarbon, totalSavedCarbon)

	}
	return totalSavedCarbonForBikes
}
func Distance() {
	oxford := haversine.Coord{Lat: 51.45, Lon: 1.15} // Oxford, UK
	turin := haversine.Coord{Lat: 45.04, Lon: 7.42}  // Turin, Italy
	mi, km := haversine.Distance(oxford, turin)
	fmt.Println("Miles:", mi, "Kilometers:", km)
}
func AgreegatedHistorics(request Request) AllAgreegatedHistoric {
	allAgreegatedHistoric := AllAgreegatedHistoric{}
	//fmt.Println(request.BikeIds)
	agreegatedHistoric := AgreegatedHistoric{"", 0, 0, 0, ""}
	dateFrom, _ := time.Parse("2006-01-02 15:04:05", request.From)
	dateTo, _ := time.Parse("2006-01-02 15:04:05", request.To)
	hrsFrom, _, _ := dateFrom.Clock()
	hrsTo, _, _ := dateTo.Clock()
	//example from 2019-10-23 05:21:50 to 2019-10-23 17:21:50
	// from hr 5 to hr 17 it is ok
	if hrsTo > hrsFrom {
		for index := hrsFrom; index < hrsTo; index++ {

			var r Request
			r.BikeIds = request.BikeIds
			r.From = strings.Split(request.From, " ")[0] + " " + strconv.Itoa(index) + ":00:00"
			r.To = strings.Split(request.To, " ")[0] + " " + strconv.Itoa(index+1) + ":00:00"
			agreegatedHistoric = getAverageDistanceDuringHour(r)
			agreegatedHistoric.Hours = index
			agreegatedHistoric.Date = string(r.From)
			allAgreegatedHistoric.ListAllAgreegatedHistoric = append(allAgreegatedHistoric.ListAllAgreegatedHistoric, agreegatedHistoric)
			fmt.Println("agreegatedHistoric : ", agreegatedHistoric)
		}
		//but if from hrs 17 on day 24 to hr 5 on day 25
		// wee need to divide it for 2 parts
	} else {
		for index := hrsFrom; index < 24; index++ {
			var r Request
			r.BikeIds = request.BikeIds
			r.From = strings.Split(request.From, " ")[0] + " " + strconv.Itoa(index) + ":00:00"
			r.To = strings.Split(request.To, " ")[0] + " " + strconv.Itoa(index+1) + ":00:00"
			agreegatedHistoric = getAverageDistanceDuringHour(r)
			agreegatedHistoric.Hours = index
			agreegatedHistoric.Date = string(r.From)
			allAgreegatedHistoric.ListAllAgreegatedHistoric = append(allAgreegatedHistoric.ListAllAgreegatedHistoric, agreegatedHistoric)
			fmt.Println("agreegatedHistoric : ", agreegatedHistoric)
		}
		for index := 1; index < hrsTo; index++ {
			var r Request
			r.BikeIds = request.BikeIds
			r.From = strings.Split(request.From, " ")[0] + " " + strconv.Itoa(index) + ":00:00"
			r.To = strings.Split(request.To, " ")[0] + " " + strconv.Itoa(index+1) + ":00:00"
			agreegatedHistoric = getAverageDistanceDuringHour(r)
			agreegatedHistoric.Hours = index
			allAgreegatedHistoric.ListAllAgreegatedHistoric = append(allAgreegatedHistoric.ListAllAgreegatedHistoric, agreegatedHistoric)
			fmt.Println("agreegatedHistoric : ", agreegatedHistoric)
		}
	}

	return allAgreegatedHistoric
}
func getAverageDistanceDuringHour(request Request) AgreegatedHistoric {
	var agreegatedHistoric AgreegatedHistoric
	var avg float64 = 0
	k := 0
	ids := ""
	index := false
	for i := 0; i < len(request.BikeIds); i++ {
		index = true
		fmt.Println(request)
		sql := `SELECT  "ID", "Bike_ID", "Latitude", "Longitude", "Sat_Count", "Speed", "Distance", "Battery", "DateTime" FROM public."Historic_Raw_Bike" where "Bike_ID" like $1 AND "DateTime"::date BETWEEN $2 AND $3 AND "DateTime"::time BETWEEN $4 AND $5`
		rows, err := database.HistoricDb.Query(sql, strconv.Itoa(request.BikeIds[i]), strings.Split(request.From, " ")[0], strings.Split(request.To, " ")[0], strings.Split(request.From, " ")[1], strings.Split(request.To, " ")[1])

		if err != nil {
			return agreegatedHistoric
		}
		defer rows.Close()

		for rows.Next() {

			bikeInfo := BikeInfo{}
			k++
			err = rows.Scan(
				&bikeInfo.Id,
				&bikeInfo.Bike_id,
				&bikeInfo.Lat,
				&bikeInfo.Long,
				&bikeInfo.Sat_count,
				&bikeInfo.Speed,
				&bikeInfo.Distance,
				&bikeInfo.Battery,
				&bikeInfo.Date,
			)
			if index {
				ids += bikeInfo.Bike_id
				index = false
			}
			//fmt.Println(bikeInfo)
			dis, _ := strconv.ParseFloat(bikeInfo.Distance, 64)
			avg = avg + dis

		}

	}
	fmt.Println("avg : ", avg)
	count := float64(k)
	agreegatedHistoric.BikeId = ids
	agreegatedHistoric.Distance = avg / count
	if count == 0 {
		agreegatedHistoric.Distance = 0
	}

	return agreegatedHistoric
}

var err error

func GetHistoric(date1 string, date2 string, conn *websocket.Conn) {
	BikesData := []BikeData{}
	var data [][]BikeData
	for i := 1; i <= 100; i++ {

		BikesData = nil
		sql := `SELECT "ID" , "Latitude" , "Longitude" , "Battery" , "Sat_Count" , "DateTime"  , "Bike_ID" FROM public."Historic_Raw_Bike" WHERE "DateTime" > $1 AND "DateTime" < $2 AND "Bike_ID" = $3` // sql statement
		if err != nil {

		}
		//parse date into go format
		date1_parse, err := time.Parse("Mon Jan 2 2006 15:04:05 ", date1)
		if err != nil {

		}
		//parse date 2 into go format
		date2_parse, err := time.Parse("Mon Jan 2 2006 15:04:05 ", date2)
		if err != nil {

		}
		//get al rows
		rows, err := database.HistoricDb.Query(sql, date1_parse, date2_parse, i)
		if err != nil {

		}

		for rows.Next() {
			bike_data := &BikeData{}
			bike_data.id = 0
			bike_data.lat = ""
			bike_data.long = ""
			bike_data.battery = ""
			bike_data.sat_count = ""
			bike_data.date = ""
			bike_data.bike_id = ""

			err = rows.Scan(&bike_data.id, &bike_data.lat, &bike_data.long, &bike_data.battery, &bike_data.sat_count, &bike_data.date, &bike_data.bike_id)
			//append data into a form of [of bikes ] [of data] [of interface of data or structure]
			if err != nil {

				continue
			}
			BikesData = append(BikesData, *bike_data)
		}
		data = append(data, BikesData)
	}
	result := []map[string]interface{}{}
	results := [][]map[string]interface{}{}
	for _, bike := range data {
		for _, mybike := range bike {
			if &mybike != nil {
				result = append(result, mybike.ToMap())
			}
		}
		results = append(results, result)
		if len(bike) != 0 {
			result = nil
		}
	}
	updateJson, _ := json.Marshal(results)

	msg := string(updateJson)

	if err = conn.WriteMessage(1, []byte("h"+msg)); err != nil {
		fmt.Println(err)
	}
}

func GetDynamic(date1 string, date2 string, conn *websocket.Conn) {
	BikesData := []BikeData{}
	var data [][]BikeData
	for i := 1; i <= 100; i++ {

		BikesData = nil
		sql := `SELECT "ID" , "Latitude" , "Longitude" , "Battery" , "Sat_Count" , "DateTime" , "Bike_ID" FROM public."Historic_Raw_Bike" WHERE "DateTime" > $1 AND "DateTime" < $2 AND "Bike_ID" = $3` // sql statement
		if err != nil {

		}
		//parse date into go format
		date1_parse, err := time.Parse("Mon Jan 2 2006 15:04:05 ", date1)
		if err != nil {

		}
		//parse date 2 into go format
		date2_parse, err := time.Parse("Mon Jan 2 2006 15:04:05 ", date2)
		if err != nil {

		}
		//get al rows
		rows, err := database.HistoricDb.Query(sql, date1_parse, date2_parse, i)
		if err != nil {

		}

		for rows.Next() {
			bike_data := &BikeData{}
			bike_data.id = 0
			bike_data.lat = ""
			bike_data.long = ""
			bike_data.battery = ""
			bike_data.sat_count = ""
			bike_data.date = ""
			bike_data.bike_id = ""

			err = rows.Scan(&bike_data.id, &bike_data.lat, &bike_data.long, &bike_data.battery, &bike_data.sat_count, &bike_data.date, &bike_data.bike_id)
			//append data into a form of [of bikes ] [of data] [of interface of data or structure]
			if err != nil {

				continue
			}
			BikesData = append(BikesData, *bike_data)
		}
		data = append(data, BikesData)
	}
	result := []map[string]interface{}{}
	results := [][]map[string]interface{}{}
	for _, bike := range data {
		for _, mybike := range bike {
			if &mybike != nil {
				result = append(result, mybike.ToMap())
			}
		}
		results = append(results, result)
		if len(bike) != 0 {
			result = nil
		}
	}
	updateJson, _ := json.Marshal(results)

	msg := string(updateJson)

	if err = conn.WriteMessage(1, []byte("DynamicMarker - "+msg)); err != nil {
		fmt.Println(err)
	}
}

func InitializeData(conn *websocket.Conn) {
	data := "["
	for i := 1; i <= 100; i++ {
		queryStmt, err := database.LiveDb.Prepare(` SELECT "ID" , "Latitude" , "Longitude" , "Battery" , "Sat_Count" , "DateTime" FROM public."Live_Raw_Bike" WHERE "Bike_ID"=$1 order by "ID" desc limit 1`)
		if err != nil {

		}
		var id int32
		var lat string
		var long string
		var battery string
		var sat_count string
		var date string
		err = queryStmt.QueryRow(i).Scan(&id, &lat, &long, &battery, &sat_count, &date)

		if id == 0 {
			if err = conn.WriteMessage(1, []byte("Device "+strconv.Itoa(i)+" IS OFF!")); err != nil {
				fmt.Println(err)

			}
		} else {
			data += "[" + strconv.Itoa(i) + "," + lat + "," + long + "]" + "," + battery

		}
	}
	data += "]"
	if err = conn.WriteMessage(1, []byte(data)); err != nil {
		fmt.Println(err)
	}

}
func (this *BikeData) ToMap() map[string]interface{} {
	result := map[string]interface{}{}
	if this.id != 0 {
		result["id"] = this.id
	}
	if this.lat != "" {
		result["lat"] = this.lat
	}
	if this.long != "" {
		result["long"] = this.long
	}
	if this.battery != "" {
		result["battery"] = this.battery
	}
	if this.sat_count != "" {
		result["sat_count"] = this.sat_count
	}
	if this.date != "" {
		result["date"] = this.date
	}
	if this.bike_id != "" {
		result["bike_id"] = this.bike_id
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
