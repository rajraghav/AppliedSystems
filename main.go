package main

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const graphDir = "data/graphs/"

type Graph struct {
	AdjacencyList map[string][]string `json:"adjacency_list"`
}

func main() {
	router := gin.Default()

	router.POST("/graph", postGraph)
	router.GET("/graph/:id/path", getShortestPath)
	router.DELETE("/graph/:id", deleteGraph)

	if err := os.MkdirAll(graphDir, os.ModePerm); err != nil {
		panic(err)
	}

	router.Run(":8080")
}

func postGraph(ctx *gin.Context) {
	var graph Graph
	if err := ctx.ShouldBindJSON(&graph); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error(err)})
		return
	}

	graphID := uuid.New().String()

	graphFile := filepath.Join(graphDir, graphID+".json")

	data, err := json.Marshal(graph)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save graph"})
		return
	}

	if err := os.WriteFile(graphFile, data, os.ModePerm); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save graph"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"graph_id": graphID})
}

func getShortestPath(ctx *gin.Context) {
	graphID := ctx.Param("id")
	start := ctx.Query("start")
	end := ctx.Query("end")

	graphFile := filepath.Join(graphDir, graphID+".json")
	data, err := os.ReadFile(graphFile)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Graph not found"})
		return
	}

	var graph Graph
	if err := json.Unmarshal(data, &graph); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read graph"})
		return
	}


	path := bfsShortestPath(graph.AdjacencyList, start, end)
	if path == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No path found"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"path": path})
	}
}

func deleteGraph(ctx *gin.Context) {
	graphID := ctx.Param("id")
	graphFile := filepath.Join(graphDir, graphID+".json")

	if _, err := os.Stat(graphFile); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Graph not found"})
		return
	}

	if err := os.Remove(graphFile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete graph"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Graph deleted"})
}