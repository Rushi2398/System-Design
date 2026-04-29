class StripeAdapter {
  constructor(stripe) {
    this.stripe = stripe;
  }

  pay(amount) {
    const res = this.stripe.createCharge(amount);
    return {
      id: res.id,
      amount: res.amount,
    };
  }
}

class PayPalAdapter {
  constructor(paypal) {
    this.paypal = paypal;
  }

  pay(amount) {
    const res = this.paypal.sendPayment(amount);
    return {
      id: res.txnId,
      amount: res.total,
    };
  }
}

module.exports = { StripeAdapter, PayPalAdapter };
