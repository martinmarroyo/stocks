FROM postgres:15
# Set up database
ENV POSTGRES_USER=stocks
ENV POSTGRES_PASSWORD=stocks
ENV POSTGRES_DB=coop-stocks
# Move data into container
COPY tables /opt/database/tables
# Load data into db
COPY ingest.sh /docker-entrypoint-initdb.d/