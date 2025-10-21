import { TradingStrategy } from "./strategyInterface.js";

export class Investor {
  constructor(name, strategy) {
    this.name = name;
    this.strategy = strategy;
  }

  update(stock, price) {
    process.stdout.write(`${this.name} received update for ${stock} → `);
    this.strategy.decideAction(stock, price);
  }

  setStrategy(strategy) {
    console.log(`\n${this.name} switched strategy.`);
    this.strategy = strategy;
  }
}
