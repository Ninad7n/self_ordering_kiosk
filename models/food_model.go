package models

type FoodModel struct {
	ID         int    `json:"ID"`
	NAME       string `json:"NAME"`
	IS_VEG     int    `json:"IS_VEG"`
	FULL_PRICE string `json:"FULL_PRICE"`
}
