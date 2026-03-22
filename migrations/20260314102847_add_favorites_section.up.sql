CREATE TABLE new_paintings (
    id INTEGER PRIMARY KEY,
    uuid TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    author TEXT NOT NULL,
    size TEXT NOT NULL,
    price INTEGER NOT NULL,
    img_url TEXT NOT NULL,
    technique TEXT NOT NULL,
    description TEXT NOT NULL,
    position INTEGER,
    sold INTEGER NOT NULL,
    printable INTEGER NOT NULL,
    copiable INTEGER NOT NULL,
    favorite INTEGER NOT NULL,
    uploaded_at DATETIME DEFAULT (datetime('now', 'utc')),
    last_edited_at DATETIME DEFAULT (datetime('now', 'utc'))
);

INSERT INTO new_paintings (uuid, name, author, size, price, img_url, technique, description, position, sold, printable, copiable, favorite, uploaded_at, last_edited_at)
SELECT uuid, name, author, size, price, img_url, technique, description, position, sold, printable, copiable, 1, uploaded_at, last_edited_at FROM paintings;

UPDATE new_paintings
SET position = rowid;

DROP TABLE paintings;
ALTER TABLE new_paintings RENAME TO paintings;
