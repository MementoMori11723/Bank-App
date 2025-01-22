package web

import (
	"bank-app/database/bank"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func get_data(url string, data []byte) (bank.Responce, error) {
  client := http.Client{
    Timeout: time.Second * 30,
  }

  req, err := http.NewRequest(http.MethodPost, baseUrl + url, bytes.NewBuffer(data))
  if err != nil {
    return bank.Responce{}, err
  }

  req.Header.Set("Content-Type", "application/json")
  
  res, err := client.Do(req)
  if err != nil {
    return bank.Responce{}, err
  }

  defer res.Body.Close()

  b, err := io.ReadAll(res.Body)
  if err != nil {
    return bank.Responce{}, err
  }

  var result bank.Responce
  err = json.Unmarshal(b, &result)
  if err != nil {
    return bank.Responce{}, err
  }

  return result, nil
}
