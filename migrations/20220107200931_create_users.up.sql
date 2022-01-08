CREATE TABLE "users" (
	"id"	INTEGER NOT NULL UNIQUE,
	"email "	TEXT NOT NULL UNIQUE,
	"encrypted_password"	TEXT NOT NULL,
	PRIMARY KEY("id")
);