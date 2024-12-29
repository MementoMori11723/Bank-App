package config

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"os"
)

type config struct {
	Database struct {
		Path string `json:"path"`
	} `json:"database"`
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Log string `json:"log"`
}

func New() (string, string) {
	conf, err := os.ReadFile("config.json")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var data config
	err = json.NewDecoder(bytes.NewReader(conf)).Decode(&data)
	if err != nil {
		slog.Error(err.Error())
	}

	file, err := os.OpenFile(data.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	w := io.MultiWriter(os.Stderr, file)
	logger := slog.New(slog.NewJSONHandler(w, nil))
	slog.SetDefault(logger)

	return data.Server.Port, data.Database.Path
}
