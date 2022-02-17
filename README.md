# api-go-template

This is a simple Go API template that uses a controller-service based model to build its routes.

## Directory tree

```
├── output/
│   ├── text_file_1.txt
│   ├── text_file_2.txt
├── routes/
│   ├── end/
│   │   ├── controller.go
│   │   ├── route.go
│   │   └── service.go
│   └── root/
│       ├── controller.go
│       ├── route.go
│       └── service.go
├── go.mod
├── go.sum
├── main.go
├── README.md
├── routes.go
```

### Output

The output directory stores all files that are created when you sent a request.

### Routes/

The routes directory stores all routes definitions and its components.

#### Controller.go

This file defines the HTTP interaction for the route. It translates and treats the Fiber context to useful information for the service. That means it provides the HTTP interface for the service.

#### Route.go

This file defines the Method, the path for the route and the Handler creation function. It exports everything needed to the main.go file to recognize this route.

#### Service.go

The guts and the raw functions that execute the operations needed for the system to work.

> :warning: Should not have any coupling to HTTP. These functions should only have the strict necessary to work.

### main.go

API initialization and API execution.

### routes.go

API route definition throughout the imports of the handlers created on the `routes/` directory.
