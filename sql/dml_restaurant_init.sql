-- =============================================
-- Author:      William Wibowo Ciptono
-- Create date: 14 Des 2022
-- Description: 
-- =============================================

insert into roles (id, role_name) values (1, "user");
insert into roles (id, role_name) values (2, "admin");

insert into users (id, full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values (1, 'Reinwald Cammack', '7211694495', 'rcammack0@discovery.com', 'rcammack0', '$2a$04$.bXHRTCpDQWR/N.BFmg.U.dd7GV6o..m0.DDISz2wM.ty9d.RVmsW', '2022-02-23', null, 1, 1);
insert into users (id, full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values (2, 'Terrel Pettiford', '9803639991', 'tpettiford1@accuweather.com', 'tpettiford1', '$2a$04$.uU.iSiCNslD0NAYjV7Duue57bdEEHliQqRucqa1.fisdGG/c.DLG', '2022-08-28', null, 5, 1);
insert into users (id, full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values (3, 'Ailee Hambly', '4906434437', 'ahambly2@patch.com', 'ahambly2', '$2a$04$IcSDw29tFuDjjGjOtwH6P.RyOTcDzLAPD5ZyI4pMvLGT3efAteuju', '2022-03-10', null, 5, 1);
insert into users (id, full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values (4, 'Jo-anne Kernar', '7251531487', 'jkernar3@paypal.com', 'jkernar3', '$2a$04$GX1pJVxU55hAE6cPZoq2h.6PDJuSKz.PdSZOgVVDb10Ersbd5iAMC', '2022-03-14', null, 1, 1);
insert into users (id, full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values (5, 'Donetta Gardener', '4814877652', 'dgardener4@youtu.be', 'dgardener4', '$2a$04$ULAmhc1a.xte8qPwcJk4a.Jdbn8GyRALa0LCC0C0cSdpE47bJFzYK', '2022-10-12', null, 4, 1);
insert into users (id, full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values (6, 'Admin Admin', '8532753943', 'admin@mail.com', 'admin', '$2a$04$MegFiHl0PLKEXaGXpjMFf.BgBDodVLflV6nsh1/N1z2nupXACEIdO', '2022-10-12', null, 4, 2);