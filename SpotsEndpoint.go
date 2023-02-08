package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "errors"
    "net/http"
)

func getSpots(context *gin.Context){
    context.IndentedJSON(http.StatusOK, spots)
}

func getSpotById(id string) (*spot, error) {
    for index, oneSpot := range spots {
        if oneSpot.Id == id {
            return &spots[index], nil
        }
    }

    s := fmt.Sprintf("Spot not found! (ID = %c)", id)
    return nil, errors.New(s)
}

func getOneSpot(context *gin.Context){
    id := context.Param("id")
    spot, err := getSpotById(id)
    errorMessage := fmt.Sprintf("Spot not found! (ID = %c)", id)

    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": errorMessage})
        return
    }

    context.IndentedJSON(http.StatusOK, spot)
}

// go run spotsEndpoint.go Spots.go
func main() {
    router :=gin.Default()
    router.GET("/spots", getSpots)
    router.GET("/spots/:id", getOneSpot)
    router.Run("localhost:8080")
}
