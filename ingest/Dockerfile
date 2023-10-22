FROM golang:1.21-bullseye

ARG DB_URL=postgresql://stocks:stocks@stocksdb:5432/coop-stocks
ARG API_KEY=demo
ENV DB_URL=${DB_URL}
ENV ALPHA_VANTAGE_API_KEY=${API_KEY}

WORKDIR /opt/stocks

COPY src src 

COPY entrypoint.sh /entrypoint.sh

RUN go build -C src/app && chmod +x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh"]