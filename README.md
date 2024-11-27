Random Hostname Generator

This Go program generates random hostnames by combining an adjective, a color, and an animal. It's a fun and unique way to create memorable names for your devices, projects, or anything else you need!
Features

    Generate single or multiple random hostnames.
    Customize the server port and API endpoint via command-line arguments.
    Expose a REST API to return hostnames in JSON format.

Installation

    Clone the repository:

git clone https://github.com/yourusername/random-hostname-generator.git
cd random-hostname-generator

Build the executable:

    go build -o random-hostname-generator

Usage
1. Command-Line Hostname Generation
Generate a Single Hostname:

Run the program without any arguments to generate one random hostname:

./random-hostname-generator

Generate Multiple Hostnames:

Use the -count flag to specify the number of hostnames to generate:

./random-hostname-generator -count=5

Example Output:

Breezy-Silver-Fox
Ancient-Golden-Cheetah
Sparkly-Blue-Rabbit
Elegant-White-Tiger
Dizzy-Red-Owl

2. Run as a Web Service

You can start the program as a web server to expose a REST API for hostname generation.
Default Behavior:

Start the server on port 8080 with the /hostname endpoint:

./random-hostname-generator

Custom Port and Endpoint:

Specify a custom port and API endpoint using the -port and -endpoint flags:

./random-hostname-generator -port=9090 -endpoint=/generate-hostname

Example Output:

Starting server on port 9090 with endpoint /generate-hostname

3. Access the API
Single Hostname:

Use curl to request a single hostname:

curl http://localhost:8080/hostname

Example Response:

{
    "hostnames": ["Elegant-Pearl-Fox"]
}

Multiple Hostnames:

Add the count query parameter to request multiple hostnames:

curl "http://localhost:8080/hostname?count=3"

Example Response:

{
    "hostnames": [
        "Brave-Golden-Tiger",
        "Cheerful-Blue-Kitten",
        "Magnificent-White-Bird"
    ]
}

How it Works

    The program combines an adjective, a color, and an animal to generate a unique hostname.
    The slices (Adjectives, Colors, and Animals) contain predefined lists of words.
    The capitalize function ensures each word starts with a capital letter.
    The shuffle function randomizes the slices to maximize uniqueness.
    An HTTP server exposes an API for hostname generation.

Examples

Here are some examples of hostnames that this program might generate:

    Sparkly-Golden-Narwhal
    Dizzy-Fuchsia-Monkey
    Fluffy-Rainbow-Kitten
    Ancient-Bronze-Dinosaur
    Bumpy-Ultraviolet-Penguin

Contributing

We welcome contributions to this project! Here are some ways to help:

    Add more adjectives, colors, or animals to the respective slices.
    Improve the code structure or add features (e.g., caching, custom word lists).
    Enhance the REST API (e.g., authentication, advanced query parameters).

License

This project is licensed under the MIT License - see the LICENSE file for details.
