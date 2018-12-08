CREATE TABLE transactions (
  id                serial                        PRIMARY KEY NOT NULL,
  tx_hash           varchar(64)                   NOT NULL UNIQUE,
  sender            varchar(40)                   NOT NULL,
  receiver          varchar(40)                   NOT NULL,
  delegate          varchar(40)                   NOT NULL,
  fee_amount        text                          NOT NULL,
  send_amount       text                          NOT NULL,
  token             text                          NOT NULL,
  nonce             integer                       NOT NULL,
  last_created      timestamp with time zone      NOT NULL DEFAULT current_timestamp,
  last_modified     timestamp with time zone      NOT NULL DEFAULT current_timestamp
);

CREATE OR REPLACE FUNCTION update_last_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_modified = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_transactions_modtime BEFORE UPDATE ON transactions FOR EACH ROW EXECUTE PROCEDURE update_last_modified_column();
