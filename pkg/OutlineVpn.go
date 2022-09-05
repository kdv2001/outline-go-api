package pkg

import (
	"fmt"
	"net/http"
)

type OutlineVpn struct {
	apiKey string
}

func CreateOutlineVpn(apiUrl string) OutlineVpn {
	return OutlineVpn{apiKey: apiUrl}
}

func (api *OutlineVpn) GetKeys() {
	//http.get
	resp, err := http.Get("https://198.244.248.183:65164/tUNDCPIJc5IlHiG-fXVGRg/access-keys/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp)
}

func (api *OutlineVpn) CreateKey() {

}

func (api *OutlineVpn) DeleteKey() {

}

func (api *OutlineVpn) RenameKey() {

}

func (api *OutlineVpn) AddDataLimit() {

}

func (api *OutlineVpn) DeleteDataLimit() {

}

func (api *OutlineVpn) GetTransferredData() {

}

func (api *OutlineVpn) GetServerInformation() {

}

func (api *OutlineVpn) SetServerName() {

}

func (api *OutlineVpn) SetHostName() {

}

func (api *OutlineVpn) GetMetricsStatus() {

}

func (api *OutlineVpn) SetPortNewForAccessKeys() {

}

func (api *OutlineVpn) SetDataLimitForAllKeys() {

}

func (api *OutlineVpn) DeleteDataLimitForAllKeys() {

}
