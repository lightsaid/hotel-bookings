## docker/up: 在后台运行 docker compose 
docker/up:
	docker-compose up -d 

## docker/down: 停止 docker compose 
docker/down:
	docker-compose down

.PHONY: docker/up docker/down