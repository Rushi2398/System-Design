import { Notifier } from "./notifier.js";

export class NotifierDecorator extends Notifier {
  constructor(notifier) {
    super();
    this.notifier = notifier;
  }

  send(message) {
    this.notifier.send(message);
  }
}
