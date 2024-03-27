Simple API service for card validation following Luhn algorithm

## Bootstrap service

```bash
cd test-task
docker build -t test-task .
docker run -p 8080:8000 -it --rm --name test-task test-task
```

## How to test

```bash
curl -X POST http://127.0.0.1:8080/api/v1/card/validate
   -H 'Content-Type: application/json'
   -d '{"number":"4111111111111111","month":"01", "year":"2021"}'
```
