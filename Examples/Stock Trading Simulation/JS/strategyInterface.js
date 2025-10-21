export class TradingStrategy {
  decideAction(stock, price) {
    throw new Error("Must Implement decideAction()");
  }
}
