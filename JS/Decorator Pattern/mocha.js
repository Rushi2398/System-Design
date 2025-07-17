import BeverageDecorator from "./beverageDecorator.js";

class Mocha extends BeverageDecorator {
  cost() {
    return super.cost() + 0.3;
  }

  description() {
    return super.description() + ", Mocha";
  }
}

export default Mocha;
