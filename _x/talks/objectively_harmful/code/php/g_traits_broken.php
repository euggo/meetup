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
trait GreetHuman {
    protected $name;

    public function greeting($name) {
        return "Hello, $name. I'm $this->name.";
    }
}

trait GreetWolf {
    protected $freq = 1;

    public function greeting($_) {
        $msg = str_repeat("woof ", $this->freq);
        return trim($msg)."!";
    }
}
// END1 OMIT

// BGN2 OMIT
class Human implements greeter {
    use GreetHuman;

    function __construct($name) {
        $this->name = $name;
    }
}
// END2 OMIT

// BGN3 OMIT
class Wolf implements greeter {
    use GreetWolf;

    function __construct($freq) {
        $this->freq = $freq;
    }
}
// END3 OMIT

// BGN4 OMIT
class Werewolf implements greeter {
    use GreetHuman, GreetWolf;

    function __construct($name, $freq) {
        $this->name = $name;
        $this->freq = $freq;
    }
}
// END4 OMIT

// BGN7 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
meet("Dan", $a, $b, $c);
// Trait method greeting has not been applied, because there
// are collisions with other trait methods on Werewolf
// END7 OMIT
