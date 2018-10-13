# Alfred

This project aims at create a bot back-end with an front-end disponible on differents services like Facebook Messenger, Google Home etc... We want to create build for Windows and Linux to have a personnal assistant. For the moment we need to decide is we gonna code it in Go or Rust but anyway in a perfect world we create bindings that allow everyone to create their back-end and the bot will have a core that manage output on all services.

# Prototype :

I have an idea but all improvements are welcome. I want to create a repository call "plugins" with all backend service. An other folder call "core" with functions that must be called in the plugins. Functions such as (send message()) which gonna work on all services.
