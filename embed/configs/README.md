<!-- PROJECT LOGO -->
<div align="center">
  <h1 align="center">Ratel</h1>
  <p align="center">
    A powerful web framework for Go developers
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

Ratel is a versatile web framework designed to streamline web development with Go.
It provides developers with a comprehensive set of tools and functionalities to simplify project setup,
management, and deployment.

Whether you're building a small web application or a large-scale project,
Ratel offers the flexibility and scalability you need to succeed.

Key Features:

- **Project Management**: Ratel offers commands for project setup, configuration, and deployment,
  making it easy to get started with your web development journey.
- **Middleware Support**: The framework includes middleware commands for managing authentication,
  logging, error handling, and request/response modification.
- **Database Integration**: Ratel provides commands for database setup, migration, seeding, querying,
  and administration, ensuring smooth integration with your chosen database system.
- **View Handling**: With Ratel, you can easily create and list views within your project,
  simplifying the process of managing your project's frontend components, the generated views are in the templ
  format used by the [a-h/templ](https://github.com/a-h/templ) templating engine.
- **Live Reloading**: Ratel supports live reloading with [Air](https://github.com/air-verse/air), allowing you to
  see changes in real-time.
- **Handler Management**: Ratel provides commands for creating, listing, handlers in your project
- **Model Management**: Ratel provides commands for creating, listing, models in your project

# And more...

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

<!-- INSTALLATION -->

## Installation

To install Ratel on your local machine, simply download the latest version of the binary here:
Website coming soon...

Alternatively, you can build Ratel from source by following these steps:

1. Clone the Ratel repository:

```bash
git clone https://github.com/adam-fraga/ratel.git
```

## Usage

Once Ratel is installed, you can start using it to develop your web projects.

### Creating a New Project

Use the `ratel project create` command to create a new project:

```bash
ratel project create my-project
```

### Initializing a New Project

```bash
go mod init "github.com/your-username/your-project"
```

### Install necessary dependencies

You need to install Air for live reloading and templ templating engine.
You can install them using the following links:

[Air](https://github.com/air-verse/air)
[Templ](https://github.com/a-h/templ)

We also recommend using a terminal multiplexer like tmux or screen to run multiple commands
in a single terminal window this will help you run the server and the air live reload at the same time
and build your typescript + tailwind files with webpack.

```bash
npm install
go get github.com/a-h/templ
go get joho/godotenv
go mod tidy
```

### Managing Middleware

Use the `ratel middleware` command to manage middleware functionalities within your project:

```bash
ratel middleware create auth
```

### View Handling

Use the `ratel view` command to handle views within your project, including creation (.templ), listing, updating, and deletion:

```bash
ratel view create-component header
ratel view list-components

ratel view create-page home
ratel view list-pages

ratel view create-template base
ratel view list-templates

# And more...
```

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Roadmap

See the open issues for a list of proposed features (and known issues) to track the development progress.

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create.
Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## License

Distributed under the MIT License. See LICENSE for more information.

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>

## Contact

Adm FRG - fragadams@gmail.com
[Discord](https://discord.gg/yourdiscord)
[Youtube](https://youtube.com/yourchannel)

Project Link: [https://github.com/adam-fraga/ratel](https://github.com/your-username/ratel)

<p align="right">(<a href="#table-of-contents">back to top</a>)</p>
```
