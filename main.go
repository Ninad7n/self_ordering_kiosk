package main

import (
	"log"
	"net/http"
	"self_ordering_kiosk/apis"
	"self_ordering_kiosk/db"
)

func main() {
	http.HandleFunc("/get_menu_list", apis.GetMenuList)
	http.HandleFunc("/add_food_to_menu", apis.AddFoodToMenu)
	http.HandleFunc("/update_price", apis.UpdatePrice)
	http.HandleFunc("/delete_food", apis.DeleteFood)
	http.HandleFunc("/create_order", apis.CreateOrder)
	http.HandleFunc("/get_order_list", apis.GetOrderList)
	http.HandleFunc("/delete_order", apis.DeleteOrder)
	http.HandleFunc("/place_order", apis.PlaceOrder)
	http.HandleFunc("/create_invoice", apis.CreateInvoice)
	http.HandleFunc("/get_invoice/", apis.GetInvoice)
	http.HandleFunc("/update_invoice_status", apis.UpdateInvoiceStatus)
	log.Fatal("Fatal : ", http.ListenAndServe(db.DbUrl, nil))
}
