# perlin noise generator
Tomas Mažvila IFF-6/13 modulio P170B328 lygiagretus programavimas inžinerinis porjektas


Programos instaliavimo instrukcija linux vartotojamas:
reikalinga: go bei git paketai, tuomet:

`mkdir ~/go`
`export GOPATH=$HOME/go`
`export PATH=$PATH:$GOPATH/bin`
`go get github.com/tomazvila/perlin`
`go get github.com/fogleman/gg`
`cd $GOPATH/src/github.com/tomazvila/perlin`
`go install`

Dabar iš bet kur galima pasileisti perlin noise generatorių naudojant komandą perlin_noise

Norint pakeisti generavimo konstantas, reikia paredaguoti kodą esantį `cd $GOPATH/src/github.com/tomazvila/perlin/main/perlin_noise.go` faile 12-19 eilutėse
