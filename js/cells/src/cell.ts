/**
 * @license Apache-2.0
 */

/**
 * Basic interface of a cell.
 */
export interface Cell {
    /**
     * Returns name of the cell.
     */
    name?: string;

    /**
     * Returns name of the cell soul.
     */
    soul?: string;

    /**
     * Returns body of the cell.
     */
    body?: Cell[];

    /**
     * Returns memory of the cell.
     */
    memory?: any;
}
