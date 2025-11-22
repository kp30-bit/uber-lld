# Uber-Ola Low-Level Design

A Go implementation of a ride-sharing system (Uber/Ola) low-level design.

## Architecture

This project follows clean architecture principles with the following structure:

- **`cmd/`** - Main application entry point
- **`internal/domain/`** - Core domain entities (Rider, Driver, Trip, Location, TripMetaData)
- **`internal/manager/`** - Manager singletons (RiderMgr, DriverMgr, TripMgr, StrategyMgr)
- **`internal/strategy/`** - Strategy pattern implementations (Pricing and Driver Matching strategies)
- **`pkg/common/`** - Common types and utilities (Rating, TripStatus, Util)
- **`pkg/interfaces/`** - Strategy interfaces and provider interfaces

## Design Patterns

- **Singleton Pattern**: Used for managers (RiderMgr, DriverMgr, TripMgr, StrategyMgr)
- **Strategy Pattern**: Used for pricing strategies and driver matching strategies
- **Thread Safety**: All managers use mutexes for concurrent access

## Running the Application

```bash
go run cmd/main.go
```

## Building

```bash
go build -o uber cmd/main.go
```

## Class Diagram

See `uber_lld.png` for the UML class diagram.
