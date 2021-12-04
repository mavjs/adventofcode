use std::fs::File;
use std::io::{BufReader, BufRead};

fn part1(vec: &Vec<Vec<i32>>, bit_length: usize) {
    let mut gamma = Vec::new();

    for bit in 0..bit_length {
        let mut bit_one = 0;
        let mut bit_zero = 0;

        for inner_vec in vec {
            let number = inner_vec[bit];

            if number == 0 {
                bit_zero = bit_zero + 1;
            } else {
                bit_one = bit_one + 1;
            }
        }
        if bit_one > bit_zero {
            gamma.push("1");
        } else {
            gamma.push("0");
        }
    }

    let epsilon = gamma.clone().into_iter().map(|x| { if x == "0" { "1" } else { "0" }}).collect::<Vec<&str>>().join("");
    let value_gamma = isize::from_str_radix(&gamma.join(""), 2).unwrap();
    let value_epsilon = isize::from_str_radix(&epsilon, 2).unwrap();

    println!("Part 1 - gamma: {} ({}), epsilon: {} ({}) - power consumption: {}", gamma.join(""), value_gamma, epsilon, value_epsilon, value_gamma * value_epsilon);
}
// TODO: actually implement part 2 logic here.
fn part2(vec: &Vec<Vec<i32>>, bit_length: usize) {
    let mut gamma = Vec::new();

    for bit in 0..bit_length {
        let mut bit_one = 0;
        let mut bit_zero = 0;

        for inner_vec in vec {
            let number = inner_vec[bit];

            if number == 0 {
                bit_zero = bit_zero + 1;
            } else {
                bit_one = bit_one + 1;
            }
        }
        if bit_one > bit_zero {
            gamma.push("1");
        } else {
            gamma.push("0");
        }
    }

    let epsilon = gamma.clone().into_iter().map(|x| { if x == "0" { "1" } else { "0" }}).collect::<Vec<&str>>().join("");
    let value_gamma = isize::from_str_radix(&gamma.join(""), 2).unwrap();
    let value_epsilon = isize::from_str_radix(&epsilon, 2).unwrap();

    println!("Part 2 - gamma: {} ({}), epsilon: {} ({}) - power consumption: {}", gamma.join(""), value_gamma, epsilon, value_epsilon, value_gamma * value_epsilon);
}

fn main() {
    let contents = File::open("data/input.txt").unwrap();
    let reader = BufReader::new(contents);
    let mut vec = Vec::new();
    let mut bit_length: usize = 0 as usize;


    for line in reader.lines() {
        let number: Vec<i32> = line.unwrap().chars().map(|c| c.to_string().parse::<i32>().unwrap()).collect();
        bit_length = number.len();
        vec.push(number);
    }
    
    part1(&vec, bit_length);

    part2(&vec, bit_length);
}