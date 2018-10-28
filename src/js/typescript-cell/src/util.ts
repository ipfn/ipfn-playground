/**
 * @license Apache-2.0
 */
import { Cell } from '@ipfn/cells';

/**
 * Converts function input parameters to cells.
 */
export function reduceCells(cells: Cell[][]): Cell[] {
  return [].concat.apply([], cells);
}
