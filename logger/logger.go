// Contains the definitions necessary to implement the application's log components
package log

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Colores disponibles para terminal
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
	MagentaBold = "\033[35;1m"
	CyanBold    = "\033[36m;1m"
)

// Level Nivel de log (None, Fatal, Error, Warn, Info, Debug)
type Level int

// Niveles de Log disponibles
const (
	NoneLevel Level = iota + 1
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// ConsoleWriter Interface define las funciones que debe implementar
// el componente que se encargue de escribir los logs en terminal
type ConsoleWriter interface {
	Print(...interface{})
}

// Config Define los parámetros de configuración del Logger de RCC
type Config struct {
	Colors bool
	Level  Level
}

var writer ConsoleWriter
var config Config

// msgColors Define el formato los mensajes en la terminal, con o sin colores
var msgColors map[string]string = map[string]string{
	"fatal":      "[FATAL] ",
	"fatalColor": RedBold + "[FATAL] " + Reset,
	"error":      "[ERROR] ",
	"errorColor": Red + "[ERROR] " + Reset,
	"warn":       "[WARN ] ",
	"warnColor":  Yellow + "[WARN ] " + Reset,
	"info":       "[INFO ] ",
	"infoColor":  Green + "[INFO ] " + Reset,
	"debug":      "[DEBUG] ",
	"debugColor": Cyan + "[DEBUG] " + Reset,
}

// init es la función que se ejecuta al cargar el package
func init() {
	writer = log.New(os.Stdout, "", log.LstdFlags)
	config = Config{
		Colors: true,
		Level:  NoneLevel,
	}
}

// StringToLevel Devuelve el Level correspondiente al string recibido por parámetro
func StringToLevel(levelString string) (*Level, error) {
	var level Level
	switch levelString {
	case "none":
		level = NoneLevel
	case "fatal":
		level = FatalLevel
	case "error":
		level = ErrorLevel
	case "warning":
		level = WarnLevel
	case "info":
		level = InfoLevel
	case "debug":
		level = DebugLevel
	default:
		return nil, errors.New("Unknown log level")
	}

	return &level, nil
}

// SetMode Establece el nivel de logs de la aplicación
func SetMode(level Level) {
	config.Level = level
}

// SetColors Establece si el log en la terminal debe ser en color
func SetColors(colors bool) {
	config.Colors = colors
}

// log Es la función privada que se encarga de imprimir el log en la terminal
func printLog(level string, args ...interface{}) {
	writer.Print(level + fmt.Sprint(args...))
}

// Info Es la función que genera logs de nivel de información
func Info(args ...interface{}) {
	if config.Level >= InfoLevel {
		var level string = msgColors["info"]
		if config.Colors {
			level = msgColors["infoColor"]
		}
		printLog(level, args...)
	}
}

// Infof Es la función que genera logs de nivel de información utilizando
// el formato del mensaje especificado
func Infof(format string, a ...interface{}) {
	if config.Level >= InfoLevel {
		var level string = msgColors["info"]
		if config.Colors {
			level = msgColors["infoColor"]
		}
		printLog(level, fmt.Sprintf(format, a...))
	}
}

// Warn Es la función que genera logs de nivel de advertencia
func Warn(args ...interface{}) {
	if config.Level >= WarnLevel {
		var level string = msgColors["warn"]
		if config.Colors {
			level = msgColors["warnColor"]
		}
		printLog(level, args...)
	}
}

// Warnf Es la función que genera logs de nivel de advertencia utilizando
// el formato del mensaje especificado
func Warnf(format string, a ...interface{}) {
	if config.Level >= WarnLevel {
		var level string = msgColors["warn"]
		if config.Colors {
			level = msgColors["warnColor"]
		}
		printLog(level, fmt.Sprintf(format, a...))
	}
}

// Error Es la función que genera logs de nivel de error
func Error(args ...interface{}) {
	if config.Level >= ErrorLevel {
		var level string = msgColors["error"]
		if config.Colors {
			level = msgColors["errorColor"]
		}
		printLog(level, args...)
	}
}

// Errorf Es la función que genera logs de nivel de error utilizando
// el formato del mensaje especificado
func Errorf(format string, a ...interface{}) {
	if config.Level >= ErrorLevel {
		var level string = msgColors["error"]
		if config.Colors {
			level = msgColors["errorColor"]
		}
		printLog(level, fmt.Sprintf(format, a...))
	}
}

// Debug Es la función que genera logs de nivel de depuración
func Debug(args ...interface{}) {
	if config.Level >= DebugLevel {
		var level string = msgColors["debug"]
		if config.Colors {
			level = msgColors["debugColor"]
		}
		printLog(level, args...)
	}
}

// Debugf Es la función que genera logs de nivel de depuración utilizando
// el formato del mensaje especificado
func Debugf(format string, a ...interface{}) {
	if config.Level >= DebugLevel {
		var level string = msgColors["debug"]
		if config.Colors {
			level = msgColors["debugColor"]
		}
		printLog(level, fmt.Sprintf(format, a...))
	}
}

// Fatal Es la función que genera logs de nivel de fatal
func Fatal(args ...interface{}) {
	if config.Level >= FatalLevel {
		var level string = msgColors["fatal"]
		if config.Colors {
			level = msgColors["fatalColor"]
		}
		printLog(level, args...)
		os.Exit(1)
	}
}

// Fatalf Es la función que genera logs de nivel de fatal utilizando
// el formato del mensaje especificado
func Fatalf(format string, a ...interface{}) {
	if config.Level >= FatalLevel {
		var level string = msgColors["fatal"]
		if config.Colors {
			level = msgColors["fatalColor"]
		}
		printLog(level, fmt.Sprintf(format, a...))
		os.Exit(1)
	}
}
