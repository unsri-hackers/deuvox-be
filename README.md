# Deuvox

Deuvox is ...

## Getting Started

To start developing this project, you need to clone this repository. After that, you can now start developing this project by run the main using this command:

```
go run cmd/main.go
```

To build the project you can use this command:

```
go build cmd/main.go
```

To run the test you can use this command:

```
go test ./...
```

It's recommended to run this project using docker.
To run this project using docker:

```
docker-compose up
```

## Folder Structure
```
    - cmd                 # This is where the main.go located
    - internal            # This folder is used to store clean architecture folder
      - app               # This folder is to initialize application (dependacy injection, router, and middleware)
      - delivery          # Handler (checking data from client in here)
      - repository        # Get data
        - inner           # Data (database, cache, api, grpc, upstream)
      - usecase           # Bussiness logic
    - pkg                 # Utility Here
```

<!-- TODO: -->
