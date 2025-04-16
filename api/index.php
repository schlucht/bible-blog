<?php 
declare(strict_types= 1);
error_reporting(E_ALL);

use Ots\Bible\Core\F;
use Phroute\Phroute\RouteCollector;
use Phroute\Phroute\Dispatcher;
    
$path = parse_url($_SERVER["REQUEST_URI"], PHP_URL_PATH);

include_once(dirname(__FILE__) . "/vendor/autoload.php");


$router = new RouteCollector();

$router->get("/api/", function () {
    return "Hello Lothar";
});

$dispatcher = new Dispatcher($router->getData());
$response = $dispatcher->dispatch($_SERVER["REQUEST_METHOD"], $path);

echo $response;

F::dd($router, true);

