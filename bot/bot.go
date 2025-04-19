package bot

import (
	"fmt"
	"math/rand"

	"github.com/Prtik12/Discord_Bot_Golang/config"
	"github.com/bwmarrin/discordgo"
)

var (
	BotId    string
	goBot    *discordgo.Session
	commands = []discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Pong!",
		},
		{
			Name:        "hello",
			Description: "Say hello to the bot",
		},
		{
			Name:        "flip",
			Description: "Flip a coin",
		},
		{
			Name:        "roll",
			Description: "Roll a dice",
		},
		{
			Name:        "joke",
			Description: "Tell a random joke",
		},
		{
			Name:        "cat",
			Description: "Get a random cat fact",
		},
		{
			Name:        "dog",
			Description: "Get a random dog fact",
		},
		{
			Name:        "quote",
			Description: "Get a motivational quote",
		},
		{
			Name:        "random",
			Description: "Get a random number between 1 and 100",
		},
		{
			Name:        "help",
			Description: "List all available commands",
		},
	}
)

func Start() {
	var err error

	// Create a new Discord session using the provided bot token.
	goBot, err = discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Debug print to confirm goBot initialization
	if goBot == nil {
		fmt.Println("Error: goBot session is nil.")
		return
	}
	fmt.Println("goBot initialized successfully.")

	// Enable message content intent
	goBot.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsDirectMessages |
		discordgo.IntentsMessageContent

	// Register event handler for Ready event to ensure the bot is fully connected
	goBot.AddHandler(readyHandler)

	// Register interaction handler
	goBot.AddHandler(interactionHandler)

	// Open the connection to Discord
	err = goBot.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running!")
}

// Ready event handler
func readyHandler(s *discordgo.Session, r *discordgo.Ready) {
	// Get bot user information to set the BotId
	u, err := s.User("@me")
	if err != nil {
		fmt.Println("Error retrieving account details:", err)
		return
	}

	// Ensure the BotId is valid
	if u == nil {
		fmt.Println("Error: Bot user is nil.")
		return
	}

	BotId = u.ID

	// Register slash commands with Discord after the bot is fully connected
	err = registerSlashCommands()
	if err != nil {
		fmt.Println("Error registering commands:", err)
		return
	}
}

// Register slash commands with Discord (global commands for all servers).
func registerSlashCommands() error {
	// Check if goBot is nil before proceeding
	if goBot == nil {
		return fmt.Errorf("goBot is nil")
	}

	// Register commands with Discord.
	for _, command := range commands {
		_, err := goBot.ApplicationCommandCreate(goBot.State.User.ID, "", &command)
		if err != nil {
			return fmt.Errorf("cannot create command %s: %v", command.Name, err)
		}
	}

	// Debug print to confirm commands registration is complete
	fmt.Println("Commands registered successfully.")
	return nil
}

// Handle interactions (slash commands)
func interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Ensure this is a valid interaction
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	// Switch based on the command name
	switch i.ApplicationCommandData().Name {
	case "ping":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "🏓 Pong!",
			},
		})
	case "hello":
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "👋 Hey there!",
			},
		})
	case "flip":
		result := "Heads"
		if rand.Intn(2) == 0 {
			result = "Tails"
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("You flipped: %s", result),
			},
		})
	case "roll":
		roll := rand.Intn(6) + 1
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("You rolled: %d", roll),
			},
		})
	case "joke":
		joke := "Why do programmers prefer dark mode? Because light attracts bugs! 🐛"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: joke,
			},
		})
	case "cat":
		catFact := "Did you know? Cats sleep for 70% of their lives!"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: catFact,
			},
		})
	case "dog":
		dogFact := "Dogs have a sense of time and can be trained to understand it!"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: dogFact,
			},
		})
	case "quote":
		quote := "The only way to do great work is to love what you do. – Steve Jobs"
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: quote,
			},
		})
	case "random":
		randomNumber := rand.Intn(100) + 1
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("Your random number is: %d", randomNumber),
			},
		})
	case "help":
		helpMessage := "Here are all the available commands:\n"
		// Dynamically generate help message
		for _, command := range commands {
			helpMessage += fmt.Sprintf("`/%s`: %s\n", command.Name, command.Description)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: helpMessage,
			},
		})
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "❓ Unknown command.",
			},
		})
	}
}
