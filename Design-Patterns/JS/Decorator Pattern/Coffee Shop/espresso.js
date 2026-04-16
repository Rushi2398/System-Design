import Beverage from "./beverage.js";

export class Espresso extends Beverage {
  cost() {
    return 1.99;
  }

  description() {
    return "Espresso";
  }
}
