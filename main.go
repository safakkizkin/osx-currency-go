package main

import (
	"encoding/json"
	"fmt"
	"github.com/caseymrm/menuet"
	"io/ioutil"
	"log"
	"net/http"
	"osx-currency-go/models"
	"time"
)

func getCurrenciesInfo() {
	_oldValue := 0.0
	for {
		if _oldValue != 0.0 {
			switch _interval {
			case models.ThirtySeconds:
				time.Sleep(time.Second * 30)
				break
			case models.AMinute:
				time.Sleep(time.Minute)
				break
			case models.TenMinutes:
				time.Sleep(time.Minute * 10)
				break
			default:
			case models.TenSeconds:
				time.Sleep(time.Second * 10)
				break
			}
		}

		resp, err := http.Get("https://freecurrencyapi.net/api/v2/latest?apikey=ae3e2a90-4cf6-11ec-baf0-fd76e1528414&base_currency=USD")
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		sb := string(body)
		var currency models.CurrencyDto
		err = json.Unmarshal([]byte(sb), &currency)
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}

		titleStr := fmt.Sprintf("USD to TRY: %f", currency.Data.TRY)
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: titleStr,
		})

		if _oldValue > (currency.Data.TRY+0.5) || _oldValue < (currency.Data.TRY-0.5) {
			_oldValue = currency.Data.TRY
			menuet.App().Notification(menuet.Notification{
				Title: fmt.Sprintf("Currency changed %f", currency.Data.TRY),
			})
		}
	}
}

var currencies = []menuet.MenuItem{
	menuet.MenuItem{
		Text: "10 secs",
		Clicked: func() {
			setInterval(models.TenSeconds)
		},
	},
	menuet.MenuItem{
		Text: "30 secs",
		Clicked: func() {
			setInterval(models.ThirtySeconds)
		},
	},
	menuet.MenuItem{
		Text: "1 min",
		Clicked: func() {
			setInterval(models.AMinute)
		},
	},
	menuet.MenuItem{
		Text: "10 minutes",
		Clicked: func() {
			setInterval(models.TenMinutes)
		},
	},
}

var intervals = []menuet.MenuItem{
	menuet.MenuItem{
		Text: "10 secs",
		Clicked: func() {
			setInterval(models.TenSeconds)
		},
	},
	menuet.MenuItem{
		Text: "30 secs",
		Clicked: func() {
			setInterval(models.ThirtySeconds)
		},
	},
	menuet.MenuItem{
		Text: "1 min",
		Clicked: func() {
			setInterval(models.AMinute)
		},
	},
	menuet.MenuItem{
		Text: "10 minutes",
		Clicked: func() {
			setInterval(models.TenMinutes)
		},
	},
}

func getCurrencies() []menuet.MenuItem {
	return currencies
}

func getIntervals() []menuet.MenuItem {
	return intervals
}

func setInterval(newIntervalValue int) {
	_interval = newIntervalValue
}

func menuItems() []menuet.MenuItem {
	return []menuet.MenuItem{
		menuet.MenuItem{
			Text:     "Interval",
			Children: getIntervals,
		},
		menuet.MenuItem{
			Text:     "Currency",
			Children: getCurrencies,
		},
	}
}

var _interval int
var _oldValue int

func main() {
	go getCurrenciesInfo()
	menuet.App().Label = "com.github.safakkizkin.osx-currency-go"
	menuet.App().Children = menuItems
	menuet.App().RunApplication()
}
