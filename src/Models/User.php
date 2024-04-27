<?php

namespace Ots\models;

use Ots\config\F;
use Ots\Database;

class User
{
    
    public function __construct()
    {
        F::dd(Database::connect(), false);
    }
    public function getUser()
    {
        return [
            'user' => 'Lothi',
            'email' => 'User@lothi.ch'
        ];
    }
}
