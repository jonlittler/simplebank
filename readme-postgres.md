# Postgres
```bash
psql -h raspi5 -p 5433 -U postgres 
psql -h raspi5 -p 5433 -U sys_techschool -d techschool

docker pull postgres:13-alpine
docker images
docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:13-alpine
docker exec -it postgres13 psql -U root
docker exec -it postgres13 /bin/sh
docker logs postgres13
docker stop postgres13
docker start postgres13
docker ps
docker rm postgres13
```

```bash
psql
\d      -- display all
\dt     -- display tables
\q      -- quit

createdb --username=root --owner=root simplebank
dropdb simplebank
psql simplebank

docker exec -it postgres13 createdb --username=root --owner=root simplebank
docker exec -it postgres13 dropdb simplebank
docker exec -it postgres13 psql -U root simplebank
```

## Table Plus
```bash
# 1. Import the TablePlus GPG key
wget -qO - https://deb.tableplus.com/apt.tableplus.com.gpg.key | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/tableplus-archive.gpg > /dev/null

# 2. Add the repository
sudo add-apt-repository "deb [arch=amd64] https://deb.tableplus.com/debian/22 tableplus main"

# 3. Install TablePlus
sudo apt update
sudo apt install tableplus
```
