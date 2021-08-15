use std::mem::size_of;

static B: [u8; 10] = [99, 97, 114, 114, 121, 116, 111, 119, 101, 108];
static C: [u8; 11] = [116, 104, 97, 110, 107, 115, 102, 105, 115, 104, 0];

fn main() {
    let a: usize = 42;
    let b: &[u8; 10] = &B;
    let c: Box<[u8]> = Box::new(C);

    println!("a (an unsigned integer):");
    println!("  location: {:p}", &a);
    println!("  size:     {:?} bytes", size_of::<usize>());
    println!("  value:    {:?}",  a);
    println!();

    println!("b (an unsigned integer):");
    println!("  location: {:p}", &b);
    println!("  size:     {:?} bytes", size_of::<&[u8; 19]>());
    println!("  value:    {:?}",  b);
    println!();

    println!("c (an unsigned integer):");
    println!("  location: {:p}", &c);
    println!("  size:     {:?} bytes", size_of::<Box<[u8]>>());
    println!("  value:    {:?}",  c);
    println!();

    println!("B (an unsigned integer):");
    println!("  location: {:p}", &B);
    println!("  size:     {:?} bytes", size_of::<[u8; 10]>());
    println!("  value:    {:?}",  B);
    println!();

    println!("C (an unsigned integer):");
    println!("  location: {:p}", &C);
    println!("  size:     {:?} bytes", size_of::<[u8; 11]>());
    println!("  value:    {:?}",  C);
    println!();
}