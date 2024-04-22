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
		ALLINone("Kyiv", chatID, bot, "TRUE.", "http://api.openweathermap.org/data/2.5/weather?q=Kyiv&units=metric&appid=%s")
	}, th.CommandEqual("Kyiv"))

	bh.Start()
}

func ALLINone(town string, chatid telego.ChatID, bot *telego.Bot, mess string, URL string) {
	// URL для запиту погоди в Києві
	apiKey := "3f7c7314bbddea4af2f8175638c88ad6"

	url := fmt.Sprintf(URL, apiKey)
	
	// Виконати GET-запит
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("Помилка під час виконання запиту: %s", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Прочитати відповідь у вигляді масиву байтів
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Помилка при читанні відповіді: %s", err)
		os.Exit(1)
	}

	// Розкодувати JSON-відповідь
	var data WeatherData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Помилка при розкодуванні JSON: %s", err)
		os.Exit(1)
	}
	TEMP := data.Main.Temp

	tempkiyv := fmt.Sprintf("Температура повітря в "+town+": %.1f°C\n", TEMP)

	message1 := tu.Message(
		chatid,
		tempkiyv,
	)

	message2 := tu.Message(
		chatid,
		mess,
	)
	//SEND FIRST MESSAGE:
	bot.SendMessage(message1)
	//SEND SECOND MESSAGE:
	bot.SendMessage(message2)

}
