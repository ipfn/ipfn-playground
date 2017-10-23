/**
 * @license Apache-2.0
 */
import * as protons from '@ipfn/protons';


/**
 * Protocol buffers messages.
 */
export const protobuf = protons(`
  syntax = "proto3";

  package ipfn;

  message Cell {
    string name = 1;
    string soul = 2;
    repeated Cell body = 3;
    repeated string bonds = 4;
    Any memory = 5;
  }

  // Messages below are subject to following copyright:
  // Protocol Buffers - Google's data interchange format
  // Copyright 2008 Google Inc.  All rights reserved.
  // https://developers.google.com/protocol-buffers/
  message Any {
    string type_url = 1;
    bytes value = 2;
  }

  message Struct {
    map<string, Value> fields = 1;
  }

  message Value {
    oneof kind {
      NullValue null_value = 1;
      double number_value = 2;
      string string_value = 3;
      bool bool_value = 4;
      Struct struct_value = 5;
      ListValue list_value = 6;
    }
  }

  enum NullValue {
    NULL_VALUE = 0;
  }

  message ListValue {
    repeated Value values = 1;
  }
`);
