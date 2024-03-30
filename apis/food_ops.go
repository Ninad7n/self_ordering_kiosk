package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"self_ordering_kiosk/db"
	"self_ordering_kiosk/models"
)

const code = 301

func GetMenuList(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	db := db.DbConn()

	rows, err := db.Query("SELECT * FROM food")
	if err != nil {
		fmt.Println("ERROR AT GetMenuList query : ", err)
	}

	defer rows.Close()
	var foodList []models.FoodModel
	for rows.Next() {
		var food models.FoodModel
		err = rows.Scan(&food.ID, &food.NAME, &food.IS_VEG, &food.FULL_PRICE)
		if err != nil {
			fmt.Println("ERROR AT GetMenuList scan : ", err)
		}
		fmt.Println("name", food.NAME)
		foodList = append(foodList, food)
	}
	foodJson, err := json.Marshal(foodList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(foodJson))
	res.Header().Set("Content-Type", "application/json")
	res.Write(foodJson)
	defer db.Close()
}

func AddFoodToMenu(resW http.ResponseWriter, req *http.Request) {
	db := db.DbConn()
	if req.Method != "POST" {
		return
	}
	response, err := db.Query(
		"INSERT INTO food (NAME, IS_VEG, FULL_PRICE) VALUES(?,?,?)",
		req.FormValue("name"),
		req.FormValue("is_veg"),
		req.FormValue("full_price"),
	)
	if err != nil {
		fmt.Println("ERROR AT GetMenuList query : ", err)
	}
	log.Println(response, "INSERTED")

	defer db.Close()
	http.Redirect(resW, req, "/get_menu_list", code)
}

func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	fPrice := r.FormValue("full_price")
	id := r.FormValue("id")
	response, err := db.Query("UPDATE food SET FULL_PRICE=? WHERE id=?", fPrice, id)
	if err != nil {
		panic(err.Error())
	}
	log.Println(response, "UPDATE")
	defer db.Close()
	http.Redirect(w, r, "/get_menu_list", code)
}

func DeleteFood(w http.ResponseWriter, r *http.Request) {
	db := db.DbConn()
	id := r.FormValue("id")
	response, err := db.Query("DELETE FROM food WHERE ID=?", id)
	if err != nil {
		panic(err.Error())
	}
	log.Println(response, "DELETE")
	defer db.Close()
	http.Redirect(w, r, "/get_menu_list", code)
}
