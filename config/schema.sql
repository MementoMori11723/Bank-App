-- Account table 
create table account (id int, first_name text, last_name text, username text, email text, password text, balance float); 

-- Transaction
create table transaction (id int, sender string, receiver string, amount float, timestamp string);
