# ZenBTC

ZenBTC is a Cosmos SDK module that provides Bitcoin bridge functionality, allowing users to mint and redeem zenBTC tokens backed by Bitcoin.

## Overview

The ZenBTC module enables:
- Bitcoin deposits and zenBTC minting
- zenBTC burning and Bitcoin redemptions
- Cross-chain transaction management
- Supply tracking and exchange rate calculations

## Project Structure

```
zenbtc/
├── api/                    # API definitions and gRPC services
├── bindings/              # Smart contract bindings
├── converter/             # Bitcoin transaction converter
├── internal/              # Internal utilities and clients
├── proto/                 # Protocol buffer definitions
├── x/zenbtc/             # Cosmos SDK module
│   ├── keeper/           # Business logic implementation
│   ├── types/            # Data structures and types
│   ├── module/           # Module configuration
│   └── testutil/         # Test utilities and mocks
└── test/                 # Integration tests
```

## Testing

### Running All Tests

```bash
# Run all tests in the project
go test ./...

# Run tests with verbose output
go test -v ./...
```

### Running Specific Test Suites

```bash
# Run keeper integration tests
go test ./x/zenbtc/keeper -v

# Run keeper integration tests with specific test suite
go test ./x/zenbtc/keeper -v -run TestKeeperTestSuite

# Run individual test methods
go test ./x/zenbtc/keeper -v -run Test_ZenbtcKeeper_GetExchangeRate

# Run converter tests
go test ./converter -v

# Run internal package tests
go test ./internal/... -v
```

### Running Tests with Coverage

```bash
# Run tests with coverage report
go test -cover ./...

# Generate detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# Run tests with coverage for specific packages
go test -cover ./x/zenbtc/keeper
go test -cover ./x/zenbtc/types
```

### Running Integration Tests

```bash
# Run the comprehensive integration test suite
go test ./x/zenbtc/keeper -v -run TestKeeperTestSuite

# This will run tests for:
# - Exchange rate calculations
# - Pending mint transactions
# - Redemptions
# - Supply management
# - Burn events
# - Parameter getters
# - Store operations
```

### Test Structure

The integration tests (`integration_test.go`) cover:

1. **Exchange Rate Tests**: Testing BTC to zenBTC exchange rate calculations
2. **Transaction Management**: Testing pending mint transactions and redemptions
3. **Supply Tracking**: Testing supply data storage and retrieval
4. **Burn Events**: Testing burn event processing
5. **Parameter Management**: Testing all parameter getter methods
6. **Store Operations**: Testing collection operations and walking functions

### Mock Dependencies

The tests use gomock for mocking dependencies:
- `MockAccountKeeper`: Mocks account operations
- `MockBankKeeper`: Mocks bank operations  
- `MockParamSubspace`: Mocks parameter storage

## Development

### Adding New Tests

When adding new functionality, follow the existing test patterns:

1. Add unit tests for individual functions
2. Add integration tests for keeper methods
3. Use the existing test suite structure
4. Include both success and error scenarios

### Test Best Practices

- Use descriptive test names
- Test both positive and negative cases
- Use table-driven tests for multiple scenarios
- Mock external dependencies
- Clean up test data after each test

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass
5. Submit a pull request

## License

[Add your license information here]
