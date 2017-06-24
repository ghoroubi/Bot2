package main

import (
	"gopkg.in/telegram-bot-api.v4"
	//"fmt"
)

func GetPlanConfirm() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}
	its_myNumber := tgbotapi.KeyboardButton{Text: ItsMyNumber}
	Change_TheNumber:=tgbotapi.KeyboardButton{Text: ChangeTheNumebr}

	commands := [][]tgbotapi.KeyboardButton{}

	row0:=[]tgbotapi.KeyboardButton{its_myNumber}
	commands=append(commands, row0)

	row1 := []tgbotapi.KeyboardButton{Change_TheNumber}
	commands = append(commands, row1)


	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}
func GetConfirmMobileToActivationKeys() tgbotapi.ReplyKeyboardMarkup{
	rep:=tgbotapi.ReplyKeyboardMarkup{}
	commands:=[][]tgbotapi.KeyboardButton{}

	itsOK:=tgbotapi.KeyboardButton{Text:Confirm}
	itsCancel:=tgbotapi.KeyboardButton{Text:ChangeNumber}
	//goHome:=tgbotapi.KeyboardButton{Text:Home}
	row1:=[]tgbotapi.KeyboardButton{itsOK}
	row2:=[]tgbotapi.KeyboardButton{itsCancel}
	commands=append(commands,row1)
	commands=append(commands,row2)
	rep.Keyboard=commands
	rep.ResizeKeyboard=true
	return rep
}


func GetHomeKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}
	firstPlan := tgbotapi.KeyboardButton{Text: Plan1}
	secondPlan:=tgbotapi.KeyboardButton{Text: Plan2}
	check_Plan := tgbotapi.KeyboardButton{Text: CheckPlan}
	contact := tgbotapi.KeyboardButton{Text: ContactUsKey}
	//temp:=tgbotapi.KeyboardButton{Text:"TEST"}

	commands := [][]tgbotapi.KeyboardButton{}

	row1 := []tgbotapi.KeyboardButton{firstPlan}
	commands = append(commands, row1)
	row2:=[]tgbotapi.KeyboardButton{secondPlan}
	commands=append(commands, row2)
	row3 := []tgbotapi.KeyboardButton{contact,check_Plan}
	commands = append(commands, row3)

	//row2 := []tgbotapi.KeyboardButton{contact, }
	//commands = append(commands, row2)

	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}

