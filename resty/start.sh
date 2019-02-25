#!/bin/sh

sed -i "s|__PASTE_CLICK_UPSTREAM__|$PASTE_CLICK_UPSTREAM|g;s|__PASTE_CLICK_FRONTEND_BUCKET__|$PASTE_CLICK_FRONTEND_BUCKET|g" \
    /usr/local/openresty/nginx/conf/sites-enabled/paste.click.conf

/usr/local/openresty/bin/openresty -g "daemon off;"
