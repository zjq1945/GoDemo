package utility

import (
	"fmt"
	"github.com/spf13/viper"
)

func GetMySqlConnectionString() string {
	var rtn string
	v := viper.New()
	v.SetConfigName("testConfig.development")
	v.SetConfigType("json")
	v.AddConfigPath("config")

	err := v.ReadInConfig()
	if err != nil {
		fmt.Println("error in read in config: ", err)
		rtn = ""
	} else {
		rtn = v.GetString("ConsumerConnectionString")
		//fmt.Println("value is:", rtn)
	}
	return rtn
}
