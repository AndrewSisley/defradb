// Copyright 2025 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package action

import (
	"context"
	"strings"
	"time"

	"github.com/sourcenetwork/immutable"
	"github.com/sourcenetwork/testo/action"
	"github.com/stretchr/testify/require"

	acpIdentity "github.com/sourcenetwork/defradb/acp/identity"
	"github.com/sourcenetwork/defradb/client"
	"github.com/sourcenetwork/defradb/internal/db"
	"github.com/sourcenetwork/defradb/tests/clients/cli"
	"github.com/sourcenetwork/defradb/tests/clients/http"
	"github.com/sourcenetwork/defradb/tests/state"
)

// CreateDoc will attempt to create the given document in the given collection
// using the set [MutationType].
type CreateDoc struct {
	stateful

	// NodeID may hold the ID (index) of a node to apply this create to.
	//
	// If a value is not provided the document will be created in all nodes.
	NodeID immutable.Option[int]

	// The identity of this request. Optional.
	//
	// If an Identity is not provided the created document(s) will be public.
	//
	// If an Identity is provided and the collection has a policy, then the
	// created document(s) will be owned by this Identity.
	//
	// Use `ClientIdentity` to create a client identity and `NodeIdentity` to create a node identity.
	// Default value is `NoIdentity()`.
	Identity immutable.Option[state.Identity]

	// Specifies whether the document should be encrypted.
	IsDocEncrypted bool

	// Individual fields of the document to encrypt.
	EncryptedFields []string

	// The collection in which this document should be created.
	CollectionID int

	// The document to create, in JSON string format.
	//
	// If [DocMap] is provided this value will be ignored.
	Doc string // there are 2 for doc save... json, and map - do we want to keep docmap? (maybe only for the embedded go varient?)

	// The document to create, in map format.
	//
	// If this is provided [Doc] will be ignored.
	// DocMap map[string]any // create/save takes a client.Document, int. tests can rely on json constructor - all other constructors just need unit tests

	// Any error expected from the action. Optional.
	//
	// String can be a partial, and the test will pass if an error is returned that
	// contains this string.
	ExpectedError string
}

type CreateDoc_Go_CollectionSave CreateDoc
type CreateDoc_HttpClient_CollectionSave CreateDoc

type CreateDoc_Go_CollectionCreate CreateDoc
type CreateDoc_HttpClient_CollectionCreate CreateDoc

type CreateDoc_Go_GQL CreateDoc
type CreateDoc_HttpClient_GQL CreateDoc

// ...
var _ action.Action = (*CreateDoc_Go_CollectionSave)(nil)
var _ action.Stateful[*state.State] = (*CreateDoc_Go_CollectionSave)(nil)

func (a *CreateDoc_Go_CollectionSave) Execute() {
	nodeIDs, nodes := getNodesWithIDs(a.NodeID, a.s.Nodes)
	for index, node := range nodes {
		nodeID := nodeIDs[index]
		collection := a.s.Nodes[nodeID].Collections[a.CollectionID]

		docs, err := parseCreateDocs(a.Doc, collection)
		if err != nil {
			return nil, err
		}

		txn := getTransaction(a.s, node, immutable.None[int](), a.ExpectedError)
		ctx := makeContextForDocCreate(a.s, db.InitContext(a.s.Ctx, txn), nodeID, a.Identity)

		docIDs := make([]client.DocID, len(docs))
		for i, doc := range docs {
			err := collection.Save(ctx, doc,
				client.CreateDocEncrypted(a.IsDocEncrypted),
				client.CreateDocWithEncryptedFields(a.EncryptedFields),
			)
			if err != nil {
				return nil, err
			}
			docIDs[i] = doc.ID()
		}

	}
}

func (a *CreateDoc_HttpClient_CollectionSave) Execute() {
	panic("todo")
}

func collectionSave(c client.Collection) {

}

func (a *CreateDoc_Go_CollectionCreate) Execute() {
	panic("todo")
}

func (a *CreateDoc_HttpClient_CollectionCreate) Execute() {
	panic("todo")
}

func (a *CreateDoc_Go_GQL) Execute() {
	panic("todo")
}

func (a *CreateDoc_HttpClient_GQL) Execute() {
	panic("todo")
}

func makeDocCreateOptions(action *CreateDoc) []client.DocCreateOption {
	return []client.DocCreateOption{
		client.CreateDocEncrypted(action.IsDocEncrypted),
		client.CreateDocWithEncryptedFields(action.EncryptedFields),
	}
}

type ToGoable interface {
	ToGo() Action
}

type ToHttpClientable interface {
	ToHttpClient() Action
}

type ToHttpRequestable interface {
	ToHttpRequest() Action
}

type ToCLIable interface {
	ToCLI() Action
}

func (a *CreateDoc) ToGo() Action {
	r := (CreateDoc_Go_CollectionSave)(*a)
	return &r
}

func (a *CreateDoc) ToHttpClient() Action {
	panic("todo")
}

func (a *CreateDoc) ToHttpRequest() Action {
	panic("todo")
}

func (a *CreateDoc) ToCLI() Action {
	panic("todo")
}

// parseCreateDocs parses and returns documents from a CreateDoc action.
func parseCreateDocs(doc string, collection client.Collection) ([]*client.Document, error) {
	switch {
	case client.IsJSONArray([]byte(doc)):
		return client.NewDocsFromJSON([]byte(doc), collection.Version())

	default:
		val, err := client.NewDocFromJSON([]byte(doc), collection.Version())
		if err != nil {
			return nil, err
		}
		return []*client.Document{val}, nil
	}
}

func getTransaction(
	s *state.State,
	db client.TxnStore,
	transactionSpecifier immutable.Option[int],
	expectedError string,
) client.Txn {
	if !transactionSpecifier.HasValue() {
		return nil
	}

	transactionID := transactionSpecifier.Value()

	if transactionID >= len(s.Txns) {
		// Extend the txn slice so this txn can fit and be accessed by TransactionId
		s.Txns = append(s.Txns, make([]client.Txn, transactionID-len(s.Txns)+1)...)
	}

	if s.Txns[transactionID] == nil {
		// Create a new transaction if one does not already exist.
		txn, err := db.NewTxn(s.Ctx, false)
		if assertError(s.T, err, expectedError) {
			txn.Discard(s.Ctx)
			return nil
		}

		s.Txns[transactionID] = txn
	}

	return s.Txns[transactionID]
}

func makeContextForDocCreate(s *state.State, ctx context.Context, nodeIndex int, identity immutable.Option[state.Identity]) context.Context {
	ctx = getContextWithIdentity(ctx, s, identity, nodeIndex)
	return ctx
}

// getContextWithIdentity returns a context with the identity for the given reference and node index.
// If the identity does not exist, it will be generated.
// The identity added to the context is prepared for a request, i.e. its [Identity.BearerToken] is set.
func getContextWithIdentity(
	ctx context.Context,
	s *state.State,
	identity immutable.Option[state.Identity],
	nodeIndex int,
) context.Context {
	return acpIdentity.WithContext(ctx, getIdentityForRequestSpecificToNode(s, identity, nodeIndex))
}

// getIdentity returns an identity for the request specific to the node.
func getIdentityForRequestSpecificToNode(
	s *state.State,
	identity immutable.Option[state.Identity],
	nodeIndex int,
) immutable.Option[acpIdentity.Identity] {
	if !identity.HasValue() {
		return acpIdentity.None
	}
	return immutable.Some(getIdentityForRequest(s, identity.Value(), nodeIndex))
}

const (
	// authTokenExpiration is the expiration time for auth tokens.
	authTokenExpiration = time.Minute * 1
)

// getIdentityForRequest returns the identity for the given reference and node index.
// It prepares the identity for a request by generating a token if needed, i.e. it will
// return an identity with [Identity.BearerToken] set.
func getIdentityForRequest(s *state.State, identity state.Identity, nodeIndex int) acpIdentity.Identity {
	identHolder := state.GetIdentityHolder(s, identity)
	ident := identHolder.Identity

	if fullIdent, ok := ident.(acpIdentity.FullIdentity); ok {
		token, ok := identHolder.NodeTokens[nodeIndex]
		if ok {
			fullIdent.SetBearerToken(token)
		} else {
			audience := getNodeAudience(s, nodeIndex)
			if documentACPType == SourceHubDocumentACPType || audience.HasValue() {
				err := fullIdent.UpdateToken(authTokenExpiration, audience, immutable.Some(s.SourcehubAddress))
				require.NoError(s.T, err)
				identHolder.NodeTokens[nodeIndex] = fullIdent.BearerToken()
			}
		}
	}
	return ident
}

func getNodeAudience(s *state.State, nodeIndex int) immutable.Option[string] {
	if nodeIndex >= len(s.Nodes) {
		return immutable.None[string]()
	}
	switch client := s.Nodes[nodeIndex].Client.(type) {
	case *http.Wrapper:
		return immutable.Some(strings.TrimPrefix(client.Host(), "http://"))
	case *cli.Wrapper:
		return immutable.Some(strings.TrimPrefix(client.Host(), "http://"))
	}

	return immutable.None[string]()
}
