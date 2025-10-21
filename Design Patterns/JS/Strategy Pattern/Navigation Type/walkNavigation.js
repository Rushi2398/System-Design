import { NavigationStrategy } from "./navigationInterface.js";

export class WalkNavigation extends NavigationStrategy {
  navigate(source, destination) {
    console.log(`Navigating from ${source} to ${destination} via Walk.`);
  }
}
