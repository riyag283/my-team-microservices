CREATE TABLE IF NOT EXISTS "team_members" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "role" varchar NOT NULL,
  "city" varchar
);
