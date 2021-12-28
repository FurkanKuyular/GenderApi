# GenderApi

This is my first go lang. project i enjoy so much :).

and this service that brings gender information by name, 

Written with **https://go.dev/**

## Installation

* Install go language from https://go.dev/dl/
* Get clone project
* Fill the env file like below

  ``
  DATABASE_NAME=gender_dictionary
  ``

  ``
  DATABASE_DSN=username:password@tcp(mysqlIp:port)
  ``

  ``
  TABLE_NAME=gender
  ``

* Run the following command in order

   ``
   1.) go run mysql/db_config.go
   ``
   
   ![image](https://user-images.githubusercontent.com/52002022/147598798-7fd8f326-cd53-4c4f-9e13-7ce08a96ad4c.png)
   
   ``
   2.) go run csv/csv_import.go
   ``
   
   ![image](https://user-images.githubusercontent.com/52002022/147598916-4bef94f2-803e-42e1-a163-3f6dd2afb9dc.png)
   
## Usage

* Run the following command

   ``
   1.) go run gender.go
   ``
   
   2.) Open your browser and request http://localhost:8081/gender?name=furkan

   ```json
    {
       "success": true,
       "payload": {
        "name": "FURKAN",
        "gender": "M",
        "country": "TR"
      }
    }
   ```


**Note:  ``name`` query parameter is required to get payload if you not gonna add ``name`` query parameter probably can get error message**

 ```json
        {
          "success": false,
          "error": {
            "message": "Name is not exist"
            }
        }
 ```



