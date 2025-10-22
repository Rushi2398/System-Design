import { NotifierDecorator } from "./notifierDecorator.js";

export class SMSNotifier extends NotifierDecorator {
  send(message) {
    this.notifier.send(message);
    console.log("Sending SMS:", message);
  }
}
