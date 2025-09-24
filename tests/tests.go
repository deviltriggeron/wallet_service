package tests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestPostNewID(t *testing.T) {
	newID, _ := uuid.NewUUID()
	jsonData := fmt.Appendf(nil, `{
		"walletId": "%s",
		"operationType": "DEPOSIT",
		"amount": 1000
	}`, newID)

	url := "http://localhost:8080/api/v1/wallet"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestPost(t *testing.T) {

}

func TestGetExistingID(t *testing.T) {
	walletId := "123e4567-e89b-12d3-a456-426614174000"
	url := "http://localhost:8080/api/v1/wallets" + walletId

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

}
