/**
 * @license Apache-2.0
 */

/**
 * Basic interface of a cell.
 */
export interface Cell {
  /**
   * Name of the cell.
   */
  name?: string;

  /**
   * Name of the cell soul.
   */
  soul?: string;

  /**
   * Memory of the cell.
   */
  memory?: any;

  /**
   * Body of the cell.
   */
  body?: Cell[];

  /**
   * Bonds of the cell.
   */
  bonds?: string[];
}
