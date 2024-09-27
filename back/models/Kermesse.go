package models

type Kermesse struct {
	Base
	Name     string `json:"name"`
	Location string `json:"location"`
	Date     string `json:"date"`
}
