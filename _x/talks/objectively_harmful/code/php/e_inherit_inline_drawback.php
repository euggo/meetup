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
// END1 OMIT

// BGN2 OMIT
class Human implements greeter, messager {
	protected $name;

    function __construct($name) {
        $this->name = $name;
    }

    public function greeting($name) {
        return "Hello, $name. I'm $this->name.";
    }

    public function message() {
        return "Nice to meet you.";
    }
}
// END2 OMIT

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

class Werewolf extends Wolf {
    function __construct($name, $freq) {
        parent::__construct($freq);
        Human::__construct($name);
    }

    public function greeting($name) {
        return parent::greeting($name) . " " . Human::greeting($name);
    }
}

// BGN2 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
meet("Dan", $a, $b, $c);
talk($a, $b, $c);
// END2 OMIT
