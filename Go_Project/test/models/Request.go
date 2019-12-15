package models

import "fmt"

type Request struct {
	From    string `json:"from"`
	To      string `json:"to"`
	BikeIds []int  `json:"ids"`
}

type TotalHours struct {
	BikeId string  `json:"bike"`
	Hours  float64 `json:"hours"`
}
type TotalSavedCost struct {
	BikeId    string  `json:"bike"`
	SavedCost float64 `json:"SavedCost"`
}
type TotalSavedCarbon struct {
	BikeId      string  `json:"bike"`
	SavedCarbon float64 `json:"SavedCarbon"`
}
type TotalHoursForBikes struct {
	ListTotalHours []TotalHours
}
type AllSavedForBike struct {
	BikeId      string  `json:"bike"`
	SavedCarbon float64 `json:"SavedCarbon"`
	SavedCost   float64 `json:"SavedCost"`
	Hours       float64 `json:"hours"`
	Distance    float64 `json:"distance"`
}
type AllSavedForBikes struct {
	ListAllSavedForBike []AllSavedForBike
}
type TotalSavedCostForBikes struct {
	ListTotalSavedCost []TotalSavedCost
}
type TotalSavedCarbonForBikes struct {
	ListTotalSavedCarbon []TotalSavedCarbon
}
type AgreegatedHistoric struct {
	BikeId   string  `json:"bike"`
	Speed    float64 `json:"speed"`
	Hours    int     `json:"hours"`
	Distance float64 `json:"distance"`
	Date     string  `json:"date"`
}
type AllAgreegatedHistoric struct {
	ListAllAgreegatedHistoric []AgreegatedHistoric
}

func (e AgreegatedHistoric) ToJSON() string {
	return fmt.Sprintf(`{"bike_id": "%s","Hours": "%d","Speed": "%f","Distance": "%f","Date":"%s"}`, e.BikeId, e.Hours, e.Speed, e.Distance, e.Date)
}

func (e AllSavedForBike) ToJSON() string {
	return fmt.Sprintf(`{"bike_id": "%s","Hours": "%f","SavedCost": "%f","SavedCarbon": "%f","Distance": "%f"}`, e.BikeId, e.Hours, e.SavedCost, e.SavedCarbon, e.Distance)
}

func (e TotalHours) ToJSON() string {
	return fmt.Sprintf(`{"bike_id": "%s","Hours": "%f"}`, e.BikeId, e.Hours)
}
func (e TotalSavedCost) ToJSON() string {
	return fmt.Sprintf(`{"bike_id": "%s","SavedCost": "%f"}`, e.BikeId, e.SavedCost)
}
func (e TotalSavedCarbon) ToJSON() string {
	return fmt.Sprintf(`{"bike_id": "%s","SavedCarbon": "%f"}`, e.BikeId, e.SavedCarbon)
}
