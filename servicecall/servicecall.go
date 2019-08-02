package servicecall

import (
	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetResponse() {
	response, err := http.Get("http://testapi01.azurewebsites.net/api/hello")
	if err != nil {
		fmt.Println("The http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
