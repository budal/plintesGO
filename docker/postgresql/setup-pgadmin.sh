#!/bin/sh

cat >/var/lib/pgadmin/servers <<EOF
{
  "Servers": {
    "1": {
      "Name": "${APP_NAME:-SisCOPOM}",
      "Group": "Servers",
      "Host": "${DB_HOST:-postgres}",
      "Port": ${DB_PORT:-5432},
      "MaintenanceDB": "postgres",
      "Username": "${DB_USERNAME:-siscopom}",
      "PassFile": "/var/lib/pgadmin/pgpass",
      "SSLMode": "prefer"
    }
  }
}
EOF

cat >/var/lib/pgadmin/pgpass <<EOF
${DB_HOST:-postgres}:${DB_PORT:-5432}:*:${DB_USERNAME:-siscopom}:${DB_PASSWORD:-password}
EOF

chmod 600 /var/lib/pgadmin/pgpass
chown pgadmin /var/lib/pgadmin/pgpass

exec /entrypoint.sh "$@"
