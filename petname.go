/*
  petname: library for generating human-readable, random names
           for objects (e.g. hostnames, containers, blobs)

  Copyright 2014 Dustin Kirkland <dustin.kirkland@gmail.com>

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

// Package petname is a library for generating human-readable, random
// names for objects (e.g. hostnames, containers, blobs).
package petname

import (
	"math/rand"
	"strings"
	"time"
)

var (
	adverbsByLength = map[int][]string{
		4: {"duly", "only"},
		5: {"badly", "daily", "early", "fully", "jolly", "newly", "oddly", "sadly", "truly"},
		6: {"barely", "deadly", "deeply", "easily", "evenly", "fairly", "firmly", "freely", "gently", "gladly", "hardly", "highly", "hugely", "humbly", "kindly", "lately", "likely", "lively", "loudly", "lovely", "mainly", "merely", "mildly", "mostly", "namely", "nearly", "neatly", "nicely", "openly", "overly", "partly", "poorly", "purely", "rarely", "really", "safely", "simply", "slowly", "solely", "subtly", "surely", "unduly", "vastly", "weekly", "wholly", "widely", "wildly", "yearly"},
		7: {"awfully", "blindly", "briefly", "broadly", "cheaply", "cleanly", "clearly", "closely", "eagerly", "equally", "exactly", "finally", "firstly", "frankly", "ghastly", "greatly", "grossly", "happily", "heavily", "ideally", "jointly", "largely", "legally", "lightly", "locally", "loosely", "luckily", "monthly", "morally", "notably", "plainly", "quickly", "quietly", "rapidly", "readily", "rightly", "roughly", "sharply", "shortly", "tightly", "totally", "usually", "utterly", "vaguely", "wrongly"},
		8: {"actively", "actually", "annually", "arguably", "brightly", "commonly", "directly", "entirely", "forcibly", "formally", "formerly", "friendly", "globally", "heartily", "honestly", "horribly", "manually", "mentally", "multiply", "mutually", "normally", "possibly", "probably", "promptly", "properly", "publicly", "randomly", "recently", "reliably", "remotely", "scarcely", "secondly", "secretly", "sensibly", "severely", "slightly", "smoothly", "socially", "steadily", "strictly", "strongly", "suddenly", "suitably", "terribly", "uniquely", "unlikely", "urgently", "usefully", "verbally", "visually"},
		9: {"adversely", "allegedly", "amazingly", "basically", "blatantly", "carefully", "centrally", "certainly", "correctly", "curiously", "currently", "eminently", "endlessly", "evidently", "extremely", "factually", "generally", "genuinely", "gradually", "hideously", "hopefully", "illegally", "immensely", "initially", "instantly", "intensely", "literally", "logically", "miserably", "naturally", "nominally", "obviously", "painfully", "partially", "perfectly", "precisely", "presently", "primarily", "privately", "radically", "regularly", "routinely", "seemingly", "seriously", "similarly", "sincerely", "specially", "strangely", "trivially", "typically", "uniformly", "violently", "virtually", "willingly"},
	}

	adjectivesByLength = map[int][]string{
		2: {"in", "on", "up"},
		3: {"ace", "apt", "big", "fit", "fun", "hip", "hot", "key", "new", "one", "pet", "pro", "set", "top"},
		4: {"able", "bold", "boss", "busy", "calm", "cool", "cute", "dear", "deep", "easy", "epic", "fair", "fast", "fine", "firm", "fond", "free", "full", "game", "glad", "good", "holy", "huge", "just", "keen", "kind", "live", "main", "many", "meet", "mint", "more", "neat", "next", "nice", "open", "pure", "rare", "real", "rich", "safe", "star", "sure", "tidy", "tops", "true", "vast", "warm", "well", "wise"},
		5: {"above", "alert", "alive", "ample", "awake", "aware", "brave", "brief", "chief", "civil", "clean", "clear", "close", "comic", "crack", "crisp", "eager", "equal", "exact", "fancy", "finer", "first", "fleet", "frank", "fresh", "funky", "funny", "grand", "great", "grown", "handy", "happy", "hardy", "ideal", "joint", "known", "large", "legal", "light", "liked", "loved", "loyal", "lucky", "major", "merry", "model", "moral", "moved", "noble", "noted", "novel", "prime", "proud", "quick", "quiet", "rapid", "ready", "right", "saved", "sharp", "smart", "solid", "sound", "still", "sunny", "super", "sweet", "tight", "tough", "valid", "vital", "vocal", "whole", "wired", "witty"},
		6: {"active", "actual", "amazed", "amused", "better", "bright", "caring", "casual", "causal", "choice", "clever", "cosmic", "cuddly", "daring", "decent", "direct", "divine", "driven", "enough", "exotic", "expert", "famous", "fluent", "flying", "gentle", "giving", "golden", "guided", "helped", "heroic", "honest", "humane", "humble", "immune", "intent", "living", "loving", "master", "mature", "mighty", "modern", "modest", "moving", "mutual", "native", "nearby", "needed", "normal", "picked", "poetic", "polite", "pretty", "prompt", "proper", "proven", "pumped", "rested", "robust", "ruling", "sacred", "saving", "secure", "select", "simple", "smooth", "social", "sought", "square", "stable", "steady", "strong", "subtle", "suited", "superb", "tender", "trusty", "unique", "united", "upward", "usable", "useful", "valued", "viable", "wanted", "worthy"},
		7: {"adapted", "allowed", "amazing", "amusing", "assured", "awaited", "beloved", "blessed", "capable", "capital", "careful", "central", "certain", "charmed", "classic", "closing", "concise", "content", "correct", "crucial", "cunning", "curious", "current", "darling", "dashing", "desired", "devoted", "diverse", "driving", "dynamic", "elegant", "eminent", "enabled", "endless", "engaged", "enjoyed", "eternal", "ethical", "evident", "evolved", "excited", "factual", "fitting", "flowing", "genuine", "glowing", "growing", "guiding", "healthy", "helpful", "helping", "hopeful", "immense", "intense", "knowing", "lasting", "leading", "legible", "lenient", "liberal", "logical", "magical", "massive", "maximum", "musical", "natural", "neutral", "notable", "optimal", "optimum", "organic", "patient", "perfect", "pleased", "popular", "precise", "premium", "present", "primary", "quality", "refined", "regular", "related", "relaxed", "renewed", "settled", "sharing", "shining", "sincere", "skilled", "smiling", "special", "stirred", "summary", "supreme", "topical", "touched", "trusted", "unified", "upright", "wealthy", "welcome", "willing", "winning", "working"},
		8: {"absolute", "accepted", "accurate", "adapting", "adequate", "adjusted", "advanced", "allowing", "apparent", "arriving", "artistic", "assuring", "balanced", "becoming", "bursting", "champion", "charming", "cheerful", "climbing", "coherent", "communal", "complete", "composed", "concrete", "creative", "credible", "deciding", "definite", "delicate", "destined", "discrete", "distinct", "dominant", "electric", "emerging", "enabling", "engaging", "enhanced", "enormous", "equipped", "evolving", "exciting", "faithful", "feasible", "flexible", "generous", "glorious", "gorgeous", "grateful", "harmless", "humorous", "immortal", "improved", "included", "infinite", "informed", "innocent", "inspired", "integral", "internal", "intimate", "inviting", "learning", "literate", "magnetic", "measured", "national", "obliging", "oriented", "outgoing", "peaceful", "pleasant", "pleasing", "polished", "positive", "possible", "powerful", "precious", "prepared", "probable", "profound", "promoted", "rational", "relative", "relaxing", "relevant", "relieved", "renewing", "resolved", "romantic", "selected", "sensible", "settling", "singular", "smashing", "splendid", "sterling", "stirring", "striking", "stunning", "suitable", "sweeping", "talented", "teaching", "thankful", "thorough", "together", "tolerant", "touching", "trusting", "ultimate", "unbiased", "uncommon", "verified", "welcomed", "wondrous", "workable"},
	}

	namesByLength = map[int][]string{
		2: {"ox"},
		3: {"ant", "ape", "asp", "bat", "bee", "boa", "bug", "cat", "cod", "cow", "cub", "doe", "dog", "eel", "eft", "elf", "elk", "emu", "ewe", "fly", "fox", "gar", "gnu", "hen", "hog", "imp", "jay", "kid", "kit", "koi", "lab", "man", "owl", "pig", "pug", "pup", "ram", "rat", "ray", "yak"},
		4: {"bass", "bear", "bird", "boar", "buck", "bull", "calf", "chow", "clam", "colt", "crab", "crow", "dane", "deer", "dodo", "dory", "dove", "drum", "duck", "fawn", "fish", "flea", "foal", "fowl", "frog", "gnat", "goat", "grub", "gull", "hare", "hawk", "ibex", "joey", "kite", "kiwi", "lamb", "lark", "lion", "loon", "lynx", "mako", "mink", "mite", "mole", "moth", "mule", "mutt", "newt", "orca", "oryx", "pika", "pony", "puma", "seal", "shad", "slug", "sole", "stag", "stud", "swan", "tahr", "teal", "tick", "toad", "tuna", "wasp", "wolf", "worm", "wren", "yeti"},
		5: {"adder", "akita", "alien", "aphid", "bison", "boxer", "bream", "bunny", "burro", "camel", "chimp", "civet", "cobra", "coral", "corgi", "crane", "dingo", "drake", "eagle", "egret", "filly", "finch", "gator", "gecko", "ghost", "ghoul", "goose", "guppy", "heron", "hippo", "horse", "hound", "husky", "hyena", "koala", "krill", "leech", "lemur", "liger", "llama", "louse", "macaw", "midge", "molly", "moose", "moray", "mouse", "panda", "perch", "prawn", "quail", "racer", "raven", "rhino", "robin", "satyr", "shark", "sheep", "shrew", "skink", "skunk", "sloth", "snail", "snake", "snipe", "squid", "stork", "swift", "swine", "tapir", "tetra", "tiger", "troll", "trout", "viper", "wahoo", "whale", "zebra"},
		6: {"alpaca", "amoeba", "baboon", "badger", "beagle", "bedbug", "beetle", "bengal", "bobcat", "caiman", "cattle", "cicada", "collie", "condor", "cougar", "coyote", "dassie", "donkey", "dragon", "earwig", "falcon", "feline", "ferret", "gannet", "gibbon", "glider", "goblin", "gopher", "grouse", "guinea", "hermit", "hornet", "iguana", "impala", "insect", "jackal", "jaguar", "jennet", "kitten", "kodiak", "lizard", "locust", "maggot", "magpie", "mammal", "mantis", "marlin", "marmot", "marten", "martin", "mayfly", "minnow", "monkey", "mullet", "muskox", "ocelot", "oriole", "osprey", "oyster", "parrot", "pigeon", "piglet", "poodle", "possum", "python", "quagga", "rabbit", "raptor", "rodent", "roughy", "salmon", "sawfly", "serval", "shiner", "shrimp", "spider", "sponge", "tarpon", "thrush", "tomcat", "toucan", "turkey", "turtle", "urchin", "vervet", "walrus", "weasel", "weevil", "wombat"},
		7: {"anchovy", "anemone", "bluejay", "buffalo", "bulldog", "buzzard", "caribou", "catfish", "chamois", "cheetah", "chicken", "chigger", "cowbird", "crappie", "crawdad", "cricket", "dogfish", "dolphin", "firefly", "garfish", "gazelle", "gelding", "giraffe", "gobbler", "gorilla", "goshawk", "grackle", "griffon", "grizzly", "grouper", "haddock", "hagfish", "halibut", "hamster", "herring", "jackass", "javelin", "jawfish", "jaybird", "katydid", "ladybug", "lamprey", "lemming", "leopard", "lioness", "lobster", "macaque", "mallard", "mammoth", "manatee", "mastiff", "meerkat", "mollusk", "monarch", "mongrel", "monitor", "monster", "mudfish", "muskrat", "mustang", "narwhal", "oarfish", "octopus", "opossum", "ostrich", "panther", "peacock", "pegasus", "pelican", "penguin", "phoenix", "piranha", "polecat", "primate", "quetzal", "raccoon", "rattler", "redbird", "redfish", "reptile", "rooster", "sawfish", "sculpin", "seagull", "skylark", "snapper", "spaniel", "sparrow", "sunbeam", "sunbird", "sunfish", "tadpole", "termite", "terrier", "unicorn", "vulture", "wallaby", "walleye", "warthog", "whippet", "wildcat"},
		8: {"aardvark", "airedale", "albacore", "anteater", "antelope", "arachnid", "barnacle", "basilisk", "blowfish", "bluebird", "bluegill", "bonefish", "bullfrog", "cardinal", "chipmunk", "cockatoo", "crayfish", "dinosaur", "doberman", "duckling", "elephant", "escargot", "flamingo", "flounder", "foxhound", "glowworm", "goldfish", "grubworm", "hedgehog", "honeybee", "hookworm", "humpback", "kangaroo", "killdeer", "kingfish", "labrador", "lacewing", "ladybird", "lionfish", "longhorn", "mackerel", "malamute", "marmoset", "mastodon", "moccasin", "mongoose", "monkfish", "mosquito", "pangolin", "parakeet", "pheasant", "pipefish", "platypus", "polliwog", "porpoise", "reindeer", "ringtail", "sailfish", "scorpion", "seahorse", "seasnail", "sheepdog", "shepherd", "silkworm", "squirrel", "stallion", "starfish", "starling", "stingray", "stinkbug", "sturgeon", "terrapin", "titmouse", "tortoise", "treefrog", "werewolf", "woodcock"},
	}
)

// Call this function once before using any other to get real random results
func NonDeterministicMode() {
	rand.Seed(time.Now().UnixNano())
}

// Adverb returns a random adverb from a list of petname adverbs.
func Adverb(maxLength int) string {
	if maxLength < 4 {
		return ""
	}

	if maxLength > 9 {
		maxLength = 9
	}

	length := rand.Intn(maxLength+1-4) + 4
	adverbs := adverbsByLength[length]

	return adverbs[rand.Intn(len(adverbs))]
}

// Adjective returns a random adjective from a list of petname adjectives.
func Adjective(maxLength int) string {
	if maxLength < 2 {
		return ""
	}

	if maxLength > 8 {
		maxLength = 8
	}

	length := rand.Intn(maxLength+1-2) + 2
	adjectives := adjectivesByLength[length]

	return adjectives[rand.Intn(len(adjectives))]
}

// Name returns a random name from a list of petname names.
func Name(maxLength int) string {
	if maxLength < 2 {
		return ""
	}

	if maxLength > 8 {
		maxLength = 8
	}

	length := rand.Intn(maxLength+1-2) + 2
	names := namesByLength[length]

	return names[rand.Intn(len(names))]
}

// Generate generates and returns a random pet name.
// It takes three parameters:  the number of words in the name, the max length of each word, and a separator token.
// If a single word is requested, simply a Name() is returned.
// If two words are requested, a Adjective() and a Name() are returned.
// If three or more words are requested, a variable number of Adverb() and a Adjective and a Name() is returned.
// If a word less than max length is not available, returns an empty string
// The separator can be any character, string, or the empty string.
func Generate(words, maxLength int, separator string) string {
	if words == 0 {
		return ""
	} else if words == 1 {
		return Name(maxLength)
	} else if words == 2 {
		return Adjective(maxLength) + separator + Name(maxLength)
	}
	var petname []string
	for i := 0; i < words-2; i++ {
		petname = append(petname, Adverb(maxLength))
	}
	petname = append(petname, Adjective(maxLength), Name(maxLength))
	return strings.Join(petname, separator)
}
