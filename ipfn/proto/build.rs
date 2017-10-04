extern crate capnpc;

fn main() {
     capnpc::CompilerCommand::new()
        .file("cell.capnp")
        .file("tensor.capnp")
        .run()
        .expect("compiling schema");
}