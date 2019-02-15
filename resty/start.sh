#!/bin/sh

sed "s;__PASTE_CLICK_UPSTREAM__;$PASTE_CLICK_UPSTREAM;" \
    < /usr/local/openresty/nginx/conf/sites-enabled/paste.click.conf \
    > /usr/local/openresty/nginx/conf/sites-enabled/paste.click.conf

/usr/local/openresty/bin/openresty -g "daemon off;"
