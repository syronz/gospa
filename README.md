# gospa

getting help from 
```
https://github.com/roberthodgen/spa-server
```

## Test proxy reverse
```bash
TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InN1cGVyIiwiaWQiOjc5LCJsYW5ndWFnZSI6Imt1IiwiY29tcGFueV9pZCI6MTAwMSwibm9kZV9pZCI6MTAxLCJleHAiOjE2MjE5MTc0NTJ9.LKu-_eauXAt7be-J8UMU_kZ0-tZbyGyMzTcck7RZEXU
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://127.0.0.1:7173/api/restapi/v1/companies/1001/roles
curl -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://127.0.0.1:8080/api/restapi/v1/companies/1001/roles
curl -H 'Accept: application/json' http://127.0.0.1:8080
```
