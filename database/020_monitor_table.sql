SET search_path TO pulse, public;

CREATE TABLE IF NOT EXISTS monitors (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    method TEXT NOT NULL,
    headers_json JSONB NOT NULL,
    interval_sec INTEGER NOT NULL,
    acceptable_http_codes INTEGER[] NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    enabled BOOLEAN NOT NULL
);
