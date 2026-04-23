class UPI {
  constructor(id) {
    this.id = id;
  }

  pay(amount) {
    return `UPI(${this.id}) paid ${amount}`;
  }
}

class Card {
  constructor(cardNumber) {
    this.cardNumber = cardNumber;
  }

  pay(amount) {
    return `Card(${this.cardNumber}) paid ${amount}`;
  }
}

module.exports = { UPI, Card };
