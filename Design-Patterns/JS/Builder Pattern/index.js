const { Stripe, PayPal } = require("./thirdParty");
const { StripeAdapter, PayPalAdapter } = require("./adapters");
const AdapterFactory = require("./factory");
const PaymentBuilder = require("./builder");
const { SuccessStrategy, FailureStrategy } = require("./strategy");
const StrategyFactory = require("./strategyFactory");

AdapterFactory.register("stripe", () => new StripeAdapter(new Stripe()));
AdapterFactory.register("paypal", () => new PayPalAdapter(new PayPal()));

StrategyFactory.register("SUCCESS", SuccessStrategy);
StrategyFactory.register("FAILED", FailureStrategy);

const provider = "stripe";

const adapter = AdapterFactory.create(provider);

const rawObject = adapter.pay(1000);

const payment = new PaymentBuilder()
  .setID(rawObject.id)
  .setAmount(rawObject.amount)
  .setStatus("FAILED")
  .build();

const status = payment.status;
const strategy = StrategyFactory.create(status);
console.log(strategy.process(payment));
