-- Create "users" table
CREATE TABLE `users` (
    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    `email` text NOT NULL,
    `first_name` text NOT NULL,
    `last_name` text NOT NULL,
    `created_at` datetime NOT NULL
);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create index "user_id_email" to table: "users"
CREATE INDEX `user_id_email` ON `users` (`id`, `email`);