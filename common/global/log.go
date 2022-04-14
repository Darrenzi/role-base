package global

import (
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
	switch strings.ToLower(Config.Logger.Level) {
	case "panic":
		{
			log.SetLevel(log.PanicLevel)
		}
	case "fatal":
		{
			log.SetLevel(log.FatalLevel)
		}
	case "error":
		{
			log.SetLevel(log.ErrorLevel)
		}
	case "info":
		{
			log.SetLevel(log.InfoLevel)
		}
	case "debug":
		{
			log.SetLevel(log.DebugLevel)
		}
	default:
		{
			log.SetLevel(log.ErrorLevel)
		}
	}

}
