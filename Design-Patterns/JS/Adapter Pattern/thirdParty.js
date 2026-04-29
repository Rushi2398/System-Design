class Stripe {
  makePayment(amount) {
    return `Stripe charged ${amount}`;
  }
}

class PayPal {
  sendPayment(amount) {
    return `PayPal charged ${amount}`;
  }
}

module.exports = { Stripe, PayPal };
