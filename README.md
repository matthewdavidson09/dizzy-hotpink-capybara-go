# Random Hostname Generator

This Go program generates random hostnames by combining an adjective, a color, and an animal. It's a fun way to create unique and memorable names for your devices or projects.

## Usage

1.  **Compile the code:**
    ```bash
    go build
    ```

2.  **Run the executable:**
    ```bash
    ./random-hostname-generator
    ```

    This will generate one random hostname.

3.  **Generate multiple hostnames:**
    ```bash
    ./random-hostname-generator -count=5
    ```

    This will generate 5 random hostnames. You can change the `count` flag to generate any number of hostnames.

## Examples

Here are some examples of hostnames that this program might generate:

*   Sparkly-Golden-Narwhal
*   Dizzy-Fuchsia-Monkey
*   Fluffy-Rainbow-Kitten
*   Ancient-Bronze-Dinosaur
*   Bumpy-Ultraviolet-Penguin

## How it Works

The program uses three constant slices of strings: `Adjectives`, `Colors`, and `Animals`. It randomly selects one item from each slice and combines them to create a hostname.

The `capitalize` function ensures that each word in the hostname starts with a capital letter.

The `shuffle` function randomizes the order of the items in the slices, ensuring that the generated hostnames are truly random.

## Contributing

Feel free to contribute to this project by adding more adjectives, colors, or animals to the respective slices. You can also improve the code or add new features.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
