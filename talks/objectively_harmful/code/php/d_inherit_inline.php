<?php

interface greeter { public function greeting($str); }

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

class Human implements greeter {
	protected $name;
    function __construct($name) { $this->name = $name; }

    public function greeting($name) { return "Hello, $name. I'm $this->name."; }
}

// BGN1 OMIT
class Wolf extends Human {
    protected $freq = 1; // OMIT
    function __construct($freq) { // OMIT
        $this->freq = $freq; // OMIT
    } // OMIT
    // ...
    public function greeting($_) { // OMIT
        $msg = str_repeat("woof ", $this->freq); // OMIT
        return trim($msg)."!"; // OMIT
    } // OMIT
}

class Werewolf extends Wolf {
    function __construct($name, $freq) {
        parent::__construct($freq); Human::__construct($name);
    }
    public function greeting($name) {
        return parent::greeting($name) . " " . Human::greeting($name);
    }
}

$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
meet("Dan", $a, $b, $c);
// END1 OMIT
