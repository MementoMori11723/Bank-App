-- Account Table
CREATE TABLE IF NOT EXISTS account (
  id TEXT PRIMARY KEY,
  first_name TEXT NOT NULL CHECK (LENGTH(first_name) > 0),
  last_name TEXT NOT NULL CHECK (LENGTH(last_name) > 0),
  username TEXT UNIQUE NOT NULL CHECK (LENGTH(username) >= 3),
  email TEXT NOT NULL CHECK (email LIKE '%@%._%'),
  password TEXT NOT NULL CHECK (LENGTH(password) >= 8),
  balance REAL NOT NULL DEFAULT 0.0 CHECK (balance >= 0),
  image_url TEXT NOT NULL DEFAULT 'https://api.dicebear.com/9.x/big-smile/svg?seed=user'
) STRICT;

-- Transaction History Table
CREATE TABLE IF NOT EXISTS history (
  id TEXT PRIMARY KEY, 
  sender TEXT NOT NULL, 
  receiver TEXT NOT NULL, 
  amount REAL NOT NULL CHECK (amount > 0), 
  timestamp TEXT NOT NULL, 
  FOREIGN KEY (sender) REFERENCES account (username) ON DELETE CASCADE, 
  FOREIGN KEY (receiver) REFERENCES account (username) ON DELETE CASCADE 
) STRICT;
