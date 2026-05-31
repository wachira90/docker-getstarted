<?php
$host = 'app_mysql';
$db   = 'testdb';
$user = 'root';
$pass = 'Example1234';
$charset = 'utf8mb4';
$dsn = "mysql:host=$host;dbname=$db;charset=$charset";
$options = [
    // Throw an exception when an error occurs (great for debugging)
    PDO::ATTR_ERRMODE            => PDO::ERRMODE_EXCEPTION,
    // Fetch results as an associative array by default
    PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC,
    // Disable emulated prepares to ensure the database handles statements securely
    PDO::ATTR_EMULATE_PREPARES   => false,
];
try {
    $pdo = new PDO($dsn, $user, $pass, $options);
    $sql = "SELECT * FROM `employees` LIMIT 50;";
    $stmt = $pdo->query($sql);
    $employees = $stmt->fetchAll();
    if ($employees) {
        echo "<h2>Employee List (Top 50)</h2>";
        echo "<table border='1' cellpadding='5' style='border-collapse: collapse;'>";
        echo "<tr>";
        foreach (array_keys($employees[0]) as $columnName) {
            echo "<th>" . htmlspecialchars($columnName) . "</th>";
        }
        echo "</tr>";
        foreach ($employees as $row) {
            echo "<tr>";
            foreach ($row as $data) {
                echo "<td>" . htmlspecialchars((string) $data) . "</td>";
            }
            echo "</tr>";
        }
        echo "</table>";
    } else {
        echo "<p>No employees found in the database.</p>";
    }
} catch (PDOException $e) {
    die("Database Error: " . $e->getMessage());
}
?>