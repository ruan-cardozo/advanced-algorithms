# Sorting Algorithms with OpenTelemetry

This project implements various sorting algorithms in Go with performance monitoring using OpenTelemetry and Jaeger.

## Prerequisites

### Go Installation
```sh
# For Ubuntu/Debian
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### Docker Installation
```sh
# For Ubuntu/Debian
sudo apt-get update
sudo apt-get install docker.io
sudo systemctl start docker
sudo systemctl enable docker
```

## Project Setup

1. Clone the repository:
```sh
git clone <repository-url>
cd advanced-algorithms
```

2. Install dependencies:
```sh
go mod download
```

3. Start Jaeger:
```sh
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.22
```

## Available Algorithms

- `bubble_sort`: Basic Bubble Sort implementation
- `bubble_sort_improved`: Optimized Bubble Sort
- `insertion_sort`: Insertion Sort
- `selection_sort`: Selection Sort
- `merge_sort`: Merge Sort
- `quick_sort`: Quick Sort
- `tim_sort`: Tim Sort
- `heap_sort`: Heap Sort

## Usage

Run algorithms with:
```sh
go run main.go <number_of_elements> <number_of_executions> <algorithm>
```

Examples:
```sh
# Single algorithm
go run main.go 1000 1 bubble_sort

# Multiple algorithms
go run main.go 500 5 bubble_sort,quick_sort

# All algorithms
go run main.go 100 3 all
```

## OpenTelemetry and Jaeger

This project uses OpenTelemetry to collect performance metrics and Jaeger for visualization.

### Accessing Jaeger UI
1. Start Jaeger using the Docker command above
2. Open `http://localhost:16686` in your browser
3. Select "advanced-algorithms" from the Service dropdown
4. Click "Find Traces" to view execution data

### Trace Information
- Total execution time
- Number of comparisons
- Number of swaps
- Total operations
- Algorithm-specific metrics

## Project Structure

```
├── algorithms/         # Sorting implementations
├── otel/              # OpenTelemetry setup
├── random_numbers/    # Number generation
├── strategy/          # Strategy pattern
├── utils/            # Helper functions
├── main.go           # Entry point
├── go.mod            # Dependencies
└── README.md         # Documentation
```