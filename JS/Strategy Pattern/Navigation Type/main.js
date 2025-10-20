import { Navigator } from "./navigator.js";
import { CarNavigation } from "./carNavigation.js";
import { BusNavigation } from "./busNavigation.js";
import { WalkNavigation } from "./walkNavigation.js";

const carNavigator = new Navigator(new CarNavigation());
const busNavigator = new Navigator(new BusNavigation());
const walkNavigator = new Navigator(new WalkNavigation());

carNavigator.navigate("Home", "Office");

busNavigator.navigate("Park", "Cafe");

walkNavigator.navigate("Airport", "Hotel");
