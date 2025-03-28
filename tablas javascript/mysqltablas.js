// Importar el módulo mysql2
const mysql = require('mysql2');

// Configurar la conexión a la base de datos
const connection = mysql.createConnection({
  host: 'localhost',         // Dirección del servidor MySQL
  user: 'root',        // Usuario de MySQL
  password: '', // Contraseña del usuario
});

// Conectar a la base de datos
connection.connect((err) => {
  if (err) {
    console.error('❌ Error al conectar a MySQL:', err);
    return;
  }
  console.log('✅ Conexión exitosa a MySQL');
});