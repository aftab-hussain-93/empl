package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aftab-hussain-93/empl/client"
	"github.com/aftab-hussain-93/empl/internal/service"
	fault "github.com/aftab-hussain-93/empl/pkg/err"
)

func main() {
	appHost := os.Getenv("APP_HOST")
	if appHost == "" {
		appHost = "app" // docker compose name
	}
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}
	url := fmt.Sprintf("http://%s:%s/api", appHost, appPort)
	http_client := client.New()
	http_client.SetBaseURL(url)

	appStartUpWait(url + "/status")

	ctx := context.Background()

	if _, err := testCreateEmployeeSuccess(ctx, http_client); err != nil {
		log.Println("create employee test failed", err.Error())
	} else {
		log.Println("create employee test passed")
	}

	if err := testCreateEmployeeInvalidName(ctx, http_client); err != nil {
		log.Println("create employee w/ invalid name test failed", err.Error())
	} else {
		log.Println("create employee w/ invalid name test passed")
	}
}

func appStartUpWait(statusCheckURL string) {
	log.Println("waiting for app to startup")
	for {
		resp, err := http.Get(statusCheckURL)
		if err == nil {
			log.Println("got status", resp.StatusCode)
			defer resp.Body.Close()
			if resp.StatusCode == 200 {
				log.Println("api is up")
				return
			}
		} else {
			log.Println("errored", err.Error())
		}

		time.Sleep(200 * time.Millisecond)
		log.Println("finding")
	}
}

func testCreateEmployeeSuccess(ctx context.Context, http_client *client.Client) (*service.Employee, error) {
	emp := map[string]any{
		"name":     "user name",
		"position": "manager",
		"salary":   50,
	}

	status, resp, err := http_client.CreateEmployee(ctx, emp)
	if err != nil {
		return nil, fmt.Errorf("error sending create employee request, err %w", err)
	}
	if status != http.StatusCreated {
		return nil, fmt.Errorf("error invalid status code, got %v, expected 201", status)
	}

	employee, ok := resp.(*service.Employee)
	if !ok {
		log.Printf("error invalid response \n%+v\n", resp)
		return nil, fmt.Errorf("error invalid response type, expected employee return type")
	}
	if employee.ID <= 0 {
		return nil, fmt.Errorf("error object not created successfully, id not found")
	}

	return employee, nil
}

func testCreateEmployeeInvalidName(ctx context.Context, http_client *client.Client) error {
	request := map[string]any{
		"position": "manager",
		"salary":   50,
	}

	status, resp, err := http_client.CreateEmployee(ctx, request)
	if err != nil {
		return fmt.Errorf("error sending create employee request, err %w", err)
	}
	if status != http.StatusBadRequest {
		return fmt.Errorf("error invalid status code, got %v, expected 400", status)
	}

	flt, ok := resp.(*fault.HTTPError)
	if !ok {
		log.Printf("error invalid response \n%+v\n", resp)
		return fmt.Errorf("error invalid response type, expected error")
	}
	if flt.Error.Code != fault.ErrBadRequest {
		return fmt.Errorf("error response bad request expected")
	}

	return nil
}
