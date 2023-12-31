init:
	docker compose build --no-cache
	docker compose up -d 
	sleep 3
	make db-migration
	make db-seeder
up:
	docker compose up -d
down:
	docker compose down --remove-orphans
db:
	docker compose exec db mysql -u sharely -D sharely -p
db-migration:
	docker compose exec server goose -dir ./database/migrations mysql "sharely:testtest@tcp(db:3306)/sharely" up
db-seeder:
	docker compose exec server goose -dir ./database/seeder mysql "sharely:testtest@tcp(db:3306)/sharely" up
act:
	act --container-architecture linux/amd64
create-model:
	cd server && sqlboiler mysql -c database/database.toml -o internal/domain/models -p models --no-tests --wipe
