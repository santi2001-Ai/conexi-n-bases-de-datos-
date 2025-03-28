// Importar la biblioteca MongoDB
const { MongoClient } = require('mongodb');

// URL de conexión (ajusta según tu configuración)
const url = 'mongodb://localhost:27017'; // Servidor local
const dbName = 'nombre_de_tu_base_de_datos'; // Reemplaza con el nombre de tu base de datos

// Crear una instancia del cliente
const client = new MongoClient(url, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

// Conectar a la base de datos
async function conectarMongoDB() {
  try {
    await client.connect();
    console.log('✅ Conexión exitosa a MongoDB');
    
    // Seleccionar la base de datos
    const db = client.db(dbName);
    
    // Aquí puedes hacer operaciones en la base de datos
    // const collection = db.collection('tu_colección'); 
    // await collection.insertOne({ nombre: 'Ejemplo' });

  } catch (error) {
    console.error('❌ Error al conectar a MongoDB:', error.message);
  } 
  
}

// Ejecutar la función de conexión
conectarMongoDB();