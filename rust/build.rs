extern crate cbindgen;

use std::env;
use std::path::Path;

fn main() {
    let crate_dir = env::var("CARGO_MANIFEST_DIR").unwrap();
    let crate_path = Path::new(&crate_dir);
    let config = cbindgen::Config::from_root_or_default(&crate_path);

  println!("cargo:rerun-if-changed=src/");

    cbindgen::generate_with_config(&crate_path, config)
      .expect("Unable to generate bindings")
      .write_to_file("../include/ipfn/rust/bindings.h");
}
