CREATE TABLE "user" (
                        "id" serial PRIMARY KEY,
                        "name" varchar not null,
                        "phone" varchar not null,
                        "password" varchar not null,
                        "email" varchar unique not null,
                        "token" varchar,
                        "is_activated" boolean default false,
                        "is_super_user" boolean default false,
                        "is_email_verified" boolean default false,
                        "branch_id" integer not null,
                        "create_date" timestamptz not null default now(),
                        "update_date" timestamptz not null default now()
);

CREATE TABLE "order" (
                         "id" serial PRIMARY KEY,
                         "user_id" integer not null,
                         "name" varchar not null,
                         "address_to" varchar not null,
                         "address_from" varchar not null,
                         "status_id" integer not null,
                         "priority_id" integer not null,
                         "assigned_to" integer,
                         "create_date" timestamptz not null default now(),
                         "update_date" timestamptz not null default now()
);

CREATE TABLE "order_equipment" (
                                   "id" serial PRIMARY KEY,
                                   "order_id" integer not null,
                                   "equipment_id" integer not null,
                                   "quantity" integer not null,
                                   "create_date" timestamptz not null default now(),
                                   "update_date" timestamptz not null default now()
);

CREATE TABLE "equipment" (
                             "id" serial PRIMARY KEY,
                             "name" varchar not null,
                             "quantity" integer not null,
                             "measure" varchar null,
                             "create_date" timestamptz not null default now(),
                             "update_date" timestamptz not null default now()
);

CREATE TABLE "status" (
                          "id" serial PRIMARY KEY,
                          "name" varchar not null,
                          "create_date" timestamptz not null default now(),
                          "update_date" timestamptz not null default now()
);

CREATE TABLE "action_type" (
                               "id" serial PRIMARY KEY,
                               "name" varchar not null
);

CREATE TABLE "log" (
                       "id" serial PRIMARY KEY,
                       "order_id" integer not null,
                       "description" varchar not null,
                       "action_type_id" integer not null,
                       "ip_address" varchar,
                       "updated_by" integer not null,
                       "create_date" timestamptz not null default now(),
                       "update_date" timestamptz not null default now()
);

CREATE TABLE "comment" (
                           "id" serial PRIMARY KEY,
                           "user_id" integer not null,
                           "order_id" integer not null,
                           "comment_text" varchar not null,
                           "create_date" timestamptz not null default now(),
                           "update_date" timestamptz not null default now()
);

CREATE TABLE "hierarchy" (
                             "id" serial PRIMARY KEY,
                             "user1_id" integer not null,
                             "user2_id" integer not null,
                             "create_date" timestamptz not null default now(),
                             "update_date" timestamptz not null default now()
);

CREATE TABLE "user_order_permission" (
                             "id" serial PRIMARY KEY,
                             "user_id" integer not null,
                             "order_id" integer not null,
                             "permission_id" integer not null,
                             "create_date" timestamptz not null default now(),
                             "update_date" timestamptz not null default now()
);

CREATE TABLE "permission" (
                              "id" serial PRIMARY KEY,
                              "name" varchar not null,
                              "create_date" timestamptz not null default now(),
                              "update_date" timestamptz not null default now()
);

CREATE TABLE "priority" (
                            "id" serial PRIMARY KEY,
                            "level" integer not null,
                            "name" varchar not null,
                            "create_date" timestamptz not null default now(),
                            "update_date" timestamptz not null default now()
);

CREATE TABLE "driver" (
                          "id" serial PRIMARY KEY,
                          "name" varchar not null,
                          "description" varchar not null,
                          "create_date" timestamptz not null default now(),
                          "update_date" timestamptz not null default now()
);

CREATE TABLE "branch" (
                          "id" serial PRIMARY KEY,
                          "name" varchar not null,
                          "create_date" timestamptz not null default now(),
                          "update_date" timestamptz not null default now()
);

CREATE TABLE "otp" (
                       "id" serial PRIMARY KEY,
                       "code" varchar not null,
                       "user_id" integer not null,
                       "type" varchar not null,
                       "create_date" timestamptz not null default now(),
                       "update_date" timestamptz not null default now()
);

ALTER TABLE "order" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("status_id") REFERENCES "status" ("id");

ALTER TABLE "order_equipment" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "order_equipment" ADD FOREIGN KEY ("equipment_id") REFERENCES "equipment" ("id");

ALTER TABLE "log" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "log" ADD FOREIGN KEY ("updated_by") REFERENCES "user" ("id");

ALTER TABLE "hierarchy" ADD FOREIGN KEY ("user1_id") REFERENCES "user" ("id");

ALTER TABLE "hierarchy" ADD FOREIGN KEY ("user2_id") REFERENCES "user" ("id");

ALTER TABLE "log" ADD FOREIGN KEY ("action_type_id") REFERENCES "action_type" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_order_permission" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_order_permission" ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "user_order_permission" ADD FOREIGN KEY ("permission_id") REFERENCES "permission" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("priority_id") REFERENCES "priority" ("id");

ALTER TABLE "order" ADD FOREIGN KEY ("assigned_to") REFERENCES "driver" ("id");

ALTER TABLE "user" ADD FOREIGN KEY ("branch_id") REFERENCES "branch" ("id");

ALTER TABLE "otp" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
