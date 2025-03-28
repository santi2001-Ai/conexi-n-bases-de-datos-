<?php
// Configurar la conexión a la base de datos
$host = "localhost";
$user = "root";
$password = "";
$dbname = ""; // Agrega el nombre de la base de datos si es necesario

// Crear conexión
$conn = new mysqli($host, $user, $password, $dbname);

// Verificar conexión
if ($conn->connect_error) {
    die("❌ Error al conectar a MySQL: " . $conn->connect_error);
}

echo "✅ Conexión exitosa a MySQL";

// Cerrar conexión (opcional)
// $conn->close();
?>
