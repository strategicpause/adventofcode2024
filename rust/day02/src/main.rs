use std::fs::File;
use std::io::{BufRead,BufReader};
use std::time::Instant;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    assert_eq!(2, run_and_measure("A", part_a, "day02/sample.txt"));
    assert_eq!(4, run_and_measure("B", part_b, "day02/sample.txt"));

    run_and_measure("A", part_a, "day02/input.txt");
    run_and_measure("B", part_b, "day02/input.txt");

    Ok(())
}

fn run_and_measure(part: &str, f: fn(BufReader<File>) -> i64, file_name: &str) -> i64 {
    let file_path = format!("../../data/{}", file_name);
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);

    let start = Instant::now();
    let answer = f(reader);
    let duration = start.elapsed();

    println!("Part {} ({}): {} ({} us)", part, file_name, answer, duration.as_micros());

    answer
}

fn part_a(reader: BufReader<File>) -> i64 {
    reader.lines()
        .map(|line| line.unwrap())
        .map(|line| {
            let x = line.split(' ').map(|c| c.parse::<i64>().unwrap()).collect::<Vec<i64>>();
            if is_safe(x) { 1 } else { 0 }
        })
        .sum()
}

fn is_safe(x: Vec<i64>) -> bool {
    (x.windows(2).all(|t| t[0] < t[1])
        || x.windows(2).all(|t| t[0] > t[1])
    ) && x.windows(2).all(|t| (t[0] - t[1]).abs() >= 1 && (t[0] - t[1]).abs() <= 3)
}

fn part_b(reader: BufReader<File>) -> i64 {
    reader.lines()
        .map(|line| line.unwrap())
        .map(|line| {
            let x = line.split(' ').map(|c| c.parse::<i64>().unwrap()).collect::<Vec<i64>>();
            if is_any_safe(x) { 1 } else { 0 }
        })
        .sum()
}

fn is_any_safe(x: Vec<i64>) -> bool {
    (0..x.len()).any(|i| {
        let mut y = x.to_vec();
        y.remove(i);
        is_safe(y)
    })
}