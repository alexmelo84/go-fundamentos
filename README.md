# Golang

Estudo do Golang pelo [curso](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/?couponCode=24T4MT92724A).

## Observações:

- na primeira linha de cada arquivo Go, devemos informar o nome do pacote. No caso do pacote principal, ele deve ser *main*:

```
package main
```

Através disso o Go saberá por onde começar a executar a aplicação;

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

Exemplo disso está no módulo *1-Pacotes*, em que no módulo *auxiliar* há o pacote *auxiliar1* há a função *Escrever* e que é executada dentro do pacote *main.go*;

- nos pacotes auxiliares, quando uma função é iniciada com letra minúscula, ela só pode ser executada dentro do próprio pacote.

Exemplo é no módulo *1-Pacotes*, onde o módulo *auxiliar* tem o pacote *auxiliar2* há a função *escrever2*. Essa função não pode ser executada no pacote principal (*main.go*), só podendo ser executada no pacote *auxiliar1*;

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

- Go trabalha com 4 tipos de números inteiros: int8, int16, int32, int64. Se omitir o tipo de inteiro na inicialização da variável, o Go considerará a arquitetura do processador para inferir o tipo de inteiro;

- tipos inteiros também podem receber números negativos;

- variáveis do tipo *char* retornam o código ASCII do caractere;

- valor zero: utilizado para variáveis inicilizadas sem valor;

- no Go podemos retornar mais de um tipo em uma função, exemplo de uma função que retornaria dois tipos de valores:

```
func exemplo(parametro1, parametro2 [tipo]) (tipo, tipo) {}

variavel1, variavel2 := exemplo(parametro1, parametro2)
```

Se for precisar de apenas um dos resultados de uma função que retorna múltiplos valores, deve-se usar o *_*:

```
func exemplo(parametro1, parametro2 [tipo]) (tipo, tipo) {}

variavel1, _ := exemplo(parametro1, parametro2)
```

- no Go não é possível fazer operações em variáveis de tipos diferentes, exemplo que dará erro:

```
var numero1 int8 = 10
var numero1 int16 = 29
soma := numero1 + numero2
fmt.Println(soma)
```

- no Go não tem herança, mas pode-se simular isso passando um tipo de struct dentro de outro struct, exemplo com o tipo *pessoa* em *estudante*:

```
type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso string
	campus string
}
```

- para utilizar ponteiros, deve adicionar um * antes do tipo da variável, exemplo:

```
var variavel1 int
var ponteiro1 *int
variavel1 = 20
ponteiro1 = &variavel1
fmt.Println(variavel3, ponteiro1, *ponteiro1)
```

Quando lemos apenas o nome do ponteiro (nesse exemplo é o *ponteiro1*) pegamos apenas a posição de memória. Se quisermos pegar o valor do ponteiro, adiciona-se o * para desreferenciar o ponteiro (nesse exemplo é **ponteiro1*);

- quando criamos um *slice* utilizando o *make*, ao estourarmos a capacidade máxima, o Go vai dobrar o tamanho original do *slice*;

- no loop que usar o *range*, podemos usar o *_* caso não precisemos de algum dado, exemplo que não usamos o índice:

```
nomes := [3]string{"Alex", "João", "Paulo"}
for _, nome := range nomes {
	fmt.Println(nome)
}
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
