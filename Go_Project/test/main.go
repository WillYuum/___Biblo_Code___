package main

import (
	"Go_Project/test/controller"
	"Go_Project/test/database"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readFile() {
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		res1 := strings.Split(scanner.Text(), ",")
		message := map[string]interface{}{
			"bike_id":   res1[0],
			"lat":       res1[1],
			"long":      res1[2],
			"battery":   res1[3],
			"sat_count": res1[4],
			"date":      res1[5],
		}

		bytesRepresentation, err := json.Marshal(message)
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Post("http://23.254.225.235:10000/add", "application/json", bytes.NewBuffer(bytesRepresentation))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp)
	}

}

func main() {
	fmt.Println("hello wolrd")
	database.InitDb()
	handleRequests()
	//readFile()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//get historic information for list bike id during timestamp
	myRouter.HandleFunc("/getHistoric", controller.HistoricSpeedAndDistanceHandler).Methods("POST")
	//add bike by sending json object in body request
	myRouter.HandleFunc("/add", controller.PostDatahandlerBodyRequest).Methods("POST")
	//for testing
	myRouter.HandleFunc("/test", controller.Test).Methods("POST")
	//get all bikes from mainDB without speed and distance
	myRouter.HandleFunc("/getAll", controller.GetAllHandler)
	//get active bikes
	myRouter.HandleFunc("/getActiveBikes", controller.ActiveBikesHandler)
	//get saved cost for list bike id during timestamp
	myRouter.HandleFunc("/savedCost", controller.SavedCostPerBikeHandler)
	//get saved CO2 for list bike id during timestamp
	myRouter.HandleFunc("/savedCarbon", controller.SavedCarbonPerBikeHandler)
	//total hours used for list bike id during timestamp
	myRouter.HandleFunc("/totalHours", controller.GetTotalHoursHandler)
	//myRouter.HandleFunc("/d", controller.HaverSineDistance)

	myRouter.HandleFunc("/saved", controller.GetAllSavedHandler)

	myRouter.HandleFunc("/ws", controller.WSEndpoint) // websocket link
	//WebSocket Used On FrontEndJs

	//add a row
	myRouter.HandleFunc("/add/{message}", controller.PostDatahandlerMessageRequest)

	//get agreegated distance and speed for list bike id during timestamp
	myRouter.HandleFunc("/agreegatedHistoric", controller.AgreegatedHistoricHandler)

	log.Fatal(http.ListenAndServe(":10000", myRouter))

}
