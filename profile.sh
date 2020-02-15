sudo sh -c "alp --sum -f /var/log/nginx/access.log > ./nginx-log"
sudo sh -c "mysqldumpslow -s t /var/log/mysql/mysql-slow.log > ./mysql-log"
