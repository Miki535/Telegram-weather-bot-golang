package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mymmrac/telego"

	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

type WeatherResponse struct {
	Data []struct {
		Temp    float64 `json:"temp"`
		Weather struct {
			Description string `json:"description"`
		} `json:"weather"`
	} `json:"data"`
}

func main() {
	botToken := "HYUvamAneTOKEN"

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

	// Kyiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Kyiv", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Kyiv&units=metric&appid=%s")
	}, th.CommandEqual("Kyiv"))

	// Ternopil tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Ternopil", chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Ternopil&country=UA&key=fb16830dddc5462c8ee0fcf5cb5db86c")
	}, th.CommandEqual("Ternopil"))

	// Odessa tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Odessa", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Odessa&units=metric&appid=%s")
	}, th.CommandEqual("Odessa"))

	// Lviv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Lviv", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Lviv&units=metric&appid=%s")
	}, th.CommandEqual("Lviv"))

	// Ivano-Frankivsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Ivano-Frankivsk", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Ivano-Frankivsk&units=metric&appid=%s")
	}, th.CommandEqual("IvanoFrankivsk"))

	// Dnipro tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Dnipro", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Dnipro&units=metric&appid=%s")
	}, th.CommandEqual("Dnipro"))

	// Kharkiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Kharkiv", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Kharkiv&units=metric&appid=%s")
	}, th.CommandEqual("Kharkiv"))

	// Khmelnytskyi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Khmelnytskyi", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Khmelnytskyi&units=metric&appid=%s")
	}, th.CommandEqual("Khmelnytskyi"))

	// Mukacheve tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Mukacheve", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mukacheve&units=metric&appid=%s")
	}, th.CommandEqual("Mukacheve"))

	// Rivne tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Rivne", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Rivne&units=metric&appid=%s")
	}, th.CommandEqual("Rivne"))

	// Poltava tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Poltava", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Poltava&units=metric&appid=%s")
	}, th.CommandEqual("Poltava"))

	// Zhytomyr tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Zhytomyr ", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Zhytomyr&units=metric&appid=%s")
	}, th.CommandEqual("Zhytomyr"))

	// Vinnytsia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Vinnytsia", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Vinnytsia&units=metric&appid=%s")
	}, th.CommandEqual("Vinnytsia"))

	// Sevastopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Sevastopol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Sevastopol&units=metric&appid=%s")
	}, th.CommandEqual("Sevastopol"))

	// Simpferopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Simferopol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Simferopol&units=metric&appid=%s")
	}, th.CommandEqual("Simferopol"))

	// Kherson tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Kherson", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Kherson&units=metric&appid=%s")
	}, th.CommandEqual("Kherson"))

	// Melitopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Melitopol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Melitopol&units=metric&appid=%s")
	}, th.CommandEqual("Melitopol"))

	// Mariupol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Mariupol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mariupol&units=metric&appid=%s")
	}, th.CommandEqual("Mariupol"))

	// Donetsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Donetsk", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Donetsk&units=metric&appid=%s")
	}, th.CommandEqual("Donetsk"))

	// Luhansk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Luhansk", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Luhansk&units=metric&appid=%s")
	}, th.CommandEqual("Luhansk"))

	// Sumy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Sumy", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Sumy&units=metric&appid=%s")
	}, th.CommandEqual("Sumy"))

	// Chernihiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Chernihiv", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Chernihiv&units=metric&appid=%s")
	}, th.CommandEqual("Chernihiv"))

	// Lutsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Lutsk", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Lutsk&units=metric&appid=%s")
	}, th.CommandEqual("Lutsk"))

	// Uzhhorod tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Uzhhorod", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Uzhhorod&units=metric&appid=%s")
	}, th.CommandEqual("Uzhhorod"))

	// Cherkasy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Cherkasy", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Cherkasy&units=metric&appid=%s")
	}, th.CommandEqual("Cherkasy"))

	// Zaporizhzhia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Zaporizhzhia", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Zaporizhzhia&units=metric&appid=%s")
	}, th.CommandEqual("Zaporizhzhia"))

	// Mykolaiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Mykolaiv", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mykolaiv&units=metric&appid=%s")
	}, th.CommandEqual("Mykolaiv"))

	// Chernivtsi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone("Chernivtsi", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Chernivtsi&units=metric&appid=%s")
	}, th.CommandEqual("Chernivtsi"))

	bh.Start()

}

func ALLINone(town string, chatid telego.ChatID, bot *telego.Bot, URL string) {
	url := URL

	resp, err := http.Get(url)
	if err != nil {
		SendMessage(chatid, bot, "error1")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		SendMessage(chatid, bot, "Ви використали весь ліміт API! Бот буде доступним через 24 години!")
		return
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		SendMessage(chatid, bot, "error2")
		return
	}

	description := weatherResponse.Data[0].Weather.Description

	tempcity := fmt.Sprintf("Температура повітря %.1f°C\n", weatherResponse.Data[0].Temp)
	message1 := tu.Message(chatid, tempcity)
	bot.SendMessage(message1)
	switch description {
	case "Clear sky":
		descriptionmessage("Чисте небо", chatid, bot)
	default:
		descriptionmessage("Нема опису погоди", chatid, bot)
	}
}

func SendMessage(chatid telego.ChatID, bot *telego.Bot, text string) {
	message := tu.Message(
		chatid,
		text)
	bot.SendMessage(message)
}

func descriptionmessage(desc string, chatid telego.ChatID, bot *telego.Bot) {
	message2 := tu.Message(chatid, fmt.Sprintf("Опис погоди:", desc))
	bot.SendMessage(message2)
}
