docker-compose exec db ./cockroach sql --insecure -e "CREATE DATABASE IF NOT EXISTS lt"
docker-compose exec db ./cockroach sql --insecure -e "CREATE TABLE IF NOT EXISTS lt.lasts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    last TIMESTAMPTZ,
    tag STRING,
    INDEX last_idx (last),
    INDEX tag_idx (tag)
);"
docker-compose exec db ./cockroach sql --insecure -e "CREATE USER IF NOT EXISTS lt;"
docker-compose exec db ./cockroach sql --insecure -e "GRANT ALL ON DATABASE lt TO lt;"
