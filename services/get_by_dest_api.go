package services

import (
    "encoding/json"
    "flight-details/models"
    "flight-details/utils"
    "fmt"
)

// QueryBuilder handles the construction of Elasticsearch queries
type QueryBuilder struct {
    query map[string]interface{}
}

// NewQueryBuilder creates a new instance of QueryBuilder
func NewQueryBuilder() *QueryBuilder {
    return &QueryBuilder{
        query: map[string]interface{}{
            "query": map[string]interface{}{
                "bool": map[string]interface{}{
                    "must": make([]map[string]interface{}, 0),
                },
            },
        },
    }
}

// AddMatchQuery adds a match query to the must clause
func (qb *QueryBuilder) AddMatchQuery(field, value string) *QueryBuilder {
    mustClauses := qb.query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{})
    matchQuery := map[string]interface{}{
        "match": map[string]interface{}{
            field: value,
        },
    }
    mustClauses = append(mustClauses, matchQuery)
    qb.query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = mustClauses
    return qb
}


// Build returns the final query as a string
func (qb *QueryBuilder) Build() (string, error) {
    queryJSON, err := json.Marshal(qb.query)
    if err != nil {
        return "", fmt.Errorf("error marshaling query: %v", err)
    }
    return string(queryJSON), nil
}

// SearchResult represents the structure of Elasticsearch response
type SearchResult struct {
    Hits struct {
        Hits []struct {
            Source models.Flight `json:"_source"`
        } `json:"hits"`
    } `json:"hits"`
}

// SearchFlights searches for flights using the modular query builder
func SearchFlights(destination, date string) ([]models.Flight, error) {
    // Create a new query builder
    queryBuilder := NewQueryBuilder()
    
    // Build the query using method chaining
    query, err := queryBuilder.
        AddMatchQuery("DestCityName", destination).
        AddMatchQuery("timestamp", date).
        Build()
    
    if err != nil {
        return nil, fmt.Errorf("error building query: %v", err)
    }

    // Execute the search
    resp, err := utils.SearchElasticsearch("kibana_sample_data_flights", query)
    if err != nil {
        return nil, err
    }

    // Parse the response
    var results SearchResult
    err = json.Unmarshal([]byte(resp), &results)
    if err != nil {
        return nil, fmt.Errorf("error unmarshaling response: %v", err)
    }

    // Extract flights from results
    flights := make([]models.Flight, len(results.Hits.Hits))
    for i, hit := range results.Hits.Hits {
        flights[i] = hit.Source
    }

    return flights, nil
}