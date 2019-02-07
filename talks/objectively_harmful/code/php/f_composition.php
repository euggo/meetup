<?php

interface greeter { public function greeting($str); }

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

interface messager { public function message(); }

function talk(messager ...$messagers) {
    foreach ($messagers as $messager) {
		echo $messager->message()."\n";
	}
}

// BGN1 OMIT
class Human implements greeter, messager { // ... }
	protected $name; // OMIT
    function __construct($name) { $this->name = $name; } // OMIT
    // OMIT
    public function greeting($name) { return "Hello, $name. I'm $this->name."; } // OMIT
    // OMIT
    public function message() { return "Nice to meet you."; } // OMIT
} // OMIT

class Wolf implements greeter { // ... }
    protected $freq = 1; // OMIT
    function __construct($freq) { $this->freq = $freq; } // OMIT
    // OMIT
    public function greeting($_) { // OMIT
        $msg = str_repeat("woof ", $this->freq); // OMIT
        return trim($msg)."!"; // OMIT
    } // OMIT
} // OMIT

class Werewolf implements greeter, messager {
    protected $human;
    protected $wolf;

    function __construct($name, $freq) {
        $this->human = new Human($name);
        $this->wolf = new Wolf($freq);
    }
    public function greeting($name) {
        return $this->wolf->greeting($name) . " " . $this->human->greeting($name);
    }
    public function message() {
        return $this->human->message();
    }
}
// END1 OMIT

$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
// BGN2 OMIT
meet("Dan", $a, $b, $c);
talk($a, $c);
// END2 OMIT
