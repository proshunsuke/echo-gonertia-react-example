compose/build:
	docker compose build --no-cache

compose/up:
	docker compose up

ssh/app:
	docker compose exec app /bin/sh

ssh/front:
	docker compose exec front /bin/sh

run/app:
	docker compose run app /bin/sh

run/front:
	docker compose run front /bin/sh

build/front:
	docker compose exec front npm run build

fmt:
	docker compose exec app go fmt ./...
