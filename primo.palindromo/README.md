# Primo e Palindromo

Foi criado uma versão Go, C e fortran eu tentei mas conheço pouco ainda só uma tentativa de como encontrar palíndromos e primos na expansão decimal do pi.

Para isto precisamos pegar esta expansão decimal e o site que está nada menos e nada mais que 31,4 trilhões de dígitos decimais da expansão decimal de pi é neste site [pi delivery](https://pi.delivery/).

Você pode fazer testes usando API REST que eles desenvolveram - [api rest pi](https://pi.delivery/#apipi_get) tem limitação da quantidade de dígitos da extensão PI em 1000 dígitos, mas o legal que pode pegar por posição da extensão pi.

```bash
$ curl 'https://api.pi.delivery/v1/pi?start=0&numberOfDigits=100'
```

É um novo recorde alcançado em março de 2022, foi calculado 100 trilhões de dígitos de pi, todos detalhes estão no site acima, vale a pena da uma conferida super interessante e vale a pena o aprendizado.

Aqui temos somente um arquivo com pix.txt com somente 98304 casas decimais da expansão do PI, para seguir, precisamos usar os arquivos disponibilizados pelo google com PI para que possamos rodar nosso algoritimo.


O site com os arquivos PI completos [estão aqui - pi 100 trilhões de dígitos](https://storage.googleapis.com/pi100t/index.html) e a página com todo explicativo de como eles calcularam 31.4 trilhões de dígitos [basta clicar aqui](https://pi.delivery/#introductionindex)

Você pode baixar os arquivos utilizando gsutil, porém cuidado são 82 TB de storage se quiser baixar tudo, 2 diretórios, 1832 arquivos.

```bash
$ gsutil -m rsync -R gs://pi100t ./

```
Aqui está o post explicadno toda arquitetura [calculando 100 trilhões dígitos de PI](https://cloud.google.com/blog/products/compute/calculating-100-trillion-digits-of-pi-on-google-cloud)

Um programa para ler o arquivo ycd é o [cruncher - linux](http://www.numberworld.org/y-cruncher/#Download), com ele você irá conseguir descompactar, visualizar etc..

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
