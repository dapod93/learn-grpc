package main

//go:generate buf format -w
//go:generate buf generate

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration --target ./database/ent ./database/ent/schema
