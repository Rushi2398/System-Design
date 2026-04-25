import eventBus from "./eventBus.js";

eventBus.subscribe("order_created", (order) => {
  console.log(`Email sent for order ${order.id}`);
});

// Inventory service
eventBus.subscribe("order_created", (order) => {
  console.log(`Inventory updated for ${order.id}`);
});

// Shipping service
eventBus.subscribe("order_created", (order) => {
  console.log(`Shipping initiated for ${order.id}`);
});
