### Description
This Demo is used to show how to implement a RESTful service by Go
### What have been shown in this demo
- How to do CRUD operations on MySql (mysql)
- How to do CRUD operations on MsSql (go-mssqldb)
- How to do CRUD operations on Rest Api (Gin) 
- How to use GORM for data access layer to MySql (gorm)
- How to consume a Restful Api
- How to cache (go-cache)
- How to use middleware to do some customized things ie. logging, authentication, caching response body
- How to read settings from config file (viper)
- How to do the unit test
### Following APIs have created for this demo
- /Consumer/GetAll                        (GET)
- /Consumer/Get/ConsumerName              (GET)
- /Consumer/Insert                        (POST)
- /Consumer/Update                        (PUT)
- /Consumer/Delete                        (DELETE)
- /Consumer/GetResponseFromOtherService   (GET)
