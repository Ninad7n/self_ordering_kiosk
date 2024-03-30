package models

type InvoiceModel struct {
	ID           int    `json:"ID"`
	TABLE_NO     int    `json:"TABLE_NO"`
	TAX          string `json:"TAX"`
	AMOUNT       string `json:"AMOUNT"`
	STATUS       string `json:"STATUS"`
	CREATED_DATE string `json:"CREATED_DATE"`
}
