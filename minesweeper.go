package main

import (
	"fmt"
	"math/rand"
	"time"
)

var symb = [12]rune {'O','1','2','3','4','5','6','7','8','M','X','.'}

func grid(dimx int, dimy int, op int) [][]int{
	mat := make([][]int,dimy)
	for i:=0; i< dimy; i++{
		mat[i] = make([]int,dimx)
	}
	for i:=0; i< dimy; i++{
		for j:=0; j< dimx; j++{
			if op==0{
				mat[i][j] = 0
			} else{
				mat[i][j] = 11
			}
		}
	}
	return mat
}

func printGrid(mat [][]int,dimx int, dimy int, nmines int, nbanners int, name string){
	fmt.Println("---------------------------------------")
	fmt.Println(name)
	fmt.Println("mines = ",nmines)
	fmt.Println("banners = ",nbanners)
	fmt.Println("---------------------------------------")
	for i:=0; i< dimy; i++{
		for j:=0; j< dimx; j++{
			fmt.Printf("%c",symb[mat[i][j]])
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------")
}

func putNumbers(mat [][]int,dimx int,dimy int) [][]int{
	for i:=0 ; i<dimy ; i++ {
		for j:=0 ; j<dimx ; j++ {
			if(symb[mat[i][j]]!='M'){
				z:=0
				for dy:=-1 ; dy<=1 ; dy++ {
					for dx:=-1 ; dx<=1 ; dx++ {
						x := j+dx
						y := i+dy
						if(((dx!=0)||(dy!=0))&&((x>=0)&&(x<dimx)&&(y>=0)&&(y<dimy))){
							if(symb[mat[y][x]]=='M'){
								z++
							}
						}
					}
				}
				mat[i][j]=z
			}
		}
	}
	return mat
}

func chooseCell(gmat [][]int,dimx int, dimy int) (int,int){
	var ( x,y int)
	for true{
		fmt.Print("x = ")
		fmt.Scan(&x)
		fmt.Print("y = ")
		fmt.Scan(&y)
		if((x>=0)&&(x<dimx)&&(y>=0)&&(y<dimx)&&((symb[gmat[y][x]]=='.')||(symb[gmat[y][x]]=='X'))){
			break
		}
	}
	return x , y
}

func game(mat [][]int,gmat [][]int,dimx int, dimy int, nmines int){
	var op int
	nbanners := 0
	nspaces := dimx*dimy
	for true{
		printGrid(gmat,dimx,dimy,nmines,nbanners,"Your Grid")
		fmt.Print("[1]choose square | [2]put banner | [3]quit = ")
		fmt.Scan(&op)
		if(op==1){
			x , y := chooseCell(gmat,dimx,dimy)
			nspaces--
			if(symb[mat[y][x]]=='M'){
				printGrid(mat,dimx,dimy,nmines,nbanners,"Solution of Grid")
				fmt.Println("You Lose.")
				break
			} else if(nspaces==nmines){
				printGrid(mat,dimx,dimy,nmines,nbanners,"Solution of Grid")
				fmt.Println("You Win.")
				break
			} else{
				gmat[y][x]=mat[y][x]
				printGrid(gmat,dimx,dimy,nmines,nbanners,"Your Grid")
			}
		} else if(op==2){
			x , y := chooseCell(gmat,dimx,dimy)
			switch symb[gmat[y][x]] {
			case 'X':
				gmat[y][x]=11
				nbanners--
			case '.':
				gmat[y][x]=10
				nbanners++
			}
		} else{
			break
		}
	}
}

func new_game(){
	dimy := 0
	dimx := 0
	var nmines int
	
	for ((dimx<4)||(dimy<4)) {
		fmt.Print("Number of Rows(minimum 4) : ")
		fmt.Scan(&dimy)
		fmt.Print("Number of Columns(minimum 4) : ")
		fmt.Scan(&dimx)
	}
	
	mat:=grid(dimx,dimy,0)
	gmat :=grid(dimx,dimy,1)
	
	for true{
		fmt.Print("Number of Mines : ")
		fmt.Scan(&nmines)
		if (nmines<=(dimx*dimy)){
			break
		}
	}
	
	for tmp := nmines; tmp>0 ; tmp--{
		var (x , y int)
		for true{
			x = rand.Intn(dimx)
			y = rand.Intn(dimy)
			if mat[y][x]!= 9{
				mat[y][x] = 9
				break
			}
		}
	}
	mat=putNumbers(mat,dimx,dimy)
	game(mat,gmat,dimx,dimy,nmines)
}

func title(){
	fmt.Println("=========================================")
	fmt.Println("!              MINESWEEPER              !")
	fmt.Println("=========================================")
}

func menu(){
	for true{
		fmt.Println("---------------------------------")
		fmt.Println("[1] - New Game")
		fmt.Println("[2] - Quit")
		fmt.Println("---------------------------------")
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			new_game()
		case 2:
			return
		}
	}
}

func main() {
	title()
	rand.Seed(time.Now().UnixNano())
	menu()
}