package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/veridian-sky/gocsf/classes"
	"github.com/veridian-sky/gocsf/objects"
)

func main() {
	fmt.Println("=== OCSF Go Package Example ===\n")

	// Example 1: Working with Objects
	fmt.Println("Example 1: Creating OCSF Objects")
	fmt.Println("----------------------------------")

	user := &objects.User{
		Name:      "john.doe",
		FullName:  "John Doe",
		EmailAddr: "john.doe@example.com",
		Uid:       "user-123",
	}

	account := &objects.Account{
		Name: "Production Account",
		Uid:  "acc-456",
		Type: "Cloud",
	}

	user.Account = account

	actor := &objects.Actor{
		User:    user,
		AppName: "MyApp",
		AppUID:  "app-789",
	}

	actorJSON, _ := json.MarshalIndent(actor, "", "  ")
	fmt.Println(string(actorJSON))
	fmt.Println()

	// Example 2: Creating an Event Class
	fmt.Println("Example 2: Creating an Account Change Event")
	fmt.Println("--------------------------------------------")

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
		Severity:    "Informational",
		SeverityID:  1,
	}

	eventJSON, _ := json.MarshalIndent(event, "", "  ")
	fmt.Println(string(eventJSON))
	fmt.Println()

	// Example 3: Creating a Process Activity Event
	fmt.Println("Example 3: Creating a Process Activity Event")
	fmt.Println("---------------------------------------------")

	processEvent := &classes.ProcessActivity{
		ActivityID:   1, // Launch
		ActivityName: "Launch",
		Time:         time.Now().Unix(),
		Process: &objects.Process{
			Name:    "python3",
			Pid:     12345,
			CmdLine: "python3 /opt/app/main.py",
			File: &objects.File{
				Name: "python3",
				Path: "/usr/bin/python3",
			},
		},
		Actor: &objects.Actor{
			User: &objects.User{
				Name: "appuser",
				Uid:  "1000",
			},
		},
		Device: &objects.Device{
			Hostname: "app-server-01",
			Uid:      "device-123",
			Type:     "Server",
		},
		CategoryUID: 1,    // System Activity
		ClassUID:    1007, // Process Activity
		Severity:    "Informational",
		SeverityID:  1,
	}

	processEventJSON, _ := json.MarshalIndent(processEvent, "", "  ")
	fmt.Println(string(processEventJSON))
}
