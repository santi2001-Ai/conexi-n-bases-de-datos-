<?php
require 'vendor/autoload.php'; // Cargar la biblioteca de MongoDB

use MongoDB\Client;

// URL de conexión
$mongoUrl = "mongodb://localhost:27017"; // Servidor local
$dbName = "nombre_de_tu_base_de_datos"; // Reemplaza con el nombre de tu base de datos

try {
    // Crear una instancia del cliente de MongoDB
    $client = new Client($mongoUrl);
    
    // Seleccionar la base de datos
    $db = $client->$dbName;

    echo "✅ Conexión exitosa a MongoDB";
    
    // Aquí puedes realizar operaciones en la base de datos
    // $collection = $db->nombre_de_tu_coleccion;
    // $insertOneResult = $collection->insertOne(['nombre' => 'Ejemplo']);

} catch (Exception $e) {
    echo "❌ Error al conectar a MongoDB: " . $e->getMessage();
}
?>
