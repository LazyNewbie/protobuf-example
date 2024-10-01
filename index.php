<?php

require_once 'vendor/autoload.php';

$message = new \Main\Event([
    'date' => date('Ymd'),
    'time' => date('His'),
    'country' => 'Portugal',
    'idLanguage' => 555,
    'siteHostName' => 'google.com',
    'idAdvertiser' => 658468435,
]);

$jsonData = $message->serializeToJsonString();    // strlen($jsonData) = 122 bytes      {"date":20241001,"time":141225,"country":"Portugal","idLanguage":555,"siteHostName":"google.com","idAdvertiser":658468435}
$binaryData = $message->serializeToString();      // strlen($binaryData) = 40 bytes     ....

$dataStream = new \Google\Protobuf\Internal\CodedInputStream($binaryData);
$restoredFromBin = new \Main\Event(null);
$restoredFromBin->parseFromStream($dataStream);

var_dump($restoredFromBin->serializeToJsonString() === $jsonData);      // true
