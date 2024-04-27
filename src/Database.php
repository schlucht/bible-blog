<?php


namespace Ots;

use PDO;
use Ots\config\F;

class Database
{   
    private $server = 'localhost';
    private $user = 'root';
    private $password = '';
    private $database = 'bible-blog';

    public static function connect()
    {
        try {
        $db = new PDO('mysql:host=' . $server . ';dbname=' . $database . ';charset=utf8', $user, $password);
        $db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);        
        return $db;
        } catch (PDOException $e) {
            echo 'Connection failed: ' . $e->getMessage();
        }
    }
}
