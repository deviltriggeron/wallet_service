package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
)

func TestDepositNonExistentID(t *testing.T) {
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

func TestDepositExistingID(t *testing.T) {
	walletId, err := createWalletsAndReturnID()
	if err != nil {
		t.Fatal(err)
	}

	jsonData := []byte(fmt.Sprintf(`{
		"walletId": "%s",
		"operationType": "DEPOSIT",
		"amount": 500
	}`, walletId))

	url := "http://localhost:8080/api/v1/wallet"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestWithDrawNonExistentID(t *testing.T) {
	newID, _ := uuid.NewUUID()
	jsonData := fmt.Appendf(nil, `{
		"walletId": "%s",
		"operationType": "WITHDRAW",
		"amount": 500
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

func TestWithDrawExistingID(t *testing.T) {
	walletId, err := createWalletsAndReturnID()
	if err != nil {
		t.Fatal(err)
	}
	jsonData := []byte(fmt.Sprintf(`{
		"walletId": "%s",
		"operationType": "WITHDRAW",
		"amount": 500
	}`, walletId))

	url := "http://localhost:8080/api/v1/wallet"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestGetNonExistentID(t *testing.T) {
	newID, _ := uuid.NewUUID()
	url := "http://localhost:8080/api/v1/wallets/" + newID.String()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestGetExistingID(t *testing.T) {
	walletId, err := createWalletsAndReturnID()
	if err != nil {
		panic(err)
	}

	url := "http://localhost:8080/api/v1/wallets/" + walletId.String()

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func TestGet(t *testing.T) {
	url := "http://localhost:8080"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func createWalletsAndReturnID() (uuid.UUID, error) {
	url := "http://localhost:8080/api/v1/createWallets"
	jsonData := []byte(fmt.Sprintf(`{
	"amount": 1000
	}`))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var respStruct struct {
		WalletID string `json:"walletId"`
		Balance  int    `json:"balance"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respStruct); err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(respStruct.WalletID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
