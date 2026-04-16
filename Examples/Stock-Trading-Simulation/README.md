# 🧩 Stock Trading Simulation System  
### Using **Strategy** and **Observer** Design Patterns

---

## 📖 Overview

This project demonstrates a **Stock Trading Simulation System** that combines two fundamental design patterns:

- **Observer Pattern** – to manage real-time stock price updates and investor notifications.  
- **Strategy Pattern** – to allow investors to dynamically switch between different trading strategies.

The system models a simple market where investors subscribe to stock updates and automatically decide whether to **buy**, **sell**, or **hold** based on their active trading strategy.

---

## 🧠 Design Summary

### 1. Observer Pattern

**Purpose:**  
Allow multiple investors (observers) to subscribe to and receive updates from stock objects (subjects).

**Roles:**
- **Subject (`Stock`)**: Maintains a list of observers and notifies them when its state (price) changes.  
- **Observer (`Investor`)**: Gets updates from the stock and reacts accordingly.

### 2. Strategy Pattern

**Purpose:**  
Encapsulate different trading behaviors and allow investors to switch strategies at runtime.

**Roles:**
- **Strategy Interface (`TradingStrategy`)**: Declares a `DecideAction()` method.
- **Concrete Strategies (`AggressiveStrategy`, `ConservativeStrategy`, etc.)**: Implement different decision-making logic.
- **Context (`Investor`)**: Uses a `TradingStrategy` object to decide how to act on updates.

---

## ⚙️ System Behavior

1. Multiple investors can **subscribe** to the same stock.
2. When a stock’s price changes, it **notifies all its observers**.
3. Each investor reacts to the update according to their **current trading strategy**.
4. Investors can **switch strategies at runtime** without modifying other parts of the system.

---

## 🧩 Example Scenario

**Stocks:** `AAPL`, `TSLA`  
**Investors:**  
- Alice (AggressiveStrategy)  
- Bob (ConservativeStrategy)

**Flow:**
1. Both subscribe to AAPL.
2. AAPL’s price changes → both receive notifications.
3. Alice buys aggressively; Bob holds conservatively.
4. Alice switches her strategy to Conservative at runtime.
5. AAPL’s price changes again → Alice’s reaction changes.
