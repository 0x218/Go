<h1> <p align="center"> <span style='font-weight:bold;align=center'>apiServer</span></p></h1>
This program is developed to demonstrate simple api testing.  


### API Testing Documentation:
---
```
GET /health → Should return "OK"
GET /users → List all users
GET /users?age=30 → Filter by age
GET /users?min_age=30 → Filter by age > 30
GET /users?max_age=30 → Filter by age <  30
GET /users?state=texas → Filter by state (case-insensitive)
POST /login
	and add below configuration under the body -> raw
	{
		"username": "admin",
		"password": "password"
	}
```

### Application execution using Go:
---
```
go run apiServer.goAPI Testing DocumentationGOOS=windows GOARCH=amd64 go build -o renderInWindows.exe main.go
```   

### Building application into an executable:
---
```
Windows: 
	go build -o apiServer.exe apiServer.go

Linux:
	set GOOS=linux
	set GOARCH=amd64
	go build -o apiServer apiServer.go
```  

#### Usage:  
---
a. Double click the apiServer (execuable)  - or use ```go run apiServer.go``` to run from the source code.
b. Open your browser and navigate to http://localhost:8080/
c. Open Postman and follow the "API Testing Documentation"
