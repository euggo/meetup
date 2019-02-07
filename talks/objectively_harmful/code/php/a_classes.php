<?php

// BGN1 OMIT
class Human {
	protected $name;
    function __construct($name) { $this->name = $name; }

    public function greeting($name) { return "Hello, $name. I'm $this->name."; }
}

class Wolf {
	protected $freq = 1;
    function __construct($freq) { $this->freq = $freq; }

    public function greeting($_) {
        $msg = str_repeat("woof ", $this->freq);
        return trim($msg)."!";
    }
}

$a = new Human("Alice");
$b = new Wolf(3);
$username = "Dan";
echo $a->greeting($username)."\n";
echo $b->greeting($username)."\n";
// END1 OMIT
