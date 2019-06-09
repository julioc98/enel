package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Login ...
func Login(cpf, id string) (string, error) {
	values := map[string]interface{}{"I_CANAL": "ZINT", "I_COBRADORA": "", "I_CPF": cpf, "I_CNPJ": "", "I_ANLAGE": id, "I_COD_SERV": "TC", "I_LISTA_INST": "X", "I_BANDEIRA": "X", "I_FBIDTOKEN": "", "I_VERTRAG": "", "I_PARTNER": "", "I_IDPERGUNTA_01": "", "I_RESPOSTA_01": "", "I_IDPERGUNTA_02": "", "I_RESPOSTA_02": "", "I_EXEC_LOGIN": "X"}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("https://portalhome.eneldistribuicaosp.com.br/api/sap/getloginv2", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return "", err
	}

	return string(resp.Header.Get("Authorization")), nil
}
