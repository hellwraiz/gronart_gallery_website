CREATE TABLE IF NOT EXISTS paintings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uuid TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    author TEXT NOT NULL,
    size TEXT NOT NULL,
    price INTEGER NOT NULL,
    img_url TEXT NOT NULL,
    technique TEXT NOT NULL,
    uploaded_at DATETIME DEFAULT (datetime('now', 'utc')),
    last_edited_at DATETIME DEFAULT (datetime('now', 'utc'))
);

CREATE INDEX IF NOT EXISTS idx_paintings_uuid ON paintings(uuid);
