package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type fn func() tgbotapi.ReplyKeyboardMarkup

func  SumPlan(chatId int64,keys fn)  {
	//mobile, err := GetMobileNumber(chatId, telegramId)
	//if err != nil {
	//	//SendError(chatId, GetHomeKeys)
	//} else {
	//	if mobile != "" {
	//		//
	//		CheckCreditByMobile(chatId, telegramId, mobile)
	//
	//	} else {
	//
	//		SendForceReply(chatId, EnterPhonePlz)
	//	}
	//}
fmt.Println(chatId,"\n",PlanConfirm)
 msg:=tgbotapi.NewMessage(chatId,PlanConfirm)

	msg.ReplyMarkup=keys()
	bot.Send(msg)
	}
/*func SumPlan(chatId int64,telegramId int,keys fn)  {
	//mobile, err := GetMobileNumber(chatId, telegramId)
	//if err != nil {
	//	//SendError(chatId, GetHomeKeys)
	//} else {
	//	if mobile != "" {
	//		//
	//		CheckCreditByMobile(chatId, telegramId, mobile)
	//
	//	} else {
	//
	//		SendForceReply(chatId, EnterPhonePlz)
	//	}
	//}
	msg:=tgbotapi.NewMessage(chatId,PlanConfirm)
	msg.ReplyMarkup=keys()
}*/
func SendVCF(chatId int64,text string){
	file:=tgbotapi.NewDocumentUpload(chatId,vcf)
	msg:=tgbotapi.NewMessage(chatId,vcfVal)
	bot.Send(file)
	bot.Send(msg)
}
func SendError(chatId int64, keys fn) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keys()
	msg.Text = SystemError
	bot.Send(msg)
}
func SendTextMessage(chatId int64, text string, keys fn) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = keys()
	bot.Send(msg)
}

func SendForceReply(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true}
	bot.Send(msg)
}

type SCharge struct {
	Id     string
	Charge int
}

func GetMobileNumber(chatId int64, telegramId int) (string, error) {
	var mobile string
	result, err := db.Query("select mobile from tbTelegramUsers where telegramId=" + strconv.Itoa(telegramId))
	if err != nil {
		return "", err
	}
	for result.Next() {
		_ = result.Scan(&mobile)
	}

	return mobile, nil
}

func CheckCreditByMobile(chatId int64, telegramId int, mobile string) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = GetHomeKeys()

	res, err := http.Get("http://api.rayanehkomak.com/crm/checkcharge?mobile=" + mobile)
	if err != nil {
		SendError(chatId, GetHomeKeys)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	var data SCharge
	json.Unmarshal(body, &data)
	msg.Text = fmt.Sprintf(CreditVal, data.Charge)
	bot.Send(msg)

}
func CheckCreditByCode(chatId int64, telegramId int, code string) {
	strmobile := GetMobileByCode(chatId, telegramId, code)
	if strmobile != "" {
		CheckCreditByMobile(chatId, telegramId, strmobile)
		UpdateCode(chatId, telegramId)
		SetMobile(chatId, telegramId, strmobile)
	}
}

func ChargeByCode(chatId int64, telegramId int, code string) {
	strmobile := GetMobileByCode(chatId, telegramId, code)
	if strmobile != "" {

		sum, err := SetRandGift(chatId, telegramId)
		if err == nil {
			RKCharge(chatId, strmobile, sum, GetGiftKeys)
			UpdateCode(chatId, telegramId)

			SetMobile(chatId, telegramId, strmobile)
		}
	}
}

func GetMobileByCode(chatId int64, telegramId int, code string) string {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = GetHomeKeys()

	// get mobile
	var strmobile, strdate string
	var id int
	result, err := db.Query("select id,mobile,date from  tbCodes where telegramId=" + strconv.Itoa(telegramId) + " and code=" + code + " and isused=0 order by id desc ")
	if err != nil {
		SendError(chatId, GetHomeKeys)
		return ""
	}

	for result.Next() {
		_ = result.Scan(&id, &strmobile, &strdate)
	}
	if strmobile == "" {
		msg.Text = CodeError
		bot.Send(msg)
		return ""
	}

	tim1, err := time.Parse(LongFormat, strdate)
	oneday := time.Hour * 24 * 1
	tim2 := tim1.Add(oneday)
	now := time.Now()
	if !(now.Format(LongFormat) > tim1.Format(LongFormat) && now.Format(LongFormat) < tim2.Format(LongFormat)) {
		msg.Text = CodeError
		bot.Send(msg)
		return ""
	}

	return strmobile
}
func UpdateCode(chatId int64, telegramId int) {

	strQ := "update tbCodes set isused=1 where telegramId=" + strconv.Itoa(telegramId)
	_, err = db.Exec(strQ)
	if err != nil {
		SendError(chatId, GetHomeKeys)
		return
	}
}

func SetMobile(chatId int64, telegramId int, strmobile string) {

	var id int
	result, err := db.Query("select id from tbTelegramUsers where telegramId=" + strconv.Itoa(telegramId))
	if err != nil {
		SendError(chatId, GetHomeKeys)
		return
	}
	for result.Next() {
		_ = result.Scan(&id)
	}
	if id == 0 {
		strQ := "insert into tbTelegramUsers(telegramId,mobile) values(" + strconv.Itoa(telegramId) + "," + strmobile + ")"
		_, err = db.Exec(strQ)
		if err != nil {
			SendError(chatId, GetHomeKeys)
			return
		}
	}
}
/* func AboutKeys(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, AboutGift)
	msg.ReplyMarkup = GetGiftKeys()
	bot.Send(msg)
}
*/
func PlanKeys(chatId int64) {
	msg := tgbotapi.NewMessage(chatId, PlanConfirm)
	msg.ReplyMarkup = GetPlanKeys()
	bot.Send(msg)
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

func SendSecurityCode(chatId int64, telegramId int, mobile string, typ string, keys fn) {
	var strDate string
	var isUsed bool
	var intCode int
	result, err := db.Query("select top 1 Date,IsUsed,Code from  tbCodes where isused=0 and telegramId=" + strconv.Itoa(telegramId) + " order by id desc ")
	if err != nil {
		SendError(chatId, keys)
		return
	}

	for result.Next() {
		_ = result.Scan(&strDate, &isUsed, &intCode)
	}

	tim1, err := time.Parse(LongFormat, strDate)
	oneday := time.Hour * 24 * 1
	tim2 := tim1.Add(oneday)
	now := time.Now()

	var code int
	if isUsed == false && now.Format(LongFormat) > tim1.Format(LongFormat) && now.Format(LongFormat) < tim2.Format(LongFormat) {
		code = intCode
	} else {
		code = random(1000, 9999)
		castedVal := time.Now().Format(LongFormat)
		strQ := "insert into tbCodes (TelegramId, Code,Date, IsUsed,Mobile) values (" + strconv.Itoa(telegramId) + "," + strconv.Itoa(code) + ",'" + castedVal + "',0,'" + mobile + "')"
		_, err = db.Exec(strQ)
		if err != nil {
			SendError(chatId, keys)
			return
		}
	}

	req, err := http.NewRequest("POST", "http://api.rayanehkomak.com/rk/sms/send?num="+mobile+"&txt="+urlencode(fmt.Sprintf(SecCode, code)), nil)
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		SendError(chatId, keys)
		return
	}

	if typ == "CheckCredit" {
		SendForceReply(chatId, EnterSecCode)
	} else if typ == "RandomCharge" {
		SendForceReply(chatId, EnterSecCode_R)

	}
}

func RKCharge(chatId int64, mobile string, charge int, keys fn) {

	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = keys()
	res, err := http.Get("http://api.rayanehkomak.com/crm/charge?mobile=" + mobile + "&charge=" + strconv.Itoa(charge))
	if err != nil {
		//hanle error
		SendError(chatId, GetGiftKeys)
		return
	}

	if res.StatusCode == 200 {
		msg.Text = fmt.Sprintf(ChargeActivated, charge, mobile)
		bot.Send(msg)
	} else {
		SendError(chatId, keys)
		return
	}
}

func DoChargeOccasion(chatId int64, telegramId int, charge int) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = GetGiftKeys()

	var lstc int
	result, err := db.Query("select top 1 TelegramId from  tbOccasionPromotion  where telegramId=" + strconv.Itoa(telegramId) + " and occasion='" + occasion + "'")
	if err != nil {
		SendError(chatId, GetGiftKeys)
		return
	}

	for result.Next() {
		_ = result.Scan(&lstc)
	}
	if lstc != 0 {
		msg.Text = OccasionError
		bot.Send(msg)
		return
	}

	_, err = db.Exec("insert into tbOccasionPromotion (TelegramId, Gift, Date,Occasion) values (" + strconv.Itoa(telegramId) + "," + strconv.Itoa(charge) + ", '" + TodayWitZeroStr() + "','" + occasion + "')")
	if err != nil {
		SendError(chatId, GetGiftKeys)
		return
	}

	msg.Text = fmt.Sprintf(ChargedSuccessfully, charge)
	bot.Send(msg)
	SendForceReply(chatId, EnterPhoneForUse)

}

func SetRandGift(chatId int64, telegramId int) (int, error) {
	r, _ := db.Query("select sum(gift) from tbRandPromotion where telegramId=" + strconv.Itoa(telegramId) + " and IsUsed=0")
	var sum int

	for r.Next() {
		_ = r.Scan(&sum)
	}

	_, err = db.Exec("update tbRandPromotion set IsUsed=1 where telegramId=" + strconv.Itoa(telegramId))
	if err != nil {
		SendError(chatId, GetGiftKeys)
		return 0, err
	}
	return sum, nil
}
func DoCharge(chatId int64, telegramId int, charge int) {
	msg := tgbotapi.NewMessage(chatId, "")
	msg.ReplyMarkup = GetGiftKeys()

	r, err := db.Query("select sum(gift) from tbRandPromotion where telegramId=" + strconv.Itoa(telegramId) + " and IsUsed=0")
	var sum int

	for r.Next() {
		_ = r.Scan(&sum)
	}

	var lstc int
	result, err := db.Query("select top 1 TelegramId from  tbRandPromotion  where telegramId=" + strconv.Itoa(telegramId) + " and date='" + TodayWitZeroStr() + "'")
	if err != nil {
		SendError(chatId, GetGiftKeys)
		return
	}

	for result.Next() {
		_ = result.Scan(&lstc)
	}

	if lstc != 0 {
		msg.Text = LotteryError + " " + fmt.Sprintf(LotterySum, sum)
		bot.Send(msg)

	} else {

		_, err = db.Exec("insert into tbRandPromotion (TelegramId, Gift, Date) values (" + strconv.Itoa(telegramId) + "," + strconv.Itoa(charge) + ", '" + TodayWitZeroStr() + "')")
		if err != nil {
			SendError(chatId, GetGiftKeys)
			return
		}
		msg.Text = fmt.Sprintf(LottSuccess, charge) + " " + fmt.Sprintf(LotterySum, sum+charge)
		bot.Send(msg)
		sum = sum + charge

	}
	if sum > maxRand {

		mobile, err := GetMobileNumber(chatId, telegramId)
		if err != nil {
			SendError(chatId, GetHomeKeys)
		} else {
			if mobile != "" {
				sum, err := SetRandGift(chatId, telegramId)
				if err == nil {

					RKCharge(chatId, mobile, sum, GetGiftKeys)
				}

			} else {
				SendForceReply(chatId, EnterPhoneForGift)

			}
		}

	}
}

// func Charge(chatId int64, mobile string, telegramId int) {
// 	msg := tgbotapi.NewMessage(chatId, "")

// 	var sumCharg = 0

// 	msg.ReplyMarkup = GetHomeKeys()

// 	var lstc int
// 	result, err := db.Query("select top 1 Cid from  tbPromotion order by Id desc")
// 	if err != nil {
// 		SendError(chatId, GetHomeKeys)
// 		return
// 	}
// 	for result.Next() {
// 		_ = result.Scan(&lstc)
// 	}

// 	var query string
// 	strDate := TodayStr()
// 	strZeroDate := TodayWitZeroStr()
// 	if giftForToday {
// 		query = "select top " + strconv.Itoa(callCount) + " Cmobile, Dur_Sec, MDate, C.Cid cid from  MohDReplicated md inner join tbCalls c on md.counter = c.Ccounter inner join tbCustomers cs on c.Ccustomer = cs.Cid where CMobile='" + mobile + "' and Mdate = '" + strZeroDate + "' and Cdate = '" + strDate + "' and c.Cid > " + strconv.Itoa(lstc)
// 	} else {
// 		query = "select top " + strconv.Itoa(callCount) + " Cmobile, Dur_Sec, MDate, C.Cid cid from  MohDReplicated md inner join tbCalls c on md.counter = c.Ccounter inner join tbCustomers cs on c.Ccustomer = cs.Cid where CMobile='" + mobile + "' and c.Cid > " + strconv.Itoa(lstc) + " order by c.Cid desc"
// 	}

// 	calls, err := db.Query(query)
// 	if err != nil {
// 		SendError(chatId, GetHomeKeys)
// 		return
// 	}
// 	for calls.Next() {
// 		var Cmobile, Mdate string
// 		var Dur_Sec, Cid int
// 		err = calls.Scan(&Cmobile, &Dur_Sec, &Mdate, &Cid)
// 		if Dur_Sec > second {
// 			dur_Call := Dur_Sec / second
// 			charg1 := dur_Call * charge

// 			b := Dur_Sec % second

// 			mcharge := ((b * charge) / second) + charg1

// 			res, err := http.Get("http://api.rayanehkomak.com/crm/charge?mobile=" + Cmobile + "&charge=" + strconv.Itoa(mcharge) + "&second=" + strconv.Itoa(Dur_Sec))
// 			if err != nil {
// 				SendError(chatId, GetHomeKeys)
// 				break
// 			}
// 			if res.StatusCode == 200 {
// 				_, err := db.Exec("insert into tbPromotion (Cid, Dur_Sec, Gift_Price, MDate, Cmobile,Telegram) values (" + strconv.Itoa(Cid) + ", '" + strconv.Itoa(Dur_Sec) + "', '" + strconv.Itoa(mcharge) + "', '" + Mdate + "', '" + mobile + "'," + strconv.Itoa(telegramId) + ")")
// 				if err != nil {
// 					SendError(chatId, GetHomeKeys)
// 					break
// 				}
// 				sumCharg = sumCharg + mcharge
// 			}
// 		}
// 	}
// 	if err == nil {
// 		if sumCharg == 0 {
// 			msg.Text = NoFund
// 		} else {
// 			msg.Text = fmt.Sprintf(ChargedSuccessfully, sumCharg)
// 		}
// 	}
// 	bot.Send(msg)
// }
