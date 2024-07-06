package zlog

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Source(err error) string {
	_, f, l, _ := runtime.Caller(0)
	return fmt.Sprintf("error: %v | source: %s | line: %d", err, f, l)
}

func Fatal(errR error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	go func() {
		file, err := os.OpenFile(
			"./log/fatal.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			slog.Warn(err.Error(), err)
		}
		defer file.Close()
		logger := zerolog.New(file).With().Timestamp().Logger()
		logger.Fatal().Err(err).Msg(errR.Error())
	}()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Fatal().Err(errR).Msg(errR.Error())
}

func Error(errR error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	go func() {
		file, err := os.OpenFile(
			"./log/error.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			slog.Warn(err.Error(), err)
		}
		defer file.Close()
		logger := zerolog.New(file).With().Timestamp().Logger()
		logger.Error().Stack().Err(errR).Msg(errR.Error())
	}()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Error().Stack().Err(errR).Msg(errR.Error())
}

func Warning(msg string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	go func() {
		file, err := os.OpenFile(
			"./log/warn.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			slog.Warn(err.Error(), err)
		}
		defer file.Close()
		logger := zerolog.New(file).With().Timestamp().Logger()
		logger.Warn().Msg(msg)
	}()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Warn().Msg(msg)
}

func Info(data interface{}, msg string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	go func() {
		file, err := os.OpenFile(
			"./log/info.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0664,
		)
		if err != nil {
			slog.Warn(err.Error(), err)
		}
		defer file.Close()
		logger := zerolog.New(file).With().Timestamp().Logger()
		logger.Warn().Msg(msg)
	}()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	dataMap := map[string]any{
		"data": data,
	}
	log.Info().Fields(dataMap).Msg(msg)
}
