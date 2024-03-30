package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"self_ordering_kiosk/db"
	"self_ordering_kiosk/models"
)

func GetOrderList(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	db := db.DbConn()

	rows, err := db.Query("SELECT * FROM orders ORDER by CREATED_DATE DESC")
	if err != nil {
		fmt.Println("ERROR AT GetMenuList query : ", err)
	}

	defer rows.Close()
	var orders []models.OrderModel
	for rows.Next() {
		var order models.OrderModel
		err = rows.Scan(&order.ID, &order.INVOICE_ID, &order.TABLE_NO, &order.CUST_MOBILE, &order.FOOD_NAME, &order.PRICE, &order.QUANTITY, &order.PAYMENT, &order.CREATED_DATE)
		if err != nil {
			fmt.Println("ERROR AT GetMenuList scan : ", err)
		}
		orders = append(orders, order)
	}
	orderJson, err := json.Marshal(orders)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(orderJson))
	res.Header().Set("Content-Type", "application/json")
	res.Write(orderJson)
	defer db.Close()
}

func GetFoodById(resW http.ResponseWriter, req *http.Request) models.FoodModel {
	db := db.DbConn()
	var food models.FoodModel
	response, err := db.Query("SELECT * FROM food WHERE ID=?", req.FormValue("id"))
	if err != nil {
		fmt.Println("ERROR AT GetFoodById query : ", err)
	}
	defer response.Close()
	for response.Next() {
		response.Scan(&food.ID, &food.NAME, &food.IS_VEG, &food.FULL_PRICE)
	}
	defer db.Close()
	return food
}

func GetOrdersByTable(resW http.ResponseWriter, req *http.Request) []models.OrderModel {
	db := db.DbConn()
	var orders []models.OrderModel
	response, err := db.Query("SELECT * FROM orders WHERE TABLE_NO=? AND PAYMENT='placed'", req.FormValue("table_no"))
	if err != nil {
		fmt.Println("ERROR AT GetBillByTable query : ", err)
	}
	defer response.Close()
	for response.Next() {
		var order models.OrderModel
		response.Scan(&order.ID, &order.INVOICE_ID, &order.TABLE_NO, &order.CUST_MOBILE, &order.FOOD_NAME, &order.PRICE, &order.QUANTITY, &order.PAYMENT, &order.CREATED_DATE)
		orders = append(orders, order)
	}
	fmt.Println("orders_data", orders)
	defer db.Close()
	return orders
}

func CreateOrder(resW http.ResponseWriter, req *http.Request) {
	db := db.DbConn()
	if req.Method != "POST" {
		return
	}
	var foodData = GetFoodById(resW, req)
	response, err := db.Query(
		"INSERT INTO orders (TABLE_NO, CUST_MOBILE, FOOD_NAME, PRICE, PAYMENT, QUANTITY) VALUES(?,?,?,?,?,?)",
		req.FormValue("table_no"),
		req.FormValue("cust_no"),
		foodData.NAME,
		foodData.FULL_PRICE,
		"in_cart",
		req.FormValue("quantity"),
	)
	if err != nil {
		fmt.Println("ERROR AT GetMenuList query : ", err)
	}
	log.Println(response, "INSERTED")

	defer db.Close()
	http.Redirect(resW, req, "/get_order_list", code)
}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	id := r.FormValue("id")
	response, err := db.Query("Update orders SET PAYMENT = 'placed' WHERE TABLE_NO = ? AND PAYMENT = 'in_cart'", id)
	if err != nil {
		panic(err.Error())
	}
	log.Println(response, "UPDATE")
	defer db.Close()
	http.Redirect(w, r, "/get_order_list", code)
}

func UpdateInvoiceId(inv_id int64, table_no string) {
	db := db.DbConn()
	response, err := db.Query("UPDATE orders SET INVOICE_ID=? WHERE PAYMENT = 'placed' AND TABLE_NO=?", inv_id, table_no)
	if err != nil {
		panic(err.Error())
	}
	log.Println(response, "UPDATE")
	defer db.Close()
}

func UpdatePayment(inv_id string) {
	db := db.DbConn()
	response, err := db.Query("UPDATE orders SET PAYMENT = 'paid' WHERE INVOICE_ID=?", inv_id)
	if err != nil {
		panic(err.Error())
	}
	log.Println(response, "UPDATE")
	defer db.Close()
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	id := r.FormValue("id")
	response, err := db.Query("DELETE FROM orders WHERE ID = ? AND PAYMENT = 'in_cart' ", id)
	if err != nil {
		panic(err.Error())
	}
	log.Println(response, "DELETE")
	defer db.Close()
	http.Redirect(w, r, "/get_order_list", code)
}
