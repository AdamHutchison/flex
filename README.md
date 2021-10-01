# Flux Framework

The Flux framework is a work in progress and has been designed as a personal starting point for building web API's.

## Modules
Flux consists of many smaller modules that handle specific pieces of the frameworks functionality. See the readme of each for documentation:

* [Flux Config](https://github.com/AdamHutchison/flux-config) - Handles the loading and retrival of configuration / env values
* [Gorilla Mux](https://github.com/gorilla/mux) - Handles routing
* [Gorilla Handlers](https://github.com/gorilla/handlers) - Adds useful middleware handlers
* [Gorilla Schema](https://github.com/gorilla/schema) - Used in request validation
* [Go Playground Validator](https://github.com/go-playground/validator) - Used in request validation
* [GORM](https://gorm.io/index.html) - ORM used to handle all things model and database related

## Routing

### Registration
Routes are registered in the `RegisterRoutes()` function contained within the `http/routes/routes.go` file. An instance of the [Gorilla Mux](https://github.com/gorilla/mux) is injected into this function and you are free to use it as you see fit.

### Handlers
Handlers are located in the `/http/handlers` directory. A base handler has been added for your convenience that implements a `Respond` method. This is a simple helper to allow you to send API responses easily. It is up to you if you use it.

### Request Validation
Request bodies and query strings can be easily validated using the the [Go Playground Validator](https://github.com/go-playground/validator) package. To add validation to a request firstly create a validator struct containing the rules you wish to apply:

```go
package validators

type HomeValidator struct {
	Count string `validate:"required,eq=40"`
}
```

You may then use the `Validate(validator interface{}, w http.ResponseWriter, r *http.Request) error` method on the `BaseHandler`:

```go
func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Validate(new(validators.HomeValidator), w, r)

	if err != nil {
		h.Error(err, w, http.StatusBadRequest)
		return
	}

	data := transformers.HomeTransformer{
		Message: "Welcome to your new flux app",
	}

	h.Respond(data, w, http.StatusOK)
}
```

### Response Transformers
Transformers are just structs that allow you to structure you api responses easily. These are located in the `http/transformers` directory. The `Response`method on the `BaseHandler` will expect a transformer struct and will automatically wrap it in the `BaseTransformer` allowing you to easily control how you response data is structured.

### Middleware
Middleware should be place in the `http/middleware` directory and global middleware should be registered within the `RegisterGlobalMiddleware()` function in the `http/middleware/middleware.go` file. Middleware registered here will be applied to all routes.

## Database
Flux uses the GORM to allow you to manage anything database related.

### Creating Models
Models are located in the `database/models` directory and are simply structs that contain a `gorm.Model` embedded struct. These are just standard GORM models and more information can be found about them in the [GORM documentation](https://gorm.io/docs/models.html).

```go
package models

import (
	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Email  string
  Name string
}
```

### Obtaining a database instance
Database credentials can be set in the .env file as follows:
```bash
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=flux
```

You can then obtain a GORM DB connection instance by calling the `database.DB()` function:

```go
func (h HomeHandler) Create(w http.ResponseWriter, r *http.Request) {
	db := database.DB()

	user := models.User{Name: "Jim", Email: "jim@test.com", Password: "Super Secure"}

	db.Create(&user)

	h.Respond(user, w, http.StatusOK)
}
```

### Migrations

There are two options for migrations in flux, models can either be auto migrated by registering them in the `RegisterAutoMigrations()` function contained within `database/migrations/migrate.go` file, e.g:

```go
func RegisterAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(new(models.User))
}
```

or you can manually add a bespoke migration by adding code to the `RegisterStandardMigrations()` within `database/migrations/migrate.go`. It is recommended that the actual code for running the migration is extracted to it's own function and then only called from this here to help keep the file maintainable over time. e.g here we have a `database/migrations/migrate.go` like so:

```go
package migrations

import (
	"github.com/AdamHutchison/flux/database"
	"github.com/AdamHutchison/flux/database/models"
)

func RunUserMigration(db *gorm.DB) {
	db.Migrator().DropColumn(&models.User{}, "password")
}
```

and we call the `RunUserMigration()` from our `RegisterStandardMigrations()` function:

```go
func RegisterStandardMigrations(db *gorm.DB) {
	RunUserMigration(db)
}
```