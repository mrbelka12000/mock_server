CREATE TABLE IF NOT EXISTS services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS handlers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    service_id INTEGER NOT NULL,
    route TEXT NOT NULL,
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS handler_cases (
     id INTEGER PRIMARY KEY AUTOINCREMENT,
     handler_id INTEGER NOT NULL,
     tag_case INTEGER DEFAULT 1,
     request_body TEXT,
     response_body TEXT,
     request_headers TEXT,
     response_headers TEXT,
     FOREIGN KEY (handler_id) REFERENCES handlers(id) ON DELETE CASCADE
);
