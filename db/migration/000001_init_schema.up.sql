CREATE TABLE "customers" (
	  "id" bigserial PRIMARY KEY,
	  "full_name" varchar NOT NULL,
	  "phone_number" varchar NOT NULL,
	  "created_at" timestamptz NOT NULL DEFAULT (now())
	);


	CREATE INDEX ON "customers" ("full_name");

