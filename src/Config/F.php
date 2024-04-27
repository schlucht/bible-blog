<?php

namespace Ots\config;

class F
{
    public static function dd($var, $die=true)
    {
        echo '<div style="color: black;width: 80%;margin: 0 auto;padding: 10px;border: 1px solid red;background: #cecece">';
            echo '<pre style="font-size: 16px">';
                print_r($var);
            echo '</pre>';
        echo '</div>';
        if ($die) {
            die();
        }
    }
}