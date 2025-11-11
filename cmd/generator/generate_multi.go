package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

// OCSF Schema structures
type Schema struct {
	Version string                 `json:"version"`
	Objects map[string]ObjectDef   `json:"objects"`
	Events  map[string]EventDef    `json:"events"`
	Types   map[string]TypeMapping `json:"types"`
}

type ObjectDef struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Caption     string                  `json:"caption"`
	Extends     string                  `json:"extends"`
	Attributes  map[string]AttributeDef `json:"attributes"`
	References  []Reference             `json:"references"`
	Constraints map[string]interface{}  `json:"constraints"`
}

type EventDef struct {
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Caption     string                  `json:"caption"`
	Category    string                  `json:"category"`
	UID         int                     `json:"uid"`
	Extends     string                  `json:"extends"`
	Attributes  map[string]AttributeDef `json:"attributes"`
	Profiles    []string                `json:"profiles"`
}

type AttributeDef struct {
	Type        string               `json:"type"`
	TypeName    string               `json:"type_name"`
	Description string               `json:"description"`
	Caption     string               `json:"caption"`
	Requirement string               `json:"requirement"`
	IsArray     bool                 `json:"is_array"`
	ObjectType  string               `json:"object_type"`
	ObjectName  string               `json:"object_name"`
	Enum        map[string]EnumValue `json:"enum"`
	Observable  int                  `json:"observable"`
	Sibling     string               `json:"sibling"`
	Profile     string               `json:"profile"`
	Group       string               `json:"group"`
	References  []Reference          `json:"references"`
}

type EnumValue struct {
	Caption     string `json:"caption"`
	Description string `json:"description"`
}

type Reference struct {
	Description string `json:"description"`
	URL         string `json:"url"`
}

type TypeMapping struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Caption     string `json:"caption"`
}

// Type groups for organizing the output
var typeGroups = map[string][]string{
	"identity": {
		"user", "account", "actor", "group", "organization", "session", "idp",
		"authorization", "auth_factor", "authentication_token", "ldap_person",
		"scim", "sso", "programmatic_credential",
	},
	"network": {
		"network_endpoint", "network_interface", "network_traffic", "network_connection_info",
		"network_proxy", "endpoint", "endpoint_connection", "dns_query", "dns_answer",
		"http_request", "http_response", "http_header", "http_cookie", "url", "api",
		"tls", "tls_extension", "dce_rpc", "autonomous_system", "rpc_interface",
		"domain_contact",
	},
	"security": {
		"attack", "malware", "threat_actor", "campaign", "finding", "finding_info",
		"tactic", "technique", "sub_technique", "kill_chain_phase", "mitigation",
		"d3fend", "d3f_tactic", "d3f_technique", "enrichment", "reputation",
		"osint", "evidence", "observables", "anomaly", "anomaly_analysis",
	},
	"vulnerability": {
		"vulnerability", "cve", "cvss", "cwe", "advisory", "affected_code",
		"affected_package", "epss", "kb_article", "remediation",
	},
	"system": {
		"device", "device_hw_info", "process", "file", "container", "cloud",
		"database", "databucket", "application", "service", "agent", "product",
		"module", "kernel", "kernel_driver", "os", "display", "keyboard_info",
		"peripheral_device", "load_balancer", "startup_item", "web_resource",
		"win/reg_key", "win/reg_value", "win/win_resource", "win/win_service",
	},
	"compliance": {
		"compliance", "assessment", "check", "cis_benchmark", "cis_benchmark_result",
		"cis_control", "cis_csc", "data_classification", "data_security", "policy",
		"baseline", "security_state", "scan", "malware_scan_info",
	},
	"common": {
		"key_value_object", "metadata", "location", "timespan", "fingerprint",
		"certificate", "digital_signature", "encryption_details", "hassh",
		"ja4_fingerprint", "san", "whois", "email", "email_auth", "extension",
		"feature", "image", "logger", "long_string", "metric", "node", "object",
		"observation", "occurrence_details", "request", "response", "rule", "table",
		"trace", "transformation_info", "vendor_attributes", "package", "sbom",
		"software_component", "span", "edge", "graph", "job", "script", "ticket",
		"trait", "environment_variable", "query_info", "query_evidence",
		"related_event", "discovery_details", "managed_entity", "resource_details",
		"port_info", "firewall_rule", "analytic", "analysis_target", "classifier_details",
		"additional_restriction", "access_analysis_result", "permission_analysis_result",
		"aircraft", "unmanned_aerial_system", "unmanned_system_operating_area",
		"identity_activity_metrics", "process_entity", "observable",
	},
}

// fetchSchemaFromAPI downloads the OCSF schema from the specified version
func fetchSchemaFromAPI(version string) ([]byte, error) {
	url := fmt.Sprintf("https://schema.ocsf.io/%s/export/schema", version)

	fmt.Printf("Fetching OCSF schema from: %s\n", url)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch schema: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("Successfully fetched schema (%d bytes)\n", len(data))
	return data, nil
}

// loadSchema loads the schema either from API or local file
func loadSchema(version string, useLocal bool) (*Schema, error) {
	var data []byte
	var err error

	if useLocal {
		fmt.Println("Loading schema from local schema.json file...")
		data, err = ioutil.ReadFile("schema.json")
		if err != nil {
			return nil, fmt.Errorf("error reading local schema file: %w", err)
		}
	} else {
		data, err = fetchSchemaFromAPI(version)
		if err != nil {
			return nil, err
		}

		// Save the fetched schema to local file for future use
		err = ioutil.WriteFile("schema.json", data, 0644)
		if err != nil {
			fmt.Printf("Warning: Could not save schema to local file: %v\n", err)
		} else {
			fmt.Println("Schema saved to schema.json for future local use")
		}
	}

	var schema Schema
	err = json.Unmarshal(data, &schema)
	if err != nil {
		return nil, fmt.Errorf("error parsing schema JSON: %w", err)
	}

	return &schema, nil
}

func main() {
	// Parse command line arguments
	version := flag.String("version", "1.6.0", "OCSF schema version to fetch (e.g., 1.6.0, 1.3.0)")
	local := flag.Bool("local", false, "Use local schema.json file instead of fetching from API")
	help := flag.Bool("help", false, "Show usage information")

	flag.Parse()

	if *help {
		fmt.Println("OCSF Go Type Generator")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Printf("  %s [options]\n", os.Args[0])
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  # Generate types from OCSF v1.6.0 (default)")
		fmt.Printf("  %s\n", os.Args[0])
		fmt.Println()
		fmt.Println("  # Generate types from OCSF v1.3.0")
		fmt.Printf("  %s -version=1.3.0\n", os.Args[0])
		fmt.Println()
		fmt.Println("  # Use local schema.json file")
		fmt.Printf("  %s -local\n", os.Args[0])
		return
	}

	// Load the schema
	schema, err := loadSchema(*version, *local)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Loaded OCSF schema version: %s\n", schema.Version)

	// Generate types in separate files
	generateTypeFiles(schema)

	fmt.Println("Generated Go structs in multiple logical files successfully!")
}

func generateTypeFiles(schema *Schema) {
	// Create a map to track which objects are assigned to which files
	objectToFile := make(map[string]string)

	for fileName, objectList := range typeGroups {
		for _, objName := range objectList {
			objectToFile[objName] = fileName
		}
	}

	// Generate each type file
	for fileName, objectList := range typeGroups {
		generateTypeFile(schema, fileName, objectList, objectToFile)
	}
}

func generateTypeFile(schema *Schema, fileName string, objectList []string, objectToFile map[string]string) {
	var sb strings.Builder

	// Package header
	sb.WriteString("// Code generated from OCSF schema. DO NOT EDIT.\n")
	sb.WriteString("package gocsf\n\n")

	// Sort object names for consistent output
	var objectNames []string
	for _, objName := range objectList {
		if _, exists := schema.Objects[objName]; exists {
			objectNames = append(objectNames, objName)
		}
	}
	sort.Strings(objectNames)

	// Check if we need imports by examining the types
	needsTime := false
	needsJSON := false

	for _, name := range objectNames {
		obj := schema.Objects[name]
		for _, attr := range obj.Attributes {
			if strings.Contains(attr.Type, "timestamp_t") || strings.Contains(attr.Type, "datetime_t") {
				needsTime = true
			}
			if strings.Contains(attr.Type, "json_t") {
				needsJSON = true
			}
		}
	}

	// Add imports only if needed
	if needsTime || needsJSON {
		sb.WriteString("import (\n")
		if needsJSON {
			sb.WriteString("\t\"encoding/json\"\n")
		}
		if needsTime {
			sb.WriteString("\t\"time\"\n")
		}
		sb.WriteString(")\n\n")
	}

	// Add category comment
	categoryName := strings.Title(fileName)
	sb.WriteString(fmt.Sprintf("// %s Types\n\n", categoryName))

	// Generate struct for each object in this category
	for _, name := range objectNames {
		obj := schema.Objects[name]
		generateObjectStruct(&sb, name, &obj, objectToFile)
	}

	// Only write file if it has content beyond header
	content := sb.String()
	if strings.Contains(content, "type ") {
		err := os.WriteFile(fileName+".go", []byte(content), 0644)
		if err != nil {
			log.Fatal("Error writing", fileName+".go:", err)
		}
		fmt.Printf("Generated %s.go with %d types\n", fileName, len(objectNames))
	}
}

func generateObjectStruct(sb *strings.Builder, name string, obj *ObjectDef, objectToFile map[string]string) {
	structName := toGoStructName(name)

	// Add comment with description
	if obj.Description != "" {
		sb.WriteString(fmt.Sprintf("// %s %s\n", structName, obj.Description))
	}

	// Start struct definition
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))

	// Sort attributes for consistent output
	var attrNames []string
	for attrName := range obj.Attributes {
		attrNames = append(attrNames, attrName)
	}
	sort.Strings(attrNames)

	// Generate fields
	for _, attrName := range attrNames {
		attr := obj.Attributes[attrName]
		generateField(sb, attrName, &attr, structName)
	}

	sb.WriteString("}\n\n")
}

func generateField(sb *strings.Builder, fieldName string, attr *AttributeDef, parentStruct string) {
	goFieldName := toGoFieldName(fieldName)
	goType := mapOCSFTypeToGo(attr, parentStruct)

	// Add field comment
	if attr.Description != "" {
		sb.WriteString(fmt.Sprintf("\t// %s\n", attr.Description))
	}

	// Generate field with JSON tag
	jsonTag := fmt.Sprintf("`json:\"%s", fieldName)
	if attr.Requirement == "optional" {
		jsonTag += ",omitempty"
	}
	jsonTag += "\"`"

	sb.WriteString(fmt.Sprintf("\t%s %s %s\n", goFieldName, goType, jsonTag))
}

func mapOCSFTypeToGo(attr *AttributeDef, parentStruct string) string {
	baseType := ""

	switch attr.Type {
	case "string_t", "hostname_t", "file_path_t", "file_name_t", "email_t", "url_t", "ip_t", "subnet_t", "mac_t", "md5_t", "sha1_t", "sha256_t", "sha512_t", "file_hash_t", "uuid_t":
		baseType = "string"
	case "integer_t", "port_t":
		baseType = "int"
	case "long_t":
		baseType = "int64"
	case "float_t":
		baseType = "float64"
	case "boolean_t":
		baseType = "bool"
	case "timestamp_t":
		baseType = "time.Time"
	case "datetime_t":
		baseType = "time.Time"
	case "json_t":
		baseType = "json.RawMessage"
	case "object_t":
		if attr.ObjectType != "" {
			targetType := toGoStructName(attr.ObjectType)
			// Use pointers for all object types to prevent potential recursion issues
			// This is safer and allows for mutual references between types
			baseType = "*" + targetType
		} else {
			baseType = "interface{}"
		}
	default:
		baseType = "interface{}"
	}

	if attr.IsArray {
		return "[]" + baseType
	}

	return baseType
}

func toGoStructName(name string) string {
	return toCamelCase(name, true)
}

func toGoFieldName(name string) string {
	return toCamelCase(name, true)
}

func toCamelCase(s string, upperFirst bool) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == '_' || r == '-' || r == '/' || r == ' '
	})

	var result strings.Builder
	for i, word := range words {
		if word == "" {
			continue
		}

		if i == 0 && !upperFirst {
			result.WriteString(strings.ToLower(word))
		} else {
			result.WriteString(strings.ToUpper(string(word[0])))
			if len(word) > 1 {
				result.WriteString(strings.ToLower(word[1:]))
			}
		}
	}
	return result.String()
}
