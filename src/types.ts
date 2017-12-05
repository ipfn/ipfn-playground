/**
 * @license Apache-2.0
 */

/**
 * Interface of synaptic types.
 */
export interface Types {
  Bool: string;
  Number: string;
  String: string;
}

/**
 * Mapping of native synaptic types to it's names.
 */
export const SynapticTypes: Types = {
  Bool: 'bool',
  Number: 'number',
  String: 'string'
};

export default SynapticTypes;
