CREATE TABLE IF NOT EXISTS balances (
    id VARCHAR(255), 
    account VARCHAR(255), 
    amount DECIMAL(10, 2)
);

INSERT INTO balances (id, account, amount) VALUES ("f0dba890-cc3e-47ef-a5eb-a14940685e8d", "979dc99d-bcde-4000-a96d-8740374735b7", 17000);
INSERT INTO balances (id, account, amount) VALUES ("6215bc5f-f125-40ca-9f25-4ba545a2d500", "02e8191c-d8c5-4bb2-bdec-abcf54cbb8b8", 22500);