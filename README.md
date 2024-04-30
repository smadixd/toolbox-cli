# Toolbox CLI

Toolbox CLI is a simple command-line tool that allows you to show ping and disk usage. It follows the Cobra CLI architecture, making it easy to add new commands and extend functionality.

## Installation

To install Toolbox CLI, you need to have Go installed on your machine. Then, run the following command:

go get -u github.com/your-username/toolbox-cli

# Usage

Once you have installed Toolbox CLI, you can use it by running the following command:

toolbox-cli [command]

## Commands

- `ping`: Show ping to a specified host.
- `disk-usage`: Show disk usage of a specified directory.

For example, to ping google.com, you can run:

toolbox-cli ping google.com

To check the disk usage of the current directory, you can run:

toolbox-cli disk-usage .

## Contributing

If you would like to contribute to Toolbox CLI, follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch to your forked repository.
5. Open a pull request to the main repository.

## License

Toolbox CLI is licensed under the MIT License. See the [LICENSE](https://github.com/smadixd/toolbox-cli/blob/main/LICENSE) file for more information.
