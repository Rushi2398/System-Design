class Stripe {
  createCharge(amount) {
    return { id: "txn_1", amount };
  }
}

class PayPal {
  sendPayment(amount) {
    return { txnId: "txn_2", total: amount };
  }
}

module.exports = { Stripe, PayPal };
