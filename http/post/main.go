package main

import (
	"bytes"
	"crypto/rc4"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	passwordByte, err := RC4Encrypt([]byte("qweASD123"), []byte("LAtZFRCFmBTJ"))
	if err == nil {
		fmt.Println(string(passwordByte))
	}
	httpPostJson()

}

func RC4Encrypt(src, key []byte) ([]byte, error) {
	rc4Cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, len(src))
	rc4Cipher.XORKeyStream(cipherText, src)
	return []byte(base64.StdEncoding.EncodeToString(cipherText)), nil
}

func httpPostJson() {
	jsonStr := []byte(`{ "username": "root", "password": "E32rw1wfssDL" }`)
	url := "http://192.168.4.176:30089/v1/user/login"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	statuscode := resp.StatusCode
	hea := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	fmt.Println(statuscode)
	fmt.Println(hea["Access-Token"])

}
