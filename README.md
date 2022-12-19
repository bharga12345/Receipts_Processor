# Receipts_Processor

# Fetch Rewards Back-end Take home assessment

Language: Go

# API endpoints

  POST /receipts/process
  
  GET /receipts/{id}/points
  
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


