class StrategyFactory {
  static registry = new Map();

  static register(type, clazz) {
    this.registry.set(type, clazz);
  }

  static create(type) {
    const Clazz = this.registry.get(type);
    if (!Clazz) throw new Error("Unknown strategy");
    return new Clazz();
  }
}

module.exports = StrategyFactory;
