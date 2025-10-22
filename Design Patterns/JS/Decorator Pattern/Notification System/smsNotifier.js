import { NotifierDecorator } from "./notifierDecorator.js";

export class SMSNotifier extends NotifierDecorator {
  send(message) {
    super.send(message);
    console.log("Sending SMS:", message);
  }
}
