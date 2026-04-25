class EventBus {
  constructor() {
    this.events = new Map();
  }

  subscribe(event, listener) {
    if (!this.events.has(event)) {
      this.events.set(event, []);
    }
    this.events.get(event).push(listener);
  }

  publish(event, data) {
    const listeners = this.events.get(event) || [];
    for (const listener of listeners) {
      setImmediate(() => listener(data));
    }
  }
}

export default new EventBus();
