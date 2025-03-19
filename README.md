# stock-analysis

# Run the following commands to start the project
depends on docker and docker-compose

```
docker compose up -d
```

# Run the following commands to stop the project
```
docker compose down
```

# Run the following commands to start the project locally
depends on go and nodejs=22.14.0
```
cd ui
npm install
npm run build
cd ..
go mod download
go run main.go
```