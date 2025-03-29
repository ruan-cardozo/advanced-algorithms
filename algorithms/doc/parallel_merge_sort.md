# Algoritmo ParallelMergeSort - Documentação

## Visão Geral
O `ParallelMergeSort` é uma implementação paralela do algoritmo Merge Sort que utiliza concorrência em Go para melhorar o desempenho em arrays grandes. Esta implementação explora o potencial do paralelismo para dividir o trabalho de ordenação em tarefas independentes que podem ser executadas simultaneamente.

## Funcionamento do Algoritmo

### Estruturas Principais
- `ParallelMergeSort`: Estrutura principal que implementa a interface de ordenação
- `sortResult`: Estrutura auxiliar para transmitir resultados entre goroutines
- `THRESHOLD`: Constante que define o tamanho mínimo do array para aplicar paralelização (500.000 elementos)

### Principais Funções

#### `Name() string`
Retorna o identificador do algoritmo para uso em logs e relatórios.

#### `Sort(arr []int) ([]int, int, int)`
Função pública que inicia a ordenação e retorna o array ordenado junto com estatísticas.

#### `SortParallel(arr []int) ([]int, int, int)`
Implementação recursiva que:
1. Verifica se o tamanho do array está abaixo do threshold
2. Divide o array em duas metades
3. Processa cada metade em goroutines separadas
4. Combina os resultados usando `simpleMerge`
5. Retorna o array ordenado e estatísticas acumuladas

#### `simpleMerge(left, right []int, result []int) (int, int)`
Função que realiza a fusão sequencial de dois arrays ordenados, contabilizando comparações e trocas.

## Canais no Go

### O que são Canais
Canais são estruturas de dados de comunicação tipadas que permitem que goroutines troquem dados de forma segura, implementando o princípio "não compartilhe memória para comunicar; em vez disso, comunique para compartilhar memória".

### Implementação no Algoritmo
```go
leftChan := make(chan sortResult)
rightChan := make(chan sortResult)

go func() {
    left, c, s := ms.SortParallel(arr[:mid])
    leftChan <- sortResult{array: left, comparisons: c, swaps: s}
}()

go func() {
    right, c, s := ms.SortParallel(arr[mid:])
    rightChan <- sortResult{array: right, comparisons: c, swaps: s}
}()

leftResult := <-leftChan
rightResult := <-rightChan
```

### Funcionamento dos Canais
1. **Criação**: Canais são criados com `make(chan TipoDosDados)`
2. **Envio**: Dados são enviados usando o operador `<-` à direita do canal
3. **Recebimento**: Dados são recebidos usando o operador `<-` à esquerda do canal
4. **Sincronização**: Canais não-bufferizados bloqueiam o remetente até que um receptor esteja pronto

## Goroutines vs Threads Tradicionais

### Goroutines
- **Leveza**: Ocupam apenas ~2KB de memória no início (podem crescer)
- **Escalonamento**: Gerenciadas pelo runtime do Go, não pelo sistema operacional
- **Multiplexação**: Múltiplas goroutines são executadas em um número limitado de threads do SO
- **Comunicação**: Projetadas para usar canais como mecanismo primário de comunicação
- **Inicialização**: Extremamente rápidas de criar (microssegundos)
- **Escala**: Aplicações Go podem executar milhares ou milhões de goroutines simultaneamente

### Threads de Sistema Operacional
- **Peso**: Requerem mais memória (~1MB por thread) e recursos do sistema
- **Escalonamento**: Gerenciadas diretamente pelo sistema operacional
- **Overhead**: Maior custo de criação e troca de contexto
- **Comunicação**: Geralmente usam memória compartilhada e locks
- **Limite**: O número prático de threads é limitado pelos recursos do sistema

### Implicações no Algoritmo
O uso de goroutines permite criar uma estrutura de execução em árvore para o Merge Sort, onde cada nível da recursão pode ser executado em paralelo com custos mínimos. Contudo, o overhead de gerenciamento de goroutines torna-se significativo com arrays pequenos, justificando o uso do threshold.

## Otimizações Implementadas

### Threshold para Controle de Paralelismo
- Arrays com tamanho ≤ THRESHOLD são processados sequencialmente
- Reduz o overhead de criação de goroutines quando o ganho seria mínimo
- Valor configurável (default: 500.000) pode ser ajustado para diferentes ambientes

### Uso de Algoritmo Sequencial para Subarrays Pequenos
- A versão sequencial do MergeSort é utilizada quando o tamanho está abaixo do threshold
- Elimina milhares de goroutines desnecessárias em arrays grandes

### Merge Sequencial
- A etapa final de merge é executada sequencialmente
- Evita complexidades de sincronização em uma etapa naturalmente sequencial

## Considerações de Desempenho
- Para arrays muito grandes (>10M elementos), o paralelismo oferece ganhos significativos
- O overhead de gerenciamento de goroutines pode anular os ganhos em arrays menores
- O algoritmo se beneficia de sistemas com múltiplos núcleos de processamento
- O threshold ideal pode variar dependendo do hardware e da carga de trabalho
