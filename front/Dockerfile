FROM hugomods/hugo:exts as builder

WORKDIR /src
COPY . .
RUN hugo --minify --gc

FROM caddy:alpine

COPY --from=builder /src/public /srv
COPY Caddyfile /etc/caddy/Caddyfile