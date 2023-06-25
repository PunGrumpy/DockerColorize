#! /usr/bin/env bash

PROJECT_NAME="Dockercolorize"
PROJECT_VERSION="2.6.0"
PROJECT_URL="https://github.com/PunGrumpy/dockercolorize"
FILE_NAME="dockercolorize"
OS_TYPE=$(uname -s)

RED=$(tput setaf 1)
GREEN=$(tput setaf 2)
YELLOW=$(tput setaf 3)
BLUE=$(tput setaf 4)
PURPLE=$(tput setaf 5)
CYAN=$(tput setaf 6)
BOLD=$(tput bold)
BEEP=$(tput blink)
RESET=$(tput sgr0)

########## Banner ##########
banner() {
    local text="${1:?Missing text}"
    echo -en "${CYAN}
██████╗  ██████╗  ██████╗██╗  ██╗███████╗██████╗  ██████╗ ██████╗ ██╗      ██████╗ ██████╗ ██╗███████╗███████╗
██╔══██╗██╔═══██╗██╔════╝██║ ██╔╝██╔════╝██╔══██╗██╔════╝██╔═══██╗██║     ██╔═══██╗██╔══██╗██║╚══███╔╝██╔════╝
██║  ██║██║   ██║██║     █████╔╝ █████╗  ██████╔╝██║     ██║   ██║██║     ██║   ██║██████╔╝██║  ███╔╝ █████╗  
██║  ██║██║   ██║██║     ██╔═██╗ ██╔══╝  ██╔══██╗██║     ██║   ██║██║     ██║   ██║██╔══██╗██║ ███╔╝  ██╔══╝  
██████╔╝╚██████╔╝╚██████╗██║  ██╗███████╗██║  ██║╚██████╗╚██████╔╝███████╗╚██████╔╝██║  ██║██║███████╗███████╗
╚═════╝  ╚═════╝  ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝╚═╝╚══════╝╚══════╝
Project: ${PROJECT_NAME} | Version: ${PROJECT_VERSION} | URL: ${PROJECT_URL}${RESET}\n\n"
    echo -e "${RED}[${RESET} ${BOLD}${YELLOW}${text}${RESET} ${RED}]${RESET}\n"
}

##########################################################

clear

##########################################################

########## Welcome ##########
banner "👋 Welcome to ${PROJECT_NAME}"

echo -en "${BLUE}
       .
      \":\"
    ___:____     |\"\\/\"|
  ,'        \`.    \\  /
  |  O        \\\___/  |
~^~^~^~^~^~^~^~^~^~^~^~^~${RESET}\n"

sleep 2
clear

##########################################################

########## Check not root ##########
banner "👮 Checking not root"

if [[ $EUID -eq 0 ]]; then
    echo -e "${RED}This script must not be run as root 👮${RESET}" 1>&2 && exit 1
    exit 1
else 
    echo -e "${GREEN}Not root 👮${RESET}"
fi

sleep 2
clear

########## Check OS is supported ##########
banner "🐧 Checking OS"

if [[ "$OS_TYPE" == "Darwin" ]]; then
    echo -e "${GREEN}OS is supported 🍎${RESET}"
elif [[ "$OS_TYPE" == "Linux" ]]; then
    echo -e "${GREEN}OS is supported 🐧${RESET}"
else
    echo -e "${RED}OS is not supported 🐀${RESET}" 1>&2 && exit 1
    exit 1
fi

sleep 2
clear

########## Check docker is installed ##########
banner "🐳 Checking Docker"

if ! [ -x "$(command -v docker)" ]; then
    echo -e "${RED}Docker is not installed 🐀${RESET}" 1>&2 && exit 1
    exit 1
else 
    echo -e "${GREEN}Docker is installed 🐳${RESET}"
fi

sleep 2
clear

########## Check docker-compose is installed ##########
banner "🐳 Checking Docker Compose"

if ! [ -x "$(command -v docker-compose)" ]; then
    echo -e "${RED}Docker Compose is not installed 🐀${RESET}" 1>&2 && exit 1
    exit 1
else 
    echo -e "${GREEN}Docker Compose is installed 🐳${RESET}"
fi

sleep 2
clear

########## Check internet connection ##########
banner "🌐 Checking Internet Connection"

if ping -q -c 1 -W 1 google.com >/dev/null; then
    echo -e "${GREEN}Internet connection is up 🟢${RESET}"
else
    echo -e "${RED}Internet connection is down 🔴${RESET}" 1>&2 && exit 1
    exit 1
fi

sleep 2
clear

##########################################################

########## Setup go ##########
banner "🐭 Setting up go"

if ! [ -x "$(command -v go)" ]; then
    echo -e "${RED}go is not installed 🐀${RESET}" 1>&2
    read -p "Do you want to install go? (y/n) " -n 1 -r
    echo ""
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        sudo apt install golang-go
    else
        echo -e "${RED}go is required for this project 🐀${RESET}" 1>&2 && exit 1
        exit 1
    fi
else 
    echo -e "${GREEN}go is installed 🐭${RESET}"
fi

sleep 2
clear

########## Setup golangci lint ##########
banner "🐭 Setting up golangci-lint"

if ! [ -x "$(command -v golangci-lint)" ]; then
    echo -e "${RED}golangci-lint is not installed 🐀${RESET}" 1>&2
    read -p "Do you want to install golangci-lint? (y/n) " -n 1 -r
    echo ""
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.1
    else
        echo -e "${RED}golangci-lint is required for this project 🐀${RESET}" 1>&2 && exit 1
        exit 1
    fi
else 
    echo -e "${GREEN}golangci-lint is installed 🐭${RESET}"
fi

sleep 2
clear

########## Setup make ##########
banner "📦 Setting up make"

if ! [ -x "$(command -v make)" ]; then
    echo -e "${RED}make is not installed 🐀${RESET}" 1>&2
    read -p "Do you want to install make? (y/n) " -n 1 -r
    echo ""
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        sudo apt install make
    else
        echo -e "${RED}make is required for this project 🐀${RESET}" 1>&2 && exit 1
        exit 1
    fi
else 
    echo -e "${GREEN}make is installed 🐭${RESET}"
fi

sleep 2
clear

##########################################################

########## Run lint ##########
banner "🐭 Running lint"

golangci-lint run

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Lint passed 🐭${RESET}"
else
    echo -e "${RED}Lint failed 🐀${RESET}" 1>&2 && exit 1
fi

sleep 2
clear

########## Run tests ##########
banner "🧪 Running tests"

go test -race -coverprofile=coverage.out -covermode=atomic -tags test ./...

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Tests passed 🧪${RESET}"
else
    echo -e "${RED}Tests failed 🐀${RESET}" 1>&2 && exit 1
fi

sleep 2
clear

########## Run build ##########
banner "🏗️ Running build"

make build

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build passed 🏗️${RESET}"
else
    echo -e "${RED}Build failed 🐀${RESET}" 1>&2 && exit 1
fi

sleep 2
clear

########## Change name file ##########
banner "📝 Changing name file"

cd bin

if [[ "$OS_TYPE" == "Darwin" ]]; then
    mv $FILE_NAME-darwin-amd64 $FILE_NAME
elif [[ "$OS_TYPE" == "Linux" ]]; then
    mv $FILE_NAME-linux-amd64 $FILE_NAME
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Name file changed 📝${RESET}"
else
    echo -e "${RED}Name file not changed 🐀${RESET}" 1>&2 && exit 1
fi

sleep 2
clear

########## Move file ##########
banner "📦 Moving file"

if [[ "$OS_TYPE" == "Darwin" ]]; then
    mv $FILE_NAME /usr/local/bin
elif [[ "$OS_TYPE" == "Linux" ]]; then
    sudo mv $FILE_NAME /usr/local/bin
fi

if [ $? -eq 0 ]; then
    echo -e "${GREEN}File moved 📦${RESET}"
else
    echo -e "${RED}File not moved 🐀${RESET}" 1>&2 && exit 1
fi

sleep 2
clear

########## Run dockercolorize ##########
banner "🐳 Running dockercolorize"

docker ps | dockercolorize

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Dockercolorize passed 🐳${RESET}"
else
    echo -e "${RED}Dockercolorize failed 🐀${RESET}" 1>&2 && exit 1
fi

sleep 2
clear

##########################################################

########## Finish ##########
banner "🎉 Finish"

echo -e "${GREEN}Dockercolorize is installed 🎉${RESET}"
echo -e "${GREEN}Run dockercolorize 🐳${RESET}"
echo -e "${GREEN}docker ps | dockercolorize${RESET}"
echo -e "${GREEN}Enjoy 🎉${RESET}"