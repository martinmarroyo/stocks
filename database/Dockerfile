FROM postgres:15
ARG PG_USER=default
ARG PG_PASS=defaultpassword
ARG PG_DB=mydefaultdb 
# Set up database
ENV POSTGRES_USER=${PG_USER}
ENV POSTGRES_PASSWORD=${PG_PASS}
ENV POSTGRES_DB=${PG_DB}
# Move data into container
COPY tables /opt/database/tables
# Load data into db
COPY initializedb.sh /docker-entrypoint-initdb.d/