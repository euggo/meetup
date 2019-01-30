<?php

interface greeter {
	public function greeting($str);
}

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

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

class Human implements greeter {
    use GreetHuman;

    function __construct($name) {
        $this->name = $name;
    }
}

class Wolf implements greeter {
    use GreetWolf;

    function __construct($freq) {
        $this->freq = $freq;
    }
}

// BGN1 OMIT
class Werewolf implements greeter {
    use GreetHuman, GreetWolf {
        GreetHuman::greeting as humanGreeting;
        GreetWolf::greeting as wolfGreeting;
    }

    function __construct($name, $freq) {
        $this->name = $name;
        $this->freq = $freq;
    }

    public function greeting($name) {
        return $this->wolfGreeting($name) . " " . $this->humanGreeting($name);
    }
}
// END1 OMIT

// BGN2 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
meet("Dan", $a, $b, $c);
// END2 OMIT
