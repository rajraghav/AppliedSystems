Running the server:
go run .
 



Client using Postman, with following details:

API Endpoints

POST /graph

Request Body:
{
    "adjacency_list": {
        "A": ["B", "C"],
        "B": ["A", "D", "E"],
        "C": ["A", "F"],
        "D": ["B"],
        "E": ["B", "F"],
        "F": ["C", "E"]
    }
}

Response
{
    "graph_id": "some-unique-id"
}


GET /graph/
Get the shortest path between two vertices.

Query Parameters:

start: The starting vertex.
end: The ending vertex.

Response:

{
    "path": ["A", "C", "F"]
}


DELETE /graph/
Delete a graph by its ID.

Response:
{
    "message": "Graph deleted"
}
