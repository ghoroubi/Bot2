package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

var db *sql.DB
var err error
var conf *viper.Viper
var bot *tgbotapi.BotAPI
var dbName, password, userId, server string
var token string
var occasion string
var maxRand int
var vcf string

func DbConnect() {

	db, err = sql.Open("mssql", "server="+server+";user id="+userId+";password="+password+";database="+dbName+"")
	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}
}

func main() {

	GetConf()
	DbConnect()

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message.Text == "/start" {
			SendTextMessage(update.Message.Chat.ID, welcome, GetHomeKeys)
		}else if update.Message.Text == ContactUsKey {
			//SendTextMessage(update.Message.Chat.ID, ContactUsVal, GetHomeKeys)
			SendVCF(update.Message.Chat.ID,vcf)
		}  else if update.Message.Text == SummerPlan {
			//SumPlan(update.Message.Chat.ID,update.Message.From.ID,GetPlanKeys)
			SumPlan(update.Message.Chat.ID,GetPlanKeys)
		} else if update.Message.Text == Home {
			SendTextMessage(update.Message.Chat.ID, HomeDirected, GetHomeKeys)



		} else if update.Message.ReplyToMessage != nil {
			if update.Message.ReplyToMessage.Text == EnterPhonePlz {

				SendSecurityCode(update.Message.Chat.ID, update.Message.From.ID, update.Message.Text, "CheckCredit", GetHomeKeys)
			} else if update.Message.ReplyToMessage.Text == EnterSecCode {
				CheckCreditByCode(update.Message.Chat.ID, update.Message.From.ID, update.Message.Text)

			} else if update.Message.ReplyToMessage.Text == EnterSecCode_R {

				ChargeByCode(update.Message.Chat.ID, update.Message.From.ID, update.Message.Text)

			} else if update.Message.ReplyToMessage.Text == EnterPhoneForUse {
				RKCharge(update.Message.Chat.ID, update.Message.Text, 7500, GetGiftKeys)
			} else if update.Message.ReplyToMessage.Text == EnterPhoneForGift {
				SendSecurityCode(update.Message.Chat.ID, update.Message.From.ID, update.Message.Text, "RandomCharge", GetGiftKeys)
			}

		}
	}
	defer db.Close()

}
