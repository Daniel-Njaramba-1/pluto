    Frontend - Svelte
customer shop - browse products, view prices, purchase
admin dashboard - product management, adjust base price, monitor model
web sockets - instant update for price changes
rest (echo) - fetch products, orders, customer data

    Backend - Golang
Retail Service (Handles business logic)
- Serves product, order, and customer data via Echo (REST API).
- Caches product prices in Redis for faster access.
- Broadcasts WebSocket updates to frontend on price changes.

Regression Service (Machine Learning-based pricing)
- Uses linear regression to adjust prices dynamically.
- Reads product sales, stock, and demand data from PostgreSQL.
- Temporarily stores calculated prices in Redis before committing to DB.

    Database - Postgres
sqlx + pg - for store functions
goose for migrations

    Cache - Redis
fast lookupd for adjusted prices
pub/sub mechanism to broadcast price updates
shared state between backend and frontend
can store temporary calculations from regression

    Data Flow
customer -> shop -> api -> service -> store -> db

    gRPC
uses http/2 with protocol buffers for service-service calls
many browsers dont support http/2 streaming
best for server environments












backend/
    cmd/
        api.go
        main.go
    internal/
        api/
            server.go
            admin/
            customer/
        cache/
            product_repository.go
            repository.go
            redis/
        config/
            config.env
            config.go
        db/
            db.go
            migrations/
            seeds/
        lib/
            generics/
            hashing/
            interfaces/
            logger/
        pkg/
            brand/
            cart/
            category/
            customer/
            order/
            payment/
            ...
        regression/
            ...
    go.mod
    go.sum
    Makefile
frontend/
    admin/
    customer/
    shared/
docker-compose.yml




