CREATE TYPE "role_enum" AS ENUM (
  'admin',
  'user'
);

CREATE TABLE IF NOT EXISTS "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "role" role_enum NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);
