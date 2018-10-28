/**
 * IPFN cells builder.
 * 
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cell';
import { cell, CellBuilder } from './builder';

export namespace synaptic {

  /**
   * Constructs a cell with `/synaptic/bool` soul.
   */
  export const bool = (name?: string): CellBuilder => cell(name).soul('/synaptic/bool');

  /**
   * Constructs a cell with `/synaptic/buffer` soul.
   */
  export const buffer = (name?: string): CellBuilder => cell(name).soul('/synaptic/buffer');

  /**
   * Constructs a cell with `/synaptic/number` soul.
   */
  // tslint:disable-next-line:variable-name
  export const number = (name?: string): CellBuilder => cell(name).soul('/synaptic/number');

  /**
   * Constructs a cell with `/synaptic/string` soul.
   */
  // tslint:disable-next-line:variable-name
  export const string = (name?: string): CellBuilder => cell(name).soul('/synaptic/string');
}

/**
 * Constructs a cell with `/core/input` soul.
 */
export const input = (body: CellBuilder[]): CellBuilder => cell().soul('/core/input').body(body);

/**
 * Constructs a cell with `/core/output` soul.
 */
export const output = (body: CellBuilder[]): CellBuilder => cell().soul('/core/output').body(body);

/**
 * Constructs a cell with `/core/method` soul.
 */
export const method = (name?: string): CellBuilder => cell(name).soul('/core/method');
