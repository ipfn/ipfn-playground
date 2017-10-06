//! IPFN cap'n proto structs and interfaces.

extern crate capnp;

pub mod cell_capnp {
    include!(concat!(env!("OUT_DIR"), "/cell_capnp.rs"));
}

pub mod tensor_capnp {
    include!(concat!(env!("OUT_DIR"), "/tensor_capnp.rs"));
}

pub use cell_capnp::*;
pub use tensor_capnp::*;


#[cfg(test)]
mod tests {
    use capnp::serialize_packed;
    use super::cell;

    #[test]
    fn it_works() {
        let mut message = ::capnp::message::Builder::new_default();

        {
            let mut cell = message.init_root::<cell::Builder>();

            {
                cell.set_name("test");
                cell.set_soul("test");
            }

            let mut cells = cell.init_cells(1);
            
            {
                let mut cell = cells.borrow().get(0);
                cell.set_soul("test2");
            }
        }

        serialize_packed::write_message(&mut ::std::io::stdout(), &message)
          .expect("serialize message");
    }
}
