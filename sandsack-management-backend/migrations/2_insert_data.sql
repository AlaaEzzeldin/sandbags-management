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

    (3, 'Mitte (LZ Hauptwache)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'mitte@user.com', '', true, false, true, 3),
    (4, 'Nord (FF Grubweg)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'nord@user.com', '', true, false, true, 3),
    (5, 'Süd (EA 7 Innstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'sued@user.com', '', true, false, true, 3),
    (6, 'Technik (EA 11)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea11_technik@user.com', '', true, false, true, 3),

    -- Einsatzabschnitt
    -- Mitte
    (7, 'EA 1 - Altstadt (LZ Ilzstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea1_altstadt@user.com', '', true, false, true, 4),
    (8, 'EA 2 - Neumarkt (LZ Hauptwache)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea2_neumarkt@user.com', '', true, false, true, 4),
    (9, 'EA 3 - Universität (FF Schalding r.d.D.)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea3_universitaet@user.com', '', true, false, true, 4),
    (10, 'EA 4 - Hacklberg/Anger (FF Hacklberg/FF Gaißa)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea4_hacklberg@user.com', '', true, false, true, 4),

    -- Nord
    (11, 'EA 5 - Hals (FF Hals)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea5_hals@user.com', '', true, false, true, 4),
    (12, 'EA 6 - Grubweg (FF Grubweg)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea6_grubweg@user.com', '', true, false, true, 4),

    -- EA 7
    (13, 'EA 7.1 Innstadt-West (LZ Innstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea71_innstadt_west_1@user.com', '', true, false, true, 5),
    (14, 'EA 7.2 Innstadt-Mitte (LZ Innstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea72_innstadt_mitte@user.com', '', true, false, true, 5),
    (15, 'EA 7.3 Innstadt-West II (LZ Innstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea73_innstadt_west_2@user.com', '', true, false, true, 5),
    (16, 'EA 7.4 Innstadt-Sandsack (LZ Innstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea74_innstadt_sandsack@user.com', '', true, false, true, 5),

    -- Unterabschnitt
    -- EA 1
    (17, 'EA 1.1 Altstadt-Ost (FF Heining)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea11_altstadt_ost@user.com', '', true, false, true, 5),
    (18, 'EA 1.2 Altstadt-Mitte (LZ Ilzstadt)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea12_altstadt_mitte@user.com', '', true, false, true, 5),
    (19, 'EA 1.3 Altstadt-West (FF Patriching)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea13_altstadt_west@user.com', '', true, false, true, 5),
    -- EA 2
    (20, 'EA 2.1 Neumarkt-Nord (LZ Hauptwache)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea21_neumarkt_nord@user.com', '', true, false, true, 5),
    (21, 'EA 2.2 Neumarkt-Süd (FF Schalding l.d.D.)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea22_neumarkt_sued@user.com', '', true, false, true, 5),
    -- EA 3
    (22, 'EA 3.1 Universität-Ost (FF Schalding r.d.D.)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea31_universitaet_ost@user.com', '', true, false, true, 5),
    (23, 'EA 3.2 Universität-West (FF Schalding r.d.D.)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea32_universitaet_west@user.com', '', true, false, true, 5),
  -- EA 4
    (24, 'EA 4.1 Hacklberg (FF Hacklberg)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea41_hacklberg@user.com', '', true, false, true, 5),
    (25, 'EA 4.2 Anger (FF Ries)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea42_anger@user.com', '', true, false, true, 5),
    -- EA 5
    (26, 'EA 5.1 Hals-West', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea51_hals_west@user.com', '', true, false, true, 5),
    (27, 'EA 5.2 Hals-Ost', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea52_hals_ost@user.com', '', true, false, true, 5),
    -- EA 6
    (28, 'EA 6.1 Grubweg (FF Grubweg)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea61_grubweg@user.com', '', true, false, true, 5),
    (29, 'EA 6.2 Ilzstadt-Ost (FF Grubweg)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea62_ilzstadt_ost@user.com', '', true, false, true, 5),
    (30, 'EA 6.3 Grubweg-Sandsack (FF Grubweg)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea63_grubweg_sandsack@user.com', '', true, false, true, 5),
    (31, 'EA ZF Werk I (WF ZF)', '1234', '$2a$14$oY1qUZ8/2agftsfoG6GTau2QeVSsqE3rTJVqxYKLuqnSjFqMzL1Oa',
     'ea6zf_werk@user.com', '', true, false, true, 5);

insert into "hierarchy" (user1_id, user2_id)
values(2,1),
      (1,3), (1,4), (1,5), (1,6),
      (3,7), (3,8), (3,9), (3,10),
      (4,11), (4,12),
      (5,13), (5,14), (5,15), (5,16),
      (7,17), (7,18), (7,19),
      (8,20), (8,21),
      (9,22), (9,23),
      (10,24), (10,25),
      (11,26), (11,27),
      (12,28), (12,29), (12,30), (12,31);



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
                                      (6, 'ASSIGNED');

insert into priority(id, level, name) values
                                          (1, 1, 'HIGH'), (2, 2, 'MIDDLE'), (3, 3, 'LOW');


insert into equipment(id, name, quantity, measure) values
    (1, 'Sandsack', 100000, 'Paletten'),
    (2, 'Spaten', 1000, 'Stück'),
    (3, 'Bindfaden', 2000, 'Meter'),
    (4, 'Wasserschlauch', 2000, 'Meter');



insert into driver(id, name, description) values
    (1, 'Denis Müller', ''),
    (2, 'Michael Sommer', '');