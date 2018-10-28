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
  Float: string;
  Double: string;
  SFixed32: string;
  Fixed32: string;
  Int32: string;
  Int64: string;
  UInt32: string;
  UInt64: string;
  UInt128: string;
  UInt256: string;
  SInt64: string;
  SInt32: string;
  Varint: string;
}

/**
 * Mapping of native synaptic types to it's names.
 */
export const SynapticTypes: Types = {
  Bool: 'bool',
  Number: 'number',
  String: 'string',
  Float: 'float',
  Double: 'double',
  SFixed32: 'sfixed32',
  Fixed32: 'fixed32',
  Int32: 'int32',
  Int64: 'int64',
  UInt32: 'uint32',
  UInt64: 'uint64',
  UInt128: 'uint128',
  UInt256: 'uint256',
  SInt64: 'sint64',
  SInt32: 'sint32',
  Varint: 'varint',
};

export default SynapticTypes;
