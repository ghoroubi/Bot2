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
var showOccasion bool
var OccasionButton, AboutGift string
var maxRand int

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
		} else if update.Message.Text == Charg7 {
			SendForceReply(update.Message.Chat.ID, EnterPhone7)
		} else if update.Message.Text == Charg15 {
			SendForceReply(update.Message.Chat.ID, EnterPhone15)
		} else if update.Message.Text == Charg30 {
			SendForceReply(update.Message.Chat.ID, EnterPhone30)
		} else if update.Message.Text == Charg45 {
			SendForceReply(update.Message.Chat.ID, EnterPhone45)
		} else if update.Message.Text == Charg90 {
			SendForceReply(update.Message.Chat.ID, EnterPhone90)
		} else if update.Message.Text == AboutGiftKey {
			AboutKeys(update.Message.Chat.ID)
		} else if update.Message.Text == AboutUsKey {
			SendTextMessage(update.Message.Chat.ID, AboutUsVal, GetHomeKeys)
		} else if update.Message.Text == ContactUsKey {
			SendTextMessage(update.Message.Chat.ID, ContactUsVal, GetHomeKeys)
		} else if update.Message.Text == BuyCharge {
			ChargeKeys(update.Message.Chat.ID)
		} else if update.Message.Text == OccasionButton {
			DoChargeOccasion(update.Message.Chat.ID, update.Message.From.ID, 7500)

		} else if update.Message.Text == CreditKey {
			mobile, err := GetMobileNumber(update.Message.Chat.ID, update.Message.From.ID)
			if err != nil {
				SendError(update.Message.Chat.ID, GetHomeKeys)
			} else {
				if mobile != "" {
					//
					CheckCreditByMobile(update.Message.Chat.ID, update.Message.From.ID, mobile)

				} else {

					SendForceReply(update.Message.Chat.ID, EnterPhonePlz)
				}
			}
		} else if update.Message.Text == Home {
			SendTextMessage(update.Message.Chat.ID, HomeDirected, GetHomeKeys)
		} else if update.Message.Text == LotteryKey {
			randC := random(50, 200)

			DoCharge(update.Message.Chat.ID, update.Message.From.ID, randC)
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
			} else if update.Message.ReplyToMessage.Text == EnterPhone7 {
				SendTextMessage(update.Message.Chat.ID, "http://api.rayanehkomak.com/rk/gateway/smp?amount=7500&mobile="+update.Message.Text, GetHomeKeys)
			} else if update.Message.ReplyToMessage.Text == EnterPhone15 {
				SendTextMessage(update.Message.Chat.ID, "http://api.rayanehkomak.com/rk/gateway/smp?amount=15000&mobile="+update.Message.Text, GetHomeKeys)
			} else if update.Message.ReplyToMessage.Text == EnterPhone30 {
				SendTextMessage(update.Message.Chat.ID, "http://api.rayanehkomak.com/rk/gateway/smp?amount=30000&mobile="+update.Message.Text, GetHomeKeys)
			} else if update.Message.ReplyToMessage.Text == EnterPhone45 {
				SendTextMessage(update.Message.Chat.ID, "http://api.rayanehkomak.com/rk/gateway/smp?amount=45000&mobile="+update.Message.Text, GetHomeKeys)
			} else if update.Message.ReplyToMessage.Text == EnterPhone90 {
				SendTextMessage(update.Message.Chat.ID, "http://api.rayanehkomak.com/rk/gateway/smp?amount=90000&mobile="+update.Message.Text, GetHomeKeys)
			}

		}
	}
	defer db.Close()

}
