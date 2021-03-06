## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [Released]
## [0.4.0] - 2022-06-17
### Added
- Comparison functions to compare types.
  - `EqualTags`.
  - `EqualItem` and `EqualItems`.
  - `EqualTransaction` and `EqualTransactions`.
- The CRUD Transaction methods.
  - `CreateTransaction` to create a new transaction.
  - `GetBankAccountTransactions` to get all the transactions for a bank account.
  - `UpdateTransaction` to update a transaction.
  - `DeleteTransaction` to delete a transaction.

### Updated
- Transaction to have an `AccountUUID` instead of `BankAccountUUID` for a more
generic description.

## [0.3.0] - 2022-06-17
### Added
- The CRUD Bank Account methods.
  - `CreateBankAccount` to create a new bank account.
  - `UpdateBankAccount` to update a bank account.
  - `DeleteBankAccount` to delete a bank account.
  - `GetUserBankAccounts` to get all the users bank accounts.
  - `GetOrganisationBankAccounts` to get all the organisations bank accounts.

## [0.2.0] - 2022-06-14
### Added
- The `Service` type
  - `SetURL` method to make the service a `microtest.Mock` interface.
- The microservice-package `NewService` function to be able to create a new 
msp instance.

## [0.1.0] - 2022-06-14
### Added
- Some basic types
  - Bank
  - BankAccount
  - Transaction
  - Item
  - Tag
## [0.0.0] - 2022-05-02
### Added
- Initial commit.


