FROM postgres:10.3

# Copy the database initialization script:
COPY up.sql /docker-entrypoint-initdb.d/1.sql

# Command to start postgres
CMD ["postgres"]