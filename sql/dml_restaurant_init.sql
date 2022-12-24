-- =============================================
-- Author:      William Wibowo Ciptono
-- Create date: 14 Des 2022
-- Description: 
-- =============================================


insert into roles (id, role_name) values (1, 'user');
insert into roles (id, role_name) values (2, 'admin');

insert into users (full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values ('Reinwald Cammack', '7211694495', 'rcammack0@discovery.com', 'rcammack0', '$2a$04$.bXHRTCpDQWR/N.BFmg.U.dd7GV6o..m0.DDISz2wM.ty9d.RVmsW', '2022-02-23', null, 1, 1);
insert into users (full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values ('Terrel Pettiford', '9803639991', 'tpettiford1@accuweather.com', 'tpettiford1', '$2a$04$.uU.iSiCNslD0NAYjV7Duue57bdEEHliQqRucqa1.fisdGG/c.DLG', '2022-08-28', null, 5, 1);
insert into users (full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values ('Ailee Hambly', '4906434437', 'ahambly2@patch.com', 'ahambly2', '$2a$04$IcSDw29tFuDjjGjOtwH6P.RyOTcDzLAPD5ZyI4pMvLGT3efAteuju', '2022-03-10', null, 5, 1);
insert into users (full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values ('Jo-anne Kernar', '7251531487', 'jkernar3@paypal.com', 'jkernar3', '$2a$04$GX1pJVxU55hAE6cPZoq2h.6PDJuSKz.PdSZOgVVDb10Ersbd5iAMC', '2022-03-14', null, 1, 1);
insert into users (full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values ('Donetta Gardener', '4814877652', 'dgardener4@youtu.be', 'dgardener4', '$2a$04$ULAmhc1a.xte8qPwcJk4a.Jdbn8GyRALa0LCC0C0cSdpE47bJFzYK', '2022-10-12', null, 4, 1);
insert into users (full_name, phone, email, username, password, register_date, profile_picture, play_attempt, role_id) values ('Admin Admin', '8532753943', 'admin@mail.com', 'admin', '$2a$04$MegFiHl0PLKEXaGXpjMFf.BgBDodVLflV6nsh1/N1z2nupXACEIdO', '2022-10-12', null, 4, 2);

insert into categories (id, category_name) values (1, 'appetizers'), (2, 'meals'), (3, 'drinks');

insert into menus (menu_name, avg_rating, number_of_favorites, price, menu_photo, category_id)
values 
('Nori', 0, 0, 10000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671772130/menu/appetizer-nori_rdfbss.png', 1),
('Edamame', 0, 0, 11000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/appetizer-edamame_d0slsa.png', 1),
('Spring Roll', 0, 0, 12000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/appetizer-spring_roll_m1wmfy.png', 1),
('Salmon Maki', 0, 0, 13000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/meals-salmon_maki_dqhlg9.png', 2),
('California Roll', 0, 0, 14000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/meals-california_roll_y0r8ca.png', 2),
('Salmon Sashimi', 0, 0, 15000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771969/menu/meals-salmon_sashimi_er9fdd.png', 2),
('Salmon Nigiri', 0, 0, 13000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771969/menu/meals-salmon_nigiri_nml4nk.png', 2),
('Sushi Platter', 0, 0, 40000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771969/menu/meals-sushi_plater_vvwyhe.png', 2),
('Lemon Tea', 0, 0, 12000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/drinks-lemon_tea_owmrxz.png', 3),
('Mint Tea', 0, 0, 12000, 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771967/menu/drinks-mint_tea_mgwedf.png', 3);
insert into favorited_menus (user_id, menu_id) values (1,1);
insert into menu_customs(menu_id, customization) values (1, 'size:big/small'), (1,'extra sauce:yes/no'), (2,'size:big/small');

insert into promotions(admin_id, name, description, promotion_photo, discount_rate, started_at, expired_at)
values 
	(1, 'California for Xmas', 'Discount 50% for California roll', 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/meals-california_roll_y0r8ca.png', 0.5, current_timestamp, current_date + INTERVAL '1 day'),
	(1, 'Drinks on the house', 'Discount 30% for all drinks', 'https://res.cloudinary.com/dgr6o89ym/image/upload/v1671771968/menu/drinks-lemon_tea_owmrxz.png', 0.3, current_timestamp, current_date + INTERVAL '2 day');
insert into promo_menus(promotion_id, menu_id)
values
	(1, 5),
	(2, 8),
	(2, 9);

insert into carts(user_id, promotion_id, menu_id, quantity, menu_option) 
values (1, 1, null, 1, '{"size":"big"}');