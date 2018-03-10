
declare module 'multicodec' {
  export function getCodec(body: Buffer): string;
  export function rmPrefix(body: Buffer): Buffer;
  export function addPrefix(codec: string, body: Buffer): Buffer;
}

declare module 'multicodec/src/varint-table' {
  const table: { [key: string]: Buffer };
  export = table;
}

declare module 'multicodec/src/util' {
  export function varintBufferEncode(body: Buffer): Buffer;
  export function varintBufferDecode(body: Buffer): Buffer;
}
