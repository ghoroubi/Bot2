package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"fmt"
)

func GetPlanKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}
	accept := tgbotapi.KeyboardButton{Text: AcceptPlan}
fmt.Print("inside of gtplan1")
	home := tgbotapi.KeyboardButton{Text: Home}

	commands := [][]tgbotapi.KeyboardButton{}

	row1 := []tgbotapi.KeyboardButton{accept}
	commands = append(commands, row1)
	fmt.Print("inside of gtplan2")
	row2 := []tgbotapi.KeyboardButton{home}
	commands = append(commands, row2)
	fmt.Print("inside of gtplan3")
	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}

func GetGiftKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}

	home := tgbotapi.KeyboardButton{Text: Home}
	lottery := tgbotapi.KeyboardButton{Text: LotteryKey}

	commands := [][]tgbotapi.KeyboardButton{}

	row1 := []tgbotapi.KeyboardButton{lottery}
	commands = append(commands, row1)

	row2 := []tgbotapi.KeyboardButton{home}
	commands = append(commands, row2)

	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}

func GetHomeKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}
	//buy := tgbotapi.KeyboardButton{Text: BuyCharge}
	credit := tgbotapi.KeyboardButton{Text: CreditKey}

	//about := tgbotapi.KeyboardButton{Text: AboutUsKey}
	contact := tgbotapi.KeyboardButton{Text: ContactUsKey}
	//aboutGift := tgbotapi.KeyboardButton{Text: AboutGiftKey}
	summerPlan:=tgbotapi.KeyboardButton{Text:SummerPlan}
	commands := [][]tgbotapi.KeyboardButton{}

	row0 := []tgbotapi.KeyboardButton{summerPlan}
	commands = append(commands, row0)

	row1 := []tgbotapi.KeyboardButton{credit, contact}
	commands = append(commands, row1)

	//row2 := []tgbotapi.KeyboardButton{contact, about}
	//commands = append(commands, row2)

	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}
