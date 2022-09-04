create schema ach_service;

create table ach_service.users (
    usr_id int,
    usr_lvl int
);

create table ach_service.user_achieves (
    uid int,
    achieve_id int,
    achieve_lvl int,
    max_lvl int,
    scan_count int,
    name text,
    last_scan timestamp,
    scanned_locs int[],
    temp_fl bool
);

create table ach_service.logs (
    date timestamp ,
    message text
)
