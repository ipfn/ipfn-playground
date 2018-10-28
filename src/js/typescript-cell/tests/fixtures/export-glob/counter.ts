/**
 * @license Apache-2.0
 */

export namespace remove {
  export function test() {

  }
}

export namespace example {

  export class Counter {
    public count: number = 0;

    /**
     * @param count Counter initial value.
     */
    constructor(count: number = 0) {
      this.count = count;
    }

    /**
     * @param count Increments counter.
     */
    increment(count: number) {
      this.count += count;
      return this.count;
    }
  }

}

function noise() { }
