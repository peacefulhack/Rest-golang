# Product API Documentation

This document provides information on the Product API endpoints and their usage.

## Table of Contents
- [Insert Product](#insert-product)
- [Sort Products](#sort-products)

## Insert Product

### Endpoint
```
POST /product/insert
```

This endpoint is used to insert a new product into the system.

#### Request Body
- `ProductName` (string): The name of the product.
- `ProductDesc` (string): A description of the product.
- `ProductPrice` (int): The price of the product.
- `ProductQuantity` (int): The quantity of the product.

#### Request Example
```json
{
  "ProductName": "Example Product",
  "ProductDesc": "This is an example product description.",
  "ProductPrice": 20,
  "ProductQuantity": 100
}
```

#### Response
Upon successful insertion of the product, the API will respond with a status code of `200 OK`.

### Sort Products

### Endpoint
```
GET /product/sort/:key/:order
```

This endpoint is used to retrieve a list of products sorted by the specified key and order.

#### Parameters
- `key` (string): The key to sort the products by. Accepted values are:
  - "name" (sort by product name)
  - "price" (sort by product price)
  - "new" (sort by product ID in ascending order)
- `order` (string): The order in which to sort the products. Accepted values are:
  - "asc" (ascending order)
  - "desc" (descending order)

#### Request Example
```
GET /product/sort/price/asc
```

#### Response
The API will respond with a list of products sorted according to the specified key and order. Each product object in the list will have the following attributes:

- `Id` (int): The unique identifier of the product.
- `ProductName` (string): The name of the product.
- `ProductDesc` (string): A description of the product.
- `ProductPrice` (int): The price of the product.
- `ProductQuantity` (int): The quantity of the product.

#### Response Example
```json
[
  {
    "Id": 1,
    "ProductName": "Product A",
    "ProductDesc": "Description for Product A",
    "ProductPrice": 25,
    "ProductQuantity": 50
  },
  {
    "Id": 2,
    "ProductName": "Product B",
    "ProductDesc": "Description for Product B",
    "ProductPrice": 30,
    "ProductQuantity": 75
  }
]
```
