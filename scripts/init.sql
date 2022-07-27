
create table log
(
    http_cf_connecting_ip varchar(100) not null,
    time_local            datetime     not null,
    method                varchar(10)  not null,
    body_bytes_sent       int unsigned not null,
    unique (http_cf_connecting_ip,
            time_local,
            method,
            body_bytes_sent)
);