function withLogging(handler) {
  return function (req) {
    console.log(`Logging Request: `, req);
    return handler(req);
  };
}

function withAuth(handler) {
  return function (req) {
    if (!req.user) {
      throw new Error("Unauthorised");
    }
    return handler(req);
  };
}

function withRateLimit(handler) {
  return function (req) {
    console.log("Rate limit check passed");
    return handler(req);
  };
}

export default { withLogging, withAuth, withRateLimit };
