CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS handlers (
    id SERIAL PRIMARY KEY,
    service_id INTEGER NOT NULL,
    route TEXT NOT NULL,
    CONSTRAINT fk_handlers_service FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS handler_cases (
    id SERIAL PRIMARY KEY,
    handler_id INTEGER NOT NULL,
    tag_case SMALLINT DEFAULT 1,
    request_body TEXT,
    response_body TEXT,
    request_headers JSONB,
    response_headers JSONB,
    CONSTRAINT fk_handler_cases_handler FOREIGN KEY (handler_id) REFERENCES handlers(id) ON DELETE CASCADE
);
