from pymongo import MongoClient

# Conectar a mongoDB
client = MongoClient("mongodb://localhost:27017/")
db = client["ecommerce3066478"]

# Crear Colecciones 
clientes = db["clientes1"]
productos = db["productos1"]
ventas = db["ventas1"]
compras = db["compras1"]
comercio = db["comercio1"]

# Insertar documentos de ejemplo 
clientes.insert_one({"nombre": "Juan Perez", "correo": "juanperez@gmail.com", "telefono": "555-1234", "direccion": "Calle Falsa 123"})
productos.insert_one({"nombre": "laptop", "descripcion": "laptop de alta gama", "precio": 1200, "stock": 10})

print("Base de datos y colecciones creadas en MongoDB")