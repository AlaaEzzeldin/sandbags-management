insert into "user" (id, name, phone, password, email, token, is_activated, is_super_user, is_email_verified, branch_id)
values(1, 'Einsatzleiter', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa', 'admin@admin.com', '', true, true, true, 2);

insert into "user" (id, name, phone, password, email, token, is_activated, is_super_user, is_email_verified, branch_id)
values(2, 'Mollnhof', '1212', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa', 'mollnhof@admin.com', '', true, false, true, 1);


insert into "hierarchy" (user1_id, user2_id)
values(1, 2);