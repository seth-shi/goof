package utils

import (
	"github.com/fatih/color"
	"os"
)

type Log struct {}

var logInstance = &Log{}

func GetLogInstance() *Log {

	return logInstance
}

func (l *Log) Info(msg string) {

	color.Green(msg)
}

func (l *Log) Error(msg string) {

	color.Red(msg)
}

func (l *Log) Waring(msg string) {

	color.Yellow(msg)
}

func (l *Log) ErrorFatal(msg string)  {

	color.Red(msg)
	os.Exit(1)
}