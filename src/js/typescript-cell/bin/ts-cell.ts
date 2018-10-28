#!/usr/bin/env node

/**
 * @license Apache-2.0
 */
import * as fs from 'fs';
import * as process from 'process';
import * as compactJSON from 'json-stringify-pretty-compact';

import { PackageParser } from '../src/package';

const parser = new PackageParser(process.cwd());

// Parse cell
const cell = parser.parse();

// Create `cell.json` file
fs.writeFileSync('cell.json', compactJSON(cell));

// Create `cell.pb` file
// fs.writeFileSync('cell.pb', protobuf.Cell.encode(cell));
