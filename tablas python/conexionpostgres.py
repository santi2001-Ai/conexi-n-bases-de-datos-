import psycopg2

# conectar a PostgresSQL
conn = psycopg2.connect(
    dbname = "postgres",
    user = "postgres",
    password = "12345",
    host = "localhost"
)
print("Conexión exitosa")
print(conn) 