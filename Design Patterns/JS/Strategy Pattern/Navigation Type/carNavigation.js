import { NavigationStrategy } from "./navigationInterface.js";

export class CarNavigation extends NavigationStrategy {
  navigate(source, destination) {
    console.log(`Navigating from ${source} to ${destination} via Car.`);
  }
}
