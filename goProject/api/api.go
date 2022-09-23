//package api
package api

import (
	"apidgraph/api/models"
	"context"
	"encoding/json"
	"github.com/dgraph-io/dgo/v200"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

//apidgraph: dependency injections
type apidgraph struct {
	dg  *dgo.Dgraph //Dgo.Dgraph.Nextxn
	ctx context.Context
}

// constructor
func NewApidgraph(cliente *dgo.Dgraph, ctx context.Context) *apidgraph {
	return &apidgraph{dg: cliente, ctx: ctx}
}

// GetAllPrograms: Show all programs objects
func (a *apidgraph) GetAllPrograms(w http.ResponseWriter, r *http.Request) {
	//for the browser interprets the response as json
	w.Header().Set("Content-Type", "application/json")
	txn := a.dg.NewTxn()     //to execute the queries in the database we create the transaction
	defer txn.Discard(a.ctx) //close routines when done
	const q = `
	{
		response(func: has(Program.name)) 
		  {
				uid	
				name: Program.name
				num1: Program.num1
				num2: Program.num2
				operation: Program.operation {
					uid 					
					add: Operation.add
			  		mult:Operation.mult
			  		sub: Operation.sub
				}
			
		  }
		}
	`

	queryResults := models.QueryResults{}
	response := models.Response{}
	//the query is made
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure query"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := json.Unmarshal(resp.GetJson(), &queryResults); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "Error validating struct"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Message = "ok"
	response.StatusCode = http.StatusOK
	response.Result = queryResults.Response
	// respond to the client
	json.NewEncoder(w).Encode(response)
}

func (a *apidgraph) GetProgram(w http.ResponseWriter, r *http.Request) {
	//the browser interprets the response as json
	w.Header().Set("Content-Type", "application/json")
	txn := a.dg.NewTxn() //para ejecutar las consultas en la base de datos
	defer txn.Discard(a.ctx)

	//the route parameter is extracted from r
	uid := chi.URLParam(r, "uid")
	variables := map[string]string{"$uid": uid} // dictionary en python
	const q = `
	query program($uid: string) {
		response(func: uid($uid)) 
		  {
        		uid
				name: Program.name
				num1: Program.num1
				num2: Program.num2
				operation: Program.operation { 
          			uid
					add: Operation.add
			  		mult:Operation.mult
			  		sub: Operation.sub
				}
			
		  }
		}
	`

	queryResults := models.QueryResults{}
	response := models.Response{}
	//the query is made
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure query"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := json.Unmarshal(resp.GetJson(), &queryResults); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure response"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "ok"
	response.StatusCode = http.StatusOK
	response.Result = queryResults.Response
	// respond to the client
	json.NewEncoder(w).Encode(response)
}

func (a *apidgraph) GetOneProgram(w http.ResponseWriter, r *http.Request) {
	//the browser interprets the response as json
	w.Header().Set("Content-Type", "application/json")
	txn := a.dg.NewTxn() //para ejecutar las consultas en la base de datos
	defer txn.Discard(a.ctx)

	//the route parameter is extracted from r
	name := chi.URLParam(r, "name")
	variables := map[string]string{"$name": name} // dictionary en python
	const q = `
	query program($name:string) {
		response(func: eq(Program.name, $name)) {
		  uid
		  name: Program.name
		  num1: Program.num1
		  num2: Program.num2
		  operation: Program.operation {
		   uid
		   add: Operation.add
		   mult:Operation.mult
		   sub: Operation.sub
		  }
	  
		}
	  }
	`
	queryResults := models.QueryResults{} //a := "" var a string
	response := models.Response{}
	//the query is made
	resp, err := txn.QueryWithVars(context.Background(), q, variables)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure query"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := json.Unmarshal(resp.GetJson(), &queryResults); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure response"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Message = "ok"
	response.StatusCode = http.StatusOK
	response.Result = queryResults.Response
	// respond to the client
	json.NewEncoder(w).Encode(response)
}

func (a *apidgraph) CreateProgram(w http.ResponseWriter, r *http.Request) {
	//for the browser interprete la respuesta como json
	w.Header().Set("Content-Type", "application/json")
	txn := a.dg.NewTxn() //para ejecutar las consultas en la base de datos
	defer txn.Discard(a.ctx)

	//the route parameter is extracted
	//mutationResult:= models.MutationResult{}//a := "" var a string
	response := models.Response{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure body post"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return

	}
	resp, err := txn.Mutate(context.Background(), &api.Mutation{CommitNow: true, SetJson: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure post mutation"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "ok"
	response.StatusCode = http.StatusOK
	//result delivers 2 uid, the second is chosen
	cont := 0
	for _, v := range resp.Uids {
		if cont == 1 {
			response.Result = models.Program{UID: v}
			break
		}
		cont++
	}

	// respond to the client
	json.NewEncoder(w).Encode(response)
}

func (a *apidgraph) UpdateProgram(w http.ResponseWriter, r *http.Request) {
	//for the browser interprete la respuesta como json
	w.Header().Set("Content-Type", "application/json")
	txn := a.dg.NewTxn() //para ejecutar las consultas en la base de datos
	defer txn.Discard(a.ctx)

	//the route parameter is extracted
	//mutationResult:= models.MutationResult{}//a := "" var a string
	response := models.Response{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure body post"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return

	}
	_, err = txn.Mutate(context.Background(), &api.Mutation{CommitNow: true, SetJson: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure post mutation"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Message = "ok"
	response.StatusCode = http.StatusOK
	// respond to the client
	json.NewEncoder(w).Encode(response)
}

func (a *apidgraph) DeleteProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	txn := a.dg.NewTxn()
	defer txn.Discard(a.ctx)

	response := models.Response{}
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure body Delete"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
		return
	}
	resp, err := txn.Mutate(context.Background(), &api.Mutation{CommitNow: true, SetJson: body})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message = "failure Delete Mutation"
		response.StatusCode = http.StatusInternalServerError
		json.NewEncoder(w).Encode(response)
	}
	response.Message = "ok"
	response.StatusCode = http.StatusOK
	//result delivers 2 uid, the second is chosen
	cont := 0
	for _, v := range resp.Uids {
		if cont == 1 {
			response.Result = models.Program{UID: v}
			break
		}
		cont++
	}

	// respond to the client
	json.NewEncoder(w).Encode(response)
}
