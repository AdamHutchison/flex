# Flux Framework

The Flux framework is a work in progress and has been designed as a personal starting point for web applications.

## Modules
Flux consists of many smaller modules that handle specific pieces of the frameworks functionality. See the readme of each for documentation:

* [Flux Config](https://github.com/AdamHutchison/flux-config) - Handles the loading and retrival of configuration / env values
* [Gorilla Mux](https://github.com/gorilla/mux) - Handles routing
* [Gorilla Handlers](https://github.com/gorilla/handlers) - Adds useful middleware handlers

## Routing

### Registration
Routes are registered in the `RegisterRoutes()` function contained within the `http/routes/routes.go` file. An instance of the [Gorilla Mux](https://github.com/gorilla/mux) is injected into this function and you are free to use it as you see fit.

### Handlers
Handlers are located in the `/http/handlers` directory. A base handler has been added for you convenience that implements a `Respond` method. This is a simple helper to allow you to send API responses easily. It is up to you if you use it.

### Middleware
Middleware should be place in the `http/middleware` directory and global middleware should be registered within the `RegisterGlobalMiddleware()` function in the `http/middleware/middleware.go` file. Middleware registered here will be applied to all routes.

