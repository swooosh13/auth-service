## How to up mongodb with SCRAM (current SCRAM-SHA-1)
---

0. Previously create .dbdata/db dir for volume

2. build mongodb container

        $ docker-compose up --build -d mongodb

3.  login into mongo (into mongo container)

        $ mongo -u mongoadm -p mongoadm

4. initialize user

        $ > load("docker-entrypoint-initdb.d/init.js")

