#!/bin/sh
sqlite3 /db/database.db <<EOF
create table status (host TEXT,desired INTEGER,interval INTEGER,method TEXT,proxy TEXT,lastCode INTEGER);
insert into status (host,desired,interval,method,proxy,lastCode) values ('https://www.google.com','200','30','GET','','200');
EOF