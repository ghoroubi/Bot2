package main

import (
	"gopkg.in/telegram-bot-api.v4"
)

func GetChargesKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}
	c7 := tgbotapi.KeyboardButton{Text: Charg7}
	c15 := tgbotapi.KeyboardButton{Text: Charg15}
	c30 := tgbotapi.KeyboardButton{Text: Charg30}
	c45 := tgbotapi.KeyboardButton{Text: Charg45}
	c90 := tgbotapi.KeyboardButton{Text: Charg90}

	home := tgbotapi.KeyboardButton{Text: Home}

	commands := [][]tgbotapi.KeyboardButton{}

	row0 := []tgbotapi.KeyboardButton{c15, c7}
	commands = append(commands, row0)

	row1 := []tgbotapi.KeyboardButton{c45, c30}
	commands = append(commands, row1)

	row2 := []tgbotapi.KeyboardButton{home, c90}
	commands = append(commands, row2)

	rep.Keyboard = commands
	rep.ResizeKeyboard = true
	return rep
}

func GetGiftKeys() tgbotapi.ReplyKeyboardMarkup {
	rep := tgbotapi.ReplyKeyboardMarkup{}
	occ := tgbotapi.KeyboardButton{Text: OccasionButton}
	home := tgbotapi.KeyboardButton{Text: Home}
	lottery := tgbotapi.KeyboardButton{Text: LotteryKey}

	commands := [][]tgbotapi.KeyboardButton{}
	if showOccasion {
		row0 := []tgbotapi.KeyboardButton{occ}
		commands = append(commands, row0)
	}

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
