/**
 * IPFN cells builder.
 * 
 * @license Apache-2.0
 */
import { Cell } from '../cell';
import { Builder, CellOrBuilder } from './builder';

/**
 * Creates a new instance of Cell Builder.
 */
export function cell(name?: string): Builder {
  return new Builder(name);
}

/**
 * Constructs a cell with `/synaptic/bool` soul.
 */
export const bool = (name: string) => cell(name).soul('/synaptic/bool');

/**
 * Constructs a cell with `/synaptic/buffer` soul.
 */
export const buffer = (name: string) => cell(name).soul('/synaptic/buffer');

/**
 * Constructs a cell with `/synaptic/number` soul.
 */
export const number = (name: string) => cell(name).soul('/synaptic/number');

/**
 * Constructs a cell with `/synaptic/string` soul.
 */
export const string = (name: string) => cell(name).soul('/synaptic/string');

/**
 * Constructs a cell with `/core/input` soul.
 */
export const input = (body: CellOrBuilder[]) => cell().soul('/core/input').body(body);

/**
 * Constructs a cell with `/core/output` soul.
 */
export const output = (body: CellOrBuilder[]) => cell().soul('/core/output').body(body);

/**
 * Constructs a cell with `/core/method` soul.
 */
export const method = (name: string) => cell(name).soul('/core/method');
