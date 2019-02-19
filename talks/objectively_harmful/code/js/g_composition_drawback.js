function meet(name, ...greeters) {
  for (var i = 0; i < greeters.length; i++) {
    console.log(greeters[i].greeting(name));
  };
}

function wrappedMeet(name, ...greeters) {
  for (var i = 0; i < greeters.length; i++) {
    console.log(greeters[i].wrappedGreeting(name));
  };
}

function Human(name) {
  this.name = name;

  this.greeting = function(name) {
    return "Hello, "+name+". I'm " + this.name + ".";
  };
}

// BGN1 OMIT
function Werewolf(name) {
  this.human = new Human(name);

  Object.assign(this, this.human);

  this.wrappedGreeting = function(name) {
    return this.human.greeting(name);
  };
}

c = new Werewolf("Carlos");

console.log("update 'human' name to 'Charlie' - call assigned, then wrapped");
c.human.name = "Charlie";
meet("Erin", c);
wrappedMeet("Frank", c);

console.log("update 'werewolf' name to 'Charles' - call assigned, then wrapped");
c.name = "Charles";
meet("Grace", c);
wrappedMeet("Heidi", c);
// END1 OMIT
