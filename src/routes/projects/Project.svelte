<script>
  import Projectpost from "./Projectpost.svelte";

  let projectPosts = [];

  let projectPostCreator = (name, badges, description, url) => {
    return {
      name,
      badges,
      description,
      url,
    };
  };
  projectPosts.push(
    projectPostCreator(
      "Personal Website",
      ["javascript", "svelte", "html", "css", "browser"],
      "Personal website to publish blogs and projects.",
      "https://231tr0n.github.io"
    ),
    projectPostCreator(
      "Proxy Server in Nodejs",
      ["javascript", "nodejs", "forward-proxy"],
      "Created a forward proxy server in nodejs.",
      "https://github.com/231tr0n/nodejs-proxy"
    ),
    projectPostCreator(
      "A simple password vault",
      ["go", "password-vault"],
      "A dead simple password vault.",
      "https://github.com/231tr0n/vault"
    ),
    projectPostCreator(
      "Cryptography based vulnerabilities",
      ["nodejs", "cryptography"],
      "Testing various methods of encryption for passing messages.",
      "https://github.com/231tr0n/crypto-project"
    ),
    projectPostCreator(
      "Discord Bot",
      ["nodejs", "discord-bot"],
      "Created a discord bot for recording updates and pollling.",
      "https://github.com/231tr0n/pentest-bot"
    ),
    projectPostCreator(
      "Dotfiles",
      ["config", "developer-setup"],
      "All configuration files of my system.",
      "https://github.com/231tr0n/config"
    ),
    projectPostCreator(
      "Custom go router",
      ["go", "http-router"],
      "A custom router with middleware option in golang.",
      "https://github.com/231tr0n/config"
    ),
    projectPostCreator(
      "Q learning for snake game",
      ["python", "snake game", "reinforcement-q-learning"],
      "Created a snake game where the snake trains to play the game by itself.",
      "https://github.com/231tr0n/reinforcement-learning-for-snake-game"
    ),
    projectPostCreator(
      "Q learning for ping pong game",
      ["python", "ping-pong game", "reinforcement-q-learning"],
      "Created a ping-pong game where the ball trains to play the game by itself.",
      "https://github.com/231tr0n/reinforcement-learning-for-pingpong"
    ),
    projectPostCreator(
      "ISL hand sign recognizing chatbot",
      [
        "go",
        "python",
        "html",
        "css",
        "javascript",
        "mediapipe",
        "tensorflow",
        "neural-networks",
      ],
      "Created a chatbot which recognizes ISL hand signs.",
      "https://github.com/231tr0n/fyp"
    ),
    projectPostCreator(
      "Hamiltonian path finder",
      ["python", "graph", "hamiltonian-path"],
      "Script to find hamiltonian paths in a graph.",
      "https://github.com/231tr0n/hamiltonian-path-in-a-graph"
    )
  );
</script>

<h1>Projects</h1>

{#each projectPosts as projectPost}
  <Projectpost {...projectPost} />
{/each}
