/**
 * @license Apache-2.0
 */

export function greet(name: string) {
  return `Hello, ${name}!`;
}

export namespace example {

  export namespace deep {
    export function simpleFunction(paramFirst: example2.Request, paramSecond?: string | number): example2.Response {
      return { code: 200 };
    }

    export function simpleRequest(singleEmbedParam: example2.NamedEmbed) {
      return { code: 200 };
    }
  }

  function unexportedFunction() {

  }

}

export namespace example2 {

  export interface Request {
    host: string;
    path?: string;
    method?: string;
    namedUnion: NamedEmbed | string;
    unstd: {
      [key: string]: NamedEmbed | string;
      namedString: string;
    };
    namedEmbed: NamedEmbed;
    optionalUnion?: NamedEmbed | string;
  }

  export interface NamedEmbed {
    embedField: string;
  }

  export interface Response {
    code: number;
  }

}

export * from './deep/func';
