
# Discord Bot in Golang 🤖

A fun and simple Discord bot written in Go.

---
<p align="center">
  <img width="600" alt="Command Output" src="https://github.com/user-attachments/assets/481c407c-4052-4a94-b426-d7ea78f7bbf5" />
</p>

---

## 🔧 Setup

1. **Clone the repository**:

```bash
git clone https://github.com/Prtik12/Discord_Bot_Golang.git
cd Discord_Bot_Golang
```

2. **Install dependencies**:

```bash
go mod tidy
```

3. **Create a `config.json` file** in the root directory:

```json
{
  "Token": "YOUR_BOT_TOKEN_HERE",
  "BotPrefix": "!"
}
```

> **Note:** Add `config.json` to your `.gitignore` to keep your bot token private.

4. **Run the bot**:

```bash
go run main.go
```

---

## 💻 Commands Available

- `/ping` – Responds with Pong!
- `/hello` – Sends a greeting
- `/flip` – Flips a coin
- `/roll` – Rolls a dice
- `/joke` – Tells a random programming joke
- `/help` – Lists all available commands

---


### Example Command Output:
<p align="center">
  <img width="600" alt="Help Command Output" src="https://github.com/user-attachments/assets/332666b7-fdfc-4f96-90ab-bbb7812636bb" />
</p>

---

## 📂 Project Structure

```
📦 Discord_Bot_Golang
├── bot
│   └── bot.go
├── config
│   └── config.go
├── config.json
├── main.go
└── README.md
```

---

## 🛠️ Built With

- [Go (Golang)](https://golang.org/)
- [DiscordGo](https://github.com/bwmarrin/discordgo)

---
```

Let me know if you'd like to add a logo, some badges, or a demo video section too!
