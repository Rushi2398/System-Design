import eventBus from "./eventBus.js";
import "./subscribers.js";

eventBus.publish("order_created", { id: "ORD123" });
