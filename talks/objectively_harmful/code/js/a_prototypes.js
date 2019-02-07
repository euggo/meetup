// BGN1 OMIT
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

a = new Human("Alice");
b = new Wolf(3);
username = "Dan";
console.log(a.greeting(username));
console.log(b.greeting(username));
// END1 OMIT
