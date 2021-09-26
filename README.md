# Flux Framework

The Flux framework is a work in progress and has been designed as a personal starting point for web applications.

## Modules
Flux consists of many smaller modules that handle specific pieces of the frameworks functionality. See the readme of each for documentation:

* [Flux Config](https://github.com/AdamHutchison/flux-config) - Handles the loading and retrival of configuration / env values
* [Gorilla Mux](https://github.com/gorilla/mux) - Handles routing
* [Gorilla Handlers](https://github.com/gorilla/handlers) - Adds useful middleware handlers
* [Gorilla Schema](https://github.com/gorilla/schema) - Used in request validation
* [Go Playground Validator](https://github.com/go-playground/validator) - Used in request validation

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

Then add a `GetValidator()` methtod to your handler that returns the validator:

```go
func (h HomeHandler) GetValidator() interface{} {
	validator := validators.HomeValidator{}

	return &validator
}
```

You may then use the `Validate(h Validateable, w http.ResponseWriter, r *http.Request) error` method on the `BaseHandler`:

```go
func (h HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.Validate(h, w, r)

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

