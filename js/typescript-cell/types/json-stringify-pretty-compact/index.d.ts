
declare function stringify(obj?: { [key: string]: any }): string;

declare namespace stringify {
  function stringify(obj?: { [key: string]: any }): string;
}

export = stringify;
