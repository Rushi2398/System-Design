import { NotifierDecorator } from "./notifierDecorator.js";

export class PushNotifier extends NotifierDecorator {
  send(message) {
    this.notifier.send(message);
    console.log("Sending Push Notification:", message);
  }
}
