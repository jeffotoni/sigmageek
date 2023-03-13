# Primo e Palindromo

Aqui temos somente um arquivo com pix.txt com somente 98304 casas decimais da expansão do PI, para seguir, precisamos usar os arquivos disponibilizados pelo google com PI para que possamos rodar nosso algoritimo.
O site do google com os arquivos [estão aqui](https://storage.googleapis.com/pi100t/index.html) e a página com todo explicativo de como eles calcularam 31.4 trilhões de dígitos [basta clicar aqui](https://pi.delivery/#introductionindex)

#### Desafio SigmaGeek com 4 fases.

Fase 1: Qual é o primeiro primo palíndromo de 9 dígitos encontrado na expansão decimal de Pi? 
318272813

Fase 2: Qual é o primeiro primo palíndromo de 21 dígitos encontrado na expansão decimal de Pi? 
151978145606541879151

Fase 3: Qual é o maior número primo palíndromo encontrado na expansão decimal de Pi? 

9609457639843489367549069 

Fase final: Qual é o número seguinte na sequência? 

961212169
102454201
337515733
347676743

```go
$ cd primo.palindromo/go/palindromo.pi
$ go run main 9
Primeiro primo palindromo encontrado: 318272813
Proximo primo palindromo encontrado: 318464813
Maior número primo palíndromo encontrado: 318272813
```

```c
$ cd primo.palindromo/c
$ gcc -o palindormo main.c -static
$ ./palindormo 9
Primeiro primo palindromo encontrado: 318272813
Maior número primo palíndromo encontrado: 318272813
```
