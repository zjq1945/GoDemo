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
	value, found := utility.GetCache("Jacky")
	if found {
		fmt.Println("in Get response, found cache ", value)
	} else {
		fmt.Println("in Get response, not found cache")

		utility.AddCache("Jacky", "hi there", 5*time.Minute)
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
		} else {
			var goodsList []Goods
			err := json.Unmarshal(body, &goodsList)
			if err != nil {
				fmt.Println("error in unmarshal Goods", err)
				return nil, err
			} else {
				return goodsList, nil
			}
		}
	}
}
