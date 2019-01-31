<?php

interface greeter {
	public function greeting($str);
}

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}

interface messager {
	public function message();
}

function talk(messager ...$messagers) {
    foreach ($messagers as $messager) {
		echo $messager->message()."\n";
	}
}

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

// BGN1 OMIT
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
// END1 OMIT

// BGN2 OMIT
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
// END2 OMIT

// BGN3 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
$c = new Werewolf("Carlos", 1);
meet("Dan", $a, $b, $c);
talk($a, $c);
// Uncaught TypeError: Argument 2 passed to talk() must implement interface messager
// END3 OMIT
