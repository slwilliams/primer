# primer
REST service calculating prime number, good for infrastructure scaling demo 

## Start

```
make run
```

## Request

To calculate the highest prime number under the default (50000000)

```
curl http://localhost:8080 
```

You can also pass max number as an argument. Note, if you pass too high of a number the HTTP request will timeout.

```
curl http://localhost:8080/5000000
```

## Response

`prime.val` in the response has the highest prime number bellow passed in argument. 

```
{
    "id": "98e26178-7f90-497c-8189-88054d8bd542",
    "prime": {
        "max": 5000000,
        "val": 4999963
    },
    "host": "machinename.local",
    "ts": "2018-02-18 23:28:37.961918 +0000 UTC"
}
```