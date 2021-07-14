# API documentation

<!-- --------------------------------- REGISTER -->

The API works through a basic user system with session tokens [uuid](https://fr.wikipedia.org/wiki/Universally_unique_identifier).
Users have the following attributes:

- _pseudo_ (string)
- _password_ (string)
- _admin_ (boolean)
- _blocked_ (boolean)

> Administrators can block or unblock accounts (including other administrator accounts).

> Privileges (admin) cannot be changed from the API.

> An administrator \[admin:abcdABCD1234@\] is created by default.

# endpoints

## Register

To access the different API services (mainly the generation of a fizzbuzz) you must be registered with a valid account.

### Request

- method: POST
- url: /register
- content-type: application/json
- input format:

```JSON5
{
    "pseudo":"my_pseudo",
    "password":"my_password"
}
```

> pseudo must be between 3 and 20 characters

> password must be between 8 and 30 characters and must contain at least one lowercase, one uppercase, one number and one special character

### Response

- content-type: application/json

> Even if the registration is a success, no session token is returned to this step.

### Errors

```JSON5
{
  "error": "reason..."
}
```

<!-- --------------------------------- LOGIN -->

## Login

Once registered, you must authenticate yourself in order to obtain a session token which will take the form of a cookie.

### Request

- method: POST
- url: /login
- content-type: application/json
- input format:

```JSON5
{
    "pseudo":"my_pseudo",
    "password":"my_password"
}
```

### Response

- content-type: application/json

> In case of success, the session token is saved in the "session" cookie.

### Errors

```JSON5
{
  "error": "reason..."
}
```

## Fizzbuzz

The heart of this API, it's up to you to build your personalized fizzbuzz!

> To access this service, you must have a valid session token from the "session" cookie.

### Request

- method: GET
- url: /fizzbuzz
- query inputs:
  - _int1_ \* (integer)
  - _int2_ \* (integer)
  - _limit_ \* (integer)
  - _str1_ \* (string > max-length: 30)
  - _str2_ \* (string > max-length: 30)

> _int1_ can't be equal to _int2_

> _int1_ and _int2_ cannot be greater than _limit_

> _str1_ and _str2_ cannot be more than 30 characters

### Response

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

### Errors

```JSON5
{
  "error": "reason..."
}
```

## Stats

Are you curious about the most popular fizzbuzzes right now?

> To access this service, you must have a valid session token from the "session" cookie.

### Request

- method: GET
- url: /stats

### Response

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

### Errors

```JSON5
{
  "error": "reason..."
}
```

## Block

As said in the introduction, administrators can block or unblock accounts.

> To access this service, you must have a valid session token from the "session" cookie with administrator privileges.

### Request

- method: PATCH
- url: /block
- content-type: application/json
- input format:

```JSON5
{
    "pseudo":"his_pseudo",
    "block_status":"true" // need to be "true" or "false"
}
```

### Response

- content-type: application/json

### Errors

```JSON5
{
  "error": "reason..."
}
```
