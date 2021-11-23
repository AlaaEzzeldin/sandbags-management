CREATE TABLE public."user" (
                        "id" serial PRIMARY KEY,
                        "name" varchar not null,
                        "phone" varchar,
                        "password" varchar not null,
                        "email" varchar not null,
                        "token" varchar,
                        "is_activated" boolean default false,
                        "is_super_user" boolean default false,
                        "is_email_verified" boolean default false,
                        "create_date" timestamptz not null default now(),
                        "update_date" timestamptz
);

CREATE TABLE public."order" (
                         "id" serial PRIMARY KEY,
                         "user_id" integer not null,
                         "name" varchar not null,
                         "address_to" varchar not null,
                         "address_from" varchar,
                         "status_id" integer not null,
                         "priority_id" integer not null,
                         "assigned_to" integer,
                         "create_date" timestamptz not null default now(),
                         "update_date" timestamptz
);

CREATE TABLE public."order_equipment" (
                                   "id" serial PRIMARY KEY,
                                   "order_id" integer not null,
                                   "equipment_id" integer not null,
                                   "quantity" integer not null default 0,
                                   "create_date" timestamptz not null default now(),
                                   "update_date" timestamptz
);

CREATE TABLE public."equipment" (
                             "id" serial PRIMARY KEY,
                             "name" varchar not null,
                             "quantity" integer not null default 0,
                             "create_date" timestamptz not null default now(),
                             "update_date" timestamptz
);

CREATE TABLE public."status" (
                          "id" serial PRIMARY KEY,
                          "name" varchar not null,
                          "create_date" timestamptz not null default now(),
                          "update_date" timestamptz
);

CREATE TABLE public."action_type" (
                               "id" serial PRIMARY KEY,
                               "name" varchar not null,
                               "create_date" timestamptz not null default now(),
                               "update_date" timestamptz
);

CREATE TABLE public."log" (
                       "id" serial PRIMARY KEY,
                       "order_id" integer not null,
                       "description" varchar,
                       "action_type_id" integer not null,
                       "ip_address" varchar,
                       "updated_by" integer,
                       "create_date" timestamptz not null default now(),
                       "update_date" timestamptz
);

CREATE TABLE public."comment" (
                           "id" serial PRIMARY KEY,
                           "user_id" integer not null,
                           "order_id" integer not null,
                           "comment_text" varchar not null,
                           "create_date" timestamptz not null default now(),
                           "update_date" timestamptz
);

CREATE TABLE public."hierarchy" (
                             "id" serial PRIMARY KEY,
                             "user1_id" integer not null,
                             "user2_id" integer not null,
                             "create_date" timestamptz not null default now(),
                             "update_date" timestamptz
);

CREATE TABLE public."user_order_permission" (
                                         "id" serial PRIMARY KEY,
                                         "user_id" integer not null,
                                         "order_id" integer not null,
                                         "permission_id" integer not null,
                                         "create_date" timestamptz not null default now(),
                                         "update_date" timestamptz
);

CREATE TABLE public."permission" (
                              "id" serial PRIMARY KEY,
                              "name" varchar not null,
                              "create_date" timestamptz not null default now(),
                              "update_date" timestamptz
);

CREATE TABLE public."priority" (
                            "id" serial PRIMARY KEY,
                            "level" integer not null,
                            "name" varchar not null,
                            "create_date" timestamptz not null default now(),
                            "update_date" timestamptz
);

CREATE TABLE public."driver" (
                          "id" serial PRIMARY KEY,
                          "name" varchar not null,
                          "description" varchar not null,
                          "create_date" timestamptz not null default now(),
                          "update_date" timestamptz
);

ALTER TABLE public."order" ADD FOREIGN KEY ("user_id") REFERENCES public."user" ("id");

ALTER TABLE public."order" ADD FOREIGN KEY ("status_id") REFERENCES public."status" ("id");

ALTER TABLE public."order_equipment" ADD FOREIGN KEY ("order_id") REFERENCES public."order" ("id");

ALTER TABLE public."order_equipment" ADD FOREIGN KEY ("equipment_id") REFERENCES public."equipment" ("id");

ALTER TABLE public."log" ADD FOREIGN KEY ("order_id") REFERENCES public."order" ("id");

ALTER TABLE public."log" ADD FOREIGN KEY ("updated_by") REFERENCES public."user" ("id");

ALTER TABLE public."hierarchy" ADD FOREIGN KEY ("user1_id") REFERENCES public."user" ("id");

ALTER TABLE public."hierarchy" ADD FOREIGN KEY ("user2_id") REFERENCES public."user" ("id");

ALTER TABLE public."log" ADD FOREIGN KEY ("action_type_id") REFERENCES public."action_type" ("id");

ALTER TABLE public."comment" ADD FOREIGN KEY ("order_id") REFERENCES public."order" ("id");

ALTER TABLE public."comment" ADD FOREIGN KEY ("user_id") REFERENCES public."user" ("id");

ALTER TABLE public."user_order_permission" ADD FOREIGN KEY ("user_id") REFERENCES public."user" ("id");

ALTER TABLE public."user_order_permission" ADD FOREIGN KEY ("order_id") REFERENCES public."order" ("id");

ALTER TABLE public."user_order_permission" ADD FOREIGN KEY ("permission_id") REFERENCES public."permission" ("id");

ALTER TABLE public."order" ADD FOREIGN KEY ("priority_id") REFERENCES public."priority" ("id");

ALTER TABLE public."order" ADD FOREIGN KEY ("assigned_to") REFERENCES public."driver" ("id");