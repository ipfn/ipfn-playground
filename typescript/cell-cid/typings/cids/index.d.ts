
declare module 'cids' {
  class CID {
    constructor(version: number, codec: any, hash: Buffer);

    toBaseEncodedString(): string;
  }
  namespace CID {
    export class CID {
      constructor(version: number, codec: any, hash: Buffer);

      toBaseEncodedString(): string;
    }
  }

  export = CID;
}

