CREATE TABLE "user" (
                        "id" serial PRIMARY KEY,
                        "name" varchar,
                        "phone" varchar,
                        "password" varchar,
                        "email" varchar unique,
                        "token" varchar,
                        "is_activated" boolean,
                        "is_super_user" boolean,
                        "is_email_verified" boolean,
                        "branch_id" integer,
                        "create_date" timestamp,
                        "update_date" timestamp
);

CREATE TABLE "order" (
                         "id" serial PRIMARY KEY,
                         "user_id" integer,
                         "name" varchar,
                         "address_to" varchar,
                         "address_from" varchar,
                         "status_id" integer,
                         "priority_id" integer,
                         "assigned_to" integer,
                         "create_date" timestamp,
                         "update_date" timestamp
);

CREATE TABLE "order_equipment" (
                                   "id" serial PRIMARY KEY,
                                   "order_id" integer,
                                   "equipment_id" integer,
                                   "quantity" integer,
                                   "create_date" timestamp,
                                   "update_date" timestamp
);

CREATE TABLE "equipment" (
                             "id" serial PRIMARY KEY,
                             "name" varchar,
                             "quantity" integer,
                             "create_date" timestamp,
                             "update_date" timestamp
);

CREATE TABLE "status" (
                          "id" serial PRIMARY KEY,
                          "name" varchar,
                          "create_date" timestamp,
                          "update_date" timestamp
);

CREATE TABLE "action_type" (
                               "id" serial PRIMARY KEY,
                               "name" varchar
);

CREATE TABLE "log" (
                       "id" serial PRIMARY KEY,
                       "order_id" integer,
                       "description" varchar,
                       "action_type_id" integer,
                       "ip_address" varchar,
                       "updated_by" integer,
                       "create_date" timestamp,
                       "update_date" timestamp
);

CREATE TABLE "comment" (
                           "id" serial PRIMARY KEY,
                           "user_id" integer,
                           "order_id" integer,
                           "comment_text" varchar,
                           "create_date" timestamp,
                           "update_date" timestamp
);

CREATE TABLE "hierarchy" (
                             "id" serial PRIMARY KEY,
                             "user1_id" integer,
                             "user2_id" integer,
                             "create_date" timestamp,
                             "update_date" timestamp
);

CREATE TABLE "user_order_permission" (
                                         "id" serial PRIMARY KEY,
                                         "user_id" integer,
                                         "order_id" integer,
                                         "permission_id" integer,
                                         "create_date" timestamp,
                                         "update_date" timestamp
);

CREATE TABLE "permission" (
                              "id" serial PRIMARY KEY,
                              "name" varchar,
                              "create_date" timestamp,
                              "update_date" timestamp
);

CREATE TABLE "priority" (
                            "id" serial PRIMARY KEY,
                            "level" integer,
                            "name" varchar,
                            "create_date" timestamp,
                            "update_date" timestamp
);

CREATE TABLE "driver" (
                          "id" serial PRIMARY KEY,
                          "name" varchar,
                          "description" varchar,
                          "create_date" timestamp,
                          "update_date" timestamp
);

CREATE TABLE "branch" (
                          "id" serial PRIMARY KEY,
                          "name" varchar,
                          "create_date" timestamp,
                          "update_date" timestamp
);

CREATE TABLE "otp" (
                       "id" serial PRIMARY KEY,
                       "code" varchar,
                       "user_id" integer,
                       "type" varchar,
                       "create_date" timestamp,
                       "update_date" timestamp
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

CREATE SEQUENCE IF NOT EXISTS order_number_id_seq
    AS INTEGER;


ALTER SEQUENCE order_number_id_seq owner TO postgres;
