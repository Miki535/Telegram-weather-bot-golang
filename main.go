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

	// Kyiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Kyiv", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Kyiv&units=metric&appid=%s")
	}, th.CommandEqual("Kyiv"))

	// Ternopil tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Ternopil", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Ternopil&units=metric&appid=%s")
	}, th.CommandEqual("Ternopil"))

	// Odessa tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Odessa", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Odessa&units=metric&appid=%s")
	}, th.CommandEqual("Odessa"))

	// Lviv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Lviv", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Lviv&units=metric&appid=%s")
	}, th.CommandEqual("Lviv"))

	// Ivano-Frankivsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Ivano-Frankivsk", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Ivano-Frankivsk&units=metric&appid=%s")
	}, th.CommandEqual("IvanoFrankivsk"))

	// Dnipro tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Dnipro", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Dnipro&units=metric&appid=%s")
	}, th.CommandEqual("Dnipro"))

	// Kharkiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Kharkiv", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Kharkiv&units=metric&appid=%s")
	}, th.CommandEqual("Kharkiv"))

	// Khmelnytskyi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Khmelnytskyi", chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Khmelnytskyi&units=metric&appid=%s")
	}, th.CommandEqual("Khmelnytskyi"))

	// Mukacheve tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Mukacheve", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mukacheve&units=metric&appid=%s")
	}, th.CommandEqual("Mukacheve"))

	// Rivne tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Rivne", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Rivne&units=metric&appid=%s")
	}, th.CommandEqual("Rivne"))

	// Poltava tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Poltava", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Poltava&units=metric&appid=%s")
	}, th.CommandEqual("Poltava"))

	// Zhytomyr tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Zhytomyr ", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Zhytomyr&units=metric&appid=%s")
	}, th.CommandEqual("Zhytomyr "))

	// Vinnytsia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Vinnytsia", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Vinnytsia&units=metric&appid=%s")
	}, th.CommandEqual("Vinnytsia"))

	// Sevastopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Sevastopol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Sevastopol&units=metric&appid=%s")
	}, th.CommandEqual("Sevastopol"))

	// Simpferopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Simferopol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Simferopol&units=metric&appid=%s")
	}, th.CommandEqual("Simferopol"))

	// Kherson tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Kherson", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Kherson&units=metric&appid=%s")
	}, th.CommandEqual("Kherson"))

	// Melitopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Melitopol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Melitopol&units=metric&appid=%s")
	}, th.CommandEqual("Melitopol"))

	// Mariupol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Mariupol", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mariupol&units=metric&appid=%s")
	}, th.CommandEqual("Mariupol"))

	// Donetsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Donetsk", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Donetsk&units=metric&appid=%s")
	}, th.CommandEqual("Donetsk"))

	// Luhansk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Luhansk", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Luhansk&units=metric&appid=%s")
	}, th.CommandEqual("Luhansk"))

	// Sumy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Sumy", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Sumy&units=metric&appid=%s")
	}, th.CommandEqual("Sumy"))

	// Chernihiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Chernihiv", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Chernihiv&units=metric&appid=%s")
	}, th.CommandEqual("Chernihiv"))

	// Lutsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Lutsk", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Lutsk&units=metric&appid=%s")
	}, th.CommandEqual("Lutsk"))

	// Uzhhorod tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Uzhhorod", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Uzhhorod&units=metric&appid=%s")
	}, th.CommandEqual("Uzhhorod"))

	// Cherkasy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Cherkasy", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Cherkasy&units=metric&appid=%s")
	}, th.CommandEqual("Cherkasy"))

	// Zaporizhzhia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Zaporizhzhia", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Zaporizhzhia&units=metric&appid=%s")
	}, th.CommandEqual("Zaporizhzhia"))

	// Mykolaiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Mykolaiv", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mykolaiv&units=metric&appid=%s")
	}, th.CommandEqual("Mykolaiv"))

	// Chernivtsi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		ALLINone("Chernivtsi", chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Chernivtsi&units=metric&appid=%s")
	}, th.CommandEqual("Chernivtsi"))

	bh.Start()

}

func ALLINone(town string, chatid telego.ChatID, bot *telego.Bot, URL string) {
	// URL для запиту погоди в town
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

	tempcity := fmt.Sprintf("Температура повітря в "+town+": %.1f°C\n", TEMP)

	message1 := tu.Message(
		chatid,
		tempcity,
	)

	messagefortemp := "Для такої температури повітря потрібно вдівати:"

	var mess string

	switch {
	case TEMP < -25:
		mess = messagefortemp + "1. Куртка або пальто: Основний елемент для захисту від холоду.Оберіть теплу куртку або пальто з достатнім утепленням.\n2. Світшот або светр: Додайте шар теплого одягу під куртку для додаткового комфорту.\n3. Джинси або штани зі штучного матеріалу: Виберіть штани, які добре утримують тепло, такі як джинси або штани з теплих матеріалів.\n4. Високі шкарпетки: Високі шкарпетки допоможуть утримати тепло в ногах.\n5. Зручні черевики або черевики на шнурках: Оберіть взуття з гарним утепленням і підошвою, щоб захистити ноги від холоду і забезпечити стійкість на слизьких дорогах.\n6. Шапка або шарф (за бажанням): Шапка допоможе утримати тепло голови, а шарф може бути корисним для захисту шиї і обличчя від вітру."
	}

	message2 := tu.Message(
		chatid,
		mess,
	)
	//SEND FIRST MESSAGE:
	bot.SendMessage(message1)
	//SEND SECOND MESSAGE:
	bot.SendMessage(message2)

}
