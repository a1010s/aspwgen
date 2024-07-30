# Password Generator in Go

## Overview

Welcome to the modern Password Generator in Go! This tool is designed to produce strong, secure passwords similar to those used by major tech companies like Apple and Google. It offers advanced features and customization options to ensure that your passwords are not only secure but also fit your specific needs.

## Why Use This Tool?

In today's digital world, creating strong and unique passwords is crucial for maintaining security. While traditional tools like `pwgen` have served their purpose, they often lack the modern features and flexibility required to meet today's security standards. My Go-based password generator offers the following benefits:

- **Modern and Flexible**: Adapted to current security practices, including adjustable group lengths, separators, and the number of special characters.
- **Customizable**: Allows you to specify the exact number of groups, length of each group, separators, and special characters.
- **Secure**: Utilizes cryptographic random number generation to ensure high-quality randomness.

## Features

- **Password Length**: Specify the total length of the password, including separators.
- **Group Length**: Define the length of each group of characters.
- **Special Characters**: Choose the number of special characters to include in the password.
- **Separators**: Select a custom separator between groups.
- **Multiple Passwords**: Generate multiple passwords in a single run.

## Installation

1. **Clone the Repository**:

   ```sh
   git clone https://github.com/a1010s/aspwgen.git
   cd aspwgen
   ```

2. Build the APP:
   `go build -o aspwgen main.go`

## Usage

```bash
# The password generator can be run from the command line with various options to tailor the output to your needs. Here are some examples:

# Generate 5 Passwords, Each 20 Characters Long, with Groups of 5 Characters and 2 Special Characters:
./password_generator -n 5 -l 20 -g 5 -c 2


# Generate 3 Passwords, Each 16 Characters Long, with Groups of 4 Characters, 4 Special Characters, and : as the Separator:
./password_generator -n 3 -l 16 -g 4 -s ":" -c 4


# Generate 2 Passwords, Each 12 Characters Long, with Groups of 3 Characters, 2 Special Characters, and . as the Separator:
./password_generator -n 2 -l 12 -g 3 -s "." -c 2


# Command-Line Options
# -n int: Number of passwords to generate. Default is 1.
# -l int: Total length of each password, including separators. Default is 16.
# -g int: Length of each group, excluding separators. Default is 4.
# -s string: Separator between groups. Default is -.
# -c int: Number of special characters in each password. Default is 4.


# Example Output
# Here’s a sample of what you might see:

# Generated Password: X5$-M4@-Gz1-B7*
```

## Why This Is Better

- aspwgen password generator incorporates modern security practices and flexibility that pwgen lacks. With configurable options for group length, separators, and special characters, it provides a level of customization and security that older tools cannot match.

- Whether you’re a developer, security professional, or just someone who values strong passwords, this tool is designed to meet the highest standards of security and usability.

## Contributing

If you have suggestions or improvements, please feel free to submit a pull request or open an issue. Your contributions are welcome!

```

```
