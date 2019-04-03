# Car-API
REST API for cars

## Set Up
run app.exe

or 
`go run app.go` / `go build app.go`


## Routes

### GET /
Retrieve all cars

Sample Response

```JSON
{
    "status": "200",
    "cars": [
        {
            "id": 1,
            "category": "Sedan",
            "make": "Toyota",
            "model": "Vitz",
            "HP": 150,
            "year": 2006
        },
        {
            "id": 2,
            "category": "Sedan",
            "make": "Mercedes-Benz",
            "model": "s550",
            "HP": 740,
            "year": 2018
        },

    ],
    "count": 2
}
```

### GET /cars/:id
Retrieve car by id

Sample Response

```JSON
{
    "status": "200",
    "Car": {
        "id": 1,
        "category": "Sedan",
        "make": "Toyota",
        "model": "Vitz",
        "HP": 150,
        "year": 2006
    }
}
```

### GET /cars/category/:category
Retrieve cars by category (sedan or SUV)
