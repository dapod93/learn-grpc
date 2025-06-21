atlas-create-manual-migration: ## Generate migration file
	atlas migrate new --dir file://database/ent/migrate/migrations $(name)

atlas-create-migration: ## Generate migration file
	go run -mod=mod database/ent/migrate/main.go $(name)

atlas-downgrade: ## Run atlas upgrade script
	atlas migrate down --dir "file://database/ent/migrate/migrations" --url 'sqlite://learn-grpc.db?_fk=1' --dev-url 'sqlite://temp?mode=memory&_fk=1'

atlas-upgrade: ## Run atlas upgrade script
	atlas migrate apply --dir "file://database/ent/migrate/migrations" --url 'sqlite://learn-grpc.db?_fk=1' 1

install-cli-helpers: ## Install clis
	curl -sSf https://atlasgo.sh | sh && \
	go install github.com/bufbuild/buf/cmd/buf@v1.55.1 && \
	go install github.com/air-verse/air@latest && \
	go install entgo.io/ent/cmd/ent@latest

new-ent-schema: ## Create new ent db schema
	ent new --target database/ent/schema $(name)

rehash-ent-migrations: ## Rehase atlas.sum
	atlas migrate hash --dir "file://database/ent/migrate/migrations"

serve: ## Serve the server
	air --build.cmd "go build -o ./bin/app main.go" --build.bin "./bin/app"
