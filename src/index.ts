/**
 * @license Apache-2.0
 */
import * as CID from 'cids';
import * as multihashes from 'multihashing-async';
import { Cell } from '@ipfn/cell';
import { promisify } from '@ipfn/util';
import { register } from '@ipfn/cell-codecs';
import { encode, codec as protobuf } from '@ipfn/cell-pb';

register(protobuf);

// console.log(multihashes)

const multihashing = promisify(multihashes);

const CID_VER = 1;
const CID_HASH = 'sha3-256';
const CID_CODEC = 'cell-pb-v1';

/**
 * Creates content ID of a Cell.
 */
export async function cellCID(cell: Cell, codec: string = CID_CODEC): Promise<CID> {
  const encodedCell = encode(cell);
  const multihash = await multihashing(encodedCell, CID_HASH);
  return new CID(CID_VER, CID_CODEC, multihash);
}
