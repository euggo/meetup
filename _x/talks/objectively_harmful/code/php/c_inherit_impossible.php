<?php

interface greeter {
	public function greeting($str);
}

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

class Human implements greeter {
	protected $name;

    function __construct($name) {
        $this->name = $name;
    }

    public function greeting($name) {
        return "Hello, $name. I'm $this->name.";
    }
}

class Wolf implements greeter {
    protected $freq = 1;

    function __construct($freq) {
        $this->freq = $freq;
    }

    public function greeting($_) {
        $msg = str_repeat("woof ", $this->freq);
        return trim($msg)."!";
    }
}

// BGN1 OMIT
class Werewolf extends Human, Wolf implements greeter {
    // syntax error, unexpected ',', expecting '{'
}
// END1 OMIT

// BGN2 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
meet("Dan", $a, $b, $c);
