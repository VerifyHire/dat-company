// tests/company_test.go

package tests

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "company-service/controllers"
    "company-service/models"
    "github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
    newCompany := models.Company{
        Name:        "Test Company",
        Description: "A test company",
    }

    // Convert the company struct to JSON
    jsonData, _ := json.Marshal(newCompany)

    req, err := http.NewRequest("POST", "/companies", json.NewDecoder(bytes.NewReader(jsonData)))
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()

    // Call the handler
    handler := http.HandlerFunc(controllers.CreateCompany)
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect
    assert.Equal(t, http.StatusCreated, rr.Code)

    var company models.Company
    err = json.NewDecoder(rr.Body).Decode(&company)
    if err != nil {
        t.Fatalf("could not decode response body: %v", err)
    }

    // Assert the company data is correct
    assert.Equal(t, "Test Company", company.Name)
    assert.Equal(t, "A test company", company.Description)
}

func TestGetCompanies(t *testing.T) {
    req, err := http.NewRequest("GET", "/companies", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(controllers.GetCompanies)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var companies []models.Company
    err = json.NewDecoder(rr.Body).Decode(&companies)
    if err != nil {
        t.Fatalf("could not decode response body: %v", err)
    }

    // Assert that the list of companies is returned
    assert.NotEmpty(t, companies)
}

func TestGetCompanyByID(t *testing.T) {
    // Create a company first
    newCompany := models.Company{
        Name:        "Test Company",
        Description: "A test company",
    }

    jsonData, _ := json.Marshal(newCompany)
    req, err := http.NewRequest("POST", "/companies", json.NewDecoder(bytes.NewReader(jsonData)))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.CreateCompany)
    handler.ServeHTTP(rr, req)

    var createdCompany models.Company
    err = json.NewDecoder(rr.Body).Decode(&createdCompany)
    if err != nil {
        t.Fatal(err)
    }

    // Test retrieving the company by ID
    req, err = http.NewRequest("GET", "/companies/"+createdCompany.ID.Hex(), nil)
    if err != nil {
        t.Fatal(err)
    }

    rr = httptest.NewRecorder()
    handler = http.HandlerFunc(controllers.GetCompanyByID)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var company models.Company
    err = json.NewDecoder(rr.Body).Decode(&company)
    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, createdCompany.ID, company.ID)
}

func TestUpdateCompany(t *testing.T) {
    // Create a new company first
    newCompany := models.Company{
        Name:        "Test Company",
        Description: "A test company",
    }

    jsonData, _ := json.Marshal(newCompany)
    req, err := http.NewRequest("POST", "/companies", json.NewDecoder(bytes.NewReader(jsonData)))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.CreateCompany)
    handler.ServeHTTP(rr, req)

    var createdCompany models.Company
    err = json.NewDecoder(rr.Body).Decode(&createdCompany)
    if err != nil {
        t.Fatal(err)
    }

    // Now update the company
    updatedCompany := models.Company{
        Name:        "Updated Test Company",
        Description: "An updated test company",
    }

    jsonData, _ = json.Marshal(updatedCompany)
    req, err = http.NewRequest("PUT", "/companies/"+createdCompany.ID.Hex(), json.NewDecoder(bytes.NewReader(jsonData)))
    if err != nil {
        t.Fatal(err)
    }

    rr = httptest.NewRecorder()
    handler = http.HandlerFunc(controllers.UpdateCompany)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var company models.Company
    err = json.NewDecoder(rr.Body).Decode(&company)
    if err != nil {
        t.Fatal(err)
    }

    assert.Equal(t, "Updated Test Company", company.Name)
    assert.Equal(t, "An updated test company", company.Description)
}

func TestDeleteCompany(t *testing.T) {
    // Create a new company first
    newCompany := models.Company{
        Name:        "Test Company",
        Description: "A test company",
    }

    jsonData, _ := json.Marshal(newCompany)
    req, err := http.NewRequest("POST", "/companies", json.NewDecoder(bytes.NewReader(jsonData)))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controllers.CreateCompany)
    handler.ServeHTTP(rr, req)

    var createdCompany models.Company
    err = json.NewDecoder(rr.Body).Decode(&createdCompany)
    if err != nil {
        t.Fatal(err)
    }

    // Now delete the company
    req, err = http.NewRequest("DELETE", "/companies/"+createdCompany.ID.Hex(), nil)
    if err != nil {
        t.Fatal(err)
    }

    rr = httptest.NewRecorder()
    handler = http.HandlerFunc(controllers.DeleteCompany)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    // Check that the company no longer exists
    req, err = http.NewRequest("GET", "/companies/"+createdCompany.ID.Hex(), nil)
    if err != nil {
        t.Fatal(err)
    }

    rr = httptest.NewRecorder()
    handler = http.HandlerFunc(controllers.GetCompanyByID)
    handler.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusNotFound, rr.Code)
}
