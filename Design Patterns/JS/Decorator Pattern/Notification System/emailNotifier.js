import { Notifier } from "./notifier.js";

export class EmailNotifier extends Notifier {
  send(message) {
    console.log("Sending Message: ", message);
  }
}
