# Synonym CLI
This is a cli tool that will give you a list of synonyms for a word
that you enter.

## Usage
`synonym-cli <word>`

The program will return a list of synonyms for the entered word on
stdout, which can be easily piped into other programs such as `shuf`.

**Note:** This project uses the *Merriam-Webster* thesaurus API. To
use this program, the environment variable `THESAURUS_API_KEY` must be
set. Instructions on obtaining an API key can be found [on the
official documentation](https://dictionaryapi.com/)
<!--  LocalWords:  stdout
 -->
