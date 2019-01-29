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

// BGN1 OMIT
class Wolf extends Human {
    protected $freq = 1;

    function __construct($freq) {
        parent::__construct("");
        $this->freq = $freq;
    }

    public function greeting($_) {
        $msg = str_repeat("woof ", $this->freq);
        return trim($msg)."!";
    }
}
// END1 OMIT

// BGN2 OMIT
class Werewolf extends Wolf {
    function __construct($name, $freq) {
        parent::__construct($freq);
        Human::__construct($name);
    }

    public function greeting($name) {
        return parent::greeting($name) . " " . Human::greeting($name);
    }
}
// END2 OMIT

$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Bob", 1);
meet("Carlos", $a, $b, $c);
