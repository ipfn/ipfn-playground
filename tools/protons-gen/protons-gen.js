const fs = require('fs');
const path = require('path');
const protons = require('@ipfn/protons');

const protoDir = path.join(process.cwd(), 'src', 'proto');

const PROTOS = ['cell'];

const readProto = name => fs.readFileSync(path.join(protoDir, `${name}.proto`));
const parseSchema = name => protons(readProto(name));

const protos = PROTOS.map(name => ({ schema: parseSchema(name), name }));
const protoFile = schema => `/**
 * @license Apache-2.0
 * @copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
 */
export default ${JSON.stringify(schema, null, 2)};
`;

protos.forEach(({ schema, name }) => {
  const filePath = path.join(process.cwd(), 'src', 'js', 'cell-pb', 'src', `${name}.pb.ts`);
  const fileData = protoFile(schema);
  fs.writeFileSync(filePath, fileData);
});
