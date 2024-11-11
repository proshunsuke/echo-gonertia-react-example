compose/build:
	docker compose build --no-cache

compose/up:
	docker compose up

shell/app:
	docker compose exec app /bin/sh

shell/front:
	docker compose exec front /bin/sh

run/app:
	docker compose run app /bin/sh

run/front:
	docker compose run front /bin/sh

build/front:
	docker compose exec front npm run build

check: check/go check/front

check/go:
	docker compose exec app go fmt ./...

check/front:
	docker compose exec front npx tsc --noEmit
	docker compose exec front npx @biomejs/biome check --write --unsafe resources

test: test/front

test/front:
	docker compose exec front-test npm run test
