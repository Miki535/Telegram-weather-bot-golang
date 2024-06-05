package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/sha3"
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
	fmt.Println("Starting Bot")
	var token string
	var apiKey string
	fmt.Println("Enter Bot Token to authenticate")
	fmt.Scan(&token)
	fmt.Println("Enter API key to authenticate")
	fmt.Scan(&apiKey)
	// Захешований токен
	tokenHash := "7845d38d32a4fefd6bcd8607b883b69c609fb1dd0cb2972ce9626cc35ed4c50a"
	ApiKeyHash := "c0e7ced8de9bed006980b7fbcb92bd58a4b7e3c1721a5a1066b6502a17bd9aac"
	data1 := []byte(token)
	hash := sha3.Sum256(data1)
	data2 := []byte(apiKey)
	hash2 := sha3.Sum256(data2)
	// Переводимо hash з типу byte до типу string
	hashString1 := hex.EncodeToString(hash[:])
	hashString2 := hex.EncodeToString(hash2[:])
	// Перевірка на справжність токену і клоча до апі бота
	if tokenHash == hashString1 && ApiKeyHash == hashString2 {
		fmt.Println("Bot is already authenticated")
	} else {
		fmt.Println("Bot is not authenticated")
		os.Exit(0)
	}

	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())

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
			"Привіт! В цьому боті ви можете дізнаватись температуру повітр'я в різних містах України! Щоб почати введіть /та назву свого міста. Також ви можете використати команду /buttons щоб використовувати кнопки за допомогою який дізнаватись температуру. Успіхів!",
		)

		bot.SendMessage(message)

	}, th.CommandEqual("start"))

	//buttons command
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("/Ternopil"),
				tu.KeyboardButton("Documentation"),
			),
		)
		message := tu.Message(
			chatId,
			"Вибeріть, що хочете обрати!",
		).WithReplyMarkup(keyboard)

		bot.SendMessage(message)

	}, th.CommandEqual("buttons"))

	// Kyiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Kyiv&units=metric&appid=%s")
	}, th.CommandEqual("Kyiv"))

	// Ternopil tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Ternopil&country=UA&key="+apiKey)
	}, th.CommandEqual("Ternopil"))

	// Odessa tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Odessa&units=metric&appid=%s")
	}, th.CommandEqual("Odessa"))

	// Lviv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Lviv&units=metric&appid=%s")
	}, th.CommandEqual("Lviv"))

	// Ivano-Frankivsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Ivano-Frankivsk&units=metric&appid=%s")
	}, th.CommandEqual("IvanoFrankivsk"))

	// Dnipro tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Dnipro&units=metric&appid=%s")
	}, th.CommandEqual("Dnipro"))

	// Kharkiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Kharkiv&units=metric&appid=%s")
	}, th.CommandEqual("Kharkiv"))

	// Khmelnytskyi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5/weather?q=Khmelnytskyi&units=metric&appid=%s")
	}, th.CommandEqual("Khmelnytskyi"))

	// Mukacheve tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mukacheve&units=metric&appid=%s")
	}, th.CommandEqual("Mukacheve"))

	// Rivne tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Rivne&units=metric&appid=%s")
	}, th.CommandEqual("Rivne"))

	// Poltava tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Poltava&units=metric&appid=%s")
	}, th.CommandEqual("Poltava"))

	// Zhytomyr tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Zhytomyr&units=metric&appid=%s")
	}, th.CommandEqual("Zhytomyr"))

	// Vinnytsia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Vinnytsia&units=metric&appid=%s")
	}, th.CommandEqual("Vinnytsia"))

	// Sevastopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Sevastopol&units=metric&appid=%s")
	}, th.CommandEqual("Sevastopol"))

	// Simpferopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Simferopol&units=metric&appid=%s")
	}, th.CommandEqual("Simferopol"))

	// Kherson tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Kherson&units=metric&appid=%s")
	}, th.CommandEqual("Kherson"))

	// Melitopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Melitopol&units=metric&appid=%s")
	}, th.CommandEqual("Melitopol"))

	// Mariupol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mariupol&units=metric&appid=%s")
	}, th.CommandEqual("Mariupol"))

	// Donetsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Donetsk&units=metric&appid=%s")
	}, th.CommandEqual("Donetsk"))

	// Luhansk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Luhansk&units=metric&appid=%s")
	}, th.CommandEqual("Luhansk"))

	// Sumy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Sumy&units=metric&appid=%s")
	}, th.CommandEqual("Sumy"))

	// Chernihiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Chernihiv&units=metric&appid=%s")
	}, th.CommandEqual("Chernihiv"))

	// Lutsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Lutsk&units=metric&appid=%s")
	}, th.CommandEqual("Lutsk"))

	// Uzhhorod tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Uzhhorod&units=metric&appid=%s")
	}, th.CommandEqual("Uzhhorod"))

	// Cherkasy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Cherkasy&units=metric&appid=%s")
	}, th.CommandEqual("Cherkasy"))

	// Zaporizhzhia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Zaporizhzhia&units=metric&appid=%s")
	}, th.CommandEqual("Zaporizhzhia"))

	// Mykolaiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "http://api.openweathermap.org/data/2.5//weather?q=Mykolaiv&units=metric&appid=%s")
	}, th.CommandEqual("Mykolaiv"))

	// Chernivtsi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		//chatID := tu.ID(update.Message.Chat.ID)
		//go ALLINone()
	}, th.CommandEqual("Chernivtsi"))
	bh.Start()
}

func ALLINone(chatid telego.ChatID, bot *telego.Bot, URL string) {
	url := URL

	resp, err := http.Get(url)
	if err != nil {
		SendMessage(chatid, bot, "Помилка при передачі данних!")
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
		SendMessage(chatid, bot, "Помилка при передачі данних!")
		return
	}

	description := weatherResponse.Data[0].Weather.Description

	tempcity := fmt.Sprintf("Температура повітря %.1f°C\n", weatherResponse.Data[0].Temp)
	message1 := tu.Message(chatid, tempcity)
	bot.SendMessage(message1)
	switch description {
	case "Broken clouds":
		go descriptionmessage("хмарність з проясненнями", chatid, bot)
	case "Clear sky":
		go descriptionmessage("чисте небо", chatid, bot)
	case "Overcast clouds":
		go descriptionmessage("суцільна хмарність", chatid, bot)
	case "Scattered clouds":
		go descriptionmessage("мінлива хмарність", chatid, bot)
	case "Drizzle":
		go descriptionmessage("мряка", chatid, bot)
	default:
		go descriptionmessage(description, chatid, bot)
	}
}

func SendMessage(chatid telego.ChatID, bot *telego.Bot, text string) {
	message := tu.Message(
		chatid,
		text)
	bot.SendMessage(message)
}

func descriptionmessage(desc string, chatid telego.ChatID, bot *telego.Bot) {
	message2 := tu.Message(chatid, fmt.Sprintf("Опис погоди...%s ", desc))
	bot.SendMessage(message2)
}
