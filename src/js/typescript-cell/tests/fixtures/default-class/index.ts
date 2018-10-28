
export default class Counter {
  public count: number = 0;

  constructor() {

  }

  increment(count: number) {
    this.count += count;
    // TODO:
    // return this;
  }

  get countGetter(): number {
    return this.count;
  }

  set countSetter(count: number) {
    this.count = count;
  }
}