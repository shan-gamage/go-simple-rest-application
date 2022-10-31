# go-simple-rest-application

1. Clone the project
2. Create database "cryptoprice" and use the migration.sql to create table
3. And run " go run . "



API Documentation


This porject will be run on port 8080


1. Get latest Price

  http://localhost:8080/getprice/last?type=BTC

Method      : GET
Parameters  : type="BTC"

RESPONSE
{
  "RATE": 0.00,
  "CreatedAt": "2006-01-02 15:04:05"
  "Message": "SUCCESS/system error"
}

2. Get price by timestamp

Method      : GET
Parameters  : timestamp="1666964013"

REQUEST

  http://localhost:8080/getprice/bytimestamp?timestamp=1666964013

RESPONSE

If exact match

 {
    "RATE": 123.32123,
    "CreatedAt": "2006-01-02 15:04:05"
    "Message": "SUCCESS/system error".
    "NearestRate" : 0.00
 }

If Nearest Match 
 {
    "RATE": 0.00,
    "CreatedAt": "2006-01-02 15:04:05"
    "Message": "SUCCESS/system error".
    "NearestRate" : 123.32123
 }
 
 3. Get Average

Method : GET

Parameters :  starttime="Start time stamp"
              endDate= "End time stamp"
              
REQUEST
  
  http://localhost:8002/getprice/average?starttime=1666964013&endtime=1666964133
  
  
RESPONSE

{
  "RATE": 0.00,
  "Message": "SUCCESS/system error"
}
       
