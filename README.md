<h1 align="center">Pelaware</h1>

# :paperclip: Introduction

Simple Ransomware that uses AES-256-GCM encryption and is writted in Go language.

:warning: WARNING: This software is made just for study purposes.

# :computer: Technologies

* [Golang](https://golang.org/)
* [Avanced Encryption Standard (AES)](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)

# :rocket: How to Run

```bash
# Clone Repository
$ git clone https://github.com/davidcbbc/pelaware
```

### Configure and Compile

```bash
# Go to folder
$ cd pelagoware

# Generate a AES crypto key
$ go run keygen/keygen.go

# In encrypter nd decrypter, set the variables cryptoKey, contact and dir

# Compile encrypter
$ cd encrypter
$ go build

# Compile decrypter
$ cd decrypter
$ go build
```

