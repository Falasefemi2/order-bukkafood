<!-- @format -->

# Restaurant Ordering System

This Go package is the main entry point for the restaurant ordering system. It handles the user interaction, prompting for the customer's name, and calling the necessary functions to display the welcome and goodbye messages, as well as to initiate the ordering process.

## Variables

### `subTotalBill`

The `subTotalBill` variable is a `float64` that stores the running subtotal bill for the customer's order. It is updated every time the customer orders an item.

### `customerOrder`

The `customerOrder` variable is a map that stores the customer's order. The keys in the map are the item names (strings), and the values are the quantities ordered for each item (integers). It is initialized as an empty map.

## Functions

### `main()`

This is the main function of the program, which serves as the entry point. It performs the following tasks:

1. Prompts the user to enter their name.
2. Calls the `welcome()` function (defined in a separate file) to display a welcome message with the customer's name.
3. Calls the `order()` function (defined in a separate file) to initiate the ordering process.
4. Calls the `bye()` function (defined in a separate file) to display a goodbye message with the customer's name.

## Usage

To run this program, simply execute the compiled binary or run the Go code with the `go run` command:

The program will prompt you to enter your name, display a welcome message, allow you to place orders, and finally show a goodbye message when you exit the ordering process.

Note that this `main.go` file depends on the `welcome()`, `order()`, and `bye()` functions defined in separate files. Make sure to include or import the necessary code for those functions.

Additionally, this implementation is a basic example and may require further enhancements or modifications to suit your specific requirements, such as handling payment, generating receipts, or integrating with a database for order management.

## Contributing

Contributions to this package are welcome. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on the project's GitHub repository.

# Welcome and Goodbye Messages

This Go package provides functions to print welcome and goodbye messages with a customer's name. The messages are designed to be displayed in a visually appealing way using ASCII art.

## Functions

### `welcome(customerName string)`

This function prints a welcome message with the provided customer name. The message is formatted using ASCII art, with the customer's name centered and surrounded by underscores and backslashes.

Example output:

Welcome John Doe
_/\\\_ Welcome to Bukka Hut! _/\\\_

                       ### `bye(customerName string)`

This function prints a goodbye message with the provided customer name. The message includes a farewell statement and a wish to see the customer again. Like the welcome message, it uses ASCII art for formatting.

Example output:
_/\\\_ Goodbye, Enjoy your meal! _/\\\_
We hope to see you again! John Doe

# Restaurant Menu

This Go package provides a simple implementation of a restaurant menu system. It defines a `Menu` struct to store information about menu items and includes a function to print the menu in a formatted table.

## Struct

### `Menu`

The `Menu` struct represents a menu item and has the following fields:

- `itemNo` (int): The serial number or identifier of the menu item.
- `itemName` (string): The name of the menu item.
- `itemPrice` (float64): The price of the menu item.

An example `Menu` struct:

```go
{1, "Jollof Rice", 20.88}

Menu
--------------------------------------------------
S.No  Item Name                      Price
--------------------------------------------------
1     Jollof Rice                     ₦20.88
2     Fried Rice                      ₦18.99
3     Chicken Curry                   ₦22.50
...
15    Mushroom Risotto                ₦20.60
--------------------------------------------------


# Restaurant Ordering System

This Go package provides a simple implementation of a restaurant ordering system. It allows customers to view the menu, select items, specify quantities, and keep track of their current order. The package also includes functionality to display the running order and calculate the subtotal bill.

## Functions

### `order()`

This function is the main entry point for the ordering system. It presents a menu to the user and allows them to select items, specify quantities, and track their order. The function uses a loop to continuously prompt the user for input until they choose to exit.

The `order()` function performs the following tasks:

1. Prints the restaurant menu by calling the `printMenu()` function (defined in a separate file).
2. Prompts the user to enter the food number of the desired item or 0 to exit.
3. If a valid food number is entered, it retrieves the corresponding item information (name and price) from the `menu` slice.
4. Prompts the user to enter the number of plates they want for the selected item.
5. If a non-zero quantity is entered, it updates the `customerOrder` map with the item name and quantity, and calculates the subtotal bill.
6. Calls the `orderTillNow()` function to print the current order details.

### `orderTillNow()`

This function prints the customer's order details, including the quantity and item name for each ordered item. The output is formatted in a table-like structure for better readability.

Example output:
```
