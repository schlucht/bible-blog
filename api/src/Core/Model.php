<?php

namespace Ots\Bible\Core;

use PDO;

abstract class Model
{
    public PDO $pdo;
    protected string $table;
    protected array $fields = [];
    
    public function __construct(PDO $pdo)
    {
        $this->pdo = $pdo;
    }  

    public function all()
    {       
    }

    public function fromID(int $id) 
    {        
    }

}