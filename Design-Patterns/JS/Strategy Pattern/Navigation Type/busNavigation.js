import { NavigationStrategy } from "./navigationInterface.js";

export class BusNavigation extends NavigationStrategy {
  navigate(source, destination) {
    console.log(`Navigating from ${source} to ${destination} via Bus.`);
  }
}
