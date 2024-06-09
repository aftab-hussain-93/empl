# Employee Management Service REST API

This is a simple Employee Management REST API Service.
Tech used - Go, Chi, Pgx, Docker, Docker Compose 

## Development Setup
Dependencies - Docker, Docker Compose, Makefile

Run development server - 
`make run`

API will be running on `http://localhost:8080`

Run unit tests -
`make test_unit`

Run e2e tests - requires docker daemon to be up
`make test_e2e` 


## Available Endpoints:

### List employees with Pagination ( GET request )
* `http://localhost:8080/api/employees?size=5&page=2`

* Response 
* * Success 
* * 200 OK  
```
    {
        "employees": [
            {
                "id": 37,
                "name": "manager 1",
                "position": "manager",
                "salary": 50000
            },
            {
                "id": 38,
                "name": "manager 2",
                "position": "manager",
                "salary": 50000
            },
            {
                "id": 4,
                "name": "trainee 1",
                "position": "trainee",
                "salary": 10000
            }
        ],
        "count": 3,
        "page": 1,
        "total": 3
    }
```
* * Failure 
* * 400 Bad Request
```
    {
        "error": {
            "code": "bad_request",
            "message": "error message"
        }
    }
```
* * 500 Internal
```
    {
        "error": {
            "code": "internal",
            "message": "Internal Server Error"
        }
    }
```

### Get employee by id: ( GET Request )
* `http://localhost:8080/api/employees/{employee_id}`

* Response 
* * Success 
* * 200 
```
    {
        "id":       123,
        "name":     "Employee name",
        "position": "trainee", 
        "salary":   42000
    }
```
* * Failure 
* * 404 Not Found
```
    {
        "error": {
            "code": "not_found",
            "message": "Record not found"
        }
    }
```

### Create Employee: ( POST Request )
* `http://localhost:8080/api/employees`

* Body
`
    {
        "name": "Employee name",
        "position": "trainee", -- enums - "manager" or "trainee"
        "salary": 42000
    }
`

* Response 
* * Success 
* * 201 Created
```
    {
        "id":       123,
        "name":     "Employee name",
        "position": "trainee", 
        "salary":   42000
    }
```
* * Failure 
* * 400 Bad Request
```
    {
        "error": {
            "code": "bad_request",
            "message": "error message"
        }
    }
```
* * 500 Internal
```
    {
        "error": {
            "code": "internal",
            "message": "Internal Server Error"
        }
    }
```

### Update Employee: ( PUT Request )
* `http://localhost:8080/api/employees/{employee_id}`

* Body
```
    {
        "name": "Employee name", // optional
        "position": "trainee",   // optional
        "salary": 42000          // optional
    }
```

* Response 
* * Success 
* * 200 OK
```
    {
        "id":       123,
        "name":     "Employee name",
        "position": "trainee", 
        "salary":   42000
    }
```
* * Failure 
* * 400 Bad Request
```
    {
        "error": {
            "code": "bad_request",
            "message": "error message"
        }
    }
```
* * 500 Internal
```
    {
        "error": {
            "code": "internal",
            "message": "Internal Server Error"
        }
    }
```

### Delete Employee: ( DELETE Request )
* `http://localhost:8080/api/employees/{employee_id}`

* Response 
* * Success 
* * 204 No Content

* * Failure 
* * 404 Not Found
```
    {
        "error": {
            "code": "not_found",
            "message": "Employee not found"
        }
    }
```
* * 500 Not Found
```
    {
        "error": {
            "code": "internal",
            "message": "Internal Server Error"
        }
    }
```