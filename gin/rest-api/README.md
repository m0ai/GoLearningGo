# Go REST

A simple REST API to upload and download files


## APIs

- `POST /api/v1/files/` Upload a file 
- `GET /api/v1/files/:name` Download a file with a name


## Quick Test 

```bash
# Terminal 1 
go run main.go 

# Terminal 2
curl -X POST http://localhost:8080/api/v1/files/ -F file=@/Users/moai/test.c
```

