package config

import (
	"bytes"
	_ "embed"
	"encoding/json"
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
	ServerUrl string `json:"server_url"`
}

var (
	//go:embed config.json
	conf []byte
	file *os.File
)

func New() (string, string, string, func()) {
	var data config
	err := json.NewDecoder(bytes.NewReader(conf)).Decode(&data)
	if err != nil {
		slog.Error(err.Error())
	}

	file, err := os.OpenFile(data.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	cleanup := func() {
		file.Close()
	}

	logger := slog.New(slog.NewJSONHandler(file, nil))
	slog.SetDefault(logger)

	return data.Server.Port, data.Database.Path, data.ServerUrl, cleanup
}
