<h1 align="center">
O projeto consiste em desenvolver uma API REST em Go.
</h1>

## Participantes

[Luciano Ribeiro](https://github.com/lucianorbr)

## Tecnologias

- [x] Go
- [x] Docker

- Framework:
- [x] Fiber


## O Projeto
- Fui contratado para desenvolver um projeto em Go, Java, Python ou Javascript (NodeJS),
  que vai identificar se uma sequência de letras é válida. 


- Você saberá se uma sequência é válida, se encontrar 2 ou mais sequências de quatro
  letras iguais em qualquer direção, horizontal, vertical ou nas diagonais.
   As letras da String só podem ser: (B, U, D, H)


<p>-> /sequence
{
"letters": ["DUHBHB", "DUBUHD", "UBUUHU", "BHBDHH", " DDDDUB", "UDBDUH"]
} </p> 

<p>A API deve retornar um json com "is_valid": boolean. Caso você identifique uma sequência
válida, deve ser true, caso identifique uma sequência inválida, deve ser false, como no
exemplo abaixo:</p>

<p>HTTP 200</p>

{"is_valid": true}

## Nível 1:
Desenvolva uma API que esteja de acordo com os requisitos propostos acima, que seja capaz
de validar uma sequência de letras válidas.


## Nível 2:
<p>
Use um banco de dados de sua preferência para armazenar as sequencias verificadas pela API. 
Esse banco deve garantir a unicidade, ou seja, apenas 1 registro por sequência. 
Disponibilizar um outro endpoint "/stats" que responde um HTTP GET. A resposta deve ser um 
Json que retorna as estatísticas de verificações de sequências, onde deve informar a 
quantidade de sequências válidas, quantidade de sequências inválidas, e a proporção de 
sequências válidas em relação ao total. Segue exemplo da resposta:
</p>

<p>
{"count_valid": 40, "count_invalid": 60: "ratio": 0.4} 
</p>

## Nível 3:
<p>
Construir um Docker composse para executar a API, para possibilitar a execução em qualquer 
ambiente. 
</p>

