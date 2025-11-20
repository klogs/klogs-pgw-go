package klogs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	allowedChars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ticksAtUnixEpoch = 621355968000000000
)

var randomizer = rand.New(rand.NewSource(1))

func UrlFriendlyRandomString(length int) string {
	chars := make([]string, length)
	for i := 0; i < length; i++ {
		chars[i] = string(allowedChars[randomizer.Intn(len(allowedChars))])
	}

	return strings.Join(chars, "")
}

func UTCTicks() int64 {
	nanos := time.Now().UTC().UnixNano()

	return nanos/100 + ticksAtUnixEpoch
}

func HMACSHA256(cipherText, password string) string {
	mac := hmac.New(sha256.New, []byte(password))
	mac.Write([]byte(cipherText))

	hash := mac.Sum(nil)

	return hex.EncodeToString(hash)
}

func JsonEncode(value interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	if value != nil {
		err := json.NewEncoder(buf).Encode(value)

		if err != nil {
			return nil, err
		}
	}

	return buf, nil
}

func IsSuccessStatusCode(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode <= http.StatusIMUsed
}
