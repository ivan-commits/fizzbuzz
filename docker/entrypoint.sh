#!/bin/bash

service php8.2-fpm restart
service nginx reload

# Keep alive
tail -f /var/log/nginx/access.log /var/log/nginx/error.log
