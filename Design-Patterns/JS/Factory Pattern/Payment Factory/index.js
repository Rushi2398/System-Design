const PaymentFactory = require("./paymentFactory");
const { UPI, Card } = require("./payments");

// Register types (usually done at startup)
PaymentFactory.register("upi", UPI);
PaymentFactory.register("card", Card);

// Create objects dynamically
const p1 = PaymentFactory.create("upi", "user@upi");
const p2 = PaymentFactory.create("card", "1234-xxxx");

console.log(p1.pay(1000));
console.log(p2.pay(2000));
