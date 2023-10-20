# SQL 迁移文件
MIGRATE_SQL=./db/migrations

# 通过执行golang-migrate/migrate代码来迁移MySQL，而不是安装migrate可执行程序
MIGRATE_CMD=github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# MySQL链接
DB_SOURCE=mysql://root:abc123@tcp(localhost:3366)/db_hotel_bookings


# ==================================================================================== #
# Hotel Booking Server
# ==================================================================================== #

## go: 下载依赖并同步到vendor
go:
	go mod tidy
	go mod vendor

## back/start: 启动后台api服务
back/start:
	go run cmd/backend/main.go

## front/start: 启动前台api服务
front/start:
	go run cmd/frontend/main.go
	
## rm/uploads: 删除上传文件
rm/uploads:
	rm -rf ./static/uploads/*

.PHONY: go back/start back/start rm/uploads

# ==================================================================================== #
# docker-compose
# ==================================================================================== #

## docker/up: 在后台运行 docker compose 
docker/up:
	docker-compose up -d 

## docker/down: 停止 docker compose 
docker/down:
	docker-compose down

.PHONY: docker/up docker/down


# ==================================================================================== #
# sqlc 生成CRUD代码
# ==================================================================================== #

## sqlc/gen: 生成CRUD代码
sqlc:
	docker run --rm -v $(PWD):/src -w /src sqlc/sqlc:1.22.0 generate

# ==================================================================================== #
# SQL 迁移
# ==================================================================================== #

## migrate/new name=$1: 创建迁移SQL文件，如 make migrate/new name=create_table_hotels
migrate/new:
	go run -tags 'mysql' ${MIGRATE_CMD} create -seq -ext=.sql -dir=${MIGRATE_SQL} ${name}

## migrate/up: 向上迁移所有
migrate/up:
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose up

## migrate/down: 向下迁移所有
migrate/down:
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose down

## migrate/up1: 向上迁移一次
migrate/up1: 
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose up 1

## migrate/down1: 向下迁移一次
migrate/down1:
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose down 1

## migrate/force version=$1: 强制迁移到指定版本
migrate/force:
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose force ${version}

## migrate/goto version=$1: 迁移到指定版本
migrate/goto:
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" -verbose goto ${version}

## migrate/fix version=$1: 向上迁移过程中如果出错，执行此命令快速修复，先强制升迁成功，再向下迁移
# migrate/fix:
# 	make migrate/force version=${version}
# 	make migrate/down1

# migrate/version: 查看当前的迁移版本
migrate/version:
	go run -tags 'mysql' ${MIGRATE_CMD} -path=${MIGRATE_SQL} -database="${DB_SOURCE}" version

.PHONY: migrate/new migrate/up migrate/up1 migrate/down migrate/down1 migrate/force migrate/goto migrate/version