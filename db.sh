#!/bin/sh
sqlite3 /db/database.db <<EOF
create table status (host TEXT,interval INTEGER,method TEXT,proxy TEXT);
EOF
#insert into status (host,interval,method,proxy) values ('https://www.google.com','30','GET','');