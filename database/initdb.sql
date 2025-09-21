-- Create schemas
CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS data;

-- AUTH SCHEMA - User authentication
CREATE TABLE auth.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- DATA SCHEMA - profiles
CREATE TABLE data.profiles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    data VARCHAR(255)
);

-- INIT DATA (For testing)
-- Insert sample user
INSERT INTO auth.users (email, username, password_hash)
VALUES ('test@test.com', 'test', '$2b$12$5LxaaJslD/iUVe5aXuBYaeKJQ0bjHc1ZWZE29hr.WOluI3J6J1yt6'); -- password is test

-- Insert sample game profile
INSERT INTO data.profiles (data)
VALUES ('testdata')