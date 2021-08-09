## REST API WITH GRPC
### How To Run
- run this command ``docker-compose up``
- if you change the ``docker-compose.yml`` change the config database on ``config/db.go``
- run the rest with ``go run api/main.go``
- run the rpc server ``go run rpc/server/server.go``

### Library Used
- Viber HTTP engine for Go
- Protobuf
- Google UUID Generator
- GRPC
- Postgres ORM

### TODO
- [ ] Create Unit Test
- [x] Create Rest API Method
- [x] Using Clean Architecture
- [x] Create Protobuf Contract
