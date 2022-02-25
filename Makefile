postgres:
	docker run --name postgres12alpine -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123 -d postgres:12-alpine
	
start:
	docker start postgres12alpine

createdb:
	docker exec -it postgres12alpine createdb --username=root --owner=root quanlynhanvien

dropdb:
	docker exec -it postgres12alpine dropdb quanlynhanvien