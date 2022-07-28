# Go_Sample_project1

This is a simple practice level web project in Go-lang, perform CRUD operation.  Databases are mimicked with the help of struct and slice.

Data relation: 'Movie' HAS a 'Director'.

'sliceMovie' is a slice of 'Movie'. 

Defining 5 routes: 
 POST /movies         - Create a movie
 GET /movies          - Read all movies
 GET /movies/<id>     - Read a specific movie
 PUT /movies/<id>     - Update a movie
 DELETE /movies/<id>  - Delete a movie
 
 Starting a web server is achieved with http.ListenAndServe()
 
