/**
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cell';
import { cell, toCells } from './builder';


/**
 * Interface of JavaScript value.
 */
export type IValue = IObject | string | number | boolean;

/**
 * Interface of JavaScript functions namespace.
 */
export interface IObject {
  [key: string]: any;
}

/**
 * Creates a Cell from JavaScript object.
 */
export function fromObject(obj: IObject): Cell {
  const body = Object.keys(obj)
    .map((name: string) => fromValue(obj[name], name));
  return { body };
}

/**
 * Creates a Cell from JavaScript value.
 */
export function fromValue(value: any, name?: string): Cell {
  if (typeof value === 'boolean') {
    return { name, value, soul: '/synaptic/bool' };
  }
  if (typeof value === 'string') {
    return { name, value, soul: '/synaptic/string' };
  }
  if (typeof value === 'number') {
    return { name, value, soul: '/synaptic/number' };
  }
  if (typeof value === 'object') {
    const body = Object.keys(value)
      .map((name: string) => fromValue(value[name], name));
    return { name, body };
  }
  return { name };
}

/**
 * Creates a JavaScript value from cell.
 */
export function fromCell(cell: Cell): IValue | undefined {
  const { name, body, value } = cell;
  if (value) {
    switch (cell.soul) {
      case '/synaptic/bool':
        return typeof value === 'boolean' ? value : value === '1';
      case '/synaptic/string':
        return value;
      case '/synaptic/number':
        // TODO: 
        return typeof value === 'number' ? value : parseInt(value, 10);
    }
  }
  if (!body) {
    return undefined;
  }
  if (name) {
    return { [name]: fromBody(body) };
  }
  return fromBody(body);
}

/**
 * Creates a JavaScript value from cell.
 */
export function fromBody(body: Cell[]): IObject {
  const result: IObject = {};
  body.forEach((cell: Cell) => {
    if (cell.name) {
      result[cell.name] = fromCell(cell);
    }
  })
  return result;
}
