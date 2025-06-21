atlas-create-manual-migration: ## Generate migration file
	atlas migrate new --dir file://database/ent/migrate/migrations $(name)

atlas-create-migration: ## Generate migration file
	go run -mod=mod database/ent/migrate/main.go $(name)

new-ent-schema: # Create new ent db schema
	ent new --target database/ent/schema $(name)

rehash-ent-migrations:
	atlas migrate hash --dir "file://database/ent/migrate/migrations"
