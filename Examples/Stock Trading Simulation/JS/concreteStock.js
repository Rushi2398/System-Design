export class Stock {
  constructor(stock, price) {
    this.stock = stock;
    this.price = price;
    this.observers = [];
  }

  register(observer) {
    this.observers.push(observer);
  }

  unregister(observer) {
    this.observers.reduce((obs) => observer !== obs);
  }

  notify() {
    this.observers.forEach((obs) => obs.update(this.stock, this.price));
  }

  setPrice(price) {
    console.log(`\n[Stock] ${this.stock} new price: ${price}`);
    this.price = price;
    this.notify();
  }
}
