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
    uploaded_at DATETIME DEFAULT (datetime('now', 'utc')),
    last_edited_at DATETIME DEFAULT (datetime('now', 'utc'))
);

INSERT INTO new_paintings (uuid, name, author, size, price, img_url, technique, description, position, sold, printable, copiable, uploaded_at, last_edited_at)
SELECT uuid, name, author, size, price, img_url, technique, '', 0, 0, 0, 0, uploaded_at, last_edited_at FROM paintings;

UPDATE new_paintings
SET position = rowid;

DROP TABLE paintings;
ALTER TABLE new_paintings RENAME TO paintings;


CREATE TRIGGER paintings_set_position
AFTER INSERT ON paintings
FOR EACH ROW
WHEN NEW.position IS NULL
BEGIN
  UPDATE paintings
  SET position = (
    SELECT COALESCE(MAX(position), 0)
    FROM paintings
    WHERE rowid != NEW.rowid
  ) + 1
  WHERE rowid = NEW.rowid;
END;
