atlas-create-manual-migration: ## Generate migration file
	atlas migrate new --dir file://database/ent/migrate/migrations $(name)

atlas-create-migration: ## Generate migration file
	go run -mod=mod database/ent/migrate/main.go $(name)

atlas-downgrade: ## Run atlas upgrade script
	atlas migrate down --dir "file://database/ent/migrate/migrations" --url 'sqlite://learn-grpc.db?_fk=1' --dev-url 'sqlite://temp?mode=memory&_fk=1'

atlas-upgrade: ## Run atlas upgrade script
	atlas migrate apply --dir "file://database/ent/migrate/migrations" --url 'sqlite://learn-grpc.db?_fk=1' 1

new-ent-schema: # Create new ent db schema
	ent new --target database/ent/schema $(name)

rehash-ent-migrations:
	atlas migrate hash --dir "file://database/ent/migrate/migrations"
