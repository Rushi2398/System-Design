import Beverage from "./beverage.js";

class Espresso extends Beverage {
  cost() {
    return 1.99;
  }

  description() {
    return "Espresso";
  }
}

export default Espresso;
