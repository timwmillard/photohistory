
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE locations (
   id uuid NOT NULL PRIMARY KEY,
   alias varchar(25),
   name varchar(100),
   latitude decimal,
   longitude decimal,
   elevation decimal
);

CREATE TABLE photos (
   id uuid NOT NULL PRIMARY KEY,
   time_taken timestamp WITH TIME ZONE,
   location_id uuid REFERENCES locations (id),
   url varchar(255),
   submitted_by varchar(255)
   latitude decimal,
   longitude decimal,
   elevation decimal
);

CREATE TABLE users (
   id uuid NOT NULL PRIMARY KEY,
   username varchar(255),
   firstname varchar(50),
   lastname varchar(50),
   email varchar(255),
   phone varchar(255)
);

