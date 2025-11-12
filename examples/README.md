# OCSF Go Package Examples

This directory contains example programs demonstrating how to use the OCSF Go package to create various types of security events.

## Available Examples

### 1. Basic Example (`basic/`)

Demonstrates fundamental usage of OCSF objects and event classes:
- Creating OCSF objects (User, Account, Actor)
- Building an Account Change event
- Creating a Process Activity event

**Run:** 
```bash
cd basic
go run main.go
```

### 2. Compliance Example (`compliance/`)

Comprehensive examples of compliance finding events across multiple frameworks:

#### Example 1: CIS AWS Foundations Benchmark
- Multi-region CloudTrail compliance check
- Resource details with configuration parameters
- Remediation guidance with references

#### Example 2: PCI DSS Password Complexity
- Password policy compliance validation
- Multiple check evaluations
- System device context

#### Example 3: Multi-Framework Encryption
- NIST SP 800-53, HIPAA, and ISO 27001 compliance
- RDS encryption at rest validation
- Critical finding with detailed risk assessment
- Healthcare data (PHI) protection requirements

**Run:**
```bash
cd compliance
go run main.go
```

## Key Features Demonstrated

### Compliance Finding Events
- **Standards Coverage**: CIS Benchmarks, PCI DSS, NIST, HIPAA, ISO 27001
- **Compliance Objects**: Controls, checks, requirements, and standards
- **Status Tracking**: Pass/Fail status with detailed status information
- **Risk Assessment**: Risk levels, scores, and detailed risk explanations
- **Remediation**: Actionable remediation steps with reference documentation
- **Resource Details**: Affected resources with configuration data
- **Control Parameters**: Specific parameter values being evaluated

### Common Patterns
- Creating event metadata (activity, category, class)
- Setting severity levels
- Adding cloud context (provider, region, account)
- Linking related objects (actor, device, resources)
- Including finding information (title, description, type)

## Event Structure

All compliance finding examples follow the OCSF ComplianceFinding class structure:

```go
&classes.ComplianceFinding{
    // Event metadata
    ActivityID:   2,  // Compliance Report
    CategoryUID:  2,  // Findings
    ClassUID:     2003, // Compliance Finding
    
    // Compliance details
    Compliance: &objects.Compliance{
        Control: "...",
        Standards: []string{...},
        Requirements: []string{...},
        Status: "Pass" or "Fail",
        Checks: []*objects.Check{...},
    },
    
    // Finding information
    FindingInfo: &objects.FindingInfo{...},
    
    // Risk assessment
    RiskLevel: "Low|Medium|High|Critical",
    RiskScore: 0-100,
    
    // Remediation guidance
    Remediation: &objects.Remediation{...},
}
```

## JSON Output

All examples output valid OCSF-compliant JSON that can be:
- Sent to SIEM systems
- Stored in security data lakes
- Processed by compliance monitoring tools
- Used in security automation workflows

## Adding Your Own Examples

To create a new example:

1. Create a new directory: `mkdir examples/myexample`
2. Create `main.go` with package main
3. Import required packages:
   ```go
   import (
       "github.com/veridian-sky/gocsf/classes"
       "github.com/veridian-sky/gocsf/objects"
   )
   ```
4. Build your event using OCSF structs
5. Marshal to JSON and output

## Learning Resources

- [OCSF Schema Documentation](https://schema.ocsf.io/)
- [CIS Benchmarks](https://www.cisecurity.org/cis-benchmarks/)
- [NIST SP 800-53](https://csrc.nist.gov/publications/detail/sp/800-53/rev-5/final)
- [PCI DSS](https://www.pcisecuritystandards.org/)
- [HIPAA Security Rule](https://www.hhs.gov/hipaa/for-professionals/security/index.html)
