# Alephium ALPH Balance Checker

This project helps find the ALPH balance of any Alephium address. I created two tools for this:

1. **TypeScript Version**: A library and command-line tool.
2. **Go Version**: A command-line tool.

Both tools have tests to make sure they work correctly.

## Installation

### TypeScript

1. **Navigate to the TypeScript directory:**

```bash
cd alephium-alph-balance/alephium-balance-ts
```

2. **Install dependencies:**

```bash
npm install
```
### Go

1. **Navigate to the Go directory:**

```bash
cd alephium-alph-balance/alephium-balance-go
```

2. **Install dependencies:**

```bash
go mod download
```
## Usage

### TypeScript Usage

1. **Run the balance checker:**

```bash
npx ts-node src/index.ts <Alephium_Address>
```

### Go Usage
1. **Build the Go executable:**

```bash
go build -o get-balance main.go
```

2. **Run the balance checker:**

```bash
./get-balance -address=<Alephium_Address>
```

## Testing

### TypeScript Tests

1. **Navigate to the TypeScript directory:**

```bash
cd alephium-alph-balance/alephium-balance-ts
```

2. **Run tests:**

```bash
npm test
```

### Go Tests

1. **Navigate to the Go directory:**

```bash
cd alephium-alph-balance/alephium-balance-go
```

2. **Run tests:**

```bash
go test
```
## References
[alephium documentation](https://alephium.org/)

## Contact
Please feel free to reach me at pandey7@buffalo.edu, if you have any questions or concerns.