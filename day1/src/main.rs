fn part1() {
    let mut total: i128 = 0;
    // Read in the test.txt file
    let input = std::fs::read_to_string("data.txt").unwrap();
    // Iterate over each line
    for line in input.lines() {
        // replace all letters with a |
        let line = line.replace(|c: char| c.is_alphabetic(), "|");
        // split on |
        let line = line.split("|");
        let mut first = "";
        let mut last = "";
        // iterate over all the elements and print them
        for element in line.clone() {
            // skip empty elements
            if element == "" {
                continue;
            }
            // set if first is empty
            if first == "" {
                // get first character
                first = &element[0..1];
            }
            // set last to the current element's last character
            last = &element[element.len() - 1..element.len()];
        }
        // concat first and last and convert to int
        let next: i128 = format!("{}{}", first, last).parse::<i128>().unwrap();
        // print next
        println!("next: {}", next);
        // add to total
        total += next;
    }
    // print total
    println!("total: {}", total);
}

fn part2() {
    // Create struct with string and i128
    struct Number {
        string: String,
        number: i128,
    }

    let mut total: i128 = 0;
    // Read in the test.txt file
    let input = std::fs::read_to_string("data.txt").unwrap();
    // Create vector of Number structs
    let mut numbers: Vec<Number> = Vec::new();
    // Populate with numbers 0-9
    for i in 1..10 {
        numbers.push(Number {
            string: format!("{}", i),
            number: i,
        });
    }
    numbers.push(Number {
        string: "one".to_string(),
        number: 1,
    }); 
    numbers.push(Number {
        string: "two".to_string(),
        number: 2,
    });
    numbers.push(Number {
        string: "three".to_string(),
        number: 3,
    });
    numbers.push(Number {
        string: "four".to_string(),
        number: 4,
    });
    numbers.push(Number {
        string: "five".to_string(),
        number: 5,
    });
    numbers.push(Number {
        string: "six".to_string(),
        number: 6,
    });
    numbers.push(Number {
        string: "seven".to_string(),
        number: 7,
    });
    numbers.push(Number {
        string: "eight".to_string(),
        number: 8,
    });
    numbers.push(Number {
        string: "nine".to_string(),
        number: 9,
    });

    // Iterate over each line
    for line in input.lines() {
        // set first_position to max
        let mut first_position = line.len() + 1;
        // set last_position to min
        let mut last_position = 0;

        let mut first = &Number {
            string: "".to_string(),
            number: 0,
        };
        let mut last = &Number {
            string: "".to_string(),
            number: 0,
        };

        // iterate over each number in the number vector
        for number in &numbers {
            // get the position of the number
            let first_location = line.find(&number.string);
            let last_location = line.rfind(&number.string);
            // if the number is not in the line, skip it
            if first_location == None {
                continue;
            }
            // if the position is less than first_position, set first_position to position
            if first_location.unwrap() <= first_position {
                first_position = first_location.unwrap();
                first = number;
            }
            // if the position is greater than last_position, set last_position to position
            if last_location.unwrap() >= last_position {
                last_position = last_location.unwrap();
                last = number;
            }
        }

        // concat first and last and convert to int
        let next: i128 = format!("{}{}", first.number, last.number).parse::<i128>().unwrap();
        // print next
        println!("next: {} - {} {} - {}", next, first.string, last.string, line);
        // add to total
        total += next;

        // stop if number not found
        if first.string == "" || last.string == "" {
            println!("number not found");
            break;
        }
    }
    // print total
    println!("total: {}", total);
}

fn main() {
    part2();
}