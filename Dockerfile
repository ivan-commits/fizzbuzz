FROM debian:stable-slim

# Lib
RUN apt-get update && apt-get install -y \
    nginx \
    php-fpm \
    php-xml \
    php-zip \
    php-intl \
    php-mbstring \
    php-curl \
    curl \
    unzip \
    git \
    && rm -rf /var/lib/apt/lists/*

# Composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# Workdir endpoint
WORKDIR /var/www/html

# Copy default conf
COPY ./docker/nginx/default.conf /etc/nginx/sites-available/default

# Entrypoint sh
COPY ./docker/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Nginx expose
EXPOSE 80

# Run entrypoint
ENTRYPOINT ["/entrypoint.sh"]
