package controller

import (
	"Go_Project/test/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// GetAllHandler calls `QueryBikes()` and convert the result as JSON
func GetAllHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var bikes models.Bikes

	bikes = models.QueryBikes()
	var out string
	out = "["
	for i := 0; i < len(bikes.ListBikes)-1; i++ {
		out = out + bikes.ListBikes[i].ToJSON() + ","
	}
	out = out + bikes.ListBikes[len(bikes.ListBikes)-1].ToJSON() + "]"
	fmt.Fprintf(w, string(out))

}
func PostDatahandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var bike models.Bike
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	bike.Id, err = strconv.Atoi(req.Form.Get("id"))
	if err == nil {

	}
	bike.Bike_id = req.Form.Get("bike_id")
	bike.Lat = req.Form.Get("lat")
	bike.Long = req.Form.Get("long")
	bike.Battery = req.Form.Get("battery")
	bike.Date = req.Form.Get("date")
	bike.Sat_count = req.Form.Get("sat_count")
	fmt.Println(bike)
	models.AddBike(bike)
	models.AddBikeToLive(bike)
	models.AddBikeToHistoric(bike)

}
func HistoricSpeedAndDistanceHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var request models.Request
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var id int
	fmt.Println(id)
	res1 := strings.Split(req.Form.Get("ids"), ",")

	for i := 0; i < len(res1); i++ {
		fmt.Println(res1[i])
		id, err = strconv.Atoi(res1[i])
		request.BikeIds = append(request.BikeIds, id)
	}
	request.From = req.Form.Get("from")
	request.To = req.Form.Get("to")

	fmt.Println(request)

	var bikesInfo models.BikesInfo

	bikesInfo = models.GetHistoricSpeedAndDistance(request)

	if len(bikesInfo.ListBikesInfo) == 0 {
		fmt.Fprintf(w, string(""))
	} else {
		var out string
		fmt.Println(bikesInfo)
		out = "["
		for i := 0; i < len(bikesInfo.ListBikesInfo)-1; i++ {
			out = out + bikesInfo.ListBikesInfo[i].ToJSON() + ","
		}
		out = out + bikesInfo.ListBikesInfo[len(bikesInfo.ListBikesInfo)-1].ToJSON() + "]"
		fmt.Fprintf(w, string(out))
	}

}

func PostDatahandlerBodyRequest(w http.ResponseWriter, req *http.Request) {

	b, err := ioutil.ReadAll(req.Body)
	var bike models.Bike

	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// // Unmarshal

	err = json.Unmarshal(b, &bike)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Println(bike)
	models.AddBike(bike)

	models.AddBikeToHistoric(bike)
	models.AddBikeToLive(bike)

}

type Message struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func Test(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	var msg Message

	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// // Unmarshal

	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(msg.Id)
	fmt.Println(msg.Name)
}

func GetTotalHoursHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var request models.Request
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var id int
	fmt.Println(id)
	res1 := strings.Split(req.Form.Get("ids"), ",")

	for i := 0; i < len(res1); i++ {
		fmt.Println(res1[i])
		id, err = strconv.Atoi(res1[i])
		request.BikeIds = append(request.BikeIds, id)
	}
	request.From = req.Form.Get("from")
	request.To = req.Form.Get("to")

	fmt.Println(request)

	var totalHoursForBikes models.TotalHoursForBikes

	totalHoursForBikes = models.GetTotalHours(request)

	if len(totalHoursForBikes.ListTotalHours) == 0 {
		fmt.Fprintf(w, string(""))
	} else {
		var out string
		fmt.Println(totalHoursForBikes)
		out = "["
		for i := 0; i < len(totalHoursForBikes.ListTotalHours)-1; i++ {
			out = out + totalHoursForBikes.ListTotalHours[i].ToJSON() + ","
		}
		out = out + totalHoursForBikes.ListTotalHours[len(totalHoursForBikes.ListTotalHours)-1].ToJSON() + "]"
		fmt.Fprintf(w, string(out))
	}

}

//active bike: sends data in last 10 sec to the server
//return all the active bikes
func ActiveBikesHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var bikes models.BikesInfo

	bikes = models.GetActiveBikes()
	var out string = ""
	if len(bikes.ListBikesInfo) > 0 {
		out = "["
		for i := 0; i < len(bikes.ListBikesInfo)-1; i++ {
			out = out + bikes.ListBikesInfo[i].ToJSON() + ","
		}
		out = out + bikes.ListBikesInfo[len(bikes.ListBikesInfo)-1].ToJSON() + "]"
	}
	fmt.Fprintf(w, string(out))
}
func GetAllSavedHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var request models.Request
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var id int
	fmt.Println(id)
	res1 := strings.Split(req.Form.Get("ids"), ",")

	for i := 0; i < len(res1); i++ {
		fmt.Println(res1[i])
		id, err = strconv.Atoi(res1[i])
		request.BikeIds = append(request.BikeIds, id)
	}
	request.From = req.Form.Get("from")
	request.To = req.Form.Get("to")
	var allSavedForBikes models.AllSavedForBikes

	allSavedForBikes = models.GetAllSaved(request)

	if len(allSavedForBikes.ListAllSavedForBike) == 0 {
		fmt.Fprintf(w, string(""))
	} else {
		var out string
		fmt.Println(allSavedForBikes)
		out = "["
		for i := 0; i < len(allSavedForBikes.ListAllSavedForBike)-1; i++ {
			out = out + allSavedForBikes.ListAllSavedForBike[i].ToJSON() + ","
		}
		out = out + allSavedForBikes.ListAllSavedForBike[len(allSavedForBikes.ListAllSavedForBike)-1].ToJSON() + "]"
		fmt.Fprintf(w, string(out))
	}

}
func SavedCostPerBikeHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var request models.Request
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var id int
	fmt.Println(id)
	res1 := strings.Split(req.Form.Get("ids"), ",")

	for i := 0; i < len(res1); i++ {
		fmt.Println(res1[i])
		id, err = strconv.Atoi(res1[i])
		request.BikeIds = append(request.BikeIds, id)
	}
	request.From = req.Form.Get("from")
	request.To = req.Form.Get("to")

	var totalSavedCostForBikes models.TotalSavedCostForBikes

	totalSavedCostForBikes = models.SavedCost(request)

	if len(totalSavedCostForBikes.ListTotalSavedCost) == 0 {
		fmt.Fprintf(w, string(""))
	} else {
		var out string
		fmt.Println(totalSavedCostForBikes)
		out = "["
		for i := 0; i < len(totalSavedCostForBikes.ListTotalSavedCost)-1; i++ {
			out = out + totalSavedCostForBikes.ListTotalSavedCost[i].ToJSON() + ","
		}
		out = out + totalSavedCostForBikes.ListTotalSavedCost[len(totalSavedCostForBikes.ListTotalSavedCost)-1].ToJSON() + "]"
		fmt.Fprintf(w, string(out))
	}
}
func SavedCarbonPerBikeHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var request models.Request
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var id int
	fmt.Println(id)
	res1 := strings.Split(req.Form.Get("ids"), ",")

	for i := 0; i < len(res1); i++ {
		fmt.Println(res1[i])
		id, err = strconv.Atoi(res1[i])
		request.BikeIds = append(request.BikeIds, id)
	}
	request.From = req.Form.Get("from")
	request.To = req.Form.Get("to")

	var totalSavedCarbonForBikes models.TotalSavedCarbonForBikes

	totalSavedCarbonForBikes = models.SavedCarbon(request)

	if len(totalSavedCarbonForBikes.ListTotalSavedCarbon) == 0 {
		fmt.Fprintf(w, string(""))
	} else {
		var out string
		fmt.Println(totalSavedCarbonForBikes)
		out = "["
		for i := 0; i < len(totalSavedCarbonForBikes.ListTotalSavedCarbon)-1; i++ {
			out = out + totalSavedCarbonForBikes.ListTotalSavedCarbon[i].ToJSON() + ","
		}
		out = out + totalSavedCarbonForBikes.ListTotalSavedCarbon[len(totalSavedCarbonForBikes.ListTotalSavedCarbon)-1].ToJSON() + "]"
		fmt.Fprintf(w, string(out))
	}
}

func PostDatahandlerMessageRequest(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	res1 := strings.Split(params["message"], ",")
	idd := strings.Split(res1[0], "\"")
	fmt.Println("res 0 ", res1[0])
	fmt.Println("last one  ", res1[len(res1)-1])
	ss := strings.Split(res1[len(res1)-1], "\"")
	bike := models.Bike{}
	fmt.Println(idd, ss)
	fmt.Println(params)
	fmt.Println(res1[0], res1[1], res1[2], res1[3], res1[4])
	bike.Bike_id = idd[0]
	bike.Lat = res1[1]
	bike.Long = res1[2]
	bike.Battery = res1[3]

	loc, _ := time.LoadLocation("Asia/Beirut")
	t := time.Now().In(loc)
	fmt.Println("Lebanon time : ", t)
	bike.Date = t.Format("2006-01-02 15:04:05")
	bike.Sat_count = ss[0]
	fmt.Println(bike)
	models.AddBike(bike)
	models.AddBikeToLive(bike)
	models.AddBikeToHistoric(bike)
	fmt.Fprintf(w, string("1"))

}
func HaverSineDistance(w http.ResponseWriter, req *http.Request) {
	var bikeInfo models.BikeInfo
	bikeInfo = models.GetLastRowFromLiveDb("1")
	if bikeInfo.Speed == "0" {
		fmt.Println("It is zero")
	} else {
		fmt.Println("Noo")
	}

}

func AgreegatedHistoricHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	var request models.Request
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	var id int
	fmt.Println(id)
	res1 := strings.Split(req.Form.Get("ids"), ",")

	for i := 0; i < len(res1); i++ {
		fmt.Println(res1[i])
		id, err = strconv.Atoi(res1[i])
		request.BikeIds = append(request.BikeIds, id)
	}
	request.From = req.Form.Get("from")
	request.To = req.Form.Get("to")

	var allAgreegatedHistoric models.AllAgreegatedHistoric

	allAgreegatedHistoric = models.AgreegatedHistorics(request)
	fmt.Println(allAgreegatedHistoric)
	if len(allAgreegatedHistoric.ListAllAgreegatedHistoric) == 0 {
		fmt.Fprintf(w, string(""))
	} else {
		var out string
		fmt.Println(allAgreegatedHistoric)
		out = "["
		for i := 0; i < len(allAgreegatedHistoric.ListAllAgreegatedHistoric)-1; i++ {
			out = out + allAgreegatedHistoric.ListAllAgreegatedHistoric[i].ToJSON() + ","
		}
		out = out + allAgreegatedHistoric.ListAllAgreegatedHistoric[len(allAgreegatedHistoric.ListAllAgreegatedHistoric)-1].ToJSON() + "]"
		fmt.Fprintf(w, string(out))
	}
}
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	fmt.Println("setup Response")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Client Connected")

	models.Reader(ws)

}
