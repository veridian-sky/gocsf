# Go OCSF Package

A Go package for working with the Open Cybersecurity Schema Framework (OCSF) v1.6.0.

## Overview

This package provides Go struct definitions for all OCSF objects, classes, and types, automatically generated from the official OCSF schema. It enables easy serialization and deserialization of OCSF-compliant cybersecurity event data in Go applications.

## Installation

```bash
go get github.com/veridian-sky/gocsf
```

## Package Structure

The OCSF schema has been converted into a clean Go package structure with proper separation of concerns:

```
gocsf/
├── objects/               # 167 OCSF object definitions
│   ├── actor.go
│   ├── user.go
│   ├── device.go
│   ├── process.go
│   ├── file.go
│   └── ... (and 162 more)
├── classes/               # 82 OCSF event class definitions  
│   ├── account_change.go
│   ├── file_activity.go
│   ├── process_activity.go
│   ├── network_activity.go
│   └── ... (and 78 more)
├── types/                 # OCSF type definitions
│   └── types.go
├── base_event/            # Base event attributes
│   └── base_event.go
├── dictionary_attributes/ # Complete attribute dictionary
│   └── dictionary_attributes.go
└── cmd/generator/         # Code generator
    └── main.go
```

Each package is self-contained:
- **`objects`** - Core OCSF object definitions (e.g., User, Device, Process, File, Network Endpoint)
- **`classes`** - OCSF event class definitions (e.g., AccountChange, FileActivity, NetworkActivity)
- **`types`** - Basic OCSF type definitions
- **`base_event`** - Base event attributes common to all OCSF events
- **`dictionary_attributes`** - Dictionary of all possible OCSF attributes

## Usage

### Working with Objects

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/veridian-sky/gocsf/objects"
)

func main() {
    // Create a User object
    user := &objects.User{
        Name:      "john.doe",
        FullName:  "John Doe",
        EmailAddr: "john.doe@example.com",
        Uid:       "user-123",
    }

    // Create an Account object  
    account := &objects.Account{
        Name: "Production Account",
        Uid:  "acc-456",
        Type: "Cloud",
    }

    // Link the account to the user
    user.Account = account

    // Create an Actor object
    actor := &objects.Actor{
        User:    user,
        AppName: "MyApp",
        AppUID:  "app-789",
    }

    // Serialize to JSON
    data, _ := json.MarshalIndent(actor, "", "  ")
    fmt.Println(string(data))
}
```

### Working with Event Classes

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
    "github.com/veridian-sky/gocsf/classes"
    "github.com/veridian-sky/gocsf/objects"
)

func main() {
    // Create an Account Change event
    event := &classes.AccountChange{
        ActivityID:   1, // Create
        ActivityName: "Create",
        Time:         time.Now().Unix(),
        Actor: &objects.Actor{
            User: &objects.User{
                Name: "admin",
                Uid:  "admin-123",
            },
        },
        User: &objects.User{
            Name: "newuser",
            Uid:  "user-456",
        },
        CategoryUID: 3,    // IAM
        ClassUID:    3001, // Account Change
    }

    // Serialize to JSON
    data, _ := json.MarshalIndent(event, "", "  ")
    fmt.Println(string(data))
}
```

## Features

- **Complete OCSF Coverage**: All OCSF v1.6.0 objects, classes, and types are supported
- **Type Safety**: Strongly typed Go structs with proper JSON serialization tags  
- **Organized Structure**: Separate packages for objects, classes, types, etc.
- **Generated Code**: Automatically generated from the official OCSF schema to ensure accuracy
- **Minimal Dependencies**: Only uses Go standard library packages

## Regenerating from Schema

The package includes a code generator that creates Go types from the OCSF schema:

```bash
# Generate from the local ocsf-schema.json file (default)
go run cmd/generator/main.go -schema=ocsf-schema.json -output=.

# Specify different paths
go run cmd/generator/main.go -schema=/path/to/schema.json -output=/path/to/output
```

The generator will:
1. Parse the OCSF schema JSON file
2. Generate Go packages for objects, classes, types, base_event, and dictionary_attributes
3. Handle imports automatically and use pointers to prevent recursive type issues
4. Format all generated code with gofmt

## OCSF Schema Version

This package is generated from OCSF schema version 1.6.0. The schema is included in this repository as `ocsf-schema.json`.

## Generated Structures

The package includes Go structs organized into the following packages:

### Objects Package (`objects/`)
Contains 167 object definitions including:
- **Identity & Access**: User, Account, Actor, Group, Session, Authorization
- **Network**: Endpoint, NetworkInterface, NetworkProxy, AutonomousSystem  
- **Security**: Attack, Malware, ThreatActor, Finding, Certificate, Fingerprint
- **System**: Device, Process, File, Container, Image, Application
- **Cloud**: Cloud, Service, Product, LoadBalancer
- **Communication**: Email, DNS, HTTP, TLS, RPC interfaces
- **Vulnerabilities**: CVE, CVSS, CWE, EPSS, Advisory
- **Compliance**: Compliance, CISBenchmark, CISControl, Assessment
- **Common**: Metadata, Location, Enrichment, Observable, Reputation

### Classes Package (`classes/`)
Contains 82 event class definitions including:
- **IAM**: AccountChange, Authentication, AuthorizeSession, GroupManagement, UserAccess
- **System Activity**: ProcessActivity, FileActivity, KernelActivity, ModuleActivity
- **Network Activity**: NetworkActivity, DNSActivity, HTTPActivity, SMBActivity, SSHActivity
- **Security**: DetectionFinding, IncidentFinding, VulnerabilityFinding, ComplianceFinding
- **Cloud**: DatastoreActivity, APIActivity
- **Application**: ApplicationLifecycle, WebResourceAccessActivity

### Other Packages
- **`types/`**: Basic OCSF type definitions  
- **`base_event/`**: Base event attributes inherited by all event classes
- **`dictionary_attributes/`**: Complete dictionary of all OCSF attributes

## Example Import Paths

```go
import (
    "github.com/veridian-sky/gocsf/objects"
    "github.com/veridian-sky/gocsf/classes"  
    "github.com/veridian-sky/gocsf/types"
    "github.com/veridian-sky/gocsf/base_event"
)
```

## License

This package is part of the Veridian Sky product line.

## Contributing

This package uses generated code. To update the OCSF schema or modify the generation process:

1. **Update the schema**: Replace `ocsf-schema.json` with a new version from [OCSF](https://schema.ocsf.io/)
2. **Regenerate code**: Run `go run cmd/generator/main.go -schema=ocsf-schema.json -output=.`
3. **Test compilation**: Run `go build ./...` to ensure everything compiles correctly
4. **Update tests**: Add or update tests as needed for new structures
