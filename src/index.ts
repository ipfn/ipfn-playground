/**
 * @license Apache-2.0
 */

export * from './codec';
export * from './codecs';
export * from './helpers';
export * from './prefix';

import { register } from './codecs';
import { codec as json } from './json';

// JSON codec `cell-json-v1`
register(json);
