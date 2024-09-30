# Golang

Estudo do Golang pelo [curso](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/?couponCode=24T4MT92724A).

## Observações:

- na primeira linha de cada arquivo Go, devemos informar o nome do pacote. No caso do pacote principal, ele deve ser *main*:

```
package main
```

Através disso o Go saberá por onde começar a executar a aplicação.

- pacotes auxiliares terão seu próprio nome:

```
package auxiliar
```

- o pacote principal (*main*) deve ter a função *main* obrigatoriamente:

```
func main() {}
```

- após criar um novo pacote (diretório), deve ser iniciar o módulo:

```
go mod init [moduleName]
```

- nos pacotes auxiliares, o nome da *func* deve ser escrito com letra inicial maiúscula, só assim ela poderá ser importada no pacote *main*.

Exemplo disso está no módulo *1-Pacotes*, em que no módulo *auxiliar* há o pacote *auxiliar1* há a função *Escrever* e que é executada dentro do pacote *main.go*.

- nos pacotes auxiliares, quando uma função é iniciada com letra minúscula, ela só pode ser executada dentro do próprio pacote.

Exemplo é no módulo *1-Pacotes*, onde o módulo *auxiliar* tem o pacote *auxiliar2* há a função *escrever2*. Essa função não pode ser executada no pacote principal (*main.go*), só podendo ser executada no pacote *auxiliar1*.

- para executar funções de pacotes, deve-se usar o texto que fica após a última */*, por exemplo, um pacote chamado *pacote/exemplo* deve ser executado como *exemplo.funcao1()*;

- utlização de pacotes externos: precisam ser instalados com o comando:

```
go get [url]
```

Por exemplo:

```
go get github.com/badoux/checkmail
```

- quando remove algum pacote, deve-se rodar o comando abaixo para limpar o projeto:

```
go mod tidy
```

## Utilização

Cada diretório equivale a 1 conceito diferente.

Há duas formas de rodar os executáveis:

- da raíz do projeto execute adicionando o nome do diretório antes do executável, exemplo:

```
./1-Pacotes/modulo1
```

- ou acesse cada diretório e execute, exemplo:

```
cd 1-Pacotes
./modulo1
```

Se fizer alteração em qualquer arquivo, há duas formas de testar a alteração:

- rodando o arquivo diretamente:

```
go run [filename].go
```

- gerando um novo executável:

```
go build
./[executavel]
```
