
declare module 'bs32' {
  export function encode(data: Buffer, pad?: boolean): string;
  export function decode(data: string, unpad?: boolean): Buffer;

  export function _encode(data: Buffer, charset?: string, pad?: boolean): string;
  export function _decode(data: string, table?: number[], unpad?: boolean): Buffer;
}
