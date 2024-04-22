# nft-markets
repository for backend services for application nft markets

## installation 
- install all dependency with typing the command below
  ```
  $ go install 
  ```
- updated all dependency with typing the command below
  ```
  $ go mod tidy
  ```
- run the application with the command below
  ```
  $ go run main.go 
  
  or

  // if you have install gin for live reload your running development
  
  $ gin -a 8080 -i run main.go 
  ```

## requirement
- go language version 1.22.1
- mongo database

## postman collection
- you can import postman collection from folder
  ```
  /postman/nft-markets.postman_collection
  ```
  
## endpoint API
- Method GET -> `http://localhost:8080/ping` : check api services health
- Method POST -> `http://localhost:8080/api/v1/nftmarkets/items/create` : create new
  ```
  curl --location 'http://localhost:8080/api/v1/nftmarkets/items/create' \
  --header 'Content-Type: application/json' \
  --data '{
      "item_name": "NFT: Photo Mountain in Indonesia",
      "rating": 5,
      "category": "photo",
      "image": "https://risyandi.com/",
      "reputation": 800,
      "price": 25000000,
      "availability": 2
  }'
  ```
- Method GET -> `http://localhost:8080/api/v1/nftmarkets/items/{id}` : get by id
  ```
  curl --location 'http://localhost:8080/api/v1/nftmarkets/items/6624b678eb1edfb91ced9293d'
  ```
- Method GET -> `http://localhost:8080/api/v1/nftmarkets/items/` : get all 
  - query parameter available :
    - ?rating=4
    - ?reputationBadge=red
    - ?availability=1
    - ?category=animation 
  ```
  curl --location 'http://localhost:8080/api/v1/nftmarkets/items/?rating=4'
  ```    
- Method PUT -> `http://localhost:8080/api/v1/nftmarkets/items/{id}`: update
  ```
  curl --location --request PUT 'http://localhost:8080/api/v1/nftmarkets/items/6625c8678675b7d624220fde' \
  --header 'Content-Type: application/json' \
  --data '{
      "item_name": "EDIT NFT: Photo Mountain in Indonesia",
      "rating": 5,
      "category": "photo",
      "image": "https://risyandi.com/",
      "reputation": 500,
      "price": 15000000,
      "availability": 4
  }'
  ```
- Method DELETE -> `http://localhost:8080/api/v1/nftmarkets/items/{id}`: update
  ```
  curl --location --request DELETE 'http://localhost:8080/api/v1/nftmarkets/items/6624b678eb1edfb91ced9293'
  ```
- Method PUT -> `http://localhost:8080/api/v1/nftmarkets/items/purchase/{id}` : purchase item
  ```
  curl --location --request PUT 'http://localhost:8080/api/v1/nftmarkets/items/purchase/6624b678eb1edfb91ced9293' \
  --header 'Content-Type: application/json' \
  --data '{
      "item_name": "Photo Sepeda",
      "rating": 5,
      "category": "sport",
      "image": "url://",
      "reputation": 6,
      "price": 10000,
      "availability": 1
  }'
  ```