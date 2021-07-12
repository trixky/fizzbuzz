# API documentation

<!-- --------------------------------- REGISTER -->

The API works through a basic user system with session tokens (uuid).
Users have the following attributes:

- pseudo (string)
- password (string)
- admin (boolean)
- blocked (boolean)

> Administrators can block or unblock accounts (including other administrator accounts).

> Privileges (admin) cannot be changed from the API.

# endpoints

## /register

To access the different API services (mainly the generation of a fizzbuzz) you must be registered with a valid account.

> Even if the registration is a success, no session token is returned to this step.

### request

- method: POST
- url: /register
- content-type: application/x-www-form-urlencoded
- input keys:
  - pseudo * (string)
  - password * (string)

### response

- content-type: application/json

### errors

```JSON5
{
  "error": "reason..."
}
```

<!-- --------------------------------- LOGIN -->

## /login

Once registered, you must authenticate yourself in order to obtain a session token which will take the form of a cookie.

### request

- method: POST
- url: /login
- content-type: application/x-www-form-urlencoded
- input keys:
  - pseudo * (string)
  - password * (string)

### response

- content-type: application/json

> In case of success, the session token is saved in the "session" cookie.

### errors

```JSON5
{
  "error": "reason..."
}
```

## /fizzbuzz

The heart of this API, it's up to you to build your personalized fizzbuzz!

### request

- method: GET
- url: /fizzbuzz
- content-type: application/x-www-form-urlencoded
- input keys:
  - int1 * (integer)
  - int2 * (integer)
  - limit * (integer)
  - str1 * (string > max-length: 30)
  - str2 * (string > max-length: 30)

> To access this service, you must have a valid session token from the "session" cookie.

### response

- content-type: application/json

```JSON5
{
    "fizzbuzz": [
        "1",
        "2",
        "fizz",
        "4",
        "buzz",
        "fizz",
        "7",
        "8",
        "fizz",
        "buzz",
        "11",
        "fizz",
        "13",
        "14",
        "fizzbuzz"
        // ...
    ]
}
```

### errors

```JSON5
{
  "error": "reason..."
}
```

or

```JSON5
{
  "errors": [
      "first reason...",
      "another reason..."
      // ...
    ]
}
```

## /stats

Are you curious about the most popular fizzbuzzes right now?

### request

- method: GET
- url: /stats

> To access this service, you must have a valid session token from the "session" cookie.

### response

- content-type: application/json

```JSON5
{
    "requests": [
        {
            "int1": 2,
            "int2": 3,
            "limit": 50,
            "str1": "pop",
            "str2": "corn",
            "request_number": 12
        },
        {
            "int1": 3,
            "int2": 5,
            "limit": 100,
            "str1": "fizz",
            "str2": "buzz",
            "request_number": 9
        },
        {
            "int1": 3,
            "int2": 10,
            "limit": 60,
            "str1": "su",
            "str2": "pra",
            "request_number": 3
        }
        // ... (10 requests maximum)
    ]
}
```

### errors

```JSON5
{
  "error": "reason..."
}
```

## /block

As said in the introduction, administrators can block or unblock accounts.

### request

- method: PATCH
- url: /block
- content-type: application/x-www-form-urlencoded
- input keys:
  - pseudo * (string)
  - block_status * (boolean > ["true","false"])

> To access this service, you must have a valid session token from the "session" cookie with administrator privileges.

### response

- content-type: application/json

### errors

```JSON5
{
  "error": "reason..."
}
```