class PaymentFactory {
  static registry = new Map();

  static register(type, clazz) {
    if (!type || typeof clazz !== "function") {
      throw new Error("Invalid registration");
    }

    if (this.registry.has(type)) {
      throw new Error(`Type ${type} already registered`);
    }

    this.registry.set(type, clazz);
  }

  static create(type, ...args) {
    const Clazz = this.registry.get(type);

    if (!Clazz) {
      throw new Error(`Unknown payment type: ${type}`);
    }

    return new Clazz(...args);
  }
}

module.exports = PaymentFactory;
