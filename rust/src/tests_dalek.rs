#[cfg(test)]
mod tests {
  use ring::signature::Ed25519KeyPair;
  use ring::{rand, signature};
  use test::Bencher;
  use untrusted;

  #[bench]
  fn bench_sign(b: &mut Bencher) {
    // Generate a key pair in PKCS#8 (v2) format.
    let rng = rand::SystemRandom::new();
    let pkcs8_bytes = Ed25519KeyPair::generate_pkcs8(&rng).unwrap();

    // Normally the application would store the PKCS#8 file persistently. Later
    // it would read the PKCS#8 file from persistent storage to use it.

    let key_pair = Ed25519KeyPair::from_pkcs8(untrusted::Input::from(&pkcs8_bytes)).unwrap();

    const MESSAGE: &'static [u8] = b"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx";

    b.iter(|| key_pair.sign(MESSAGE));
  }

  #[bench]
  fn bench_verify(b: &mut Bencher) {
    // Generate a key pair in PKCS#8 (v2) format.
    let rng = rand::SystemRandom::new();
    let pkcs8_bytes = Ed25519KeyPair::generate_pkcs8(&rng).unwrap();
    let key_pair = Ed25519KeyPair::from_pkcs8(untrusted::Input::from(&pkcs8_bytes)).unwrap();

    const MESSAGE: &'static [u8] = b"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx";
    let sig = key_pair.sign(MESSAGE);

    let peer_public_key_bytes = key_pair.public_key_bytes();
    let sig_bytes = sig.as_ref();

    // Verify the signature of the message using the public key. Normally the
    let peer_public_key = untrusted::Input::from(peer_public_key_bytes);
    let msg = untrusted::Input::from(MESSAGE);
    let sig = untrusted::Input::from(sig_bytes);
    b.iter(|| signature::verify(&signature::ED25519, peer_public_key, msg, sig).unwrap());
  }
}
