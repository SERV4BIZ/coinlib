package coinlib

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/SERV4BIZ/gfp/jsons"
)

// GetPrice is get cryto market price by coinlib.com
func GetPrice(txtKey string, txtCurrency string, txtSymbol string) (float64, error) {
	url := "https://coinlib.io/api/v1/coin?key=" + txtKey + "&pref=" + txtCurrency + "&symbol=" + txtSymbol
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return -1, err
	}

	res, err := client.Do(req)
	if err != nil {
		return -1, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return -1, err
	}
	jsoRes, errRes := jsons.ObjectString(string(body))
	if errRes != nil {
		return -1, errRes
	}

	f, err := strconv.ParseFloat(jsoRes.GetString("price"), 64)
	if err != nil {
		return -1, err
	}

	return f, nil
}
