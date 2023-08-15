# GoLang-based Limit Order Book

## Synopsis

This project is an experiment to understand and implement a GoLang-based Limit Order Book with concurrent processing of several transactions using GoRoutines. It is inspired by the research paper [&#34;Exploiting Concurrency in Domain-Specific Data Structures: A Concurrent Order Book and Workload Generator for Online Trading&#34; (Riviere E., 2016)](https://www.researchgate.net/publication/308086656_Exploiting_Concurrency_in_Domain-Specific_Data_Structures_A_Concurrent_Order_Book_and_Workload_Generator_for_Online_Trading). The project aims to implement batch and asynchronous processing through the feature branch development workflow.

## Features

- Order Placement: Users can place Buy or Sell orders specifying the volume and price.
- Order Types: The system handles different order types - Market, Limit, and Cancel.
- Asynchronous Processing: GoRoutines are used for efficient concurrent processing.
- Order Matching: The system matches orders based on the order type and price, adhering to price-time priority.
- Order Cancellation: Orders can be cancelled by providing the unique Order ID.

## Quick Start

1. Clone this repo:
   ```bash
   git clone https://github.com/vats98754/go-order-book.git
   ```
2. Run the main.go file:
   ```bash
   go run main.go
   ```

## Code Structure

The project is structured into three main Go files:

- `orderbook.go`: This file holds the core logic of managing the Limit Order Book.
- `tradeprocessor.go`: It manages the order queue and processing trades using GoRoutines.
- `main.go`: It's the entry point of the project and combines everything together.

There are additional Go files to allow for custom features (like a research-based )

## Contributions

This project follows the feature branch development workflow. Feel free to create a new branch for each feature or bug fix you're working on. Once you're done with your changes, open a pull request and mention one of the core contributors for a review.

## Future Work

This project is a simplified version of a real-world limit order book and has a few limitations. Future improvements include error handling, dealing with partial order matches, and more advanced order routing.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
