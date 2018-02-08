package main

type Record struct{
	Id int `json:"id"`
	Item string `json:"item"`
	Amount float32 `json:"amount"`
	Start_date string `json:"start_date"`
	End_date string `json:"end_date"`
	Update_at string `json:"updated_at"`
}

type Records []Record
