bl-local:
	cd ./build && docker compose build --no-cache

up-local:
	cd ./build && docker compose up

inBackend:
	docker exec -it boxkeyper-backend bash

InDb: 
	docker exec -it boxkeyper_db bash

genBuf: 
	buf generate proto