import { Observer } from "./Observer.js";

export class ConcreteObserver extends Observer {
  constructor(name) {
    super();
    this.name = name;
  }

  update(data) {
    console.log(`${this.name} received : ${data}`);
  }
}
