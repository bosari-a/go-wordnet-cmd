# Dictionary command-line tool

## Description

This is a command-line dictionary based on the wordnet database which uses my dictionary API: [repo link](https://github.com/bosari-a/go-wordnet?tab=readme-ov-file#wordnet-based-dictionary-api).

## Installation

1. Go to the [releases section](https://github.com/bosari-a/go-wordnet-cmd/releases/tag/v1.0.0) of this repo and download the binaries:
    - `wordnet` (with no extension) for linux based systems.
    - `wordnet.exe` for windows amd64.

2. Add the wordnet database to `$PATH`:
    - On linux you can run export or add an export statement to your `.bashrc` file.
    For example:
    ```sh
    export WORDNET_DICT_PATH="/home/path_to/dict/"
    ```
    The variable **must** be called `WORDNET_DICT_PATH`.
    
    - On windows use system environment variables:
    ![environment variables on windows](
        https://raw.githubusercontent.com/bosari-a/go-wordnet-cmd/main/imgs/windowsenv.png
    )

## Usage:

Command:
```sh
wordnet <word>
```

Example:
![example of using wordnet command](
        https://raw.githubusercontent.com/bosari-a/go-wordnet-cmd/main/imgs/wordnetcmd.png
    )

