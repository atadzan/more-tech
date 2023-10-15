CREATE TABLE IF NOT EXISTS atms (
    id serial,
    address varchar(255) not null,
    latitude float not null,
    longtitude float not null,
    all_day  boolean,
)
CREATE TABLE IF NOT EXISTS atm_services (
    id serial not null,
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
    rub_support_activity varchar(255),
)

CREATE TABLE IF NOT EXISTS offices (
    id bigserial not null,
    sale_point varchar(255),
    address varchar(255),
    schedule_id bigint references office_schedule(id),
    rko varchar(255) not null,
    type varchar(255) not null,
    suo_availability varchar(255),
    sale_point_format varchar(255),
    has_ramp varchar(255),
    latitude decimal not null,
    longtitude decimal not null,
    metro_station varchar(255),
    kep boolean
)

CREATE TABLE IF NOT EXISTS office_schedule(
    id bigserial not null,
    day_id smallint references days(id),
    open_hours varchar(255) not null,
    individual_open_hours varchar(255) not null
)

CREATE TABLE If NOT EXISTS days (
    id smallserial not null,
    title_en varchar(255) not null,
    title_ru varchar(255) not null
)
