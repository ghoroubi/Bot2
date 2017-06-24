package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
	"strconv"
	//	"time"
	//	"log"
	//	"net/url"
)

type fn func() tgbotapi.ReplyKeyboardMarkup

//If User Choose to check his/her active plan
func CheckMyPlan(chatId int64, tid int, keyboard fn) {
	fmt.Println("Your Plan  is your Mobile:::")
	mob := GetMobile(chatId, tid)
	if mob == "" {
		SendTextMessage(chatId, YouHaveNotPackage, GetHomeKeys)
		return
	} else {
		res := Check(mob,chatId)
		msg := tgbotapi.NewMessage(chatId, res)
		msg.ReplyMarkup = keyboard()
		bot.Send(msg)
	}
}

func GetUserInfo(chatId int64, telegramId int) {
	m := GetMobile(chatId, telegramId)
	if m == "" {
		TUser.TelegramID = telegramId
		return
	}
	TUser.TelegramID = telegramId
	TUser.MobileNumber = m
}

func SecCodeReview(chatId int64, code string) {
	intCode, err := strconv.Atoi(code)
	fmt.Println(intCode, ",User:", TUser.MobileNumber)
	if err != nil {
		SendError(chatId, GetHomeKeys)
	}
	if intCode != SecurityCode {
		SendForceReply(chatId, WrongCode)
		return
	}

	fmt.Println(TUser.MobileNumber)
	//CheckMobileNumber()
	tid := TUser.TelegramID
	tmob := TUser.MobileNumber
	NewUserQuery := fmt.Sprintf("INSERT INTO tbcodes(telegramid,mobile) VALUES (%d , '%s')", tid, tmob)
	db.Exec(NewUserQuery)
	af :=true
CheckMobileNumber(tmob,chatId,&af)
}
func CheckMobileNumber(mobile string, chatId int64, allowFlag *bool) {
	fmt.Println("Mobile:", mobile, "is in checking progress")
	activePlan := Check(mobile, chatId)
	TUser.MobileNumber=mobile
	NewUserQuery := fmt.Sprintf("INSERT INTO tbcodes(telegramid,mobile) VALUES (%d , '%s')", mobile,TUser.TelegramID)
	db.Exec(NewUserQuery)
	if activePlan=="ok" {
	*allowFlag=true
	}else if activePlan != "NoPackage" {
		*allowFlag = false
		fmt.Println("NO ACTIVE PLAN", activePlan)
		SendTextMessage(chatId, activePlan, GetHomeKeys)
	} else {
		*allowFlag = true
	}
}
func CallAPI(chatId int64, tid int, mob string) {

	urlToSend := fmt.Sprintf("payment.rayanehkomak.ir/rk/gateway/smp?type=callPackage&mobile=%s&pid=%s&tid=%d", mob, PID, tid)
	//toSendUrl:=url+urlencode(tempUrl)
	msg := tgbotapi.NewMessage(chatId, urlToSend)
	msg.ReplyMarkup = GetHomeKeys()
	bot.Send(msg)
	fmt.Println("Calling API", tid, ":", mob)
	NewUserQuery := fmt.Sprintf("INSERT INTO tbcodes(telegramid,mobile) VALUES (%d , '%s')", tid, mob)
	db.Exec(NewUserQuery)

}
func SendVCF(chatId int64, text string) {
	line1:="☎  ☎  021-7129 ☎  ☎ \n"
	line2:="❇️ شماره تماس با کارشناسان رایانه کمک:   ۰۲۱۷۱۲۹ \n"
	msg:=line1+line2+vcfVal
	msg2Send := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msg2Send)

	file := tgbotapi.NewDocumentUpload(chatId, vcf)
	bot.Send(file)
}
func SendError(chatId int64, keyboard fn) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keyboard()
	msg.Text = SystemError
	bot.Send(msg)
}
func SendTextMessage(chatId int64, text string, keys fn) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = keys()
	bot.Send(msg)
}

func SendForceReply(chatId int64, text string) {
	fmt.Println("ForceReply: ", text)
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true}
	bot.Send(msg)
}

func GetMobile(chatId int64, telegramId int) string {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = GetHomeKeys()
	// get mobile
	var strmobile string
	//var id int
	err := db.Get(&strmobile, "select mobile from tbCodes where telegramId="+strconv.Itoa(telegramId))
	fmt.Println(strmobile)
	if err != nil {
		//SendError(chatId, GetHomeKeys)
		strmobile = ""
		return ""
	}
	return strmobile
}

func urlencode(s string) (result string) {
	for _, c := range s {
		if c <= 0x7f { // single byte
			result += fmt.Sprintf("%%%X", c)
		} else if c > 0x1fffff { // quaternary byte
			result += fmt.Sprintf("%%%X%%%X%%%X%%%X",
				0xf0+((c&0x1c0000)>>18),
				0x80+((c&0x3f000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else if c > 0x7ff { // triple byte
			result += fmt.Sprintf("%%%X%%%X%%%X",
				0xe0+((c&0xf000)>>12),
				0x80+((c&0xfc0)>>6),
				0x80+(c&0x3f),
			)
		} else { // double byte
			result += fmt.Sprintf("%%%X%%%X",
				0xc0+((c&0x7c0)>>6),
				0x80+(c&0x3f),
			)
		}
	}

	return result
}

func SendSecCode(chatId int64, mobile string) {
	var code int

	code = random(1000, 9999)
	fmt.Println(code)
	SecurityCode = code
	//m:="09364921604"
	req, err := http.NewRequest("POST", "http://api.rayanehkomak.com/rk/sms/send?num="+mobile+"&txt="+urlencode(fmt.Sprintf(SecCode, code)), nil)
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		SendError(chatId, GetHomeKeys)
		return
	}
}
func Check(mobile string, chatId int64) string {
	var planStatus string

	res, err := http.Get("http://api.rayanehkomak.com/crm/customer/callpackages?mobile=" + mobile)

	if err != nil {
		//Panic(err)
		fmt.Println("Error Occured!")
	}
	body, err := ioutil.ReadAll(res.Body)
	var customer Customer
	json.Unmarshal(body, &customer)

	x := len(customer.PurchasedCallPackages)
	if x == 0 && UserWantToView {

		l1 := "⚠️ کاربر گرامی برای شماره %s هیچ بسته ای در سیستم ثبت نشده است "
		l2 := "برای فعال سازی لطفا یکی از بسته های موجود را انتخاب نمایید"
		planStatus = fmt.Sprintf(l1+l2+Title, mobile)
	}else if x==0 && UserWantToView==false{
	planStatus="ok"
	}else if x!= 0 && UserWantToView {
		ShowPlanDetails(mobile,chatId,customer)
	}else if x!=0 && UserWantToView==false {
		SendTextMessage(chatId,YouHaveActivePlan,GetHomeKeys)
	}

	return planStatus
}

//end of Check
func ShowPlanDetails(mobile string, chatId int64, c Customer) {
	var Plan_Days int
	var planStatus string
	Plan_Days = c.PurchasedCallPackages[0].DaysToExpire
	fmt.Println("Days2Expire:", Plan_Days)
	//Plan_Describtion: = c.PurchasedCallPackages[0].CallPackage.description
	fmt.Println("describtion: ",c.PurchasedCallPackages[0].CallPackage.description )

	//end of switch
	line2 := fmt.Sprintf(" ✅ بسته فعلی برای شماره موبایل  %s به مدت %d روز دیگر اعتبار دارد",mobile,Plan_Days)
	line2=line2+"\n 💠💠💠💠💠💠💠💠💠💠💠💠💠💠💠💠💠"
	switch UserWantToView {
	case true:
		masterLine := "\n" + line2
		planStatus = masterLine
	case false:
		masterline := YouHaveActivePlan + "\n" + line2
		planStatus = masterline
	}
	fmt.Println(planStatus)
	msg:=tgbotapi.NewMessage(chatId,planStatus)
	msg.ReplyMarkup=GetHomeKeys()
	bot.Send(msg)
}
func ShowInvoice(chatId int64, pid string, mobile string) {
	var strPid string
	NewUserQuery := fmt.Sprintf("INSERT INTO tbcodes(telegramid,mobile) VALUES (%d , '%s')", int(chatId), mobile)
	db.Exec(NewUserQuery)
	switch pid {
	case PackageOneID:
		strPid = "خرید بسته نامحدود یک ساله \n"
		t2 := fmt.Sprintf("برای شماره موبایل : %s", mobile)
		fee2 := "قیمت بسته : ۸۸۰۰۰ تومان"

		l1:="موافقتنامه: https://goo.gl/pPzj1z "
		l2:="کیفیت خدمت: سطح۱ "
		tax := "مبلغ نهایی(+۹٪ ارزش افزوده): 95,920 تومان "
		ifconfirm2 := "جهت تایید و ادامه بر روی لینک زیر کلیک کنید: \n "
		final2 := strPid + t2 + "\n" + fee2 + "\n" + l1+ "\n"+l2+"\n" +tax + "\n" + ifconfirm2
		msg := tgbotapi.NewMessage(chatId, final2)
		/////////
		bot.Send(msg)
	case PackageTwoID:
		strPid = "خرید بسته نامحدود ۳ ماهه \n"
		t2 := fmt.Sprintf("برای شماره موبایل : %s", mobile)
		fee2 := "قیمت بسته : ۵۵۰۰۰ تومان"
		l1:="موافقتنامه: https://goo.gl/pPzj1z "
		l2:="کیفیت خدمت: سطح۱ "
		tax := "مبلغ نهایی(+۹٪ ارزش افزوده): 59,950 تومان "
		ifconfirm2 := "جهت تایید و ادامه بر روی لینک زیر کلیک کنید: \n "
		final2 := strPid + t2 + "\n" + fee2 + "\n" +l1+"\n"+l2+"\n"+ tax + "\n" + ifconfirm2
		msg2 := tgbotapi.NewMessage(chatId, final2)
		msg2.ReplyMarkup = GetHomeKeys()
		bot.Send(msg2)

	}
	//	strMobile:="شماره موبایل:"
	CallAPI(chatId, int(chatId), mobile)
}
func MobileIsValid(mobile string)bool{
	var IsMobile bool
	lenFlag:=len(mobile)==11
	fmt.Println("len:",len(mobile))
	firstDigitFlag:=string(mobile[0]) == "0"
	fmt.Println("firstDigit:",string(mobile[0]))
	secDigitFlag:=string(mobile[1])=="9"
	fmt.Println("SecondDigit",string(mobile[1]))
	IsMobile=lenFlag && firstDigitFlag && secDigitFlag
	//IsMobile=firstDigitFlag && secDigitFlag
	//fdigit:=
	return IsMobile
}
