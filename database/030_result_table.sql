SET search_path TO pulse, public;

CREATE TABLE IF NOT EXISTS results (
    id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    monitor_id UUID NOT NULL,
    result_status monitor_status NOT NULL,
    http_status INTEGER NOT NULL,
    latency INTEGER NOT NULL,
    error TEXT,
    checked_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_results_monitor
        FOREIGN KEY (monitor_id)
        REFERENCES monitors (id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_results_monitor_id ON results (monitor_id);
