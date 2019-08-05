package servicecall

import (
	"demo/utility"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func GetResponse() ([]Goods, error) {
	value, found := utility.GetCache("goodsList")
	if found {
		fmt.Println("hit in cache ")
		return value.([]Goods), nil
	}
	response, err := http.Get("http://testapi01.azurewebsites.net/api/hello")
	if err != nil {
		fmt.Println("The http request failed with error %s\n", err)
		return nil, errors.New("got an error in the service call")
	} else {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error read body", err)
			return nil, err
		}
		var goodsList []Goods
		err = json.Unmarshal(body, &goodsList)
		if err != nil {
			fmt.Println("error in unmarshal Goods", err)
			return nil, err
		}
		utility.AddCache("goodsList", goodsList, 5*time.Minute)
		fmt.Println("get list via service call")
		return goodsList, nil

	}
}
