class SuccessStrategy {
  process(payment) {
    return `Payment ${payment.id} SUCCESS`;
  }
}

class FailureStrategy {
  process(payment) {
    return `Payment ${payment.id} FAILED`;
  }
}

module.exports = { SuccessStrategy, FailureStrategy };
