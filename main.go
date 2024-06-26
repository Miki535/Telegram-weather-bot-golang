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

var UrL string

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
			"Привіт! В цьому боті ви можете дізнаватись температуру повітр'я в різних містах України! Щоб почати введіть /та назву свого міста. Також ви можете використати команду /buttons щоб використовувати кнопки за допомогою який дізнаватись температуру. Якщо ви загубились ,то використайте команду /help. Успіхів!",
		)

		bot.SendMessage(message)

	}, th.CommandEqual("start"))

	//buttons command
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)

		keyboard := tu.Keyboard(
			tu.KeyboardRow(
				tu.KeyboardButton("/Ternopil"),
				tu.KeyboardButton("/Kyiv"),
				tu.KeyboardButton("/Odessa"),
				tu.KeyboardButton("/Lviv"),
				tu.KeyboardButton("/IvanoFrankivsk"),
				tu.KeyboardButton("/Dnipro"),
				tu.KeyboardButton("/Kharkiv"),
				tu.KeyboardButton("/Khmelnytskyi"),
				tu.KeyboardButton("/Mukacheve"),
				tu.KeyboardButton("/Rivne"),
				tu.KeyboardButton("/Poltava"),
				tu.KeyboardButton("/Zhytomyr"),
				tu.KeyboardButton("/Vinnytsia"),
				tu.KeyboardButton("/Sevastopol"),
				tu.KeyboardButton("/Simferopol"),
				tu.KeyboardButton("/Kherson"),
				tu.KeyboardButton("/Melitopol"),
				tu.KeyboardButton("/Mariupol"),
				tu.KeyboardButton("/Donetsk"),
				tu.KeyboardButton("/Luhansk"),
				tu.KeyboardButton("/Sumy"),
				tu.KeyboardButton("/Chernigiv"),
				tu.KeyboardButton("/Lutsk"),
				tu.KeyboardButton("/Cherkasy"),
				tu.KeyboardButton("/Zaporizhzhia"),
				tu.KeyboardButton("/Mykolaiv"),
				tu.KeyboardButton("/Chernivtsi"),
			),
		)
		message := tu.Message(
			chatId,
			"Вибeріть, що хочете обрати!",
		).WithReplyMarkup(keyboard)

		bot.SendMessage(message)

	}, th.CommandEqual("buttons"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)
		bot.SendMessage(tu.Message(chatId, "Всі доступні команди: Міста: /Ternopil, /Kyiv, /Odessa, /Kharkiv, /Khmelnytskyi, /Mukacheve, /Rivne, /Poltava, /Zhytomyr, /Vinnytsia, /Sevastopol, /Simferopol, /Kherson , /Melitopol, /Mariupol, /Donetsk, /Luhansk, /Sumy, /Chernihiv, /Lutsk, /Cherkasy, /Zaporizhzhia, /Mykolaiv, /Chernivtsi, Для розробників: /opensourse, Інші: /start, /help, /buttons"))
	}, th.CommandEqual("help"))

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)
		bot.SendMessage(tu.Message(chatId, "Цей проект є опен-сорс! Тому тримайте github лінк!"))
		bot.SendMessage(tu.Message(chatId, "https://github.com/Miki535/Telegram-weather-bot-golang"))
	}, th.CommandEqual("opensourse"))

	// Kyiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Kyiv&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Kyiv"))

	// Ternopil tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Ternopil&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Ternopil"))

	//TernoSaw
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Kyiv&country=UA&key="+apiKey, "TernoSaw")
	}, th.CommandEqual("TernoSaw"))

	// Odessa tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Odessa&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Odessa"))

	// Lviv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Lviv&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Lviv"))

	// Ivano-Frankivsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=&country=UA&key="+apiKey, "")
	}, th.CommandEqual("IvanoFrankivsk"))

	// Dnipro tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Dnipro&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Dnipro"))

	// Kharkiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Kharkiv&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Kharkiv"))

	// Khmelnytskyi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Khmelnytskyi&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Khmelnytskyi"))

	// Mukacheve tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Mukacheve&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Mukacheve"))

	// Rivne tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Rivne&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Rivne"))

	// Poltava tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Poltava&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Poltava"))

	// Zhytomyr tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Zhytomyr&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Zhytomyr"))

	// Vinnytsia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Vinnytsia&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Vinnytsia"))

	// Sevastopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Sevastopol&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Sevastopol"))

	// Simpferopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Simferopol&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Simferopol"))

	// Kherson tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Kherson&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Kherson"))

	// Melitopol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Melitopol&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Melitopol"))

	// Mariupol tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Mariupol&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Mariupol"))

	// Donetsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Donetsk&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Donetsk"))

	// Luhansk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Luhansk&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Luhansk"))

	// Sumy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Sumy&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Sumy"))

	// Chernihiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Chernihiv&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Chernihiv"))

	// Lutsk tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Lutsk&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Lutsk"))

	// Uzhhorod tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Uzhhorod&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Uzhhorod"))

	// Cherkasy tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Cherkasy&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Cherkasy"))

	// Zaporizhzhia tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Zaporizhzhia&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Zaporizhzhia"))

	// Mykolaiv tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Mykolaiv&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Mykolaiv"))

	// Chernivtsi tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Chernivtsi&country=UA&key="+apiKey, "")
	}, th.CommandEqual("Chernivtsi"))
	bh.Start()

	// Warsaw tempereature information
	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatID := tu.ID(update.Message.Chat.ID)
		go ALLINone(chatID, bot, "https://api.weatherbit.io/v2.0/current?city=Warsaw&country=UA&key="+apiKey, "")
	})
}

func ALLINone(chatid telego.ChatID, bot *telego.Bot, URL string, test string) {
	url := URL
	if test != "" {
		UrL = "https://api.weatherbit.io/v2.0/current?city=Chernivtsi&country=UA&key="
	} else {
		UrL = url
	}
	//hello
	resp, err := http.Get(UrL)
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
	case "Few clouds":
		go descriptionmessage("трішки хмаринок у небі", chatid, bot)
	default:
		go descriptionmessage(description, chatid, bot)
	}

	// Warsaw and Ternopil temperature diferents
	if test != "" {
		Warsaw := weatherResponse.Data[0].Temp
		Ternopil := weatherResponse.Data[0].Temp
		bot.SendMessage(tu.Message(chatid, fmt.Sprint(Warsaw, Ternopil)))
		if Warsaw < Ternopil {
			result := fmt.Sprintf("In Ternopil tempereature higher on", Ternopil-Warsaw)
			bot.SendMessage(tu.Message(chatid, result))
		} else if Warsaw > Ternopil {
			result := fmt.Sprintf("In Warsaw temperature higher on", Ternopil-Warsaw)
			bot.SendMessage(tu.Message(chatid, result))
		} else {
			bot.SendMessage(tu.Message(chatid, "Warsaw = Ternopil"))
		}
	} else {
		return
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
