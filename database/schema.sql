CREATE TABLE IF NOT EXISTS email_records (
    id SERIAL PRIMARY KEY,
    recipient TEXT,
    subject TEXT, 
    body TEXT
);