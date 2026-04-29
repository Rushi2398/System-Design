class StripeAdapter {
  constructor(stripe) {
    this.stripe = stripe;
  }

  pay(amount) {
    return this.stripe.makePayment(amount);
  }
}

class PayPalAdapter {
  constructor(paypal) {
    this.paypal = paypal;
  }

  pay(amount) {
    return this.paypal.sendPayment(amount);
  }
}

module.exports = { StripeAdapter, PayPalAdapter };
