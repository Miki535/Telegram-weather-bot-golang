package main

import (
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"
	"io/ioutil"
	"net/http"
	"os"

	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

// Структура для розкодування JSON-відповіді
type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	botToken := "6498753115:AAEEl4xSkVGvuiXLgdbWGk-KNeEcecT8tLc"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)

		message := tu.Message(
			chatID,
			"Привіт! В цьому боті ви можете дізнаватись температуру повітр'я в різних містах України! Щоб почати введіть /та назву свого міста. Успіхів!",
		)

		bot.SendMessage(message)

	}, th.CommandEqual("start"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)

		apiKey := "3f7c7314bbddea4af2f8175638c88ad6"
		// URL для запиту погоди в Києві
		url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=Kyiv&units=metric&appid=%s", apiKey)

		// Виконати GET-запит
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Помилка під час виконання запиту: %s", err)
			return
		}
		defer response.Body.Close()

		// Прочитати відповідь у вигляді масиву байтів
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Помилка при читанні відповіді: %s", err)
			return
		}

		// Розкодувати JSON-відповідь
		var data WeatherData
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Printf("Помилка при розкодуванні JSON: %s", err)
			return
		}
		temp := data.Main.Temp
		SENDMESS(temp, "Києві", chatID, bot, "ТЕСТ УСПІШНО ПРОЙДЕНО!")

	}, th.CommandEqual("Kyiv"))

	bh.Start()
}

func SENDMESS(TEMP float64, town string, chatid telego.ChatID, bot *telego.Bot, mess string) {
	tempkiyv := fmt.Sprintf("Температура повітря в "+town+": %.1f°C\n", TEMP)

	message1 := tu.Message(
		chatid,
		tempkiyv,
	)

	message2 := tu.Message(
		chatid,
		mess,
	)

	bot.SendMessage(message1)

	bot.SendMessage(message2)

}
