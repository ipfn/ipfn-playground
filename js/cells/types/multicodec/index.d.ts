
declare module 'multicodec' {
  export function getCodec(body: Buffer): string;
  export function rmPrefix(body: Buffer): Buffer;
  export function addPrefix(prefix: string, body: Buffer): Buffer;
}
