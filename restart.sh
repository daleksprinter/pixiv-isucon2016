make
sudo sh -c "echo > /var/log/mysql/mysql-slow.log"
sudo sh -c "echo > /var/log/nginx/access.log"
sudo systemctl restart isu-go.service
sudo systemctl restart nginx

