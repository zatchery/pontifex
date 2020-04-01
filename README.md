# Pontifex

Go based project to play around with the Pontifex Cypher also know as the [Solitaire Cypher](https://en.wikipedia.org/wiki/Solitaire_(cipher)). You can use this binary or a normal deck of cards to encrypt and decrypt messages. If you are using a deck of cards in conjuction with this remember that programmers start counting at 0.


# Usage

To download the binary run:
```go get github.com/zatchery/pontifex```
Or download the source and build it using:
```go build```

## Key Stream files

By default pontifex looks for a key stream file at ~/.pontifex; You can override which keyfile you want to use with the -s/--keystreamfile option. There is an example key file in the root directory of the project that looks like this:
```yaml
streams:
- name: A_Key
stream: ["AD", "2D","3D","4D","5D","6D","7D","8D","9D","10D","JD","QD","KD","AC","2C","3C","4C","5C","6C","7C","8C","9C","10C","JC","QC","KC","AH","2H","3H","4H","5H","6H","7H","8H","9H","10H","JH","QH","KH","AS","2S","3S","4S","5S","6S","7S","8S","9S","10S","JS","QS","KS", "XY", "XZ"]
- name: B_Key
stream: ["2D","3D","4D","5D","6D","7D","8D","9D","10D","JD","QD","KD","AD","2C","3C","4C","5C","6C","7C","8C","9C","10C","JC","QC","KC","AC","2H","3H","4H","5H","6H","7H","8H","9H","10H","JH","QH","KH","AH","2S","3S","4S","5S","6S","7S","8S","9S","10S","JS","QS","KS","AS", "XY", "XZ"]
```
I would advise against using these as they are basically in sequential order. There isn't a lot of key verification as of this moment so make sure you are playing with a full deck when you input your key. The format of the card is number/J,Q,K,A followed by the first letter of the suit capitalized. You will need both of the jokers. The colored joker is XY, the b/w joker is XZ. 

## Encrypt

To encrypt a message run:
```pontifex encrypt -s A_Key HELLOWORLD```
You should create your own key file and unique key

## Decrypt
To decrypt a message run:
```pontifex decrypt -s A_Key RIXNJDQUGC```
You need to use the same key used to encrypt it.
