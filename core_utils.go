package main

import (
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
func GetConf() {
	conf = viper.New()

	conf.SetConfigName("conf")
	conf.AddConfigPath(".")
	conf.SetConfigType("yml")
	err := conf.ReadInConfig()
	if err != nil {
		panic("could not configure app")
	}

	maxRand = conf.GetInt("public.MaxRandCharge")

	token = conf.GetString("public.Token")
	occasion = conf.GetString("occasion.Key")
	showOccasion = conf.GetBool("occasion.Visible")
	OccasionButton = conf.GetString("occasion.Button")
	AboutGift = conf.GetString("occasion.AboutGift")

	dbName = conf.GetString("db.Name")
	password = conf.GetString("db.Password")
	userId = conf.GetString("db.UserId")
	server = conf.GetString("db.Server")

}

func TodayStr() string {
	year, m, d := time.Now().Date()
	strDate := strconv.Itoa(year) + "/" + strconv.Itoa(int(m)) + "/" + strconv.Itoa(d)
	return strDate
}

func TodayWitZeroStr() string {
	year, m, d := time.Now().Date()
	m_ := strconv.Itoa(int(m))
	if len(m_) == 1 {
		m_ = "0" + m_
	}
	d_ := strconv.Itoa(d)

	if len(d_) == 1 {
		d_ = "0" + d_
	}
	strDate := strconv.Itoa(year) + "/" + m_ + "/" + d_
	return strDate

}
