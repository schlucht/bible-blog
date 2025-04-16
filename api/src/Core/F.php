<?php

namespace OTS\Bible\Core;

class F {

    public static function dd(mixed $data, bool $die = false): void
    {
        echo "<div>"; 
        echo "<pre>";
         print_r($data);
        echo "</pre>";
        echo "</div>";
        if ($die) {
            die("End script");
        }
    }
}