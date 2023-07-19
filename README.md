# TEST CACHE  (golang)

## General description

This web application is built using GO and provides the ability to cache items. Only for educational proposal practice.

## Usage

```
 go run main.go
 ``````

# Send local request to populate cache

```
http://127.0.0.1:8090/q=test
``````

# Dump cache contents
```
http://127.0.0.1:8090/dump
```

# Response dump content log in console.
```
L: Request=https://postman-echo.com/get?q=test
L: Itm Cached: false Id: 1569778511168729065 Addr: 0xc0002040a0 Time: 531.628044ms
L: Itm Cached: true Id: 1569778511168729065 Addr: 0xc0002040a0 Time: 4.552µs
```

