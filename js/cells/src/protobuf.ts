/**
 * @license Apache-2.0
 */
import protons from '@ipfn/protons';

// TODO: typings

/**
 * Protocol buffers messages.
 */
export const protobuf = protons(`
  syntax = "proto3";

  package ipfn;

  message Cell {
    string name = 1;
    string soul = 2;
    bytes memory = 3;
    repeated string bonds = 4;
    repeated Cell body = 5;
  }
`);
