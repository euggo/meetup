function meet(name, greeters) {
  for (var i = 0; i < greeters.length; i++) {
    console.log(greeters[i].greeting(name));
  };
}

function Human(name) {
  this.name = name;

  this.greeting = function(name) {
    return "Hello, "+name+". I'm " + this.name + ".";
  };
}

function Wolf(freq) {
  this.freq = freq;

  this.greeting = function() {
  msg = "woof ".repeat(this.freq).trim();
    return msg + "!"
  };
}

// BGN1 OMIT
function Werewolf(name, freq) {
  this.human = new Human(name);
  this.wolf = new Wolf(freq);

  Object.assign(this, this.human);
  Object.assign(this, this.wolf);

  this.greeting = function(name) {
    return this.wolf.greeting(name) + " " + this.human.greeting(name);
  };
}
// END1 OMIT

// BGN2 OMIT
a = new Human("Alice");
b = new Wolf(3);
c = new Werewolf("Carlos", 1);
meet("Dan", [a, b, c]);
// END2 OMIT
