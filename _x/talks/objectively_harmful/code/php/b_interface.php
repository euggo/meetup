<?php

// BGN1 OMIT
interface greeter { public function greeting($str); }

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

class Human implements greeter {
    // ...
	protected $name; // OMIT
    function __construct($name) { $this->name = $name; } // OMIT
    // OMIT
    public function greeting($name) { return "Hello, $name. I'm $this->name."; } // OMIT
}

class Wolf implements greeter {
    // ...
    protected $freq = 1; // OMIT
    function __construct($freq) { $this->freq = $freq; } // OMIT
    // OMIT
    public function greeting($_) { // OMIT
        $msg = str_repeat("woof ", $this->freq); // OMIT
        return trim($msg)."!"; // OMIT
    } // OMIT
}

$a = new Human("Alice");
$b = new Wolf(3);
meet("Dan", $a, $b);
// END1 OMIT
