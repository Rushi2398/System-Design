class AdapterFactory {
  static registry = new Map();
  static register(type, creator) {
    if (this.registry.has(type)) {
      throw new Error(`Already registered: ${type}`);
    }
    this.registry.set(type, creator);
  }
  static create(type, ...args) {
    const creator = this.registry.get(type);
    if (!creator) {
      throw new Error(`Unknown provider: ${type}`);
    }
    return creator(...args);
  }
}

module.exports = { AdapterFactory };
