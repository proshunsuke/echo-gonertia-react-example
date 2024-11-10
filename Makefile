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

fmt: fmt/go fmt/front

fmt/go:
	docker compose exec app go fmt ./...

fmt/front:
	docker compose exec front npx @biomejs/biome check --write resources

test: test/front

test/front:
	docker compose exec front-test npm run test
