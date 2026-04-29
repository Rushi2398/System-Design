const { Stripe, PayPal } = require("./thirdParty.js");
const { StripeAdapter, PayPalAdapter } = require("./adapter.js");
const { AdapterFactory } = require("./factory.js");

AdapterFactory.register("stripe", () => {
  return new PayPalAdapter(new PayPal());
});

AdapterFactory.register("paypal", () => {
  return new PayPalAdapter(new PayPal());
});

const provider = "stripe";

const payment = AdapterFactory.create(provider);

console.log(payment.pay(1000));
