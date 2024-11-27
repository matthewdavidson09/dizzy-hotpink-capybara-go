package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Adjectives = []string{
	"attractive", "bald", "beautiful", "rare", "clean", "dazzling", "lucky", "elegant",
	"fancy", "fit", "fantastic", "glamorous", "gorgeous", "handsome", "long", "magnificent",
	"muscular", "plain", "able", "quaint", "scruffy", "innocent", "short", "skinny",
	"acrobatic", "tall", "proper", "alert", "lone", "agreeable", "ambitious", "brave",
	"calm", "delightful", "eager", "faithful", "gentle", "happy", "jolly", "kind",
	"lively", "nice", "obedient", "polite", "proud", "silly", "thankful", "winning",
	"witty", "wonderful", "zealous", "expert", "amateur", "clumsy", "amusing", "vast",
	"fierce", "real", "helpful", "itchy", "atomic", "basic", "mysterious", "blurry",
	"perfect", "best", "powerful", "interesting", "decent", "wild", "jovial", "genuine",
	"broad", "brisk", "brilliant", "curved", "deep", "flat", "high", "hollow",
	"low", "narrow", "refined", "round", "shallow", "square", "steep", "straight",
	"wide", "big", "colossal", "clever", "gigantic", "great", "huge", "immense",
	"large", "little", "mammoth", "massive", "micro", "mini", "petite", "puny",
	"scrawny", "small", "polished", "teeny", "tiny", "crazy", "dancing", "custom",
	"faint", "harsh", "formal", "howling", "loud", "melodic", "noisy", "upbeat",
	"quiet", "dandy", "raspy", "rhythmic", "daring", "zany", "digital", "dizzy",
	"exotic", "fun", "furry", "hidden", "ancient", "brief", "early", "fast",
	"future", "late", "modern", "old", "prehistoric", "zesty", "rapid", "slow",
	"swift", "young", "acidic", "bitter", "cool", "creamy", "keen", "tricky",
	"fresh", "special", "unique", "hot", "magic", "main", "nutty", "mythical",
	"ripe", "wobbly", "salty", "savory", "sour", "spicy", "bright", "stale",
	"sweet", "tangy", "tart", "rich", "rural", "urban", "breezy", "bumpy",
	"chilly", "cold", "cuddly", "damaged", "damp", "restless", "dry", "flaky",
	"fluffy", "merry", "icy", "shiny", "melted", "joyous", "rough", "shaggy",
	"sharp", "radiant", "sticky", "strong", "soft", "uneven", "warm", "feisty",
	"cheery", "energetic", "abundant", "macho", "glorious", "mean", "quick", "precise",
	"stable", "spare", "sunny", "trendy", "shambolic", "striped", "boxy", "generous",
	"tame", "joyful", "festive", "bubbly", "soaring", "orbiting", "sparkly", "smooth",
	"docile", "original", "electric", "funny", "passive", "active", "cheesy", "blunt",
	"dapper", "bent", "curly", "oblong", "sneaky", "overt", "careful", "jumpy",
	"bouncy", "recumbent", "cheerful", "droll", "odd", "suave", "sleepy",
}

// Colors is a constant slice of color-related words.
var Colors = []string{
	"white", "pearl", "alabaster", "snowy", "ivory", "cream", "cotton", "chiffon",
	"lace", "coconut", "linen", "bone", "daisy", "powder", "frost", "porcelain",
	"parchment", "velvet", "tan", "beige", "macaroon", "hazel", "felt", "metal",
	"sand", "sepia", "latte", "vinyl", "glass", "hazelnut", "canvas", "wool",
	"yellow", "golden", "daffodil", "flaxen", "butter", "lemon", "mustard", "tartan",
	"blue", "cloth", "fiery", "banana", "plastic", "dijon", "honey", "blonde",
	"pineapple", "orange", "tangerine", "marigold", "cider", "rusty", "ginger", "tiger",
	"bronze", "fuzzy", "opaque", "clay", "carrot", "corduroy", "ceramic", "marmalade",
	"amber", "sandstone", "concrete", "red", "cherry", "hemp", "merlot", "garnet",
	"crimson", "ruby", "scarlet", "burlap", "brick", "bamboo", "mahogany", "blood",
	"sangria", "berry", "currant", "blush", "candy", "lipstick", "pink", "rose",
	"fuchsia", "punch", "watermelon", "rouge", "coral", "peach", "strawberry", "rosewood",
	"lemonade", "taffy", "bubblegum", "crepe", "hotpink", "purple", "mauve", "violet",
	"boysenberry", "lavender", "plum", "magenta", "lilac", "grape", "eggplant", "eggshell",
	"iris", "heather", "amethyst", "raisin", "orchid", "mulberry", "carbon", "slate",
	"sky", "navy", "indigo", "cobalt", "cedar", "ocean", "azure", "cerulean",
	"spruce", "stone", "aegean", "denim", "admiral", "sapphire", "arctic", "green",
	"chartreuse", "juniper", "sage", "lime", "fern", "olive", "emerald", "pear",
	"mossy", "shamrock", "seafoam", "pine", "mint", "seaweed", "pickle", "pistachio",
	"basil", "brown", "coffee", "chrome", "peanut", "carob", "hickory", "wooden",
	"pecan", "walnut", "caramel", "gingerbread", "syrup", "chocolate", "tortilla", "umber",
	"tawny", "brunette", "cinnamon", "glossy", "teal", "grey", "shadow", "graphite",
	"iron", "pewter", "cloud", "silver", "smoke", "gauze", "ash", "foggy",
	"flint", "charcoal", "pebble", "lead", "tin", "fossilized", "black", "ebony",
	"midnight", "inky", "oily", "satin", "onyx", "nylon", "fleece", "sable",
	"jetblack", "coal", "mocha", "obsidian", "jade", "cyan", "leather", "maroon",
	"carmine", "aqua", "chambray", "holographic", "laurel", "licorice", "khaki", "goldenrod",
	"malachite", "mandarin", "mango", "taupe", "aquamarine", "turquoise", "vermilion", "saffron",
	"cinnabar", "myrtle", "neon", "burgundy", "tangelo", "topaz", "wintergreen", "viridian",
	"vanilla", "paisley", "raspberry", "tweed", "pastel", "opal", "menthol", "champagne",
	"gunmetal", "infrared", "ultraviolet", "rainbow", "mercurial", "clear", "misty", "steel",
	"zinc", "citron", "cornflower", "lava", "quartz", "honeysuckle", "chili",
}

// Animals is a constant slice of animal names.
var Animals = []string{
	"alligator", "bee", "bird", "camel", "cat", "cheetah", "chicken", "cow", "dog",
	"corgi", "eagle", "elephant", "fish", "fox", "toad", "giraffe", "hippo",
	"kangaroo", "kitten", "lobster", "monkey", "octopus", "pig", "puppy",
	"rabbit", "rat", "scorpion", "seal", "sheep", "snail", "spider", "tiger",
	"turtle", "newt", "frog", "albatross", "crow", "falcon", "sparrow", "hawk",
	"jay", "mockingbird", "owl", "parrot", "peacock", "penguin", "pheasant",
	"piranha", "raven", "robin", "rooster", "stork", "swallow", "turkey", "vulture",
	"woodpecker", "butterfly", "carp", "crab", "eel", "goldfish", "shark", "shrimp",
	"trout", "ant", "aphid", "beetle", "caterpillar", "dragonfly", "fly",
	"grasshopper", "ladybug", "millipede", "moth", "wasp", "anteater",
	"antelope", "armadillo", "bat", "beaver", "bull", "chimpanzee",
	"dachshund", "deer", "dolphin", "moose", "gazelle", "goat", "hamster",
	"hare", "hedgehog", "horse", "hyena", "lion", "llama", "mammoth",
	"mongoose", "mouse", "otter", "panda", "platypus", "pony",
	"porcupine", "raccoon", "reindeer", "rhino", "sloth", "squirrel", "weasel",
	"wolf", "zebra", "boa", "chameleon", "gecko", "iguana", "python",
	"salamander", "whale", "lemur", "rook", "koala", "donkey", "ferret",
	"tardigrade", "orca", "narwhal", "worm", "hornet", "viper", "jaguar",
	"panther", "bobcat", "leopard", "osprey", "cougar", "dalmatian",
	"terrier", "duck", "chipmunk", "loris", "poodle", "orangutan",
	"meerkat", "huskie", "bison", "chinchilla", "coyote", "crane",
	"dinosaur", "griffin", "yeti", "troll", "seahorse", "walrus",
	"boar", "alpaca", "manatee", "condor", "cobra", "locust", "mandrill",
	"oyster", "quail", "sardine", "ram", "wallaby", "buffalo", "goblin",
	"tuna", "mustang",
}

// Capitalizes the first letter of a string
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + word[1:]
}

// Shuffles a slice
func shuffle(slice []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Generates a random hostname by picking one random item from each slice
func generateHostname() string {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	// Pick random elements from each slice
	adjective := capitalize(Adjectives[rand.Intn(len(Adjectives))])
	color := capitalize(Colors[rand.Intn(len(Colors))])
	animal := capitalize(Animals[rand.Intn(len(Animals))])

	// Combine into a hostname
	return fmt.Sprintf("%s-%s-%s", adjective, color, animal)
}

func main() {
	// Define a command-line argument for the number of names to generate
	count := flag.Int("count", 1, "Number of hostnames to generate")
	flag.Parse()

	// Validate the argument
	if *count <= 0 {
		fmt.Println("Please specify a positive number for the count.")
		os.Exit(1)
	}

	// Shuffle the arrays (optional, for randomness)
	shuffle(Adjectives)
	shuffle(Colors)
	shuffle(Animals)

	// Generate and print hostnames
	for i := 0; i < *count; i++ {
		fmt.Println(generateHostname())
	}
}
