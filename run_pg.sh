docker run --name golang \
 -p 5433:5432 \
 -e POSTGRES_DATABASE=golang \
 -e POSTGRES_USER=golang \
 -e POSTGRES_PASSWORD=golang \
 -e POSTGRES_ROOT_PASSWORD=golang \
 -d postgres:latest
