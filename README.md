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
  