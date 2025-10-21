import { TradingStrategy } from "./strategyInterface.js";

export class AggressiveStrategy extends TradingStrategy {
  decideAction(stock, price) {
    console.log(`[Aggressive] ${stock} price: ${price} → BUY immediately!`);
  }
}

export class ConservativeStrategy extends TradingStrategy {
  decideAction(stock, price) {
    if (price < 100) {
      console.log(
        `[Conservative] ${stock} price: ${price} → BUY (undervalued)`
      );
    } else {
      console.log(`[Conservative] ${stock} price: ${price} → HOLD`);
    }
  }
}
