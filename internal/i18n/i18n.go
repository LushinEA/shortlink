package i18n

import "github.com/cloudfoundry/jibber_jabber"

type LocaleString struct {
	En string `json:"en"`
	Ru string `json:"ru"`
}

var currentLocale string

func init() {
	userLanguage, _ := jibber_jabber.DetectLanguage()
	SetLocale(userLanguage)
}

func _v(en, ru string) LocaleString {
	return LocaleString{En: en, Ru: ru}
}

func SetLocale(locale string) {
	switch locale {
	case "ru":
		currentLocale = "ru"
	case "en":
		currentLocale = "en"
	default:
		currentLocale = "en"
	}
}

func GetLocale() string {
	switch currentLocale {
	case "ru":
		return "ru"
	default:
		return "en"
	}
}

func (locStr LocaleString) Locale() string {
	return locStr.TempLocale(currentLocale)
}

func (locStr LocaleString) TempLocale(loc string) string {
	switch loc {
	case "ru":
		return locStr.Ru
	default:
		return locStr.En
	}
}

var AppName = LocaleString{
	En: "shortlink",
	Ru: "shortlink",
}

var AppUsage = LocaleString{
	En: "A simple URL shortening service",
	Ru: "Простой сервис для сокращения URL",
}

var StartCommandUsage = LocaleString{
	En: "Start the shortlink server",
	Ru: "Запустить сервер коротких ссылок",
}

var FlagHostUsage = LocaleString{
	En: "Database host",
	Ru: "Хост базы данных",
}

var FlagPortUsage = LocaleString{
	En: "Database port",
	Ru: "Порт базы данных",
}

var FlagUserUsage = LocaleString{
	En: "Database user",
	Ru: "Пользователь базы данных",
}

var FlagPasswordUsage = LocaleString{
	En: "Database password",
	Ru: "Пароль базы данных",
}

var FlagDBNameUsage = LocaleString{
	En: "Database name",
	Ru: "Имя базы данных",
}

var FlagSSLModeUsage = LocaleString{
	En: "Database SSL mode",
	Ru: "Режим SSL базы данных",
}

var FlagLangUsage = LocaleString{
	En: "Language (e.g., en, ru)",
	Ru: "Язык (например, en, ru)",
}

var ConnectedToDatabase = LocaleString{
	En: "Successfully connected to database",
	Ru: "Успешно подключено к базе данных",
}
