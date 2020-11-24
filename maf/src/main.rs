use std::collections::HashMap;
// use std::ops::{Deref, DerefMut};
use std::fmt;

/// Return the GCD of a, b
fn gcd(a: u64, b: u64) -> u64 {
    let (mut a, mut b) = (a, b);
    while b != 0 {
        let old_a = a;
        a = b;
        b = old_a % b;
    }
    a
}

/// Return whether a number is prime. This is O(2^n).
fn is_prime(n: u64) -> bool {
    for i in 2..n {
        if n % i == 0 {
            return false
        }
    }
    true
}

struct Factors(pub HashMap<u64, u64>);

impl fmt::Display for Factors {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut s = String::from("");
        for (i, v) in self.0.iter() {
            s.push_str(i.to_string().as_str());
            s.push('^');
            s.push_str(v.to_string().as_str());
            s.push_str(" * ");
        }
        write!(f, "{}", s)
    }
}

/// Return a factors type. Probably don't even want a hashmap here.
fn factor(n: u64) -> Factors {
    let mut factors: HashMap<u64, u64> = HashMap::default();
    let mut p: u64 = n;
    let mut i: u64 = 2;

    while p != 1{
        if p % i == 0 {
            if let Some(c) = factors.get(&i) {
                factors.insert(i, c+1);
            } else {
                factors.insert(i, 1);
            }
            p /= i;
        } else {
            i += 1;
        }
    }

    Factors(factors)
}

fn fermat_prime(n: u64) -> bool {
    let u = exp(2, n-1, n);
    u == 1
}


// Return b^k ... but can't handle overflow
fn exp(b: u64, k: u64, m: u64) -> u64 {


    let mut total = 1u64;
    let mut pow = b;
    let mut mask = 1u64;

    for i in 1..64 {

        if pow == 0 {
            continue;
        }

        if std::u64::MAX / pow < pow {
            break;
        }

        if mask & k != 0 {
            total = (total*pow) % m;
        }

        pow = (pow * pow) % m;
        mask *= 2;
    }

    total
}



fn main() {
    // println!("{}", gcd(15, 5));
    // println!("{}", gcd(12, 8));
    // println!("{}", gcd(2*27, 3*27));


    for i in 15..100000 {
         if is_prime(i) != fermat_prime(i) {
             println!("i={} => {} != {}", i, "-", fermat_prime(i));
         }
    }


    /*
    println!("{}", exp(3, 1, 2000));
    println!("{}", exp(3, 2, 2000));
    println!("{}", exp(5, 7, 2000));
    println!("{}", exp(9, 24, 1021921));
    println!("{}", exp(3, 5, 2000));
    */
    // println!("{}", exp(3, 8+4+2+1));
    // println!("{}", factor(26));
    // println!("{}", factor(3*5*29*30));
    assert!(!is_prime(27));
    assert!(gcd(15, 5) == 5)
}
