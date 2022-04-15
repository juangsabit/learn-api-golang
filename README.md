# Basic CRUD Restful API with Golang
![Go Language](http://marcio.io/img/gopher.png?raw=true)
this api does a simple CRUD operations on a single table mysql Database .<br/> <br/>
this is build on top off gin framwork  and GORM  library <br/>

[link to GORM!](https://github.com/go-gorm/gorm/) <br/>
[link to gin!](https://github.com/gin-gonic/gin)  <br/>

the api use different HTTP methods ( GET, POST, PUT, PATCH, and DELETE ) <br>

to make this work correctly you need to :
* you need to clone gin and gorm <br/>
* you need to have mysql installed in your machine ( install xampp if you don't have ) <br/>
* create a new empty database and call it 'api-go' <br/><br/>

to add the gin framwork (copy and past this into ur terminal)
```Go
go get "https://github.com/gin-gonic/gin"
```
to add the GORM library (copy and past this into ur terminal)
```GO
go get "https://github.com/go-gorm/gorm"
```
______________________________________________________________________________________

Listening and serving HTTP on :8888
* `SINGLE GET   /v1/api/book/:id`             --> main.getById
* ` GET ALL     /v1/api/books`                --> main.getAll
* ` POST        /v1api/book`                  --> main.add
* ` PUT         /v1api/book/:id`              --> main.update
* ` DELETE      /v1api/book/:id`              --> main.delete

______________________________________________________________________________________

Thank you..~
