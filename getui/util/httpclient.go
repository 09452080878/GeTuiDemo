package getui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string, authToken string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header["authtoken"] = append(req.Header["authtoken"], authToken)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", ContentTypeJson)

	client := &http.Client{
		Timeout: DefaultConnectionTimeout * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func BytePost(url string, authToken string, bodyByte []byte) (string, error) {
	body := bytes.NewBuffer(bodyByte)

	req, err := http.NewRequest("POST", url, body)
	req.Header["authtoken"] = append(req.Header["authtoken"], authToken)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", ContentTypeJson)

	client := &http.Client{
		Timeout: DefaultConnectionTimeout * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func Delete(url string, authToken string, bodyByte []byte) (string, error) {

	body := bytes.NewBuffer(bodyByte)

	req, err := http.NewRequest("DELETE", url, body)
	req.Header["authtoken"] = append(req.Header["authtoken"], authToken)
	req.Header.Add("Charset", CHARSET)
	req.Header.Add("Content-Type", ContentTypeJson)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func GetBody(parmar interface{}) ([]byte, error) {

	body, err := json.Marshal(parmar)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}
