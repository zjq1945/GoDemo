package consumer

import (
	"demo/db/mysql"
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
	var conumer jackymysql.Consumer
	err := json.Unmarshal(consumerBytes, &conumer)
	if err != nil {
		fmt.Println("error in unmarsha", err)
	}
	// fmt.Println("after unmarshal:", conumer)
	err = jackymysql.InsertConsumer(conumer)
	if err != nil {
		fmt.Println("error to insert:", err)
		return err
	}
	return nil
}
