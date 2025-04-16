<?php 
namespace Ots\Bible\Core;

use PDO;

class Database
{
    
    public function __construct(
        private string $host, 
        private string $db, 
        private string $username, 
        private string $password)
    {
    }

    public function connect():PDO {
        $str = "mysql:host=$this->host;port=3306;dbname=$this->db";
        $pdo = new PDO($str, $this->username, $this->password);
        $pdo->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);
        $pdo->setAttribute(PDO::ATTR_DEFAULT_FETCH_MODE, PDO::FETCH_ASSOC);
        $pdo->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);
        return $pdo;
    }
}