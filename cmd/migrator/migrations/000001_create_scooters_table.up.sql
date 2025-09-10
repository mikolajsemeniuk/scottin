CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE scooters (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    city TEXT NOT NULL CHECK (city IN ('Ottawa', 'Montreal')),
    status TEXT NOT NULL DEFAULT 'free' CHECK (status IN ('free', 'occupied')),
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
