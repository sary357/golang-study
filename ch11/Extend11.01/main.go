package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type OrderDetail struct {
	ItemName string `json:"itemname"`
	Desc     string `json:"desc,omitempty"`
	Qty      int    `json:"qty"`
	Price    int    `json:"price"`
}

type ShipTo struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

type Order struct {
	Total       int           `json:"total"`
	Paid        bool          `json:"paid"`
	Fragile     bool          `json:"fragile,omitempty"`
	OrderDetail []OrderDetail `json:"orderdetail,omitempty"`
}

type CustomerOrder struct {
	UserName string `json:"username"`
	ShipTo   ShipTo `json:"shipto"`
	Order    Order  `json:"order"`
}

func AddItem(c *CustomerOrder, item OrderDetail) {
	c.Order.OrderDetail = append(c.Order.OrderDetail, item)
}

func Total(c *CustomerOrder, paid bool) {
	sum := 0

	// calculate the total sum of order in OrderDetail
	for _, i := range c.Order.OrderDetail {
		sum += i.Price * i.Qty
		if IsFragile(i) {
			c.Order.Fragile = true
		}

	}
	c.Order.Total = sum

	c.Order.Paid = paid
}

func IsFragile(item OrderDetail) bool {
	// if item name in OrderDetail contains the string "Glass" or "glass"
	// then return true
	// else return false
	if strings.Index(item.ItemName, "Glass") >= 0 || strings.Index(item.ItemName, "glass") >= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	customerData := CustomerOrder{}
	jsonStr := []byte(`
	{
		"username":"blackhat", 
		"shipto":{
			"street":"Sulphur Springs Rd", 
			"city":"Park City", "state":"VA",
			"zipcode":12345
		}, 
		"order":{
			"paid":false,
			"orderdetail":[] 
		}
	}
	`)
	if err := json.Unmarshal(jsonStr, &customerData); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	item1 := OrderDetail{
		ItemName: "A Guide to the World of zeros and ones",
		Desc:     "book",
		Qty:      3,
		Price:    50,
	}

	item2 := OrderDetail{
		ItemName: "Final Fantasy The Zodiac Age",
		Desc:     "Nintendo Switch Game",
		Qty:      1,
		Price:    50,
	}

	item3 := OrderDetail{
		ItemName: "Crystal Drinking Glass",
		Qty:      11,
		Price:    25,
	}
	AddItem(&customerData, item1)
	AddItem(&customerData, item2)
	AddItem(&customerData, item3)
	Total(&customerData, true)

	fmtString, err2 := json.MarshalIndent(customerData, "", "    ")
	if err2 != nil {
		fmt.Println(err2)
		os.Exit(1)
	}
	fmt.Println(string(fmtString))
}
