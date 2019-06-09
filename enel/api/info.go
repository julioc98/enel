package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Info ...
func Info(token string) (string, error) {

	client := &http.Client{}

	values := map[string]interface{}{"I_CANAL": "ZINT", "I_COD_SERV": "TC", "I_SSO_GUID": ""}

	jsonValue, _ := json.Marshal(values)

	url := "https://portalhome.eneldistribuicaosp.com.br/api/sap/portalinfo"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Add("authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	return string(body), nil
}
