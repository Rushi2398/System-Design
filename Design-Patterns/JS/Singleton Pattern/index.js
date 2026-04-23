class Config {
  constructor() {
    if (Config.instance) {
      return Config.instance;
    }
    this.Port = 8080;
    Config.instance = this;
  }
}

const c1 = new Config();
const c2 = new Config();

Object.freeze(c1);
Object.freeze(c2);

console.log(c1 == c2);
console.log(c1.Port);
