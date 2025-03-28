// Importar el módulo pg
const { Client } = require('pg');

// Configurar la conexión a la base de datos
const client = new Client({
  user: 'postgres',        // Reemplaza con tu usuario de PostgreSQL
  host: 'localhost',         // Reemplaza con la dirección del servidor
  database: 'template1', // Reemplaza con el nombre de la base de datos
  password: '12345', // Reemplaza con tu contraseña
  port: 5432,                // Puerto por defecto de PostgreSQL
});

// Conectar a la base de datos
client.connect()
  .then(() => {
    console.log('✅ Conexión exitosa a PostgreSQL');
  })
  .catch((err) => {
    console.error('❌ Error al conectar a PostgreSQL:', err);
  });