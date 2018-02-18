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
   * Value of the cell.
   */
  value?: any;

  /**
   * Bonds of the cell.
   */
  bonds?: Bond[];

  /**
   * Body of the cell.
   */
  body?: Cell[];
}

/**
 * Basic interface of a bond.
 */
export interface Bond {
  /**
   * Name of the bond.
   */
  name?: string;

  /**
   * Kind of a bond (input, output).
   */
  kind?: string;

  /**
   * Source of a bond.
   */
  from?: string;

  /**
   * Target of a bond.
   */
  to?: string;
}
