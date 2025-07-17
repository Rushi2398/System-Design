import Espresso from "./espresso.js";
import Milk from "./milk.js";
import Mocha from "./mocha.js";

let drink = new Espresso();
console.log(`${drink.description()}: $${drink.cost().toFixed(2)}`);

drink = new Milk(drink);
console.log(`${drink.description()}: $${drink.cost().toFixed(2)}`);

drink = new Mocha(drink);
console.log(`${drink.description()}: $${drink.cost().toFixed(2)}`);
