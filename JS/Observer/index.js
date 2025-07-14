import { Subject } from "./Subject.js";
import { ConcreteObserver } from "./ConcreteObserver.js";

const subject = new Subject();

const observer1 = new ConcreteObserver("Observer 1");
const observer2 = new ConcreteObserver("Observer 2");

subject.attach(observer1);
subject.attach(observer2);

subject.notifyObservers("First Notification");

subject.remove(observer1);

subject.notifyObservers("Second Notification");
