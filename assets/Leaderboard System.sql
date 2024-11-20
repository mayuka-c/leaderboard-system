CREATE TYPE "gender_t" AS ENUM (
  'Male',
  'Female'
);

CREATE TABLE "players" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" bigint NOT NULL,
  "join_date" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "profiles" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" citext UNIQUE NOT NULL,
  "age" int NOT NULL,
  "gender" gender_t NOT NULL,
  "player_id" bigint,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "games" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "leaderboards" (
  "game_id" bigint,
  "player_id" bigint,
  "score" bigint,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  PRIMARY KEY ("game_id", "player_id")
);

COMMENT ON TABLE "players" IS 'Stores Players Login info';

COMMENT ON TABLE "profiles" IS 'Stores Players Profile details';

COMMENT ON TABLE "games" IS 'Stores Games info';

COMMENT ON TABLE "leaderboards" IS 'Stores Leaderboard info for all games';

ALTER TABLE "profiles" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id");

ALTER TABLE "leaderboards" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id");

ALTER TABLE "leaderboards" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id");
