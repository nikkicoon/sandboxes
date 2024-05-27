package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// struct for the blocks
type Block struct {
	// position of the data record in the blockchain
	Index int
	// automatically determined, the time the data is written
	Timestamp string
	// pulse rate
	BPM int
	// SHA256 identifier of this data record
	Hash string
	// SHA256 identifier of the previous record
	PrevHash string
}

// the blockchain itself, a slice of Block
var Blockchain []Block

// create a SHA256 hash of Block data, returns the SHA256 hash as string
func calculateHash(block Block) string {
	// concatenate Index, Timestamp, BPM, and PrevHash of the
	// Block we provided as argument
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// generate a new Block
func generateBlock(oldBlock Block, BPM int) (Block, error){
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash (newBlock)

	return newBlock, nil
}

// Block Validation
// done by checking if the Index incremented as expected and by
// checking if PrevHash is the same as the Hash of the previous Block.
// Double check the hash of the current block by running calculateHash
// again on the current Block. Return a bool true if it passes all
// of these checks.
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index + 1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

// Problem handling.
// If two nodes added blocks to their chain and we receive them both.
// which one do we pick as source of truth?
// The longest chain is chosen.
// Explanation:
// Two well meaning nodes may simply have different chain lengths,
// so naturally the longer one will be the most up to date and have
// the latest blocks. So let's make sure the new chain we're taking
// in is longer than the current chain we have. If it is, we can
// overwrite our chain with the new one that has the new block(s).
//
// compare the length of the slices of the chains to accomplish this
func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

// convenience functions
// web server
func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on ", os.Getenv("ADDR"))
	s := &http.Server {
		Addr: ":" + httpAddr,
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

// Write back the full blockchain in JSON format
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// For the POST request. This allows us to simple send a POST request
// with the body {"BPM":50} and the handler will fill in the rest
// of the block for us.
type Message struct {
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	newBlock, err := generateBlock(Blockchain[len(Blockchain) - 1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if isBlockValid(newBlock, Blockchain[len(Blockchain) - 1]) {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

// when our POST requests are successgful or unsuccessful, we want to
// be alerted accordingly. This wrapper function let us know what
// happened.
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// wire it all together
func main() {
	// load the '.env' file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// separation of concerns for the genesis block, so that
	// we don't have to write it in the webserver logic.
	go func() {
		t := time.Now()
		// patient 0, the initial Block.
		genesisBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(genesisBlock)
		Blockchain = append(Blockchain, genesisBlock)
	}()
	log.Fatal(run())
}
