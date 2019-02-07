<?php

interface greeter {
    public function greeting($str);
}

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

// BGN1 OMIT
interface messager {
    public function message();
}

function talk(messager ...$messagers) {
    foreach ($messagers as $messager) {
		echo $messager->message()."\n";
	}
}

class Human implements greeter, messager {
	protected $name; // OMIT
    function __construct($name) { $this->name = $name; } // OMIT
    // OMIT
    public function greeting($name) { return "Hello, $name. I'm $this->name."; } // OMIT
    // ...
    public function message() { return "Nice to meet you."; }
}

class Wolf extends Human { // ... }
    protected $freq = 1; // OMIT
    function __construct($freq) { // OMIT
        $this->freq = $freq; // OMIT
    } // OMIT
    // OMIT
    public function greeting($_) { // OMIT
        $msg = str_repeat("woof ", $this->freq); // OMIT
        return trim($msg)."!"; // OMIT
    } // OMIT
} // OMIT

class Werewolf extends Wolf { // ... }
    function __construct($name, $freq) { // OMIT
        parent::__construct($freq); Human::__construct($name); // OMIT
    } // OMIT
    // OMIT
    public function greeting($name) { // OMIT
        return parent::greeting($name) . " " . Human::greeting($name); // OMIT
    } // OMIT
} // OMIT
// END1 OMIT

$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
// BGN2 OMIT
meet("Dan", $a, $b, $c);
talk($a, $b, $c);
// END2 OMIT
