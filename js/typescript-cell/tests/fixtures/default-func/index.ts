
export default function greet();
export default function greet(name: number);
export default function greet(name: string);
export default function greet(name?: number | string) {
  return "hello world";
}

function priv() { }
