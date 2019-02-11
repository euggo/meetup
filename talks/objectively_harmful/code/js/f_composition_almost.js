function meet(name, greeters) {
  for (var i = 0; i < greeters.length; i++) {
    console.log(greeters[i].greeting(name));
  };
}

function talk(messagers) {
  for (var i = 0; i < messagers.length; i++) {
    console.log(messagers[i].message())
  }
}

function Human(name) {
  this.name = name;

  this.greeting = function(name) {
    return "Hello, "+name+". I'm " + this.name + ".";
  };

  this.message = function() {
    return "Nice to meet you.";
  }
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

  Object.assign(this, this.human, this.wolf);

  this.greeting = function(name) {
    return this.wolf.greeting(name) + " " + this.human.greeting(name);
  };
}
// END1 OMIT

a = new Human("Alice");
b = new Wolf(3);
c = new Werewolf("Carlos", 1);
// BGN2 OMIT
meet("Dan", [a, b, c]);
talk([a, c]);
// END2 OMIT
