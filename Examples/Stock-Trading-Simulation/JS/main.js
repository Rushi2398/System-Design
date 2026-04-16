import { Stock } from "./concreteStock.js";
import { Investor } from "./concreteInvestor.js";
import {
  AggressiveStrategy,
  ConservativeStrategy,
} from "./concreteStrategy.js";

const aapl = new Stock("AAPL", 150);
const tsla = new Stock("TSLA", 200);

const alice = new Investor("Alice", new AggressiveStrategy());
const bob = new Investor("Bob", new ConservativeStrategy());

aapl.register(alice);
aapl.register(bob);

aapl.setPrice(180);
aapl.setPrice(90);

alice.setStrategy(new ConservativeStrategy());
aapl.setPrice(85);

tsla.register(alice);
tsla.setPrice(300);
