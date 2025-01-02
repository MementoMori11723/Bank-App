-- Account table
create table if not exists account (
  id text primary key, 
  first_name text,
  last_name text, 
  username text unique, 
  email text unique, 
  password text, 
  balance real
);

-- Transaction table
create table if not exists transaction (
  id text primary key,
  sender text not null,
  receiver text not null,
  amount real, 
  timestamp text,
  foreign key (sender) references account (id),
  foreign key (receiver) references account (id)
);
