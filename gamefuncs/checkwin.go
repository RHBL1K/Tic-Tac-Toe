package gamefuncs


func Wincheck(board *[3][3]string, choice1 string, choice2 string,Flag *bool, Flag2 *bool, Flagdraw *bool){

		if board[0][0]==choice1 && board[0][1]==choice1 && board[0][2]==choice1{
			*Flag =true
			
		//тут слева направо  2 строка
		}else if board[1][0]==choice1 && board[1][1]==choice1 && board[1][2]==choice1{
			*Flag =true
			
		//тут слева направо  3 строка
		}else if board[2][0]==choice1 && board[2][1]==choice1 && board[2][2]==choice1{
			*Flag =true
		
		//тут сверху вниз  1 столбик
		}else if board[0][0]==choice1 && board[1][0]==choice1 && board[2][0]==choice1{
			*Flag =true
		
		//тут сверху вниз  2 столбик
		}else if board[0][1]==choice1 && board[1][1]==choice1 && board[2][1]==choice1{
			*Flag =true
	
		//тут сверху вниз  3 столбик
		}else if board[0][2]==choice1 && board[1][2]==choice1 && board[2][2]==choice1{
			*Flag =true

		//первая диагональ сверху вниз справо на лево
		}else if board[0][0]==choice1 && board[1][1]==choice1 && board[2][2]==choice1{
			*Flag =true
		
		//вторая диагональ снизу вверх справо на лево
		}else if board[2][0]==choice1 && board[1][1]==choice1 && board[0][2]==choice1{
			*Flag =true
		
		}
		





		if board[0][0]==choice2 && board[0][1]==choice2 && board[0][2]==choice2{
			*Flag2 =true
		
		//тут слева направо  2 строка
		}else if board[1][0]==choice2 && board[1][1]==choice2 && board[1][2]==choice2{
			*Flag2 =true
			
		//тут слева направо  3 строка
		}else if board[2][0]==choice2 && board[2][1]==choice2 && board[2][2]==choice2{
			*Flag2 =true
		
		//тут сверху вниз  1 столбик
		}else if board[0][0]==choice2 && board[1][0]==choice2 && board[2][0]==choice2{
			*Flag2 =true
	
		//тут сверху вниз  2 столбик
		}else if board[0][1]==choice2 && board[1][1]==choice2 && board[2][1]==choice2{
			*Flag2 =true
		
		//тут сверху вниз  3 столбик
		}else if board[0][2]==choice2 && board[1][2]==choice2 && board[2][2]==choice2{
			*Flag2 =true
		
		//первая диагональ сверху вниз справо на лево
		}else if board[0][0]==choice2 && board[1][1]==choice2 && board[2][2]==choice2{
			*Flag2 =true
		
		//вторая диагональ снизу вверх справо на лево
		}else if board[2][0]==choice2 && board[1][1]==choice2 && board[0][2]==choice2{
			*Flag2 =true
	
		}



		
		if !*Flag && !*Flag2 { // Если никто не выиграл
			isDraw := true
			// Проверяем, есть ли еще пустые клетки
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if board[i][j] == " " {
						isDraw = false
						break
					}
				}
			}
			*Flagdraw = isDraw
		}
		


}