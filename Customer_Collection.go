package main

import (
//	"encoding/json"
//	"fmt"
//	"net/url"
)

type Body struct {
	Customer       Customer           `json:"customer"`
	AvailablePkges []AvailablePackage `json:"availablePkges"`
}
type Customer struct {
	Mobile       string        `json:"mobile"`
	PurchasedCallPackages []PurchasedCallPackage `json:"callPackages"`
	Charge       int           `json:"charge"`
	TelegramID   int           `json:"telegramid"`
	Numbers      []string      `json:"numbers"`
	Id           string        `json:"id"`
}
type CallPackage struct {
	_id string `json:"_id"`
	description string `json:"description"`
}
type PurchasedCallPackage struct {
	CallPackage         CallPackage `json:"callPackage"`
	RemainedTalkSeconds int    `json:"remainedTalkSeconds"`
	Expired             bool   `json:"expired"`
	//	PurchaseDate        time.Time `json:"purchaseDate"`
	DaysToExpire int `json:"daysToExpire"`
}
type AvailablePackage struct {
	Name string `json:"name"`
	//Description string `json:"description"`
	TalkDuration int    `json:"talkDuration"`
	Price        int    `json:"price"`
	Url          string `json:"url"`
}
type TelegramUser struct {
	TelegramID   int    `json:"telegram_id"`
	MobileNumber string `json:"mobile_number"`
	PackageID string
}



/*


func URLEncode()string{

	m := map[string]interface{}{"type": "callPackage", "mobile": "md5_base_16", "pid": "1473655478000","tid":""}
	json_str, _ := json.Marshal(m)
	fmt.Println(string(json_str[:]))

	values := url.Values{"para": {string(json_str[:])}}

	fmt.Println(values.Encode())


}*/
func PackWelcome()string {
	l1 := "  "

	l2 := "با فعالسازی بسته های نامحدود رایانه کمک هیچ هزینه مضاعفی بر روی قبض تلفن شما منظور نمی گردد."

	PWelcome := l1 + "\n" + l2
	return PWelcome
}
