FROM nginx

# Remove default configuration
RUN rm -rf /usr/share/nginx/html/ \
    /etc/nginx/conf.d/default.conf

# Create our own configuration
COPY nginx.conf /etc/nginx/conf.d/main.conf
