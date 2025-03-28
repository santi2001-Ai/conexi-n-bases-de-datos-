<?php
$host = "localhost"; // Servidor de PostgreSQL
$port = "5432"; // Puerto predeterminado de PostgreSQL
$dbname = "template1"; // Cambia esto por el nombre de tu base de datos
$user = "postgres"; // Usuario de PostgreSQL
$password = "12345"; // Contraseña del usuario

try {
    // Crear conexión con PDO
    $dsn = "pgsql:host=$host;port=$port;dbname=$dbname";
    $pdo = new PDO($dsn, $user, $password, [PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION]);

    echo "✅ Conexión exitosa a PostgreSQL";

    // Cerrar conexión (opcional)
    // $pdo = null;
} catch (PDOException $e) {
    echo "❌ Error al conectar a PostgreSQL: " . $e->getMessage();
}
?>
