insert into branch(id, name) values
                                 (1, 'Mollnhof'),
                                 (2, 'Einsatzleiter'),
                                 (3, 'Hauptabschnitt'),
                                 (4, 'Einsatzabschnitt'),
                                 (5, 'Unterabschnitt');

insert into status(id, name) values
                                 (1, 'ANSTEHEND'),
                                 (2, 'STORNIERT'),
                                 (3, 'WEITERGELEITET BEI EINSATZABSCHNITT'),
                                 (4, 'WEITERGELEITET BEI HAUPTABSCHNITT'),
                                 (5, 'ABGELEHNT BEI EINSATZABSCHNITT'),
                                 (6, 'ABGELEHNT BEI HAUPTABSCHNITT'),
                                 (7, 'ABGELEHNT BEI EINSATZLEITER'),
                                 (8, 'AKZEPTIERT'),
                                 (9, 'AUF DEM WEG'),
                                 (10, 'GELIEFERT');

insert into "user"
(id, name, phone, password, email, token, is_activated, is_super_user, is_email_verified, branch_id)
values
    (1, 'Einsatzleiter', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'admin@admin.com', '', true, true, true, 2),
    (2, 'Mollnhof', '1212', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'mollnhof@user.com', '', true, false, true, 1),
    (3, 'Mitte Hauptabschnitt', '092019', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'mitte_hauptabschnitt@user.com', '', true, false, true, 3),
    (4, 'EA 1 - Altstadt', '122342', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea1_altstadt@user.com', '', true, false, true, 4),
    (5, 'EA 1.1 Altstadt - Ost', '132324', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea11_altstadt_ost@user.com', '', true, false, true, 5),
    (6, 'EA 1.2 Altstadt - Mitte', '2425', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea12_altstadt_mitte@user.com', '', true, false, true, 5),
    (7, 'EA 2 - Neumarkt', '12322', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea2_neumarkt@user.com', '', true, false, true, 3);

insert into "hierarchy" (user1_id, user2_id)
values(2, 1), (1, 3), (3, 4), (3, 7), (4, 5), (4, 6);


insert into permission(id, name) values
                                     (1, 'CAN VIEW'),
                                     (2, 'CAN CONFIRM DELIVERY'),
                                     (3, 'CAN EDIT'),
                                     (4, 'CAN DECLINE'),
                                     (5, 'CAN ACCEPT'),
                                     (6, 'CAN COMMENT'),
                                     (7, 'CAN ASSIGN TO');


insert into action_type(id, name) values
                                      (1, 'CREATED'),
                                      (2, 'EDITED'),
                                      (3, 'COMMENTED'),
                                      (4, 'ACCEPTED'),
                                      (5, 'DECLINED'),
                                      (6, 'ASSIGNED'),
                                      (7, 'CONFIRMED DELIVERY');

insert into priority(id, level, name) values
                                          (1, 1, 'HIGH'), (2, 2, 'MIDDLE'), (3, 3, 'LOW');


insert into equipment(id, name, quantity) values
    (1, 'Sandsack', 100000);


create sequence order_number_id_seq
    as integer;

alter sequence order_number_id_seq owner to postgres;
