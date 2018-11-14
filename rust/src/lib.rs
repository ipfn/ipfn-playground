//! Rust C bindings.

use std::os::raw::c_char;
use std::ffi::{CStr, CString};

#[repr(C)]
/// Normal example docs.
pub struct Normal {
    x: i32,
    y: i32,
}

#[no_mangle]
/// Addition example docs.
pub extern fn addition(a: i32, b: i32) -> i32 {
    a + b
}

#[no_mangle]
/// Normalize example docs.
pub extern fn normalize(n: Normal) -> i32 {
    n.x + n.y
}

#[no_mangle]
pub unsafe extern fn theme_song_generate(name: *const c_char) -> *mut c_char {
    let name_str = CStr::from_ptr(name);
    let mut song = String::from("💣 ");
    song.push_str(name_str.to_str().unwrap());
    song.push_str(" Batman! 💣");

    let c_str_song = CString::new(song).unwrap();
    c_str_song.into_raw()
}

#[no_mangle]
pub extern fn theme_song_free(s: *mut c_char) {
    unsafe {
        if s.is_null() { return }
        CString::from_raw(s)
    };
}

#[allow(dead_code)]
pub extern fn fix_linking_when_not_using_stdlib() { panic!() }
