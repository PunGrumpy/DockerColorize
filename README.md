<div align="center">

# Dockercolorize 🐳🌈

Enhancing Docker output with vibrant colors. Dockercolorize is a command line tool designed to add colors to Docker outputs, making them more visually organized and readable.

</div>

<div align="center">

![GitHub Actions](https://github.com/PunGrumpy/dockercolorize/actions/workflows/go.yml/badge.svg?branch=main)
![Release](https://img.shields.io/github/v/release/PunGrumpy/dockercolorize)
![License](https://img.shields.io/github/license/PunGrumpy/dockercolorize)
![Downloads](https://img.shields.io/github/downloads/PunGrumpy/dockercolorize/total)

</div>

## 👨‍💻 Installation

### 🍺 Homebrew

```bash
# Tap the formula repository
brew tap PunGrumpy/formulas

# Check tap it's working
brew tap-info PunGrumpy/formulas
brew search PunGrumpy/formulas

# Install the formula
brew install dockercolorize
```

#### 🪄 Aliases

For a smoother user experience, use these [aliases](.github/bash/aliases.sh) to replace the original Docker commands.

## ⚙️ Configuration

Locate the configuration file in `~/.config/dockercolorize/config.json` and edit the color scheme.

```json
// Default color scheme
{
  "color": {
    "reset": "\u001b[0m",
    "black": "\u001b[0;30m",
    "darkGray": "\u001b[1;30m",
    "red": "\u001b[0;31m",
    "lightRed": "\u001b[1;31m",
    "green": "\u001b[0;32m",
    "lightGreen": "\u001b[1;32m",
    "brown": "\u001b[0;33m",
    "yellow": "\u001b[1;33m",
    "blue": "\u001b[0;34m",
    "lightBlue": "\u001b[1;34m",
    "purple": "\u001b[0;35m",
    "lightPurple": "\u001b[1;35m",
    "cyan": "\u001b[0;36m",
    "lightCyan": "\u001b[1;36m",
    "lightGray": "\u001b[0;37m",
    "white": "\u001b[1;37m"
  }
}
```

## 📚 Usage

#### 💡 docker images

```bash
di # alias
```

```bash
docker images [--format] | docker-color-output
```

![docker images](https://user-images.githubusercontent.com/5787193/93581956-7ae7f580-f9aa-11ea-8f81-d6922e1ca892.png)

#### 💡 docker ps

```bash
dps # alias
```

```bash
docker ps [-a] [--format] | docker-color-output
```

![docker ps](https://user-images.githubusercontent.com/5787193/93581144-69521e00-f9a9-11ea-86bb-c23d7879c689.png)

#### 💡 docker compose ps

⚠️ The latest version works with docker-compose `2.x`.

```bash
dcps # alias
```

```bash
docker compose ps | docker-color-output
```

![docker compose ps](https://user-images.githubusercontent.com/5787193/93630916-7267dd00-f9f3-11ea-9521-e69152fa86f1.png)

#### 💡 docker stats

⚠️ For the best experience, use the `--no-stream` flag. The `--no-stream` flag is not supported in Docker Compose.

```bash
dstats # alias
```

```bash
docker stats --no-stream | docker-color-output
```

#### 💡 docker history

⚠️ It's cannot be used with aliases because it's a subcommand.

```bash
docker history [container] [--format] | docker-color-output
```
