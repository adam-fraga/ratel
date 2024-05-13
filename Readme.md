# Documentation

Each folder containing in cmd folder is a top of level command (excludig the root command).
if the file.go has the same name of the folder then the file.go is the root command of the palet.

Each file.go is a main palet that contains the main commands that display their options and purpose,
and a function that add their subcommands.

Following that logic we can implement a command line interface that is easy to understand and to maintain
with several levels of nested commands.

Viper is used to manage the configuration of the application, and cobra is used to manage the commands.

### TODO

##### COMMANDS

- [ ] Add a command to generate handlers
- [ ] Add a command to generate models

<!-- LATER -->

- [ ] Add a command to generate the proto files
- [ ] Add a command to compile the proto files
- [ ] Add a command to generate the grpc server

##### IMPLEMENT LOGIC (HANDLER) FOR COMMANDS

- [ ] Implement the logic for the commands that generates the project
      (The project should be generated following the structure the json data folders and files)
      (Project name should be asked to the user)
      (Database should be asked to the user)

- [ ] Implement the logic for the commands that generates the cache
      (Start by choosing between simple cache like Go cache package or Valkey)
      (Write logic to run the cache client and server)
      (Write logic to launch basic cache commands)
      (Add more to this todo section later...)

- [ ] Implement the logic for the commands that interact with postgres database
      (Write logic to connect to the database)
      (Write logic to launch basic database commands)

- [ ] Implement the logic for the commands that generates the middleware
      (Write logic to create the middleware files)

- [ ] Implement the logic for the commands that generates the views with templ
      (Write logic to create the views files)

- [ ] Implement the logic for the commands that generates the handlers
      (Write logic to create the handlers files)

- [ ] Implement the logic for the commands that generates the models
      (Write logic to create the models files)
