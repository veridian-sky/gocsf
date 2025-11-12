package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/veridian-sky/gocsf/classes"
	"github.com/veridian-sky/gocsf/objects"
)

func main() {
	fmt.Println("=== OCSF Compliance Finding Example ===")
	fmt.Println()

	// Example 1: CIS Benchmark Compliance Check
	fmt.Println("Example 1: CIS AWS Foundations Benchmark Compliance Check")
	fmt.Println("----------------------------------------------------------")

	cisComplianceEvent := &classes.ComplianceFinding{
		// Event metadata
		ActivityID:   2, // Compliance Report
		ActivityName: "Compliance Report",
		CategoryUID:  2,    // Findings
		ClassUID:     2003, // Compliance Finding
		Severity:     "Medium",
		SeverityID:   3,
		Time:         time.Now().Unix(),

		// Cloud context
		Cloud: &objects.Cloud{
			Provider: "AWS",
			Region:   "us-east-1",
			Account: &objects.Account{
				Name: "Production",
				Uid:  "123456789012",
			},
		},

		// Compliance details
		Compliance: &objects.Compliance{
			Control: "Ensure CloudTrail is enabled in all regions",
			Standards: []string{
				"CIS AWS Foundations Benchmark v1.4.0",
				"NIST SP 800-53 Rev. 5",
			},
			Requirements: []string{
				"CIS AWS Foundations Benchmark - 2.1",
				"NIST AC-17(2)",
			},
			Status:       "Fail",
			StatusID:     3, // Fail
			StatusDetail: "CloudTrail is not enabled in all regions",
			Checks: []*objects.Check{
				{
					Name:       "Ensure CloudTrail is enabled in all regions",
					Desc:       "AWS CloudTrail is a web service that records AWS API calls for your account and delivers log files to you. The recorded information includes the identity of the API caller, the time of the API call, the source IP address of the API caller, and the request parameters.",
					Uid:        "2.1",
					Version:    "1.4.0",
					Standards:  []string{"CIS AWS Foundations Benchmark v1.4.0"},
					Status:     "Fail",
					StatusID:   3,
					Severity:   "Level 1",
					SeverityID: 3,
				},
			},
			ControlParameters: []*objects.KeyValueObject{
				{
					Name:  "multiRegionTrailEnabled",
					Value: "false",
				},
				{
					Name:  "requiredRegions",
					Value: "us-east-1,us-west-2,eu-west-1",
				},
			},
		},

		// Finding information
		FindingInfo: &objects.FindingInfo{
			Title: "CloudTrail Not Enabled in All Regions",
			Desc:  "CloudTrail multi-region trail is not enabled, which reduces visibility into API activity across all AWS regions.",
			Types: []string{"Configuration Issue"},
			Uid:   "finding-001",
		},

		// Affected resources
		Resources: []*objects.ResourceDetails{
			{
				Name: "Default CloudTrail",
				Type: "AWS::CloudTrail::Trail",
				Uid:  "arn:aws:cloudtrail:us-east-1:123456789012:trail/default",
				Data: map[string]interface{}{
					"IsMultiRegionTrail": false,
					"HomeRegion":         "us-east-1",
				},
			},
		},

		// Remediation guidance
		Remediation: &objects.Remediation{
			Desc: "Enable CloudTrail in all regions by creating a multi-region trail or updating the existing trail configuration.",
			References: []string{
				"https://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-create-and-update-a-trail.html",
			},
		},

		// Risk assessment
		RiskLevel:   "Medium",
		RiskLevelID: 2,
		RiskScore:   65,
		RiskDetails: "Lack of multi-region CloudTrail increases the risk of undetected unauthorized access and makes forensic investigation more difficult.",
	}

	cisJSON, _ := json.MarshalIndent(cisComplianceEvent, "", "  ")
	fmt.Println(string(cisJSON))
	fmt.Println()

	// Example 2: PCI DSS Compliance Check
	fmt.Println("Example 2: PCI DSS Password Complexity Compliance Check")
	fmt.Println("--------------------------------------------------------")

	pciComplianceEvent := &classes.ComplianceFinding{
		// Event metadata
		ActivityID:   2, // Compliance Report
		ActivityName: "Compliance Report",
		CategoryUID:  2,    // Findings
		ClassUID:     2003, // Compliance Finding
		Severity:     "High",
		SeverityID:   4,
		Time:         time.Now().Unix(),

		// Device being assessed
		Device: &objects.Device{
			Name:     "payment-server-01",
			Hostname: "payment-server-01.internal",
			Uid:      "i-0123456789abcdef0",
			Type:     "Server",
			Os: &objects.Os{
				Name:    "Ubuntu Linux",
				Version: "20.04 LTS",
			},
		},

		// Compliance details
		Compliance: &objects.Compliance{
			Control: "Implement password complexity requirements",
			Standards: []string{
				"PCI DSS v3.2.1",
			},
			Requirements: []string{
				"PCI DSS Requirement 8.2.3 - Passwords must have minimum length of 7 characters",
				"PCI DSS Requirement 8.2.3 - Passwords must contain both numeric and alphabetic characters",
			},
			Status:       "Pass",
			StatusID:     1, // Pass
			StatusDetail: "All password complexity requirements are met",
			Checks: []*objects.Check{
				{
					Name:       "Password minimum length",
					Desc:       "Verify that passwords meet minimum length requirements",
					Uid:        "8.2.3.a",
					Standards:  []string{"PCI DSS v3.2.1"},
					Status:     "Pass",
					StatusID:   1,
					Severity:   "High",
					SeverityID: 4,
				},
				{
					Name:       "Password complexity",
					Desc:       "Verify that passwords contain both numeric and alphabetic characters",
					Uid:        "8.2.3.b",
					Standards:  []string{"PCI DSS v3.2.1"},
					Status:     "Pass",
					StatusID:   1,
					Severity:   "High",
					SeverityID: 4,
				},
			},
			ControlParameters: []*objects.KeyValueObject{
				{
					Name:  "minlen",
					Value: "8",
				},
				{
					Name:  "dcredit",
					Value: "-1",
				},
				{
					Name:  "ucredit",
					Value: "-1",
				},
			},
		},

		// Finding information
		FindingInfo: &objects.FindingInfo{
			Title: "Password Complexity Requirements Met",
			Desc:  "System password policies meet PCI DSS requirements for minimum length and complexity.",
			Types: []string{"Compliance Check"},
			Uid:   "finding-002",
		},

		// Risk assessment (low risk since compliant)
		RiskLevel:   "Low",
		RiskLevelID: 1,
		RiskScore:   15,
		RiskDetails: "System is compliant with PCI DSS password requirements.",
	}

	pciJSON, _ := json.MarshalIndent(pciComplianceEvent, "", "  ")
	fmt.Println(string(pciJSON))
	fmt.Println()

	// Example 3: NIST 800-53 Multi-Framework Compliance
	fmt.Println("Example 3: Multi-Framework Encryption Compliance Check")
	fmt.Println("-------------------------------------------------------")

	multiFrameworkEvent := &classes.ComplianceFinding{
		// Event metadata
		ActivityID:   2, // Compliance Report
		ActivityName: "Compliance Report",
		CategoryUID:  2,    // Findings
		ClassUID:     2003, // Compliance Finding
		Severity:     "Critical",
		SeverityID:   5,
		Time:         time.Now().Unix(),

		// Cloud context
		Cloud: &objects.Cloud{
			Provider: "AWS",
			Region:   "us-east-1",
			Account: &objects.Account{
				Name: "Healthcare Production",
				Uid:  "987654321098",
			},
		},

		// Compliance details - Multiple frameworks
		Compliance: &objects.Compliance{
			Control: "Implement encryption for data at rest",
			Standards: []string{
				"NIST SP 800-53 Rev. 5",
				"HIPAA Security Rule",
				"ISO/IEC 27001:2013",
			},
			Requirements: []string{
				"NIST SC-28 - Protection of Information at Rest",
				"HIPAA 164.312(a)(2)(iv) - Encryption and Decryption",
				"ISO/IEC 27001 - A.10.1.1 Policy on the use of cryptographic controls",
			},
			Status:       "Fail",
			StatusID:     3, // Fail
			StatusDetail: "RDS database instance does not have encryption enabled",
			Category:     "Data Security",
			Checks: []*objects.Check{
				{
					Name:       "RDS Encryption at Rest",
					Desc:       "Verify that RDS instances have encryption enabled to protect sensitive data",
					Uid:        "SC-28",
					Standards:  []string{"NIST SP 800-53 Rev. 5", "HIPAA", "ISO/IEC 27001"},
					Status:     "Fail",
					StatusID:   3,
					Severity:   "Critical",
					SeverityID: 5,
				},
			},
		},

		// Finding information
		FindingInfo: &objects.FindingInfo{
			Title: "RDS Database Encryption Not Enabled",
			Desc:  "Production RDS database storing PHI (Protected Health Information) does not have encryption at rest enabled, violating HIPAA, NIST, and ISO 27001 requirements.",
			Types: []string{"Security Misconfiguration", "Compliance Violation"},
			Uid:   "finding-003",
		},

		// Affected resources
		Resources: []*objects.ResourceDetails{
			{
				Name: "patient-records-db",
				Type: "AWS::RDS::DBInstance",
				Uid:  "arn:aws:rds:us-east-1:987654321098:db:patient-records-db",
				Data: map[string]interface{}{
					"Engine":           "postgres",
					"EngineVersion":    "13.7",
					"StorageEncrypted": false,
					"MultiAZ":          true,
					"DBInstanceClass":  "db.r5.large",
				},
			},
		},

		// Remediation guidance
		Remediation: &objects.Remediation{
			Desc: "Enable encryption at rest for the RDS instance. Note: This requires creating a new encrypted instance and migrating data, as encryption cannot be enabled on existing instances.",
			References: []string{
				"https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.html",
				"https://aws.amazon.com/blogs/database/encrypt-an-amazon-rds-for-postgresql-db-instance/",
			},
		},

		// Risk assessment
		RiskLevel:   "Critical",
		RiskLevelID: 4,
		RiskScore:   95,
		RiskDetails: "Unencrypted database containing PHI poses significant risk of data breach and regulatory penalties under HIPAA. Immediate remediation required.",

		// Actor (who ran the compliance check)
		Actor: &objects.Actor{
			User: &objects.User{
				Name: "compliance-scanner",
				Uid:  "scanner-service-001",
				Type: "Service Account",
			},
		},
	}

	multiFrameworkJSON, _ := json.MarshalIndent(multiFrameworkEvent, "", "  ")
	fmt.Println(string(multiFrameworkJSON))
}
