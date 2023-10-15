CREATE TABLE IF NOT EXISTS atms (
    id serial primary key,
    address varchar(255) not null,
    latitude float not null,
    longitude float not null,
    all_day  boolean
);
CREATE TABLE IF NOT EXISTS atm_services (
    id serial not null primary key,
    atm_id integer references atms(id),
    wheelchair_capability varchar(255),
    wheelchair_activity varchar(255),
    blind_capability varchar(255),
    blind_activity varchar(255),
    nfc_capability varchar(255),
    nfc_activity varchar(255),
    qr_capability varchar(255),
    qr_activity varchar(255),
    usd_support_capability varchar(255),
    usd_support_activity varchar(255),
    charge_rub_capability varchar(255),
    charge_rub_activity varchar(255),
    eur_support_capability varchar(255),
    eur_support_activity varchar(255),
    rub_support_capability varchar(255),
    rub_support_activity varchar(255)
);

CREATE TABLE IF NOT EXISTS office_schedule(
    id bigserial not null primary key,
    day_id smallint references days(id),
    open_hours varchar(255) not null,
    individual_open_hours varchar(255) not null
);

CREATE TABLE IF NOT EXISTS offices (
    id bigserial not null primary key,
    sale_point varchar(255),
    address varchar(255),
    schedule_id bigint references office_schedule(id),
    rko varchar(255) not null,
    type varchar(255) not null,
    suo_availability varchar(255),
    sale_point_format varchar(255),
    has_ramp varchar(255),
    latitude decimal not null,
    longitude decimal not null,
    metro_station varchar(255),
    kep boolean
);



CREATE TABLE If NOT EXISTS days (
    id smallserial not null primary key,
    title_en varchar(255) not null,
    title_ru varchar(255) not null
);

INSERT INTO days VALUES (1,'monday', 'пн'),
                        (2,'tuesday', 'вт'),
                        (3,'wednesday', 'ср'),
                        (4,'thursday', 'чт'),
                        (5,'friday', 'пт'),
                        (6,'saturday', 'сб'),
                        (7,'sunday', 'вс');
INSERT INTO atms VALUES (1, 'ул. Богородский Вал, д. 6, корп. 1',55.802432, 37.704547, false),
                        (2, 'ул. Нижняя Красносельская, д. 43',55.773763, 37.675002, false),
                        (3, 'ул. Авиаконструктора Миля, д. 4, корп. 1',55.686137, 37.849832, false),
                        (4, 'Кадашевская наб., д. 36, стр. 1',55.745726, 37.625702, false),
                        (5, 'Новочеркасский б-р, д. 44',55.646284, 37.737542, false);
INSERT INTO atm_services VALUES
                             (1, 1, 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'SUPPORTED', 'AVAILABLE',
                              'UNSUPPORTED', 'UNAVAILABLE',  'SUPPORTED', 'AVAILABLE'),
                             (2, 2, 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'SUPPORTED', 'AVAILABLE',
                              'UNSUPPORTED', 'UNAVAILABLE',  'SUPPORTED', 'AVAILABLE'),
                             (3, 3, 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'SUPPORTED', 'AVAILABLE',
                              'UNSUPPORTED', 'UNAVAILABLE',  'SUPPORTED', 'AVAILABLE'),
                             (4, 4, 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'SUPPORTED', 'AVAILABLE',
                              'UNSUPPORTED', 'UNAVAILABLE',  'SUPPORTED', 'AVAILABLE'),
                             (5, 5, 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNKNOWN', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'UNSUPPORTED', 'UNAVAILABLE', 'SUPPORTED', 'AVAILABLE',
                              'UNSUPPORTED', 'UNAVAILABLE',  'SUPPORTED', 'AVAILABLE');
INSERT INTO office_schedule VALUES (1, 1, '09:00-18:00', '09:00-20:00'),
                                   (2, 2, '09:00-18:00', '09:00-20:00'),
                                   (3, 3, '09:00-18:00', '09:00-20:00'),
                                   (4, 4, '09:00-18:00', '09:00-20:00'),
                                   (5, 5, '09:00-17:00', '09:00-20:00'),
                                   (6, 6, 'выходной', '10:00-17:00'),
                                   (7, 7, 'выходной', 'выходной');
INSERT INTO offices VALUES (1, 'ДО «Солнечногорский» Филиала № 7701 Банка ВТБ (ПАО)', '141506, Московская область, г. Солнечногорск, ул. Красная, д. 60',
                           1, 'есть РКО', 'Да (Зона Привилегия)','Y', 'Универсальный', '',56.184479, 36.984314, true),
                        (2, 'ДО «На Баранова» Филиала № 7701 Банка ВТБ (ПАО)', '141500, Московская область, г. Солнечногорск, ул. Баранова, д. 1, 1-й этаж',
                           2, 'нет РКО', 'Да','', 'Стандарт',  '', 56.183239, 36.9757, false);



