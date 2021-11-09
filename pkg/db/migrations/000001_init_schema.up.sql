CREATE TABLE IF NOT EXISTS "user" (
  "id" VARCHAR(36) PRIMARY KEY,
  "email" VARCHAR(128) NOT NULL,
  "password" VARCHAR(256) NOT NULL,
  "verified" BOOLEAN NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "password" (
  "id" VARCHAR(36) PRIMARY KEY,
  "user_id" VARCHAR(36),
  "password" VARCHAR(256) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "profile" (
  "id" VARCHAR(36) PRIMARY KEY,
  "user_id" VARCHAR(36),
  "fullname" VARCHAR(128) NOT NULL,
  "picture" VARCHAR(128) NOT NULL DEFAULT '',
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE IF NOT EXISTS "session" (
  "jti" VARCHAR(255) PRIMARY KEY,
  "user_id" VARCHAR(36),
  "client" VARCHAR(128) NOT NULL,
  "ip" VARCHAR(128) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "password" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "profile" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "session" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

CREATE INDEX ON "user" ("email");

CREATE INDEX ON "password" ("user_id");

CREATE INDEX ON "profile" ("user_id");

CREATE INDEX ON "session" ("user_id");

COMMENT ON COLUMN "user"."password" IS 'bcrypt';

COMMENT ON COLUMN "user"."verified" IS 'true verify, false not verify';

COMMENT ON COLUMN "user"."created_at" IS 'full RFC3339 format';

COMMENT ON COLUMN "user"."updated_at" IS 'full RFC3339 format';

COMMENT ON COLUMN "user"."deleted_at" IS 'full RFC3339 format';

COMMENT ON COLUMN "password"."password" IS 'bcrypt';

COMMENT ON COLUMN "password"."created_at" IS 'full RFC3339 format';

COMMENT ON COLUMN "profile"."created_at" IS 'full RFC3339 format';

COMMENT ON COLUMN "profile"."updated_at" IS 'full RFC3339 format';

COMMENT ON COLUMN "session"."created_at" IS 'full RFC3339 format';