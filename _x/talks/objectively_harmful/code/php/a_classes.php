<?php

// BGN1 OMIT
class Human {
	protected $name;

    function __construct($name) {
        $this->name = $name;
    }

    public function greeting($name) {
        return "Hello, $name. I'm $this->name.";
    }
}
// END1 OMIT

// BGN2 OMIT
class Wolf {
	protected $freq = 1;

    function __construct($freq) {
        $this->freq = $freq;
    }

    public function greeting($_) {
        $msg = str_repeat("woof ", $this->freq);
        return trim($msg)."!";
    }
}
// END2 OMIT

// BGN3 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
$username = "Carlos";
echo $a->greeting($username)."\n";
echo $b->greeting($username)."\n";
// END3 OMIT
