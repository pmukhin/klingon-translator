[![Build Status](https://travis-ci.org/pmukhin/klingon-translator.svg?branch=master)](https://travis-ci.org/pmukhin/klingon-translator)

# klingon-translator
`klingon-translator` transliterates latin-scripted names in Klingon to the Klingon script.

## Architecture
### Transliteration
Input is lexed by a `Lexer` to find a letter-by-letter equivalent.

### Fetching the species
There's a `stapi` package which contains a `stapi.Client`. `stapi.Client` contains access to entity-specific clients.
To search for characters, for instance, we have to do the following:

```go
client := stapi.New(...)
charactersClient := client.Characters()
```

### Client
According to usage in the project, every client (in our case only one client which is `CharactersClient`) contains methods `Search` and `Get` and they take a searched name and a specific character UID perspectively.

```go
searchResult, err := client.Characters().Search(..)
getResult, err := client.Characters().Get(...)
```
### Normalize method
In case of a Get method there's `.normalize` method call. It's required because some of characters don't have species specified explicilty by the API. `.normalize` just checks if there's at least one species, and if there's no, it appends the default species.

I assume that the default species is `Human` because the given example `Uhura` is a human, but there's no explicit declaration of it given by the API.
