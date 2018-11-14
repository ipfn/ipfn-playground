/**
 * @license Apache-2.0
 * @copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
 */
export default {
  "syntax": 3,
  "package": "ipfn",
  "imports": [],
  "enums": [],
  "messages": [
    {
      "name": "Cell",
      "enums": [],
      "extends": [],
      "messages": [],
      "fields": [
        {
          "name": "name",
          "type": "string",
          "tag": 1,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        },
        {
          "name": "soul",
          "type": "string",
          "tag": 2,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        },
        {
          "name": "value",
          "type": "bytes",
          "tag": 3,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        },
        {
          "name": "bonds",
          "type": "Bond",
          "tag": 4,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": true,
          "options": {}
        },
        {
          "name": "body",
          "type": "Cell",
          "tag": 5,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": true,
          "options": {}
        }
      ],
      "extensions": null,
      "id": "Cell"
    },
    {
      "name": "Bond",
      "enums": [],
      "extends": [],
      "messages": [],
      "fields": [
        {
          "name": "name",
          "type": "string",
          "tag": 1,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        },
        {
          "name": "kind",
          "type": "string",
          "tag": 2,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        },
        {
          "name": "from",
          "type": "string",
          "tag": 3,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        },
        {
          "name": "to",
          "type": "string",
          "tag": 4,
          "map": null,
          "oneof": null,
          "required": false,
          "repeated": false,
          "options": {},
          "packed": false
        }
      ],
      "extensions": null,
      "id": "Bond"
    }
  ],
  "options": {
    "go_package": "github.com/ipfn/ipfn/go/cellpb;cellpb"
  },
  "extends": []
};
