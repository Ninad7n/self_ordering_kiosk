package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"self_ordering_kiosk/db"
	"self_ordering_kiosk/models"
	"strconv"
)

func GetInvoice(resW http.ResponseWriter, req *http.Request) {

	id := req.URL.Path[len("/get_invoice/"):]
	fmt.Println("reqID : ", id)
	db := db.DbConn()
	var response *sql.Rows
	if id == "" {
		rows, err := db.Query("SELECT * FROM invoice")
		if err != nil {
			fmt.Println("GetInvoice Query : ", err)
		}
		response = rows
	} else {
		rows, err := db.Query("SELECT * FROM invoice WHERE ID=?", id)
		if err != nil {
			fmt.Println("GetInvoice Query : ", err)
		}
		response = rows
	}

	defer response.Close()
	var invoiceList []models.InvoiceModel
	for response.Next() {
		var invoice models.InvoiceModel
		err := response.Scan(&invoice.ID, &invoice.TABLE_NO, &invoice.TAX, &invoice.AMOUNT, &invoice.STATUS, &invoice.CREATED_DATE)
		if err != nil {
			fmt.Println("GetInvoice Scan : ", err)
		}
		invoiceList = append(invoiceList, invoice)
	}
	resjson, err := json.Marshal(invoiceList)
	if err != nil {
		fmt.Println("GetInvoice Marshal : ", err)
	}
	resW.Header().Set("Content-Type", "application/json")
	resW.Write(resjson)
	defer db.Close()
}

func CreateInvoice(resW http.ResponseWriter, req *http.Request) {

	orderList := GetOrdersByTable(resW, req)
	ordervalue := 0

	for x := 0; x < len(orderList); x++ {
		local, err := strconv.Atoi(orderList[x].PRICE)
		if err != nil {
			fmt.Println("ERROR AT GetBill query : ", err)
		}
		ordervalue = ordervalue + (local * orderList[x].QUANTITY)
	}

	taxValue := (10 * ordervalue) / 100

	db := db.DbConn()
	var table_no = req.FormValue("table_no")
	response, err := db.Exec("INSERT INTO invoice (TABLE_NO, TAX, AMOUNT) VALUES(?,?,?)", req.FormValue("table_no"), strconv.Itoa(taxValue), strconv.Itoa(taxValue+ordervalue))
	if err != nil {
		fmt.Println(err)
	}
	id, err := response.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Rows affected : ", id)
	UpdateInvoiceId(id, table_no)
	data := "/get_invoice"
	http.Redirect(resW, req, data, code)
	defer db.Close()
}

func UpdateInvoiceStatus(resW http.ResponseWriter, req *http.Request) {
	db := db.DbConn()
	var inv_id = req.FormValue("invoice_id")
	response, err := db.Query("UPDATE invoice SET STATUS = 'paid' WHERE ID=?", inv_id)
	if err != nil {
		panic(err.Error())
	}
	UpdatePayment(inv_id)
	log.Println(response, "UpdateInvoiceStatus", inv_id)
	data := "/get_invoice/" + inv_id
	http.Redirect(resW, req, data, code)
	defer db.Close()
}
