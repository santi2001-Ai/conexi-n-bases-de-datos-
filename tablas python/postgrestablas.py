import psycopg2

# conectar a PostgresSQL
conn = psycopg2.connect(
    dbname="postgres",
    user="postgres",
    password="12345",
    host="localhost"
)
conn.autocommit = True 
cursor = conn.cursor() 

# Crear base de datos 
cursor.execute("CREATE DATABASE ecommerce306")
conn.close()  # Cerrar conexion para conectar a la nueva base 


# conectar a la nueva base de datos
conn = psycopg2.connect(
    dbname="ecommerce306",
    user="postgres",
    password="12345",
    host="localhost"
)  
cursor = conn.cursor()

# Crear tablas
tables = {
    "clientes": """
        CREATE TABLE IF NOT EXISTS clientes (
            id SERIAL PRIMARY KEY,
            nombre VARCHAR(255),
            correo VARCHAR(255) UNIQUE,
            telefono VARCHAR(20),
            direccion TEXT
        )
    """,
    "productos": """
        CREATE TABLE IF NOT EXISTS productos (
            id SERIAL PRIMARY KEY,
            nombre VARCHAR(255),
            descripcion TEXT,
            precio DECIMAL(10,2),
            stock INT
        )
    """,
    "ventas": """
        CREATE TABLE IF NOT EXISTS ventas (
            id SERIAL PRIMARY KEY,
            cliente_id INT,
            fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            total DECIMAL(10,2),
            FOREIGN KEY (cliente_id) REFERENCES clientes(id)
        )
    """,
    "detalle_ventas": """
        CREATE TABLE IF NOT EXISTS detalle_ventas (
        id SERIAL PRIMARY KEY,
        venta_id INT REFERENCES ventas(id),
        producto_id INT REFERENCES productos(id),
        cantidad INT,
        subtotal DECIMAL(10,2)
        )
    """,
    "compras": """
        CREATE TABLE IF NOT EXISTS compras (
            id SERIAL PRIMARY KEY,
            fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            total DECIMAL(10,2)
        )
    """,
    "detalle_compras": """
        CREATE TABLE IF NOT EXISTS detalle_compras (
            id SERIAL PRIMARY KEY,
            compra_id INT REFERENCES compras(id),
            producto_id INT REFERENCES productos(id),
            cantidad INT,
            costo DECIMAL(10,2)
        )
    """
}

for name, query in tables.items():
    cursor.execute(query)
    print(f"Tabla {name} creada.") 

conn.commit()
cursor.close()
conn.close()  # Cerrar conexion a la base de datos