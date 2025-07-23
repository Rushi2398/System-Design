import { ModernFurnitureFactory } from "./modern.js";
import { VictorianFurnitureFactory } from "./victorian.js";

function furnish(factory) {
  const chair = factory.createChair();
  const sofa = factory.createSofa();
  chair.sit();
  sofa.lieDown();
}

console.log("== Modern Furniture ==");
furnish(new ModernFurnitureFactory());

console.log("\n== Victorian Furniture ==");
furnish(new VictorianFurnitureFactory());
