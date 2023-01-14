# Basic api in Golang, using Gin and Gorm
---
---
## Structure and technologies used
### What purpose?
> This project's main objective is to learn the Golang programming language and its packages.
> Initially creating a simple API and augmenting it with logs, tests, and endpoint documentation.
### Technologies used:
- [Golang](https://go.dev/learn/)
- [Docker](https://docs.docker.com/get-started/)
- [ShelScript](https://www.shellscript.sh/)
- [PostgreSQL](https://www.postgresql.org/about/)
- [PGAdmin](https://www.pgadmin.org/docs/)
- [Swagger](https://swagger.io/solutions/api-documentation/)

### Basic informations:
> This project contains 7 API's endpoints, 6 of which are directly linked with a CRUD process to a postgres database.

> For basic interaction, you can use Swagger (via browser) or Insomnia, importing the yaml that contains the initial for app.

> The postgres database and the pgadmin are already configured, ready for use.

> Some test cases were created that can be started with the ```go test``` command.

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
>The following private addresses are fixed in the Docker compose configuration, and give access to project options:

- ``` http://172.33.0.4:9000/api/v1/ ``` grants access to API and Swagger endpoints.
- ``` http://172.33.0.3/ ``` grants access to PGAdmin that already has the configuration with the database. Use for login in pgadmin: "admin@foobar.com" and "123456" and for confirmation database access password is "foobar".
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
- ```GET http://172.33.0.4:9000/api/v1/foobar/mock``` To Shown Mock
- ```POST http://172.33.0.4:9000/api/v1/foobar/``` To Create Foobar
- ```GET http://172.33.0.4:9000/api/v1/foobar/``` To Shown all Foobar
- ```GET http://172.33.0.4:9000/api/v1/foobar/:id``` To Shown Foobar by ID
- ```GET http://172.33.0.4:9000/api/v1/foobar/:reg``` To Shown Foobar by Registration
- ```PATCH http://172.33.0.4:9000/api/v1/foobar/:id``` To Edit Foobar by ID
- ```DELETE http://172.33.0.4:9000/api/v1/foobar/:id``` To Delete Foobar by ID
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
