use ferris_says::say; // from the previous step
use std::io::{stdout, BufWriter};
#[path = "basic_samples/samples.rs"] mod samples;
#[path = "basic_samples/struct.rs"] mod structs;


fn main() {
    // let stdout = stdout();
    // let message = String::from("Hello fellow Rustaceans!");
    // let width = message.chars().count();

    // let mut writer = BufWriter::new(stdout.lock());
    // say(&message, width, &mut writer).unwrap();

    samples::printsample();

    samples::arrays_slices();

    structs::structure_sample();
}