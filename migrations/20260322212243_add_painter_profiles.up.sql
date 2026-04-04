CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    uuid TEXT NOT NULL UNIQUE,
    username TEXT NOT NULL,
    role TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    pfp TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at DATETIME DEFAULT (datetime('now', 'utc')),
    last_edited_at DATETIME DEFAULT (datetime('now', 'utc'))
);
