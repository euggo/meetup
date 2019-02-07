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

function Human(name) {
  this.name = name; // OMIT
  // OMIT
  this.greeting = function(name) { // OMIT
    return "Hello, "+name+". I'm " + this.name + "."; // OMIT
  }; // OMIT
  // ...
  // OMIT
  this.message = function() {
    return "Nice to meet you.";
  }
}
// END1 OMIT

function Wolf(freq) {
  this.freq = freq;

  this.greeting = function() {
  msg = "woof ".repeat(this.freq).trim();
    return msg + "!"
  };
}

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

a = new Human("Alice");
b = new Wolf(3);
c = new Werewolf("Carlos", 1);
// BGN2 OMIT
meet("Dan", [a, b, c]);
talk([a, c]);
// END2 OMIT
