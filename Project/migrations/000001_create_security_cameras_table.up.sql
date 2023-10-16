CREATE TABLE if not exists security_cameras (
     id BIGSERIAL PRIMARY KEY,
     created_at TIMESTAMP(0) with time zone NOT NULL default now(),
     manufacturer VARCHAR(255) NOT NULL,
     storage_capacity INT NOT NULL,
     location VARCHAR(255),
     resolution VARCHAR(255) NOT NULL,
     field_of_view REAL NOT NULL,
     recording_duration INT NOT NULL,
     power_source VARCHAR(255)
);