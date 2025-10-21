export class Navigator {
  constructor(strategy) {
    this.strategy = strategy;
  }

  navigate(source, destination) {
    if (!this.strategy) {
      console.log("No Navigation Strategy Set");
    }
    this.strategy.navigate(source, destination);
  }
}
