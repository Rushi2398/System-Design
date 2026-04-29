class PaymentBuilder {
  constructor() {
    this.payment = {};
  }
  setID(id) {
    this.payment.id = id;
    return this;
  }
  setAmount(amount) {
    this.payment.amount = amount;
    return this;
  }
  setStatus(status) {
    this.payment.status = status;
    return this;
  }

  build() {
    if (!this.payment.id) throw new Error("Invalid Payment");
    return Object.freeze(this.payment);
  }
}

module.exports = PaymentBuilder;
