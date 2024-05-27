package blockchain

// proof of work

// take the data from the block

// create a counter (nonce) whcih starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s

const Difficulty = 12

type ProofOfWOrk struct {
	// Block inside the blockchain
	Block *Block
	// A number which meets the requirement 1, which gets derived
	// from Difficulty-
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// Lsh shorthand for Leftshift
	target.Lsh(target, uint(256 - Difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// replaces our previous "DeriveHash" function
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// run the function "Run" on the ProofOfWork
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	// infinite loop
	nonce := 0

	for nonce < math.MaxInt64 {
		// prepare data
		data := pow.InitData(nonce)
		// take all that data and hash it into sha256 format
		hash := sha256.Sum256(data)
		// debug output
		fmt.Printf("\r%x", hash)

		// convert hash into big integer
		intHash.SetBytes(hash[:])

		// compare proof of work target with our new
		// bigInt version of our hash
		if intHash.Cmp(pow.Target) == -1 {
			// our hash is less than the target we are
			// looking for, break out of loop
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return none, hash[:]
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
