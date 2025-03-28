require 'mongo'

# Conectar a MongoDB
client = Mongo::Client.new('mongodb://localhost:27017/')

# Seleccionar la base de datos
db = client[:ecommerce3066478]

# Crear colecciones
clientes = db[:clientes1]
productos = db[:productos1]
ventas = db[:ventas1]
compras = db[:compras1]
comercio = db[:comercio1]

# Insertar documentos de ejemplo
clientes.insert_one({ nombre: "Juan Perez", correo: "juanperez@gmail.com", telefono: "555-1234", direccion: "Calle Falsa 123" })
productos.insert_one({ nombre: "laptop", descripcion: "laptop de alta gama", precio: 1200, stock: 10 })

puts "Base de datos y colecciones creadas en MongoDB"
