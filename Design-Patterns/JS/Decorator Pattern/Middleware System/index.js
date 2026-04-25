import baseHandler from "./baseHandler.js";
import decorators from "./decorators.js";

const handlers = decorators.withLogging(
  decorators.withAuth(decorators.withRateLimit(baseHandler)),
);

console.log(handlers({ user: "Rushikesh" }));
