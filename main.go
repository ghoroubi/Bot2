package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_    "database/sql"
	"fmt"
	//_ "github.com/denisenkom/go-mssqldb"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	//"strconv"
	//	"strconv"
)
var schema=`
  CREATE TABLE IF NOT EXISTS tbcodes (
    id SERIAL PRIMARY KEY NOT NULL,
 	telegramid int  NOT NULL,
 	mobile text NOT NULL
);`
var TelegramID int
var db *sqlx.DB
var err error
var conf *viper.Viper
var bot *tgbotapi.BotAPI
//var dbName, password, userId, server string
var token string
var UserWantToView bool
//var maxRand int
var vcf string
var TUser TelegramUser
var SecurityCode int
var PID string

func DbConnect() {
	//query:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",DBHost,DBUser,DBPassword,DBName)

	db, err = sqlx.Connect("postgres", "host=localhost user=postgres password=123456 dbname=SystemBank sslmode=disable")
	if err != nil {
		fmt.Println("From Open() attempt: " + err.Error())
	}
	db.MustExec(schema)
}

//watowato_bot
func main() {

	GetConf()
	DbConnect()
	UserWantToView = true
	//_token :="AAF04OXimpY6CLZjxAQOoFKWrsVQuWjB1zI"
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		fmt.Println(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {
		if update.Message != nil {
			TUser = TelegramUser{}
			fmt.Println("Normal Message", update.Message.Text)
			TUser.TelegramID = update.Message.From.ID
			TUser.MobileNumber = GetMobile(update.Message.Chat.ID, TUser.TelegramID)
			fmt.Printf("tid:%d ___ mobile:%s",TUser.TelegramID,TUser.MobileNumber)
			switch update.Message.Text {

			case "/start":
				TUser.TelegramID = update.Message.From.ID
				TUser.MobileNumber = GetMobile(update.Message.Chat.ID, TUser.TelegramID)
				fmt.Println(TUser.MobileNumber)
				SendTextMessage(update.Message.Chat.ID, welcome+"\n"+PackageWelcome, GetHomeKeys)
				//SimpleMessageSend(update.Message.Chat.ID,FirstStart)
				fmt.Println("mobile: ", TUser.MobileNumber, "\n", "TelID: ", update.Message.From.ID, "\n ChatID", update.Message.Chat.ID)

			case ContactUsKey:
				TUser.TelegramID = update.Message.From.ID
				TUser.MobileNumber = GetMobile(update.Message.Chat.ID, TUser.TelegramID)
				SendVCF(update.Message.Chat.ID, vcf)

			case Plan1:
				UserWantToView = false
				TUser.TelegramID = update.Message.From.ID
				TUser.MobileNumber = GetMobile(update.Message.Chat.ID, TUser.TelegramID)
				PID = PackageOneID
				if TUser.MobileNumber =="" || TUser.MobileNumber==ItsMyNumber{

					SendForceReply(update.Message.Chat.ID, EnterNumberForCheck_Buy)
				}else{
					tmsg:=fmt.Sprintf(ShowMobile,TUser.MobileNumber)
					msg:=tgbotapi.NewMessage(update.Message.Chat.ID,tmsg)
					msg.ReplyMarkup=GetPlanConfirm()
					bot.Send(msg)
				}
			case Plan2:
				UserWantToView = false
				TUser.TelegramID = update.Message.From.ID
				TUser.MobileNumber = GetMobile(update.Message.Chat.ID, TUser.TelegramID)
				PID = PackageTwoID
				if TUser.MobileNumber =="" || TUser.MobileNumber==ItsMyNumber{

					SendForceReply(update.Message.Chat.ID, EnterNumberForCheck_Buy)
				}else{
					tmsg:=fmt.Sprintf(ShowMobile,TUser.MobileNumber)
					msg:=tgbotapi.NewMessage(update.Message.Chat.ID,tmsg)
					msg.ReplyMarkup=GetPlanConfirm()
					bot.Send(msg)
				}
				//SendForceReply(update.Message.Chat.ID, EnterNumberForCheck_Buy)
				/*GetUserInfo(update.Message.Chat.ID,TelegramID)
				SumPlan(update.Message.Chat.ID, TelegramID,GetConfirmMobileToActivationKeys)*/
			case Home:
				TelegramID = update.Message.From.ID
				//GetUserInfo(update.Message.Chat.ID, TelegramID)
				SendTextMessage(update.Message.Chat.ID, HomeDirected, GetHomeKeys)

			case CheckPlan:
				UserWantToView = true
				TUser.TelegramID = update.Message.From.ID
				TUser.MobileNumber = GetMobile(update.Message.Chat.ID, TUser.TelegramID)
				if TUser.MobileNumber=="" || TUser.MobileNumber==ItsMyNumber{
					SendForceReply(update.Message.Chat.ID,EnterNumberForCheck)
				}else {
					tmsg := fmt.Sprintf(ShowMobile, TUser.MobileNumber)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, tmsg)
					msg.ReplyMarkup = GetConfirmMobileToActivationKeys()
					bot.Send(msg)
				}
			case Confirm:
				af:=true
				NewUserQuery := fmt.Sprintf("UPDATE tbcodes SET mobile='%s' WHERE telegramid=%d", TUser.MobileNumber,update.Message.From.ID)
				db.Exec(NewUserQuery)
				CheckMobileNumber(TUser.MobileNumber,update.Message.Chat.ID,&af)
				/*SendSecCode(update.Message.Chat.ID,TUser.MobileNumber)
				SendForceReply(update.Message.Chat.ID, SendCodeNotification)*/
			case ChangeNumber:
				SendForceReply(update.Message.Chat.ID,EnterNumberForCheck)
			case Cancel:
				TelegramID = update.Message.From.ID
				SendTextMessage(update.Message.Chat.ID, HomeDirected, GetHomeKeys)
			case ItsMyNumber:
				af := true
				if TUser.MobileNumber=="" || TUser.MobileNumber==ItsMyNumber {
					SendForceReply(update.Message.Chat.ID, EnterNumberForCheck_Buy)
				}else{
					CheckMobileNumber(TUser.MobileNumber, update.Message.Chat.ID, &af)

					fmt.Println(af, "FlagStatus")
					if af == true {
						ShowInvoice(update.Message.Chat.ID, PID, TUser.MobileNumber)

					}
				}
					/*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
					bot.Send(msg)*/
					//SendForceReply(update.Message.Chat.ID,MobileIsWrong)
			case ChangeTheNumebr:
				SendForceReply(update.Message.Chat.ID,EnterNumberForCheck_Buy)
			}
		}
		if update.Message.ReplyToMessage != nil {
			fmt.Println("Replay Message:", update.Message.ReplyToMessage.Text)
			switch update.Message.ReplyToMessage.Text {

			case EnterNumberForCheck_Buy:
				flag := true
				if MobileIsValid(update.Message.Text) {
					TUser.MobileNumber=update.Message.Text
					TUser.TelegramID=update.Message.From.ID
					NewUserQuery := fmt.Sprintf("UPDATE tbcodes SET mobile='%s' WHERE telegramid=%d", TUser.MobileNumber,update.Message.From.ID)
					db.Exec(NewUserQuery)
					CheckMobileNumber(update.Message.Text, update.Message.Chat.ID, &flag)
					fmt.Println(flag, "FlagStatus")
					if flag == true {
						ShowInvoice(update.Message.Chat.ID, PID, TUser.MobileNumber)//TUser wa update.message.text
						//TUser.MobileNumber=update.Message.Text
					}
				}else {
					//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
					/*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
					bot.Send(msg)*/
					SendForceReply(update.Message.Chat.ID,MobileIsWrong)
				}
			case MobileIsWrong:
				flag := true

				if MobileIsValid(update.Message.Text) {
					TUser.MobileNumber=update.Message.Text
					TUser.TelegramID=update.Message.From.ID
					NewUserQuery := fmt.Sprintf("UPDATE tbcodes SET mobile='%s' WHERE telegramid=%d", TUser.MobileNumber,update.Message.From.ID)
					db.Exec(NewUserQuery)
					CheckMobileNumber(update.Message.Text, update.Message.Chat.ID, &flag)
					//TUser.MobileNumber=update.Message.Text
					fmt.Println(flag, "FlagStatus")
					if flag == true {
						ShowInvoice(update.Message.Chat.ID, PID, update.Message.Text)
						//TUser.MobileNumber=update.Message.Text
					}
				}else {
					//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
					/*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
					bot.Send(msg)*/
					SendForceReply(update.Message.Chat.ID,MobileIsWrong)
				}
			case MobileIsWrong_CheckingCase:
				if MobileIsValid(update.Message.Text) {
					TUser.MobileNumber=update.Message.Text
					TUser.TelegramID=update.Message.From.ID
					NewUserQuery := fmt.Sprintf("UPDATE tbcodes SET mobile='%s' WHERE telegramid=%d", TUser.MobileNumber,update.Message.From.ID)
					db.Exec(NewUserQuery)
					SendSecCode(update.Message.Chat.ID, update.Message.Text)
					//TUser.MobileNumber = update.Message.Text
					SendForceReply(update.Message.Chat.ID, SendCodeNotification)
					//CheckMobileNumber(update.Message.Text, update.Message.Chat.ID)
				}else{
					//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
					/*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
					bot.Send(msg)*/
					SendForceReply(update.Message.Chat.ID,MobileIsWrong_CheckingCase)
				}
			case EnterNumberForCheck:
				if MobileIsValid(update.Message.Text) {
					TUser.MobileNumber=update.Message.Text
					TUser.TelegramID=update.Message.From.ID
					NewUserQuery := fmt.Sprintf("UPDATE tbcodes SET mobile='%s' WHERE telegramid=%d", TUser.MobileNumber,update.Message.From.ID)
					db.Exec(NewUserQuery)
					SendSecCode(update.Message.Chat.ID, update.Message.Text)
				//	TUser.MobileNumber = update.Message.Text
					SendForceReply(update.Message.Chat.ID, SendCodeNotification)
					//CheckMobileNumber(update.Message.Text, update.Message.Chat.ID)
				}else{
					//SendTextMessage(update.Message.Chat.ID,MobileIsWrong,nil)
					/*msg:=tgbotapi.NewMessage(update.Message.Chat.ID,MobileIsWrong)
					bot.Send(msg)*/
					SendForceReply(update.Message.Chat.ID,MobileIsWrong_CheckingCase)
				}
			case SendCodeNotification:
				SecCodeReview(update.Message.Chat.ID, update.Message.Text)

			case WrongCode:
				SecCodeReview(update.Message.Chat.ID,update.Message.Text)
				/*SendSecCode(update.Message.Chat.ID, update.Message.Text)
				SendForceReply(update.Message.Chat.ID, SendCodeNotification)*/
			}
		}
	}
	defer db.Close()
}
