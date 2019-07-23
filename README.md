# cookie


## Request:
```
curl -i 'localhost:8080/cookie?toto=titi&tata=tutu' 
```
## Response:
```
HTTP/1.1 200 OK
Set-Cookie: sessionToken=82f34ca4-6c35-44a7-a033-73eff87e09ca; toto=titi; tata=tutu
Date: Tue, 23 Jul 2019 21:00:55 GMT
Content-Length: 91
Content-Type: text/plain; charset=utf-8

{"session":"82f34ca4-6c35-44a7-a033-73eff87e09ca","options":{"tata":"tutu","toto":"titi"}}
```

