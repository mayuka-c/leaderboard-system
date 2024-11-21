CREATE TYPE "gender_t" AS ENUM (
  'Male',
  'Female'
);

CREATE TABLE "players" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" int64 NOT NULL,
  "join_date" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "profiles" (
  "id" bigserial PRIMARY KEY, 
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "email" citext NOT NULL,
  "age" int NOT NULL,
  "gender" gender_t NOT NULL,
  "player_id" bigint NOT NULL UNIQUE,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()'

  CONSTRAINT proper_email CHECK(email ~ '^[\w]+\@\w{0,6}\.\w{2,4}$')
);

CREATE TABLE "games" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL
);

CREATE TABLE "leaderboards" (
  "game_id" bigint NOT NULL,
  "player_id" bigint NOT NULL,
  "score" bigint,
  "updated_at" timestamptz NOT NULL DEFAULT 'now()',
  PRIMARY KEY ("game_id", "player_id")
);

COMMENT ON TABLE "players" IS 'Stores Players Login info';

COMMENT ON TABLE "profiles" IS 'Stores Players Profile details';

COMMENT ON TABLE "games" IS 'Stores Games info';

COMMENT ON TABLE "leaderboards" IS 'Stores Leaderboard info for all games';

ALTER TABLE "profiles" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id") ON DELETE CASCADE;

ALTER TABLE "leaderboards" ADD FOREIGN KEY ("game_id") REFERENCES "games" ("id") ON DELETE CASCADE;

ALTER TABLE "leaderboards" ADD FOREIGN KEY ("player_id") REFERENCES "players" ("id") ON DELETE CASCADE;
