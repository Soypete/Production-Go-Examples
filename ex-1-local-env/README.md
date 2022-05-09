# Setup Your Local Dev Environment (Linters)

There are lots of things we can do to utilize IDE, text-editor, dev tools and pipelines to encourge production of high quality code. The go binary originally shipped with a build in set of linting tools: `go-lint`, `go-fmt`, `go-vet`. Linting has always been very important to the Go community, and there have been many iterations of community driven linters. Today the most popular linter is `Golangci-Lint`. This linter is meant to be used to verify code quality and was designed to be used as part of your Conituous Integration pipeline.

## Download golang-ci-lint

Visit the [Golangci-lint](https://golangci-lint.run/usage/install/) site and install the binary on your local system. From the home directory of this repo run the command `golangci-lint run` and confirm the setup works. 

## Select Linters you want

There many linters included as part of the Golangci-lint binary. The list of linters and description of what the linters check. Some linters overlap in functionality. Go thorugh the list and select some linters you want to apply. Use the command-line flags to run these against the `main.go` in the exercise . Here is an example of the command:

```bash
golangci-lint run --enable deadcode errcheck govet'
```
## Add to IDE

Adding a linter to your developer environment (IDE or text editor) can same your alot of hassel while you are writing software. There are instructions on the [golangci-lint docs](https://golangci-lint.run/usage/integrations/) to add linting to editor of choice. Follow the instructions to set up linting on your text editor. If the settings is available enable linting on save.

Below are examples of my personal linting configurations.
### VSCODE

![vscode coming](img/vscode.png)

### Vim

![vim-go config](img/vim.png)
[link to dotfiles](https://github.com/Soypete/dotfiles/blob/main/vim/vimrc)

### Additional exercise: 
You can create a `golangci-lint.yml` to store the configurations of our linters. An example file can be found [here](ex-1-local-env/.golangci-lint.yml). Additional configurations can also be found in [this file](https://github.com/golangci/golangci-lint/blob/master/.golangci.example.yml). And a `golangci-lint.yml` file to any local go repo you have and run the `golangci-lint run` command. Change the linter configs and try again.

