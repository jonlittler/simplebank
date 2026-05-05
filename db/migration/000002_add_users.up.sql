CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00+00',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") 
 REFERENCES "users" ("username") 
 DEFERRABLE INITIALLY IMMEDIATE;

-- CREATE UNIQUE INDEX ON "accounts" ("owner", "currency");
ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" 
     UNIQUE (owner, currency);

/* set seed data */
insert into users (username, hashed_password, full_name, email)
values ('paj', 'password', 'Pj & Apple', 'pjapp@gmail.com');

insert into accounts (id, owner, balance, currency) 
values (1, 'paj', 50000000, 'USD');

select setval('accounts_id_seq', 1, true);
