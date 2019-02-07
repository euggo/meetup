function Human(name) {
  this.name = name;

  this.greeting = function(name) {
    return "Hello, "+name+". I'm " + this.name + ".";
  };
};

function Wolf(freq) {
  this.freq = freq;

  this.greeting = function() {
    msg = "woof ".repeat(this.freq).trim();
    return msg + "!"
  };
}

// BGN1 OMIT
function meet(name, greeters) {
  for (var i = 0; i < greeters.length; i++) {
    console.log(greeters[i].greeting(name));
  };
}

a = new Human("Alice");
b = new Wolf(3);
meet("Dan", [a, b]);
// END1 OMIT
