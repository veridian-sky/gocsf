# Go OCSF Package

A Go package for working with the Open Cybersecurity Schema Framework (OCSF) v1.6.0.

## Overview

This package provides Go struct definitions for all OCSF objects and events, automatically generated from the official OCSF schema. It enables easy serialization and deserialization of OCSF-compliant cybersecurity event data in Go applications.

## Installation

```bash
go get github.com/veridian-sky/gocsf
```

## Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/veridian-sky/gocsf"
)

func main() {
    // Create a User object
    user := gocsf.User{
        Name:         "john.doe",
        FullName:     "John Doe",
        EmailAddr:    "john.doe@example.com",
        Uid:          "user-123",
        Type:         "Regular",
        TypeId:       1,
        DisplayName:  "John Doe",
    }

    // Create an Account object  
    account := gocsf.Account{
        Name:   "Production Account",
        Uid:    "acc-456",
        Type:   "Cloud",
        TypeId: 2,
    }

    // Set the user's account
    user.Account = account

    // Create an Actor object
    actor := gocsf.Actor{
        User: user,
        AppName: "MyApp",
        AppUid:  "app-789",
    }

    // Serialize to JSON
    data, err := json.MarshalIndent(actor, "", "  ")
    if err != nil {
        panic(err)
    }

    fmt.Println("OCSF Actor JSON:")
    fmt.Println(string(data))
}
```

## Features

- **Complete OCSF Coverage**: All OCSF v1.6.0 objects and events are supported
- **Type Safety**: Strongly typed Go structs with proper JSON serialization tags
- **Zero Dependencies**: Only uses Go standard library packages
- **Generated Code**: Automatically generated from the official OCSF schema to ensure accuracy

## Generating Types from Different OCSF Versions

The package includes a dynamic generator that can fetch and generate Go types from any OCSF schema version:

```bash
# Generate from the latest OCSF v1.6.0 (default)
go run cmd/generator/generate_multi.go

# Generate from a specific version
go run cmd/generator/generate_multi.go -version=1.3.0

# Use local schema.json file instead of fetching
go run cmd/generator/generate_multi.go -local

# Show help
go run cmd/generator/generate_multi.go -help
```

The generator will:
1. Fetch the specified OCSF schema from `https://schema.ocsf.io/<version>/export/schema`
2. Save it locally as `schema.json` for future use
3. Generate organized Go types across multiple logical files
4. Handle imports automatically and prevent recursive type issues

## Package Structure

The OCSF types are organized into logical groups for better maintainability:

- **`identity.go`** - Identity and Access Management types (User, Account, Actor, etc.)
- **`network.go`** - Network and Communication types (Endpoint, DNS, HTTP, TLS, etc.)
- **`security.go`** - Security and Threat types (Attack, Malware, ThreatActor, etc.)
- **`vulnerability.go`** - Vulnerability and Advisory types (CVE, CVSS, CWE, etc.)
- **`system.go`** - System and Infrastructure types (Device, Process, File, Container, etc.)
- **`compliance.go`** - Compliance and Assessment types (Compliance, Check, Benchmark, etc.)
- **`common.go`** - Common utility types (Metadata, Location, Fingerprint, etc.)
- **`cmd/generator/`** - Code generator for creating the structured types
- **`cmd/example/`** - Example usage demonstrating the package functionality

## Running the Example

```bash
go run cmd/example/main.go
```

## OCSF Schema Version

This package is generated from OCSF schema version 1.6.0. The schema source can be found at:
https://schema.ocsf.io/1.6.0/export/schema

## Generated Structures

The package includes Go structs organized by domain:

- **Identity & Access**: User, Account, Actor, Group, Session, Authorization
- **Network & Communication**: Endpoint, DNS queries/responses, HTTP requests/responses, TLS details
- **Security & Threats**: Attack patterns, Malware, Threat actors, Findings, MITRE ATT&CK techniques
- **Vulnerabilities**: CVE records, CVSS scoring, CWE classifications, Security advisories
- **System & Infrastructure**: Devices, Processes, Files, Containers, Cloud resources, Applications
- **Compliance & Assessment**: Compliance frameworks, Security checks, Benchmarks (CIS, etc.)
- **Common Utilities**: Metadata, Locations, Fingerprints, Certificates, and shared data types

## License

This package is part of the Veridian Sky product line.

## Contributing

This package uses generated code organized into logical groups. To update the OCSF schema version or modify the generation process:

1. Update the schema: Download a new schema.json from OCSF
2. Modify the generator: Edit `cmd/generator/generate_multi.go` to adjust type groupings
3. Regenerate types: Run `go run cmd/generator/generate_multi.go`
4. Test compilation: Run `go build .` to ensure everything compiles correctly
