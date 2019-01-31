function meet(name, greeters) {
  for (var i = 0; i < greeters.length; i++) {
    console.log(greeters[i].greeting(name));
  };
}

// BGN1 OMIT
function talk(messagers) {
  for (var i = 0; i < messagers.length; i++) {
    console.log(messagers[i].message())
  }
}
// END1 OMIT

// BGN2 OMIT
function Human(name) {
  this.name = name;

  this.greeting = function(name) {
    return "Hello, "+name+". I'm " + this.name + ".";
  };

  this.message = function() {
    return "Nice to meet you.";
  }
}
// END2 OMIT

function Wolf(freq) {
  this.freq = freq;

  this.greeting = function() {
  msg = "woof ".repeat(this.freq).trim();
    return msg + "!"
  };
}

// BGN1 OMIT
function Werewolf(name, freq) {
  var human = new Human(name);
  var wolf = new Wolf(freq);

  Object.setPrototypeOf(wolf, human);
  Object.setPrototypeOf(this, wolf);

  this.greeting = function(name) {
    var parent = Object.getPrototypeOf(this);
    var grandp = Object.getPrototypeOf(parent);
    return parent.greeting(name) + " " + grandp.greeting(name);
  };
}
// END1 OMIT

// BGN2 OMIT
a = new Human("Alice");
b = new Wolf(3);
c = new Werewolf("Carlos", 1);
meet("Dan", [a, b, c]);
talk([a, c]);
// END2 OMIT
