
package blockmanager

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateMiningVerify(t *testing.T) {
	// initialize block manager
	bm := Blockmanager{}

	fmt.Println("Generating a genesis block (CREATE transaction)...")

	currTime := time.Now()
	genesisBlock := bm.Genesis()
	fmt.Printf("Genesis Block: %+v\n", genesisBlock)

	// 1. GENERATE TEST
	fmt.Printf("--------------\n")

	// "chain" and generate a new block
	t2 := Transaction{
		TransactionType: Create,
		TimeTransacted:  int64(time.Now().Unix()),
	}
	fmt.Printf("Generating a new EXCHANGE block on top of genesis block\n")
	block1 := bm.GenerateBlock(genesisBlock, t2)
	fmt.Printf("Block 1: %+v\n", block1)

	// try to generate a new invalid block
	fmt.Printf("Generating a new EXCHANGE (same index but faulty hash)\n")

	// the fact that this block is generated 5 seconds later will result in a different hash
	currTime = time.Now().Local().Add(time.Second * time.Duration(5))
	faultyBlock := Block{}
	faultyBlock = Block{
		Index:            1,
		Timestamp:        currTime.String(),
		Hash:             block1.Hash,
		PrevHash:         genesisBlock.Hash,
		BlockTransaction: t2,
	}
	fmt.Printf("Faulty Block 1: %+v\n", faultyBlock)

	// 2. VALIDATION TEST
	fmt.Printf("--------------\n")
	ok := bm.IsBlockValid(block1, genesisBlock)
	fmt.Printf("Is block 1 the correct successor to genesis block? (should be true): %t\n", ok)

	ok = bm.IsBlockValid(faultyBlock, genesisBlock)
	fmt.Printf("Is (faulty) block 1 the correct successor to genesis block? (should be false): %t\n", ok)

}

func TestUID(t *testing.T) {
	id := generateUID()
	fmt.Println("uid", id)
}

// if a transaction struct can succesfully pass the function parameter type check without panic, then the test is valid
func acceptTransaction(transaction Transaction) {
	fmt.Printf("Transaction type: %+v \n", transaction.TransactionType)
	fmt.Printf("Time transacted: %+v \n", transaction.TimeTransacted)
	fmt.Printf("Transaction: %+v \n", transaction)
}

func TestCreateTransaction(t *testing.T) {
	// initialize block manager
	bm := Blockmanager{}

	trans := bm.BuildCreateTransaction("Diamonds", "l33t")
	acceptTransaction(trans)
	fmt.Println("User l33t created Diamonds")
	fmt.Println()
}

func TestExchangeTransaction(t *testing.T) {
	// initialize block manager
	bm := Blockmanager{}

	trans := bm.BuildExchangeTransaction("Diamonds", "l33t", "l33a", "mike")
	acceptTransaction(trans)
	fmt.Println("Exchange diamonds from user l33t to l33a")
	fmt.Println()
}

func TestConsumeTransaction(t *testing.T) {
	// initialize block manager
	bm := Blockmanager{}

	trans := bm.BuildConsumeTransaction("Eggs", "l337")
	acceptTransaction(trans)
	fmt.Println("User l337 consumed Eggs")
	fmt.Println()
}

func TestMakeTransaction(t *testing.T) {
	// initialize block manager
	bm := Blockmanager{}

	inputItemNames := []string{"Eggs", "Milk", "Icing"}
	inputItemIds := []string{"123Eggs", "23Milkz", "icIng143"}
	trans := bm.BuildMakeTransaction(inputItemNames, inputItemIds, "Cake", "l337")
	acceptTransaction(trans)
	fmt.Println("User l337 made Cake from [Eggs, Milk, Icing]")
	fmt.Println()
}

func TestSplitTransaction(t *testing.T) {
	// initialize block manager
	bm := Blockmanager{}

	outputItemNames := []string{"1/2 Sushi roll", "1/2 Sushi roll"}
	destinationUserIds := []string{"l33a", "l33b"}
	trans := bm.BuildSplitTransaction("Sushi roll", "14SuZhiRoll", outputItemNames, "l337", destinationUserIds)
	acceptTransaction(trans)
	fmt.Println("User l337 split Sushi roll with User l33a and User l33b")
	fmt.Println()
}

func TestItemHistory(t *testing.T){
	bm := Blockmanager{}

	var bcServerTest []Block
	genesisBlock := bm.Genesis()
	bcServerTest = append(bcServerTest, genesisBlock)

	create_trans := CreateTransaction{
		OriginUserId: "-1",
		DestinationUserId: "shreya1",
		ItemId: "gucci_test",
		ItemName: "Gucci bag",
	}

	create_gucci := Transaction{
		TransactionType: Create,
		TimeTransacted:  int64(time.Now().Unix()),
		Cr: create_trans,
	}

	block1 := bm.GenerateBlock(genesisBlock, create_gucci)
	bcServerTest = append(bcServerTest, block1)

	exchange_trans := ExchangeTransaction{
		OriginUserId: "shreya1",
		DestinationUserId: "shreya2",
		ItemId: "gucci_test",
		ItemName: "Gucci bag",
	}

	exchange_gucci := Transaction{
		TransactionType: Exchange,
		TimeTransacted:  int64(time.Now().Unix()),
		Ex: exchange_trans,
	}

	block2 := bm.GenerateBlock(block1, exchange_gucci)
	bcServerTest = append(bcServerTest, block2)

	create_trans2 := CreateTransaction{
		OriginUserId: "-1",
		DestinationUserId: "shreya1",
		ItemId: "scarf_test",
		ItemName: "Scarf",
	}

	create_scarf := Transaction {
		TransactionType: Create,
		TimeTransacted: int64(time.Now().Unix()),
		Cr: create_trans2,
	}

	block3 := bm.GenerateBlock(block2, create_scarf)
	bcServerTest = append(bcServerTest, block3)

	gucci_transactions := bm.GetItemHistory("gucci_test", bcServerTest)

	if len(gucci_transactions) == 2 && len(bcServerTest) == 4 {
		fmt.Println("We gucci ;)")
	}
}


func TestItemsOfOwner(t *testing.T){
	bm := Blockmanager{}

	var bcServerTest []Block
	genesisBlock := bm.Genesis()
	bcServerTest = append(bcServerTest, genesisBlock)

	create_trans := CreateTransaction{
		OriginUserId: "-1",
		DestinationUserId: "shreya1",
		ItemId: "gucci_test",
		ItemName: "Gucci bag",
	}

	create_gucci := Transaction{
		TransactionType: Create,
		TimeTransacted:  int64(time.Now().Unix()),
		Cr: create_trans,
	}

	block1 := bm.GenerateBlock(genesisBlock, create_gucci)
	bcServerTest = append(bcServerTest, block1)

	exchange_trans := ExchangeTransaction{
		OriginUserId: "shreya1",
		DestinationUserId: "shreya2",
		ItemId: "gucci_test",
		ItemName: "Gucci bag",
	}

	exchange_gucci := Transaction{
		TransactionType: Exchange,
		TimeTransacted:  int64(time.Now().Unix()),
		Ex: exchange_trans,
	}

	block2 := bm.GenerateBlock(block1, exchange_gucci)
	bcServerTest = append(bcServerTest, block2)

	create_trans2 := CreateTransaction{
		OriginUserId: "-1",
		DestinationUserId: "shreya1",
		ItemId: "scarf_test",
		ItemName: "Scarf",
	}

	create_scarf := Transaction {
		TransactionType: Create,
		TimeTransacted: int64(time.Now().Unix()),
		Cr: create_trans2,
	}

	block3 := bm.GenerateBlock(block2, create_scarf)
	bcServerTest = append(bcServerTest, block3)

	fmt.Println(bm.GetItemsOfOwner("shreya1", bcServerTest))
}
