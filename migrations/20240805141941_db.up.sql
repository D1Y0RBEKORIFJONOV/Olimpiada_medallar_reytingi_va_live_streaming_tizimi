CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS medals(
    medalid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    countryid TEXT,
    type TEXT,
    eventid TEXT,
    athleteid TEXT
);
CREATE TABLE IF NOT EXISTS rank(
    countryid TEXT,
    gold INT,
    silver INT,
    bronze INT,
    score INT
)