package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type SecretKeyResponse struct {
	Key string `json:"key"`
}

var secretKey int64

func init() {
	secretKey = rand.Int63()
  log.Println("Secret Key in init function:", secretKey)
}

func GetSecretKey() string {
	return hex.EncodeToString([]byte(fmt.Sprintf("%d", secretKey)))
}

func createHash(data string) string {
	secretHash := []byte(fmt.Sprintf("%d", secretKey))
	h := hmac.New(sha256.New, secretHash)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateToken() string {
	randString := uuid.New().String()
	token := createHash(randString)
	return fmt.Sprintf("%s:%s", token, randString)
}

func validateToken(token string) bool {
	splitToken := strings.Split(token, ":")
	if len(splitToken) != 2 {
		return false
	}
	tokenHash := splitToken[0]
	checkHash := createHash(splitToken[1])
	return hmac.Equal([]byte(tokenHash), []byte(checkHash))
}

func BaseURL(baseURL string) {
	if baseURL != "" {
		req, err := http.NewRequest("GET", baseURL+"/health", nil)
		if err != nil {
			slog.Error(err.Error())
		}

    req.Header.Set("X-Request-Type", "secret")

		client := &http.Client{
			Timeout: time.Second * 10,
		}

		resp, err := client.Do(req)
		if err != nil {
			slog.Error(err.Error())
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error(err.Error())
		}

		var secretKeyResponse SecretKeyResponse
		err = json.Unmarshal(body, &secretKeyResponse)
		if err != nil {
			slog.Error(err.Error())
      os.Exit(1)
		}

		byteRes, err := hex.DecodeString(secretKeyResponse.Key)
		if err != nil {
			slog.Error(err.Error())
		}

		secretKey, err = strconv.ParseInt(string(byteRes), 10, 64)
		if err != nil {
			slog.Error(err.Error())
		}
    log.Println("Secret Key in BaseURL function:", secretKey)
	}
}
