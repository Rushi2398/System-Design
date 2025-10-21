import BeverageDecorator from "./beverageDecorator.js";

class Milk extends BeverageDecorator {
  cost() {
    return super.cost() + 0.2;
  }

  description() {
    return super.description() + ", Milk";
  }
}

export default Milk;
