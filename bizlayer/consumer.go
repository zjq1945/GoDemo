package consumer

import (
	"demo/db/mysql"
	"demo/servicecall"
	"encoding/json"
	"errors"
	"fmt"
)

func GetAll() []jackymysql.Consumer {

	return jackymysql.GetConsumers()
}

func Get(name string) (jackymysql.Consumer, bool) {

	consumer, result := jackymysql.GetConsumer(name)
	return consumer, result
}

func Delete(name string) error {

	_, result := Get(name)
	if result {
		jackymysql.DeleteConsumer(name)
		return nil
	} else {
		return errors.New("not found, deleting is failed")
	}
}

func Insert(consumerBytes []byte) error {
	var consumer jackymysql.Consumer
	err := json.Unmarshal(consumerBytes, &consumer)
	if err != nil {
		fmt.Println("error in unmarsha", err)
	}
	// fmt.Println("after unmarshal:", conumer)
	err = jackymysql.InsertConsumer(consumer)
	if err != nil {
		fmt.Println("error to insert:", err)
		return err
	}
	return nil
}

func Update(consumerBytes []byte) error {
	var consumer jackymysql.Consumer
	err := json.Unmarshal(consumerBytes, &consumer)
	if err != nil {
		fmt.Println("error in unmarsha", err)
	}
	err = jackymysql.UpdateConsumer(consumer)
	if err != nil {
		fmt.Println("error in update", err)
		return err
	} else {
		return nil
	}
}

func GetResponseFromOutsideService() ([]servicecall.Goods, error) {
	rtn, err := servicecall.GetResponse()
	if err != nil {
		return nil, errors.New("service call failed")
	} else {
		return rtn, nil
	}
}
