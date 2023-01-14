# Basic api in Golang, using Gin and Gorm
---
---
## Structure and technologies used
### What purpose?
>The main objective of this project is to learn the Go programming language and explore its packages available in the documentation. Initially, with the creation of a simple API, and later making it more complex, inserting logs, tests and nested endpoint documentation.

### Technologies used:
- [Golang](https://go.dev/learn/)
- [Docker](https://docs.docker.com/get-started/)
- [ShelScript](https://www.shellscript.sh/)
- [PostgreSQL](https://www.postgresql.org/about/)
- [PGAdmin](https://www.pgadmin.org/docs/)
- [Swagger](https://swagger.io/solutions/api-documentation/)

### Basic informations:
> This project contains 7 API's endpoints, 6 of which are directly linked with a CRUD process to a postgres database.

> For basic interaction, you can use Swagger (via browser) or Insomnia (importing the yaml that contains the essential config to init the use).

> The postgres database and the pgadmin are already configured, ready for use.

> Some test cases were created that can be started with the command ```go test```.

> The Project has a mirrored structure in an MVC pattern.
---
---
## How to use
### What do I need to have installed on my computer?
> Docker and Docker Compose installed in your machine

> Any Linux operating system or Linux subsystem

- [Docker](https://docs.docker.com/get-started/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### How to run the project on my computer?
> The initial command can be the following:
```sh
chmod +x start.sh
./start.sh
```
> After the first run, it may be necessary to run an increment to clean up the database directory for new run's:
```sh
sudo rm -Rf postgres-data && ./start.sh
```
### How to access the available features?
>The following private addresses are fixed in the Docker-Compose configuration, and give access to project options:

- ``` http://172.33.0.4:9000/api/v1/ ``` grants access to API and Swagger endpoints.
- ``` http://172.33.0.3/ ``` grants access to PGAdmin that already has the configuration with the database. Use for login in pgadmin: "admin@foobar.com" and "123456" and for confirmation to access database, the password is "foobar".
- ``` 172.33.0.2 ``` is the IP address of the database, accessible with the credentials foobar:foobar:foobar on port 5432.
### Example of PGAdmin access:
> [pgadmin/login](http://172.33.0.3/login)
![](/img/pgadmin.png)
### Example of use with Swagger:
> [/api/v1/swagger/index.html](http://172.33.0.4:9000/api/v1/swagger/index.html)
![](/img/swagger.png)
### Example of use with Insomnia:
> In application options, from the Document or Collection name dropdown menu, select Import/Export. Select an option from the Import Data dropdown menu. Import the file "Insomnia_api-go-gin.yaml".
> For more information: [Insomnia Documentation](https://docs.insomnia.rest/insomnia/import-export-data)
![](/img/insomnia.png)
### Example of "Mock" endpoint access:
> [/api/v1/foobar/mock](http://172.33.0.4:9000/api/v1/foobar/mock)
![](/img/mock.png)
### Test generation examples
```sh
go test
```
![](/img/test.png)
---
---
## More API information
### List of endpoints
- ```GET http://172.33.0.4:9000/api/v1/foobar/mock``` To shown mock
- ```POST http://172.33.0.4:9000/api/v1/foobar/``` To create Foobar
- ```GET http://172.33.0.4:9000/api/v1/foobar/``` To shown all Foobar
- ```GET http://172.33.0.4:9000/api/v1/foobar/:id``` To shown Foobar by ID
- ```GET http://172.33.0.4:9000/api/v1/foobar/:reg``` To shown Foobar by Registration
- ```PATCH http://172.33.0.4:9000/api/v1/foobar/:id``` To edit Foobar by ID
- ```DELETE http://172.33.0.4:9000/api/v1/foobar/:id``` To delete Foobar by ID
- ```GET http://172.33.0.4:9000/api/v1/swagger/index.html``` Swagger
---
---
## References
### This project was created using the following Golang packages
- [Gin](https://pkg.go.dev/github.com/gin-gonic/gin)
- [GORM](https://pkg.go.dev/gorm.io/gorm)
- [Logrus](https://pkg.go.dev/github.com/sirupsen/logrus)
- [Testify](https://pkg.go.dev/github.com/stretchr/testify)
- [Testing](https://pkg.go.dev/testing)
- [SwagGo](https://pkg.go.dev/github.com/swaggo/swag)
- [ValidatorV2](https://pkg.go.dev/gopkg.in/validator.v2@v2.0.1)
---
---
_end of README.md_
