use std::fs::File;
use std::io::{BufRead, BufReader};

/*
fn num_in_board(boards: &Vec<Vec<i32>>, num: i32) -> Option<(usize, usize)> {
    let mut board_num: usize = 0;
    let mut board_col: usize = 0;

    for (index, board) in boards.into_iter().enumerate() {
        for element in 0..board.len() {
            if board[element] == num {
                return Some((index, element))
            }
        }
    }
    return None
}

*/

fn move_card(boards: Vec<Vec<u8>>, draw: u8) {
    for (p, board) in boards.as_ref().clone().iter().enumerate() {
        for (q, element) in board.iter().enumerate() {
            if element == &draw {
                boards.clone().get_mut(p);
            }
        }
    }
}
fn part1(draws: Vec<u8>, boards: Vec<Vec<u8>>) {
    for draw in draws {
        move_card(boards, draw)
    }
}

fn main() {
    let contents = File::open("data/test_input.txt").unwrap();
    let reader = BufReader::new(contents);

    let mut lines = reader.lines();
    let draw_num = lines
        .next()
        .unwrap()
        .map(|line| {
            line.split(",")
                .map(|element| element.parse::<u8>().unwrap())
                .collect::<Vec<u8>>()
        })
        .expect("Failed to create draw numbers from file contents");
    println!("Draw numbers: {:?}", draw_num);

    //let mut lines = lines.flat_map(Result::ok);
    let boards: Vec<Vec<u8>> = lines
        .by_ref()
        .filter(|line| !line.as_ref().unwrap().is_empty())
        .map(|line| {
            line.as_ref()
                .unwrap()
                .split_whitespace()
                .map(|element| element.parse::<u8>().unwrap())
                .collect::<Vec<u8>>()
        })
        .collect::<Vec<Vec<u8>>>();
    println!("Boards: {:?}", boards);

    //part1(draw_num, boards);
}
