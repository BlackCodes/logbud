package example

import (
	"io/ioutil"
	"log"

	"github.com/rs/zerolog"
	lgrus "github.com/sirupsen/logrus"
)

var l1 *zerolog.Logger

func Logrus()  {
	lgrus.SetReportCaller(true)
	lgrus.SetOutput(ioutil.Discard)
	lgrus.Info("this is logrus")
}

func zeroLogs() {
	if l1 == nil {
		logger := zerolog.New(ioutil.Discard)
		_l := zerolog.New(logger)
		l1 = &_l
	}
	l1.Info().Str("aa", "bb").Msg("hello good")
}

func zeroLogsCaller() {
	zl := zerolog.New(ioutil.Discard)
	zl.Info().Caller(3).Str("aa", "bb").Msg("Hello good")
}



func sysLog() {
	log.SetOutput(ioutil.Discard)
	log.Println("[../example/example_head.go:32:2] this sys println")
}

func specLog() {
	log.Fatal(2)
}
