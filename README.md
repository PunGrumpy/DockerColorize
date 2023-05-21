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

⚠️ Currently, Dockercolorize is not available on `homebrew` or `apt-get`. Please proceed with manual installation.

#### 🪄 Aliases

For a smoother user experience, use these [aliases](bash/aliases.sh).

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
