-- Account Table
CREATE TABLE IF NOT EXISTS account (
  id TEXT PRIMARY KEY, -- Unique identifier for each account
  first_name TEXT NOT NULL CHECK (LENGTH(first_name) > 0), -- First name cannot be empty
  last_name TEXT NOT NULL CHECK (LENGTH(last_name) > 0), -- Last name cannot be empty
  username TEXT UNIQUE NOT NULL CHECK (LENGTH(username) >= 3), -- Username must be unique and at least 3 characters
  email TEXT UNIQUE CHECK (email LIKE '%@%._%'), -- Email must be unique and follow a basic email pattern
  password TEXT NOT NULL CHECK (LENGTH(password) >= 8), -- Password must be at least 8 characters
  balance REAL NOT NULL DEFAULT 0.0 CHECK (balance >= 0) -- Balance defaults to 0.0 and must not be negative
) STRICT;

-- Transaction History Table
CREATE TABLE IF NOT EXISTS history (
  id TEXT PRIMARY KEY, -- Unique identifier for each transaction
  sender TEXT NOT NULL, -- Sender account ID
  receiver TEXT NOT NULL, -- Receiver account ID
  amount REAL NOT NULL CHECK (amount > 0), -- Amount must be greater than zero
  timestamp TEXT DEFAULT (DATETIME('now')), -- Timestamp defaults to the current date and time
  FOREIGN KEY (sender) REFERENCES account (id) ON DELETE CASCADE, -- Cascade deletes for consistency
  FOREIGN KEY (receiver) REFERENCES account (id) ON DELETE CASCADE -- Cascade deletes for consistency
) STRICT;
