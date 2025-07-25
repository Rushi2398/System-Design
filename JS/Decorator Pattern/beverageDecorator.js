import Beverage from "./beverage.js";

class BeverageDecorator extends Beverage {
  constructor(beverage) {
    super();
    this.beverage = beverage;
  }

  cost() {
    return this.beverage.cost();
  }

  description() {
    return this.beverage.description();
  }
}

export default BeverageDecorator;
