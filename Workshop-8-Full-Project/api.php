<?php
// ob_clean();
header('Content-Type: application/json; charset=utf-8');
$response = [
    "status" => "success",
    "code" => 200,
    "message" => "Data retrieved successfully",
    "data" => [
        "user" => [
            "id" => 101,
            "name" => "John Doe",
            "email" => "john@example.com"
        ]
    ]
];
echo json_encode($response);
exit;
