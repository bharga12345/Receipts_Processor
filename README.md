# Receipts_Processor

# Fetch Rewards Back-end Take home assessment

Language: Go

# API endpoints

  POST /receipts/process
  
      curl --location --request POST 'localhost:8080/receipts/process' \
          --header 'Content-Type: application/json' \
          --data-raw '{
              "retailer": "Walgreens",
              "purchaseDate": "2022-01-02",
              "purchaseTime": "08:13",
              "total": "2.65",
              "items": [
                  {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
                  {"shortDescription": "Dasani", "price": "1.40"}
              ]
          }'
 
  
  GET /receipts/{id}/points  
  
      curl --location --request GET 'localhost:8080/receipts/2d451eec-2bdf-485f-a4c9-7c1eec51c319/points'
  
  payload example-1:
  
{
    "retailer": "Target",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "13:13",
    "total": "1.25",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"}
    ]
}
  
OUTPUT:  

{
    "points": 31
}

payload example-2:

{
    "retailer": "Walgreens",
    "purchaseDate": "2022-01-02",
    "purchaseTime": "08:13",
    "total": "2.65",
    "items": [
        {"shortDescription": "Pepsi - 12-oz", "price": "1.25"},
        {"shortDescription": "Dasani", "price": "1.40"}
    ]
}

OUTPUT:

{
    "points": 15
}


