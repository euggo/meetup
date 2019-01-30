<?php

// BGN1 OMIT
interface greeter {
	public function greeting($str);
}

function meet($name, greeter ...$greeters) {
    foreach ($greeters as $greeter) {
		echo $greeter->greeting($name)."\n";
	}
}
// END1 OMIT

// BGN2 OMIT
class Human implements greeter {
// END2 OMIT
	protected $name;

    function __construct($name) {
        $this->name = $name;
    }

    public function greeting($name) {
        return "Hello, $name. I'm $this->name.";
    }
}

// BGN3 OMIT
class Wolf implements greeter {
// END3 OMIT
    protected $freq = 1;

    function __construct($freq) {
        $this->freq = $freq;
    }

    public function greeting($_) {
        $msg = str_repeat("woof ", $this->freq);
        return trim($msg)."!";
    }
}

// BGN3 OMIT
$a = new Human("Alice");
$b = new Wolf(3);
meet("Dan", $a, $b);
// END3 OMIT
