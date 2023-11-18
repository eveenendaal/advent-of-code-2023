FROM mcr.microsoft.com/vscode/devcontainers/universal:2

COPY bash_aliases.sh /home/codespace/.bash_aliases

RUN chown codespace:codespace /home/codespace/.bash_aliases

ENTRYPOINT ["/opt/tools/startingpoint.sh", "built-in"]
