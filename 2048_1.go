/*
	The classic 2048 game, made in GoLang.

	Author:	Rupal Barman
	Date:	11/12/2016
	Email:	rupalbarman@gmail.com
	GitHub:	github.com/rupalbarman

	--------------------------------------------------------------------------------------

	-Focuses on simplicity, though the trade off is its efficiency, but
		it is necessary in order to make it complelety modular.
	-The modularity of simple actions like swipe and shift has been 
		greatly taken advantage of. (See swipeVeritical() ) 

	--------------------------------------------------------------------------------------
	

	Use only with permission.
*/

/* 
	ISSUES:
	SOLVED) Does not add in the form 0404-->swipe left-->8000 (Instead it does: 0404->4400)
		  Only adds up adjacent tiles
	2) Does not keep score, not yet.
*/

package main 

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)


func displayBoard(board [][]int) {
	for i:=0; i< len(board); i++{
		for j:=0; j<len(board[i]); j++{
			if(board[i][j]!= 0){
				fmt.Printf("|\t%d\t", board[i][j])
			}else{
				fmt.Printf("|\t \t")
			}
		}
		fmt.Printf("|\n\n")
	}
	fmt.Println("-----------------------------------------------------------------")
}

func shiftLeft(board [][]int){
	count:= 0
	for i:=0; i<len(board); i++{
		for j:=0; j<len(board[i]); j++{
				if board[i][j]!=0{
					board[i][count]= board[i][j]
					count+=1
				}
			}
		for count< len(board[i]) {
				board[i][count]= 0
				count+=1
		}
		count=0
	}
}

func shiftRight(board [][]int) {
	removeZeros:= func (slice []int, s int) []int {
		return append(slice[:s], slice[s+1:]...)
	}

	shiftSliceRight:= func (slice []int, numOfZeros int) []int{
		slice_temp:= make([]int, numOfZeros)
		slice_temp= append(slice_temp, slice[0:]...)		//the 3 dots makes it an array, 
															//ie. so that we append a fixed size element(array)
		return slice_temp
	}

	countZeros:=0

	for i:=0; i<len(board);i++{
		for j:=0; j<len(board[i]); j++{
			if board[i][j]==0{
				board[i]= removeZeros(board[i], j)
				countZeros+=1
				j-=1			//extra decrement each time to get accurate position and operation (misc)
			}
		}
		board[i]= shiftSliceRight(board[i], countZeros)
		countZeros=0
	}
}

func swipeHorizontal(board [][]int, left bool){
	if(left){
		shiftLeft(board)
	}else{
		shiftRight(board)
	}

	for i:=0; i<len(board); i++{
		for j:=0; j<len(board[i]); j++{
			if board[i][j]!=0 && j+1<len(board[i]){
				if(board[i][j]==board[i][j+1]){
					board[i][j]+= board[i][j+1]
					board[i][j+1]=0
				}
			}
			
		}
	}
	if(left){
		shiftLeft(board)
	}else{
		shiftRight(board)
	}
}

func shiftUp(board_transpose [][]int) {
	shiftLeft(board_transpose)
}

func shiftDown(board_transpose [][]int) {
	shiftRight(board_transpose)
}

func swipeVertical(board [][]int, up bool) {
	board_transpose:= make([][]int, len(board))
	for i:=0;i<len(board); i++{
		board_transpose[i]= make([]int, len(board[i]))
	}

	// Transposing it ( swapping rows and colums), so we can operate 
	// the swipe and shift operation up using swipe and shift left
	for i:=0; i<len(board); i++{
		for j:=0; j<len(board[i]); j++{
			board_transpose[i][j]= board[j][i]
		}
	}

	if(up){
		shiftUp(board_transpose)
	}else{
		shiftDown(board_transpose)
	}

	swipeHorizontal(board_transpose, true)

	if(up){
		shiftUp(board_transpose)
	}else{
		shiftDown(board_transpose)
	}

	// Transposing it again to get orginal matrix/ board
	for i:=0; i<len(board); i++{
		for j:=0; j<len(board[i]); j++{
			board[i][j]= board_transpose[j][i]
		}
	}
}

func driverMode() {
	fmt.Println("Driver Mode")

	board:= make([][]int, 4)

	for i:=0; i<len(board); i++{
		board[i]= make([]int, 4)
	}

	board[0][0]=2
	board[0][1]=4
	board[0][3]=2
	board[0][2]=2
	board[0][3]=2
	board[2][3]=2
	board[1][3]=2
	board[3][1]=2
	fmt.Println("Initial arrangement")
	//INIT
	displayBoard(board)
	//LEFT
	swipeHorizontal(board, true)
	displayBoard(board)
	//UP
	swipeVertical(board, true)
	displayBoard(board)
	//RIGHT
	swipeHorizontal(board, false)
	displayBoard(board)
	//DOWN
	swipeVertical(board, false)
	displayBoard(board)
}

func main() {

	//	Uncomment the following line to use the Driver/ debug mode.
	//	Once you uncomment the following line, make sure..
	//	you comment everything below the driverMode().

	//driverMode()

	rand.Seed(time.Now().UnixNano())
	board:= make([][]int, 4)

	for i:=0; i<len(board); i++{
		board[i]= make([]int, 4)
	}

	var key []byte= make([]byte, 1)
	//os.Stdin.Read(key)
	//fmt.Println("got byte", key, string(key))

	generateNewTwo:= func (rows, cols int) {
		new:
			x:= rand.Intn(rows)
			y:= rand.Intn(cols)

		if board[x][y]!=0{
			goto new
		}
		board[x][y]=2
	}

	//	Basic Help text. Add any other info you like.
	helpText:= func() {
		fmt.Println("Press W A S D. To exit, press E")
	}

	//	Initiate the Board with some random 2's to start the game with.
	generateNewTwo(4,4)
	generateNewTwo(4,4)
	generateNewTwo(4,4)
	generateNewTwo(4,4)
	displayBoard(board)
	helpText()

	//	The main game loop. After each action, a new 2 is generated..
	//	at a position randomely decided.
	for{
		os.Stdin.Read(key)
		switch{
		case string(key)== "w":
			//UP
			swipeVertical(board, true)
			generateNewTwo(4,4)
			displayBoard(board)
		case string(key)== "a":
			//LEFT
			swipeHorizontal(board, true)
			generateNewTwo(4,4)
			displayBoard(board)
		case string(key)== "s":
			//DOWN
			swipeVertical(board, false)
			generateNewTwo(4,4)
			displayBoard(board)
		case string(key)== "d":
			//Right
			swipeHorizontal(board, false)
			generateNewTwo(4,4)
			displayBoard(board)
		case string(key)== "e":
			return
		}	
	}

}