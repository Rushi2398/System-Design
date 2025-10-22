import { EmailNotifier } from "./emailNotifier.js";
import { PushNotifier } from "./pushNotifier.js";
import { SMSNotifier } from "./smsNotifier.js";

let notifier = new EmailNotifier();

notifier = new PushNotifier(notifier);
notifier = new SMSNotifier(notifier);

notifier.send("Hello, Rushikesh");
