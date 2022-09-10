package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type OutlineVpn struct {
	apiKey     string
	httpClient http.Client
}

func CreateOutlineVpn(apiUrl string) OutlineVpn {
	newClient := http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	return OutlineVpn{apiKey: apiUrl, httpClient: newClient}
}

func (api *OutlineVpn) GetKeys() (AccessKeys, error) {
	response, err := api.httpClient.Get(api.apiKey + "/access-keys/")
	if err != nil {
		fmt.Println(err)
		return AccessKeys{}, err
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return AccessKeys{}, err
	}
	var keys AccessKeys
	err = json.Unmarshal(all, &keys)
	if err != nil {
		fmt.Println(err)
		return AccessKeys{}, err
	}
	fmt.Println(keys)
	return keys, nil
}

func (api *OutlineVpn) CreateKey(keyName string) (Key, error) {
	response, err := api.httpClient.Post(api.apiKey+"/access-keys/", "", nil)
	if err != nil {
		return Key{}, err
	}
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return Key{}, err
	}
	key := Key{}
	err = json.Unmarshal(all, &key)
	if err != nil {
		return Key{}, err
	}
	err = api.RenameKey(key.Id, keyName)
	if err != nil {
		return Key{}, err
	}
	return key, nil
}

func (api *OutlineVpn) DeleteKey(keyId string) error {
	request, err := http.NewRequest(http.MethodDelete, api.apiKey+"/access-keys/"+keyId, nil)
	if err != nil {
		return err
	}
	_, err = api.httpClient.Do(request)
	if err != nil {
		return err
	}
	return nil
}

func (api *OutlineVpn) RenameKey(keyId string, keyName string) error {
	newName := Key{Name: keyName}
	body, err := json.Marshal(newName)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/access-keys/"+keyId+"/name", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	fmt.Println(request)
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) AddDataLimit(keyId string, byteLimit int) error {
	limit := Limit{Bytes{Bytes: byteLimit}}
	body, err := json.Marshal(limit)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/access-keys/"+keyId+"/data-limit", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) DeleteDataLimit(keyId string) error {
	request, err := http.NewRequest(http.MethodDelete, api.apiKey+"/access-keys/"+keyId+"/data-limit", nil)
	if err != nil {
		return err
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) GetTransferredData() error {
	response, err := api.httpClient.Get(api.apiKey + "/metrics/transfer")
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) GetServerInformation() (Server, error) {
	response, err := api.httpClient.Get(api.apiKey + "/server")
	if err != nil {
		return Server{}, err
	}
	datByte, err := io.ReadAll(response.Body)
	if err != nil {
		return Server{}, err
	}
	var serverInformation Server
	err = json.Unmarshal(datByte, &serverInformation)
	if err != nil {
		return Server{}, err
	}
	fmt.Println(response)
	return serverInformation, nil
}

func (api *OutlineVpn) SetServerName(serverName string) error {
	dataByte, err := json.Marshal("{\"name\":\"" + serverName + "\"}")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/name", bytes.NewBuffer(dataByte))
	if err != nil {
		return nil
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) SetHostName(hostName string) error {
	dataByte, err := json.Marshal("{\"hostName\":\"" + hostName + "\"}")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/name", bytes.NewBuffer(dataByte))
	if err != nil {
		return nil
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) GetMetricsStatus() (bool, error) {
	response, err := api.httpClient.Get(api.apiKey + "/metrics/enabled")
	if err != nil {
		return false, err
	}
	dataByte, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}
	stat := Status{}
	err = json.Unmarshal(dataByte, &stat)
	if err != nil {
		return false, err
	}
	return stat.metricksStatus, nil
}

func (api *OutlineVpn) SetMetricsStatus(status bool) error {
	dataByte, err := json.Marshal("{\"metricsEnabled\":" + strconv.FormatBool(status) + "}")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/metrics/enabled", bytes.NewBuffer(dataByte))
	if err != nil {
		return err
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) SetPortNewForAccessKeys(port string) error {
	dataByte, err := json.Marshal("{\"port\":\"" + port + "\"}")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/server/port-for-new-access-keys", bytes.NewBuffer(dataByte))
	if err != nil {
		return err
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) SetDataLimitForAllKeys(bytesLimit int) error {
	dataByte, err := json.Marshal("{\"limit\":" + strconv.Itoa(bytesLimit) + "}")
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPut, api.apiKey+"/server/access-key-data-limit", bytes.NewBuffer(dataByte))
	if err != nil {
		return err
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

func (api *OutlineVpn) DeleteDataLimitForAllKeys() error {
	request, err := http.NewRequest(http.MethodDelete, api.apiKey+"/server/access-key-data-limit", nil)
	if err != nil {
		return err
	}
	response, err := api.httpClient.Do(request)
	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}
