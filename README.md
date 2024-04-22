# nft-markets
repository for backend services for application nft markets, backend services using technology stack go language, gin, and mongo database.

## Installation 
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

## Requirement
- Go language version 1.22.1
- Mongodb database

## Postman collection / API documentation
Postman is an application that allows the testing of web APIs. A comprehensive set of tools that help accelerate the API Lifecycle - from design, testing, documentation, and mocking to discovery.
- you can import postman collection from folder here
  ```
  /postman/nft-markets.postman_collection
  ```
  
## Endpoint API
- Get status API services health  
  **Request**  
  Method [`GET`] `/ping`
  ```
  $ curl --location 'http://localhost:8080/ping'
  ```

- Create new item of nft  
  **Request**  
  Method [`POST`] `/api/v1/nftmarkets/items/create` 
  ```
  $ curl --location 'http://localhost:8080/api/v1/nftmarkets/items/create' \
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

- Get data item of nft by id  
  **Request**  
  Method [`GET`] `/api/v1/nftmarkets/items/{id}` 
  ```
  $ curl --location 'http://localhost:8080/api/v1/nftmarkets/items/6624b678eb1edfb91ced9293d'
  ```
  
- Get all data item of nft  
  **Request**  
  Method [`GET`] `/api/v1/nftmarkets/items/`
  - query parameter available :
    - ?rating=4
    - ?reputationBadge=red
    - ?availability=1
    - ?category=animation 
  ```
  $ curl --location 'http://localhost:8080/api/v1/nftmarkets/items/?rating=4'
  ```    

- Update data item of nft by id  
  **Request**  
  Method [`PUT`] `/api/v1/nftmarkets/items/{id}`
  ```
  $ curl --location --request PUT 'http://localhost:8080/api/v1/nftmarkets/items/6625c8678675b7d624220fde' \
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

- Delete data item of nft by id  
  **Request**  
  Method [`DELETE`] `/api/v1/nftmarkets/items/{id}`
  ```
  $ curl --location --request DELETE 'http://localhost:8080/api/v1/nftmarkets/items/6624b678eb1edfb91ced9293'
  ```

- Update data availability if purchase item of nft by id  
  **Request**  
  Method [`PUT`] `/api/v1/nftmarkets/items/purchase/{id}`
  ```
  $ curl --location --request PUT 'http://localhost:8080/api/v1/nftmarkets/items/purchase/6624b678eb1edfb91ced9293' \
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