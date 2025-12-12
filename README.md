# Language Learning App

This is a personal project to develop a simple platform that helps me practice Korean and Japanese, and tracks and customize my progress using AI agents. Initially, I was doing this in python using more SOTA AI frameworks but after some thoughts the AI agents I need for this project were not as complicated and can be implemented easily in another language through simple HTTP requests, so python's library felt like an overkill. Thus, I change the development to Golang. As a plus, the resulting binary file and docker image are much much smaller.

# Run me

To run this application, run the following command:

```sh
make run
```

If you run it with the default configuration, it will try to use a llama cpp server running on your local (on port `8080`). So you will need to install [`llama.cpp`](https://github.com/ggml-org/llama.cpp) and download a model to use this tool. I have tested it with [`gemma3`](https://huggingface.co/unsloth/gemma-3-4b-it-GGUF) and works fine with it.

# Build and run Docker Image

To build the docker image for this application, run the following command:
```sh
make docker
```

Then you can start a container of this image running the following command:
```sh
make run-docker
```

