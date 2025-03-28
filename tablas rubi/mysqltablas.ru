require 'mysql2'

# Conectar a MySQL
client = Mysql2::Client.new(
  host: "localhost",
  username: "root",
  password: ""
)

# Crear la base de datos si no existe
client.query("CREATE DATABASE IF NOT EXISTS ecommerce30664")

# Conectar a la base de datos creada
client.query("USE ecommerce30664")

# Crear tablas
tables = {
  "clientes" => "
    CREATE TABLE IF NOT EXISTS clientes (
      id INT AUTO_INCREMENT PRIMARY KEY,
      nombre VARCHAR(255),
      correo VARCHAR(255) UNIQUE,
      telefono VARCHAR(20),
      direccion TEXT
    )
  ",
  "productos" => "
    CREATE TABLE IF NOT EXISTS productos (
      id INT AUTO_INCREMENT PRIMARY KEY, 
      nombre VARCHAR(255),
      descripcion TEXT,
      precio DECIMAL(10,2),
      stock INT
    )
  ",
  "ventas" => "
    CREATE TABLE IF NOT EXISTS ventas (
      id INT AUTO_INCREMENT PRIMARY KEY,
      cliente_id INT,
      fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      total DECIMAL(10,2),
      FOREIGN KEY (cliente_id) REFERENCES clientes(id)
    )
  ",
  "detalle_ventas" => "
    CREATE TABLE IF NOT EXISTS detalle_ventas (
      id INT AUTO_INCREMENT PRIMARY KEY,
      venta_id INT,
      producto_id INT,
      cantidad INT,
      subtotal DECIMAL(10,2),
      FOREIGN KEY (venta_id) REFERENCES ventas(id),
      FOREIGN KEY (producto_id) REFERENCES productos(id)
    )
  ",
  "compras" => "
    CREATE TABLE IF NOT EXISTS compras (
      id INT AUTO_INCREMENT PRIMARY KEY,
      fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      total DECIMAL(10,2)
    )
  ",
  "detalle_compras" => "
    CREATE TABLE IF NOT EXISTS detalle_compras (
      id INT AUTO_INCREMENT PRIMARY KEY,
      compra_id INT,
      producto_id INT,
      cantidad INT,
      costo DECIMAL(10,2),
      FOREIGN KEY (compra_id) REFERENCES compras(id),
      FOREIGN KEY (producto_id) REFERENCES productos(id)
    )
  "
}

tables.each do |name, query|
  client.query(query)
  puts "Tabla #{name} creada."
end

client.close
