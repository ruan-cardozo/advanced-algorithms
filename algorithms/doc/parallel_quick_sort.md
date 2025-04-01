# Parallel Quick Sort

## Introdução

O **Parallel Quick Sort** é uma versão paralela do algoritmo Quick Sort, projetada para aproveitar o paralelismo em sistemas multicore. Ele divide o array em partições e processa essas partições simultaneamente usando goroutines no Go, reduzindo o tempo total de execução para arrays grandes.

## Funcionamento

O **Parallel Quick Sort** segue a mesma lógica básica do Quick Sort:
1. Escolhe um pivô.
2. Particiona o array em duas partes:
   - Elementos menores que o pivô.
   - Elementos maiores ou iguais ao pivô.
3. Ordena recursivamente as duas partições.

A diferença principal é que, no **Parallel Quick Sort**, as partições são processadas em paralelo usando goroutines, desde que o tamanho do subarray seja maior que um **threshold** definido.

### Estrutura do Algoritmo

1. **Particionamento**:
   - O array é dividido em duas partições com base no pivô.
   - Durante o particionamento, são realizadas comparações e trocas para reorganizar os elementos.

2. **Paralelismo**:
   - Se o tamanho do subarray for maior que o **threshold**, duas goroutines são criadas para processar as partições esquerda e direita em paralelo.
   - Caso contrário, o subarray é ordenado sequencialmente.

3. **Sincronização**:
   - Um `sync.WaitGroup` é usado para garantir que todas as goroutines terminem antes que o algoritmo continue.

### Threshold

O **threshold** é um valor que define o tamanho mínimo do subarray para que o paralelismo seja aplicado. Para subarrays menores que o **threshold**, o algoritmo usa o Quick Sort sequencial, evitando o overhead de criar goroutines desnecessárias.

## Implementação

### Código Principal

```go
const THRESHOLD = 1000

type ParallelQuickSort struct{}

func (pq ParallelQuickSort) Name() string {
    return "Parallel Quick Sort"
}

func (pq ParallelQuickSort) Sort(arr []int) ([]int, int, int) {
    return pq.SortParallel(arr)
}

func (pq ParallelQuickSort) SortParallel(arr []int) ([]int, int, int) {
    var wg sync.WaitGroup
    comparisons := 0
    swaps := 0

    pq.quickSortParallel(arr, 0, len(arr)-1, &wg, &comparisons, &swaps)
    wg.Wait()

    return arr, comparisons, swaps
}