
export namespace simple {

  export function greet();
  export function greet(name: number);
  export function greet(name: string);
  export function greet(name?: number | string) {
    return "hello world";
  }

  export function helloWorld() {
    return "hello world";
  }

  function noise() { }

  namespace priv { }
}
