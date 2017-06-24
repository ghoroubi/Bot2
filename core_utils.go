package main

import (
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)
//var DBEngine,DBName,DBUser,DBHost,DBPassword string



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
	token = conf.GetString("Public.Token")
	vcf=conf.GetString("file.VCFile")

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
