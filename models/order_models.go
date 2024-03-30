package models

type OrderModel struct {
	ID           int    `json:"ID"`
	INVOICE_ID   *int   `json:"INVOICE_ID"`
	TABLE_NO     int    `json:"TABLE_NO"`
	CUST_MOBILE  string `json:"CUST_MOBILE"`
	FOOD_NAME    string `json:"FOOD_NAME"`
	PRICE        string `json:"PRICE"`
	PAYMENT      string `json:"PAYMENT"`
	QUANTITY     int    `json:"QUANTITY"`
	CREATED_DATE string `json:"CREATED_DATE"`
}
