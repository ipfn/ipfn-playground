

#[no_mangle]
pub extern fn addition(a: i32, b: i32) -> i32 {
    a + b
}

#[allow(dead_code)]
pub extern fn fix_linking_when_not_using_stdlib() { panic!() }
