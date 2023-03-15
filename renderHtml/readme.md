<h1> <p align="center"> <span style='font-weight:bold;align=center'>renderHTML</span></p></h1>
This is developed as a simple fun program.  
The html files are created by keeing <b>automation exercise</b> in mind.
The <b>Go-lang</b> program was developed, so that I feel a hosted site.


### For users without Go-lang:
---
For those who wish to use this site/program but do not have a go-lang compiler, I built [renderInWindows.exe](https://github.com/0x218/Go/blob/master/renderHtml/renderInWindows.exe) for Windows-64 bit.  
```
GOOS=windows GOARCH=amd64 go build -o renderInWindows.exe main.go
```   

#### Usage:  
---
a. Double click the renderInWindows.exe file  
b. Open your browser  
c. Type http://localhost:3000/


### For users with Go-lang compiler:
---
* Run the program
```
go run main.go
```  
Then you open browser and navigate to http://localhost:3000/

* In case you wish to create output file as renderInWindows.exe:-
```
go build main.go -o renderInWindows.exe
```   
