CREATE DATABASE santa;

-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       user_id TEXT NOT NULL,
                       name TEXT NOT NULL,
                       username TEXT NOT NULL,
                       groups_member TEXT[] DEFAULT '{}'
);

CREATE TABLE groups (
                        id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                        name TEXT NOT NULL,
                        group_id UUID DEFAULT uuid_generate_v4(),
                        members TEXT[] DEFAULT '{}',
                        admin TEXT
);

CREATE TABLE group_members (
                               group_id UUID REFERENCES groups(id) ON DELETE CASCADE,
                               user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                               PRIMARY KEY (group_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS group_members;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
